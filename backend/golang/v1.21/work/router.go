package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"

	"idv/chris/errs"
	"idv/chris/utils"
)

// ValidateUserName 驗證使用者名稱
func ValidateUserName(username string) bool {
	regex := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9]{3,7}$`)
	return regex.MatchString(username)
}

// ValidatePassword 驗證密碼
func ValidatePassword(password string) bool {
	regex := regexp.MustCompile(`^[a-zA-Z0-9]{3,7}$`)
	return regex.MatchString(password)
}

// ValidateNickname 驗證暱稱 (包含繁體中文)
func ValidateNickname(nickname string) bool {
	regex := regexp.MustCompile(`^[a-zA-Z\p{Han}][a-zA-Z0-9\p{Han}]{1,10}$`)
	return regex.MatchString(nickname)
}

func router_api(g *gin.Engine) {
	sub := g.Group("/api/user")
	sub.POST("/register", func(c *gin.Context) {
		const msg = "api.register"

		var e error
		defer func() {
			if e != nil {
				comps.Logger().Error(msg, zap.Error(e))
				c.JSON(
					http.StatusOK,
					HttpResponse{
						Code: e.Error(),
					},
				)
				return
			}
			c.JSON(
				http.StatusOK,
				HttpResponse{
					Code: string(OK),
				},
			)
		}()

		var param = map[string]string{}
		e = c.BindJSON(&param)
		if e != nil {
			e = fmt.Errorf(string(errs.E0000))
			return
		}

		comps.Logger().Debug(msg, zap.Any("body", param))

		uname := param["name"]
		upw := param["password"]
		unickname := param["nickname"]

		if !ValidateUserName(uname) {
			e = fmt.Errorf(string(errs.E0001))
			return
		}

		if !ValidatePassword(upw) {
			e = fmt.Errorf(string(errs.E0002))
			return
		}

		if !ValidateNickname(unickname) {
			e = fmt.Errorf(string(errs.E0003))
			return
		}

		ukey := utils.MD5Encode(strings.Join([]string{uname, upw}, "|"))
		e = comps.Register(uname, upw, unickname, ukey, c.ClientIP())
		if e != nil {
			return
		}
	})

	sub.POST("/login", func(c *gin.Context) {
		const msg = "api.login"

		var e error
		var content = map[string]interface{}{}
		defer func() {
			if e != nil {
				comps.Logger().Error(msg, zap.Error(e))
				c.JSON(
					http.StatusOK,
					HttpResponse{
						Code: e.Error(),
					},
				)
				return
			}
			c.JSON(
				http.StatusOK,
				HttpResponse{
					Code:    string(OK),
					Content: content,
				},
			)
		}()

		var param = map[string]string{}
		e = c.BindJSON(&param)
		if e != nil {
			e = fmt.Errorf(string(errs.E0000))
			return
		}

		comps.Logger().Debug(msg, zap.Any("body", param))

		uname := param["name"]
		upw := param["password"]

		if !ValidateUserName(uname) {
			e = fmt.Errorf(string(errs.E0001))
			return
		}

		if !ValidatePassword(upw) {
			e = fmt.Errorf(string(errs.E0002))
			return
		}

		ukey := utils.MD5Encode(strings.Join([]string{uname, upw}, "|"))

		u, e := comps.Login(ukey)
		if e != nil {
			return
		}

		var token string
		token, e = utils.RSAEncode([]byte(ukey))
		if e != nil {
			comps.Logger().Error(msg, zap.Error(e))
			e = fmt.Errorf(string(errs.E0004))
			return
		}

		comps.Logger().Debug(msg, zap.Any("user", u))

		comps.SetToken(token, ukey, time.Minute*5)

		content["Token"] = token

		if u.ULv != "0" {
			return
		}
		content["Menu"] = map[string]string{
			"/admin/shop/add.html":  "增加店家",
			"/admin/shop/list.html": "店家清單",
		}
	})

	sub = g.Group("/api/admin/shop")
	sub.POST("/add", func(c *gin.Context) {
		const msg = "api.shop.add"

		var e error
		defer func() {
			if e != nil {
				comps.Logger().Error(msg, zap.Error(e))
				c.JSON(
					http.StatusOK,
					HttpResponse{
						Code: e.Error(),
					},
				)
				return
			}
			c.JSON(
				http.StatusOK,
				HttpResponse{
					Code: string(OK),
				},
			)
		}()

		var param = map[string]interface{}{}
		e = c.BindJSON(&param)
		if e != nil {
			e = fmt.Errorf(string(errs.E0000))
			return
		}
		comps.Logger().Debug(msg, zap.Any("body", param))

		name := param["shop_name"].(string)
		mobile := param["shop_mobile"].(string)

		id := strings.ReplaceAll(uuid.NewString(), "-", "_")
		e = comps.AddShop(id, name, mobile)
		if e != nil {
			return
		}

		for key, value := range param["list"].([]interface{}) {
			tmp := value.(map[string]interface{})
			tmp["id"] = strings.ReplaceAll(uuid.NewString(), "-", "_")
			fmt.Println(key, value)
		}
		var output []byte
		output, e = json.Marshal(param["list"])
		if e != nil {
			e = fmt.Errorf(string(errs.E0000))
			return
		}
		e = utils.FileWrite("./asset/shop", fmt.Sprintf("shop_%v", id), output)
		if e != nil {
			e = fmt.Errorf(string(errs.E0007))
			return
		}

	})
	sub.POST("/get", func(c *gin.Context) {
		const msg = "api.shop.get"

		var e error
		var content = map[string]interface{}{}
		defer func() {
			if e != nil {
				comps.Logger().Error(msg, zap.Error(e))
				c.JSON(
					http.StatusOK,
					HttpResponse{
						Code: e.Error(),
					},
				)
				return
			}
			c.JSON(
				http.StatusOK,
				HttpResponse{
					Code:    string(OK),
					Content: content,
				},
			)
		}()

		list, e := comps.GetShop()
		if e != nil {
			return
		}

		content["list"] = list
	})
}

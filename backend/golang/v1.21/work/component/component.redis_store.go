package component

import (
	"fmt"
	"time"

	"go.uber.org/zap"

	"idv/chris/errs"
)

/*
ErrorNil

@return {true:空值錯誤, false:非空值錯誤}
*/
func ErrorRedisNil(e error) bool {
	return e.Error() == "redis: nil"
}

type RedisStore interface {
	SetToken(token, ukey string, expiration time.Duration) (e error)
	DelToken(token string) (e error)
	GetToken(token string) (ukey string, e error)
}

func (s *store) SetToken(token, ukey string, expiration time.Duration) (e error) {
	const msg = "set.token"
	cmd := s.rc.Set(token, ukey, expiration)
	e = cmd.Err()
	if e != nil {
		s.Logger().Error(msg, zap.Error(e))
		e = fmt.Errorf(string(errs.E2000))
	}
	return e
}

func (s *store) DelToken(token string) (e error) {
	const msg = "del.token"
	cmd := s.rc.Del(token)
	e = cmd.Err()
	if e != nil {
		s.Logger().Error(msg, zap.Error(e))
		e = fmt.Errorf(string(errs.E2001))
	}
	return e
}

func (s *store) GetToken(token string) (ukey string, e error) {
	const msg = "check.token"
	cmd := s.rc.Get(token)
	e = cmd.Err()
	if e != nil {
		s.Logger().Error(msg, zap.Error(e))
		e = fmt.Errorf(string(errs.E2002))
		return
	}
	ukey = cmd.Val()
	e = nil
	return
}

package component

import (
	"fmt"
	"time"

	"go.uber.org/zap"

	"idv/chris/errs"
)

type SQLStore interface {
	Register(name, password, nickname, ukey, ip string) (e error)
	Login(ukey string) (u user, e error)
	IsAdmin(ukey string) (u user, e error)
	AddShop(uuid, name, mobile string) (e error)
	GetShop() (list []Shop, e error)
}

func (s *store) Register(name, password, nickname, ukey, ip string) (e error) {
	const msg = "sql.register"
	query := "INSERT INTO `User` (`UName`, `UPassword`, `UNickname`, `UKey`, `LastIP`, `CreatedAt`) VALUES (?, ?, ?, ?, ?, ? )"

	_, e = s.db.Exec(
		query,

		name,
		password,
		nickname,
		ukey,
		ip,
		time.Now().UTC(),
	)
	if e != nil {
		s.Logger().Error(msg, zap.Error(e))
		e = fmt.Errorf(string(errs.E1000))
		return
	}

	return
}

func (s *store) Login(ukey string) (u user, e error) {
	const msg = "sql.login"

	query := "SELECT count(*), `UNickname`, `ULv` FROM `User` WHERE `UKey` = ?"
	row := s.db.QueryRow(query, ukey)

	total := 0
	e = row.Scan(
		&total,
		&u.UNickname,
		&u.ULv,
	)
	if e != nil {
		s.Logger().Error(msg, zap.Error(e))
		e = fmt.Errorf(string(errs.E1001))
		return
	}
	if total == 0 {
		e = fmt.Errorf(string(errs.E1001))
		return
	}
	return
}

func (s *store) IsAdmin(ukey string) (u user, e error) {
	const msg = "sql.is_admin"

	query := "SELECT count(*), `ULv` FROM `User` WHERE `UKey` = ?"
	row := s.db.QueryRow(query, ukey)

	total := 0
	e = row.Scan(
		&total,
		&u.ULv,
	)
	if e != nil {
		s.Logger().Error(msg, zap.Error(e))
		e = fmt.Errorf(string(errs.E1002))
		return
	}
	if total == 0 {
		e = fmt.Errorf(string(errs.E1002))
		return
	}
	return
}

func (s *store) AddShop(uuid, name, mobile string) (e error) {
	const msg = "sql.add.admin"

	query := "INSERT INTO `Shop` (`UUID`,`Name`, `Mobile`, `Actived`) VALUES (?, ?, ?, ?)"

	_, e = s.db.Exec(
		query,
		uuid,
		name,
		mobile,
		1,
	)

	if e != nil {
		s.Logger().Error(msg, zap.Error(e))
		e = fmt.Errorf(string(errs.E1003))
		return
	}

	return
}

func (s *store) GetShop() (list []Shop, e error) {
	const msg = "sql.get.shop"

	query := "SELECT `UUID`, `Name`, `Mobile`, `Actived` FROM `Shop` WHERE `Actived`='1'"
	rows, e := s.db.Query(query)

	for rows.Next() {
		var item Shop
		e = rows.Scan(
			&item.UUID,
			&item.Name,
			&item.Mobile,
			&item.Actived,
		)
		if e != nil {
			s.Logger().Error(msg, zap.Error(e))
			e = fmt.Errorf(string(errs.E1004))
			return
		}
		list = append(list, item)
	}

	if e != nil {
		s.Logger().Error(msg, zap.Error(e))
		e = fmt.Errorf(string(errs.E1004))
		return
	}

	return
}

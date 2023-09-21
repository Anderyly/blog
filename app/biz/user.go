/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2023
 */

package biz

import (
	"blog/ay"
	"blog/dal/model"
	"blog/dal/query"
)

type user struct {
	*base
}

func (u user) ListByIds(ids []int64) (list []*model.Users, err error) {
	qe := query.Use(ay.Db).Users
	if list, err = qe.WithContext(u.ctx).Where(qe.Uid.In(ids...)).Find(); err != nil {
		return
	}
	return
}

func (u user) GetById(userId int64) (user *model.Users, err error) {
	qe := query.Use(ay.Db).Users
	if user, err = qe.WithContext(u.ctx).Where(qe.Uid.Eq(userId)).Take(); err != nil {
		return
	}
	return
}

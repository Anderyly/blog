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

type views struct {
	*base
}

func (con views) Create(data *model.Views) (err error) {
	db := ay.Begin()
	defer func() {
		if err != nil {
			db.Rollback()
			return
		}
		db.Commit()
	}()

	qe := query.Use(db).Views
	if err = qe.WithContext(con.ctx).Create(data); err != nil {
		return
	}

	if data.Cid != 0 {
		var count int64
		if count, err = con.CountByCidAndIp(data.Cid, data.Ip); err != nil {
			return
		}
		if count > 1 {
			return
		}
		if err = con.biz.Contents.AddViews(db, data.Cid); err != nil {
			return
		}
	}

	return

}

func (con views) CountByCidAndIp(contentsId int64, ip string) (count int64, err error) {
	qe := query.Use(ay.Db).Views
	if count, err = qe.WithContext(con.ctx).Where(qe.Ip.Eq(ip), qe.Cid.Eq(contentsId)).Count(); err != nil {
		return
	}
	return
}

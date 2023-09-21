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

type templates struct {
	*base
}

func (con templates) Get(name string) (res *model.Templates, err error) {
	qe := query.Use(ay.Db).Templates
	if res, err = qe.WithContext(con.ctx).Where(qe.K.Eq(name)).Take(); err != nil {
		return
	}
	return
}

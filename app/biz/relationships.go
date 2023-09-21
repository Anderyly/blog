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

type relationships struct {
	*base
}

func (con relationships) List(contentIds []int64) (list []*model.Relationships, err error) {
	qe := query.Use(ay.Db).Relationships
	if list, err = qe.WithContext(con.ctx).Where(qe.Cid.In(contentIds...)).Find(); err != nil {
		return
	}
	return
}

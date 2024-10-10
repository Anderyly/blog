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

type links struct {
	*base
}

func (con links) List() (list []*model.Links, err error) {
	qe := query.Use(ay.Db).Links
	if list, err = qe.WithContext(con.ctx).Find(); err != nil {
		return
	}
	return

}

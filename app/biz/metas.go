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

type metas struct {
	*base
}

func (con metas) ListCategoryByIds(ids []int64) (list []*model.Metas, err error) {
	qe := query.Use(ay.Db).Metas
	if list, err = qe.WithContext(con.ctx).Where(qe.Mid.In(ids...)).Order(qe.Sort.Desc()).Find(); err != nil {
		return
	}
	return
}

func (con metas) ListCategory() (list []*model.Metas, err error) {
	qe := query.Use(ay.Db).Metas
	if list, err = qe.WithContext(con.ctx).Where(qe.Type.Eq(int(model.MetasTypeCategory))).Order(qe.Sort.Desc()).Find(); err != nil {
		return
	}
	return
}

// 获取Tag
func (con metas) ListTag() (list []*model.Metas, err error) {
	qe := query.Use(ay.Db).Metas
	if list, err = qe.WithContext(con.ctx).Where(qe.Type.Eq(int(model.MetasTypeTag))).Order(qe.Sort.Desc()).Find(); err != nil {
		return
	}
	return
}

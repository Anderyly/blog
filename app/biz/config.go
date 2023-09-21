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

type config struct {
	*base
}

func (c config) List() (res map[string]string, err error) {
	qe := query.Use(ay.Db).Config
	var list []*model.Config
	if list, err = qe.WithContext(c.ctx).Find(); err != nil {
		return
	}

	res = map[string]string{}
	for _, v := range list {
		res[v.K] = v.V
	}
	return
}

func (c config) Get(k string) (res *model.Config, err error) {
	qe := query.Use(ay.Db).Config
	if res, err = qe.WithContext(c.ctx).Where(qe.K.Eq(k)).Take(); err != nil {
		return
	}
	return
}

func (c config) Set(k, v string) (err error) {
	var res *model.Config
	qe := query.Use(ay.Db).Config
	if res, err = ay.IgnoreNotFoundReturn(c.Get(k)); err != nil {
		return
	}
	if res == nil {
		if err = qe.WithContext(c.ctx).Create(&model.Config{
			K: k,
			V: v,
		}); err != nil {
			return
		}
	} else {
		if _, err = qe.WithContext(c.ctx).Where(qe.K.Eq(k)).Updates(&model.Config{V: v}); err != nil {
			return
		}
	}
	return
}

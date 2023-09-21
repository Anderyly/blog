/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2023
 */

package custom

import (
	"blog/app/biz"
	"blog/ay"
	"blog/dal/model"
	"context"
	"github.com/gin-gonic/gin"
	"html/template"
)

func Category(r *gin.Engine) {
	r.SetFuncMap(template.FuncMap{
		// 获取分类
		"ListCategory": func() (list []*model.Metas) {
			b := biz.NewBiz(context.Background())
			var err error
			list, err = b.Metas.ListCategory()
			if err != nil {
				ay.Logger.Error(err.Error())
			}
			return
		},
	})
}

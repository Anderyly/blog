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

func Contents(r *gin.Engine) {
	r.SetFuncMap(template.FuncMap{
		// 获取最新文章
		"ListContentsNew": func(limit int) (list []*model.Contents) {
			b := biz.NewBiz(context.Background())
			var err error
			list, err = b.Contents.ListNew(limit)
			if err != nil {
				ay.Logger.Error(err.Error())
			}
			return
		},
		// 获取热门文章
		"ListContentsHot": func(limit int) (list []*model.Contents) {
			b := biz.NewBiz(context.Background())
			var err error
			list, err = b.Contents.ListHot(limit)
			if err != nil {
				ay.Logger.Error(err.Error())
			}
			return
		},
	})
}

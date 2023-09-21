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

func Comments(r *gin.Engine) {
	r.SetFuncMap(template.FuncMap{
		// 获取最新评论
		"ListCommentsNew": func(limit int) (list []*model.Comments) {
			b := biz.NewBiz(context.Background())
			var err error
			list, err = b.Comments.ListNew(limit, b.Comments.WithSetContents())
			if err != nil {
				ay.Logger.Error(err.Error())
			}
			return
		},
	})
}

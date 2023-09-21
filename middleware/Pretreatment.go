/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package middleware

import (
	"blog/app/biz"
	"blog/app/controllers"
	"blog/ay"
	"context"
	"github.com/gin-gonic/gin"
)

func Pretreatment() gin.HandlerFunc {
	return func(c *gin.Context) {
		b := biz.NewBiz(context.Background())
		siteConf, err := b.Config.List()
		if err != nil {
			ay.Logger.Error(err.Error())
		}
		controllers.SiteConf = siteConf
		c.Next()
	}
}

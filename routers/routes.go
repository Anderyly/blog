package routers

import (
	"blog/app/biz"
	"blog/ay"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Instance(r *gin.Engine) {
	IndexRouter(r)
	SitemapRouter(r)

	b := biz.NewBiz(context.Background())
	template, err := b.Config.Get("template")
	if err != nil {
		ay.Logger.Error(err.Error())
	}
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, template.V+"/404", gin.H{})
	})

}

/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2023
 */

package index

import (
	"blog/app/biz"
	"blog/ay"
	"blog/dal/model"
	"blog/app/controllers"
	"blog/utils"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type IndexController struct {
}

func (con IndexController) Index(c *gin.Context) {
	queryPage := c.Param("page")
	queryKeywords := c.Query("s")
	queryCategory := c.Param("category")
	queryTag := c.Param("tag")

	if queryPage == "" {
		queryPage = "0"
	}
	page, err := strconv.Atoi(queryPage)
	if err != nil {
		ay.Logger.Error(err.Error())
		return
	}

	key := queryTag
	if queryKeywords != "" {
		key = queryKeywords
	}

	args := biz.ListContentsArgs{
		Page:     page,
		Keyword:  key,
		Category: queryCategory,
	}
	b := biz.NewBiz(context.Background())
	list, count, err := b.Contents.List(args, b.Contents.WithSetCategory(), b.Contents.WithSetUser())
	if err != nil {
		ay.Logger.Error(err.Error())
		return
	}

	if page == 0 {
		page = 1
	}

	c.HTML(http.StatusOK, "default/index", gin.H{
		"siteConf": controllers.SiteConf,
		"contents": list,
		"pages":    utils.Page(page, int(count), 10),
		"page":     page,
		"category": queryCategory,
		"keywords": queryKeywords,
	})
}

func (con IndexController) Archives(c *gin.Context) {
	querySlug := c.Param("slug")
	querySlug = strings.ReplaceAll(querySlug, ".html", "")

	b := biz.NewBiz(context.Background())
	res, err := ay.IgnoreNotFoundReturn(b.Contents.GetBySlug(querySlug, b.Contents.WithSetCategory(), b.Contents.WithSetUser()))
	if err != nil {
		ay.Logger.Error(err.Error())
		return
	}

	if res == nil {
		c.HTML(http.StatusOK, "default/404", gin.H{
			"siteConf": controllers.SiteConf,
		})
		return
	}

	go func() {
		if err = b.Views.Create(&model.Views{
			Cid:       res.Cid,
			Agent:     c.Request.UserAgent(),
			Ip:        c.ClientIP(),
			CreatedAt: time.Now().Unix(),
		}); err != nil {
			ay.Logger.Error(err.Error())
			return
		}
	}()

	//lib.NewJson(c).Success(res)
	//return

	c.HTML(http.StatusOK, "default/archives", gin.H{
		"siteConf": controllers.SiteConf,
		"contents": res,
		"category": res.Category.Slug,
	})
}

/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2023
 */

package routers

import (
	"blog/app/controllers/index"
	"github.com/gin-gonic/gin"
)

func SitemapRouter(r *gin.Engine) {

	router := r.Group("/")
	router.GET("sitemap.xml", index.SitemapController{}.Index)
	router.GET("post-sitemap.xml", index.SitemapController{}.Post)
	router.GET("page-sitemap.xml", index.SitemapController{}.Page)
	router.GET("category-sitemap.xml", index.SitemapController{}.Category)
	router.GET("post_tag-sitemap.xml", index.SitemapController{}.Tag)
	//router.GET("author-sitemap.xml", index.SitemapController{}.Author)
}

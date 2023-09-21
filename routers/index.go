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

func IndexRouter(r *gin.Engine) {

	router := r.Group("/")
	router.GET("", index.IndexController{}.Index)
	router.GET("page/:page", index.IndexController{}.Index)
	router.GET("category/:category", index.IndexController{}.Index)
	router.GET("archives/tag/:tag", index.IndexController{}.Index)
	
	router.GET("archives/:slug", index.IndexController{}.Archives)
}

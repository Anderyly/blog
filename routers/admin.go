/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2023
 */

package routers

import (
	"blog/app/controllers/admin"
	"github.com/gin-gonic/gin"
)

func AdminRouter(r *gin.Engine) {

	router := r.Group("/admin")
	router.GET("login", admin.LoginController{}.Login)
	router.POST("login", admin.LoginController{}.Login)
}

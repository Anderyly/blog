package middleware

import (
	"github.com/gin-gonic/gin"
)

func Instance(r *gin.Engine) {
	//r.Use(Logger(ay.Logger))
	//r.Use(Recovery(ay.Logger, true))
	r.Use(Cors())
	r.Use(Pretreatment())
}

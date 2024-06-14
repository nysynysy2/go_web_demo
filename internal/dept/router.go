package dept

import "github.com/gin-gonic/gin"

func Routes(r *gin.Engine) {
	dr := r.Group("/depts")
	dr.GET("", GetAll)
	dr.POST("", Create)
	dr.GET("/:id", GetById)
	dr.PUT("", Update)
}

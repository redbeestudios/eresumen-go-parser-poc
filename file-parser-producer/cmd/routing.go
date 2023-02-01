package cmd

import (
	"github.com/gin-gonic/gin"
)

type request struct {
	Path string `json:"path"`
}

func InitRoutes(dependencies *Dependencies) *gin.Engine {
	r := gin.New()
	r.POST("/eResumen/processMastercard", func(c *gin.Context) {
		var r request
		if err := c.BindJSON(&r); err != nil {
			c.JSON(400, "Bad Request")
			return
		}
		path := r.Path
		dependencies.MasterCardController.ProcessMastercard(c, path)
		c.JSON(200, "Ok")
		return
	})

	return r
}

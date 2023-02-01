package cmd

import (
	"github.com/gin-gonic/gin"
	dto "github.com/redbeestudios/go-parser-poc/commons/dto/mastercard"
)

func InitRoutes(dependencies *Dependencies) *gin.Engine {
	r := gin.New()
	r.GET("/eResumen/:id", func(c *gin.Context) {
		id := c.Param("id")
		response, err := dependencies.Controller.GetResumen(c, id)
		if err != nil {
			c.JSON(500, "Message: ocurrio un error")
			return
		}
		c.JSON(200, response)
		return
	})
	r.POST("/eResumen", func(c *gin.Context) {
		var resumenDto dto.ResumenAEnviar
		if err := c.ShouldBindJSON(&resumenDto); err != nil {
			c.JSON(400, "Bad Request")
			return
		}

		if err := dependencies.Controller.SaveResumen(c, &resumenDto); err != nil {
			c.JSON(500, "Ocurrio un error")
		}

	})
	return r
}

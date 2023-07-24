package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thormesfap/jornada-milhas/controllers"

)

func HandleRequest() {
	r := gin.Default()
	r.GET("/api/depoimentos", controllers.TodosDepoimentos)
	r.POST("/api/depoimentos", controllers.CriaDepoimento)
	r.GET("/api/depoimentos-home", controllers.DepoimentosHome)
	r.GET("/api/depoimentos/:id", controllers.RetornaDepoimento)
	r.PATCH("/api/depoimentos/:id", controllers.EditaDepoimento)
	r.DELETE("/api/depoimentos/:id", controllers.DeletaDepoimento)
	r.Run()
}

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thormesfap/jornada-milhas/controllers"

)

func HandleRequest() {
	r := gin.Default()
	depoimentos := r.Group("/api/depoimentos")
	{
		depoimentos.GET("/", controllers.TodosDepoimentos)
		depoimentos.POST("/", controllers.CriaDepoimento)
		depoimentos.GET("/:id", controllers.RetornaDepoimento)
		depoimentos.PATCH("/:id", controllers.EditaDepoimento)
		depoimentos.POST("/:id", controllers.AdicionaFotoAoDepoimento)
		depoimentos.DELETE("/:id", controllers.DeletaDepoimento)
	}
	r.GET("/api/depoimentos-home", controllers.DepoimentosHome)
	destinos := r.Group("/api/destinos")
	{
		destinos.GET("/", controllers.TodosDestinos)
		destinos.POST("/", controllers.CriaDestino)
		destinos.GET("/:id", controllers.RetornaDestino)
		destinos.PATCH("/:id", controllers.EditaDestino)
		destinos.DELETE("/:id", controllers.DeletaDestino)
	}

	r.Run()
}

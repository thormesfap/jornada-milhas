package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thormesfap/jornada-milhas/database"
	"github.com/thormesfap/jornada-milhas/models"
)

func TodosDestinos(c *gin.Context) {

	var d []models.Destino
	nome := c.Query("nome")
	if nome == "" {
		database.DB.Find(&d)
	} else {
		database.DB.Where(&models.Destino{Nome: nome}).Find(&d)
		if len(d) == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"mensagem": "Nenhum destino encontrado com nome \"" + nome + "\""})
			return
		}
	}
	c.JSON(http.StatusOK, d)
}

func RetornaDestino(c *gin.Context) {

	id := c.Params.ByName("id")
	var d models.Destino
	database.DB.First(&d, id)
	if d.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Destino não encontrado"})
		return
	}
	c.JSON(http.StatusOK, d)
}
func CriaDestino(c *gin.Context) {
	var destino models.Destino
	if err := c.ShouldBindJSON(&destino); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	if err := models.ValidateDestino(&destino); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&destino)
	c.JSON(http.StatusOK, destino)
}
func DeletaDestino(c *gin.Context) {
	id := c.Params.ByName("id")
	var p models.Destino
	database.DB.Delete(&p, id)
	c.JSON(http.StatusOK, gin.H{"message": "Destino apagado com sucesso"})
}
func EditaDestino(c *gin.Context) {

	id := c.Params.ByName("id")
	var d models.Destino
	database.DB.First(&d, id)
	if d.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Destino não encontrado"})
		return
	}
	if err := c.ShouldBindJSON(&d); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	if err := models.ValidateDestino(&d); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Model(&d).UpdateColumns(d)
	c.JSON(http.StatusOK, d)
}

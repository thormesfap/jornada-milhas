package controllers

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/thormesfap/jornada-milhas/database"
	"github.com/thormesfap/jornada-milhas/models"
)

func TodosDepoimentos(c *gin.Context) {

	var d []models.Depoimento
	database.DB.Find(&d)
	c.JSON(http.StatusOK, d)
}

func DepoimentosHome(c *gin.Context) {
	var d []models.Depoimento
	database.DB.Order("created_at DESC").Limit(3).Find(&d)
	c.JSON(http.StatusOK, d)
}
func RetornaDepoimento(c *gin.Context) {

	id := c.Params.ByName("id")
	var d models.Depoimento
	database.DB.First(&d, id)
	if d.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Depoimento não encontrado"})
		return
	}
	c.JSON(http.StatusOK, d)
}
func CriaDepoimento(c *gin.Context) {
	var depoimento models.Depoimento
	if err := c.ShouldBindJSON(&depoimento); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	depoimento.Foto = ""
	if err := models.ValidateDepoimento(&depoimento); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&depoimento)
	c.JSON(http.StatusOK, depoimento)
}

func AdicionaFotoAoDepoimento(c *gin.Context) {
	id := c.Params.ByName("id")
	var d models.Depoimento
	database.DB.First(&d, id)
	if d.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Depoimento não encontrado"})
		return
	}
	folder := "./public/imagens/depoimentos/"
	file, err := c.FormFile("imagem")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	extensao := filepath.Ext(file.Filename)
	novoNome := uuid.New().String() + extensao
	if err := c.SaveUploadedFile(file, folder + novoNome); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Não foi possível salvar o arquivo: " + err.Error(),
		})
		return
	}
	d.Foto = folder + novoNome
	database.DB.Save(d)
	c.JSON(http.StatusOK, gin.H{
		"mensagem": "Foto salva com sucesso.",
	})

}

func DeletaDepoimento(c *gin.Context) {
	id := c.Params.ByName("id")
	var p models.Depoimento
	database.DB.Delete(&p, id)
	c.JSON(http.StatusOK, gin.H{"message": "Depoimento apagado com sucesso"})
}
func EditaDepoimento(c *gin.Context) {

	id := c.Params.ByName("id")
	var d models.Depoimento
	database.DB.First(&d, id)
	if d.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Depoimento não encontrado"})
		return
	}
	var dAtualizado models.Depoimento
	if err := c.ShouldBindJSON(&dAtualizado); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	if err := models.ValidateDepoimento(&dAtualizado); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Model(&d).UpdateColumns(dAtualizado)
	c.JSON(http.StatusOK, d)
}

package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/thormesfap/jornada-milhas/database"
	"github.com/thormesfap/jornada-milhas/models"
	"github.com/thormesfap/jornada-milhas/controllers"
)

var ID int

func SetupDasRotasDeTeste() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas
}

func TestListandoTodosDepoimentos(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaDepoimentoMock()
	defer DeletaDepoimentoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/depoimentos", controllers.TodosDepoimentos)
	req, _ := http.NewRequest("GET", "/depoimentos", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t,http.StatusOK, resposta.Code)
}

func TestBuscaDepoimentoPorIDHandler(t *testing.T){
	database.ConectaComBancoDeDados()
	CriaDepoimentoMock()
	defer DeletaDepoimentoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/depoimentos/:id", controllers.RetornaDepoimento)
	req, _ := http.NewRequest("GET", "/depoimentos/" + strconv.Itoa(ID), nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t,http.StatusOK, resposta.Code)
}

func TestAtualizaDepoimentoPorIDHandler(t *testing.T){
	database.ConectaComBancoDeDados()
	CriaDepoimentoMock()
	defer DeletaDepoimentoMock()
	r := SetupDasRotasDeTeste()
	r.PATCH("/depoimentos/:id", controllers.EditaDepoimento)
	depoimento := models.Depoimento{Autor:"Autor de Teste para Edição", Foto:"89", Depoimento:"Depoimento alterado com sucesso através do teste"}
	body, _ := json.Marshal(depoimento)
	path := "/depoimentos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("PATCH", path, bytes.NewBuffer(body))
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t,http.StatusOK, resposta.Code)
	var DepoimentoMock models.Depoimento
	json.Unmarshal(resposta.Body.Bytes(), &DepoimentoMock)
	assert.Equal(t, "Autor de Teste para Edição", DepoimentoMock.Autor)
	assert.Equal(t, "89", DepoimentoMock.Foto)
	assert.Equal(t, "Depoimento alterado com sucesso através do teste", DepoimentoMock.Depoimento)
}

func TestCriaDepoimentoHandler(t *testing.T){
	database.ConectaComBancoDeDados()
	r := SetupDasRotasDeTeste()
	r.POST("/depoimentos", controllers.CriaDepoimento)
	depoimento := models.Depoimento{Autor:"Autor de Teste para Criação", Foto:"189", Depoimento:"Depoimento criado com sucesso através do teste"}
	body, _ := json.Marshal(depoimento)
	path := "/depoimentos"
	req, _ := http.NewRequest("POST", path, bytes.NewBuffer(body))
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t,http.StatusOK, resposta.Code)
	var DepoimentoMock models.Depoimento
	json.Unmarshal(resposta.Body.Bytes(), &DepoimentoMock)
	assert.NotEqual(t, 0, DepoimentoMock.ID)
	assert.Equal(t, "Autor de Teste para Criação", DepoimentoMock.Autor)
	assert.Equal(t, "189", DepoimentoMock.Foto)
	assert.Equal(t, "Depoimento criado com sucesso através do teste", DepoimentoMock.Depoimento)
	if DepoimentoMock.ID != 0{
		ID = int(DepoimentoMock.ID)
		defer DeletaDepoimentoMock()
	}
}

func TestDeletaDepoimentoPorIDHandler(t *testing.T){
	database.ConectaComBancoDeDados()
	CriaDepoimentoMock()
	r := SetupDasRotasDeTeste()
	r.DELETE("/depoimentos/:id", controllers.DeletaDepoimento)
	req, _ := http.NewRequest("DELETE", "/depoimentos/" + strconv.Itoa(ID), nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t,http.StatusOK, resposta.Code)
}


func CriaDepoimentoMock(){
	depoimento := models.Depoimento{Autor:"Autor de Teste", Depoimento:"Depoimento de Teste para criação", Foto:"53"}
	database.DB.Create(&depoimento)
	ID = int(depoimento.ID)
}

func DeletaDepoimentoMock(){
	var depoimento models.Depoimento
	database.DB.Delete(&depoimento, ID)
}
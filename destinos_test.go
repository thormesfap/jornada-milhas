package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thormesfap/jornada-milhas/database"
	"github.com/thormesfap/jornada-milhas/models"
	"github.com/thormesfap/jornada-milhas/controllers"
)


func TestListandoTodosDestinos(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaDestinoMock()
	defer DeletaDestinoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/destinos", controllers.TodosDestinos)
	req, _ := http.NewRequest("GET", "/destinos", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t,http.StatusOK, resposta.Code)
}

func TestBuscaDestinoPorIDHandler(t *testing.T){
	database.ConectaComBancoDeDados()
	CriaDestinoMock()
	defer DeletaDestinoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/destinos/:id", controllers.RetornaDestino)
	req, _ := http.NewRequest("GET", "/destinos/" + strconv.Itoa(ID), nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t,http.StatusOK, resposta.Code)
}

func TestBuscaDestinoPorNome(t *testing.T){
	database.ConectaComBancoDeDados()
	CriaDestinoMock()
	defer DeletaDestinoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/destinos", controllers.TodosDestinos)
	req, _ := http.NewRequest("GET", "/destinos?nome=Destino Mockado", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t,http.StatusOK, resposta.Code)
	req, _ = http.NewRequest("GET", "/destinos?nome=Nome de Destino Inexistente 4as87d9as", nil)
	resposta = httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusNotFound, resposta.Code)
}

func TestAtualizaDestinoPorIDHandler(t *testing.T){
	database.ConectaComBancoDeDados()
	CriaDestinoMock()
	defer DeletaDestinoMock()
	r := SetupDasRotasDeTeste()
	r.PATCH("/destinos/:id", controllers.EditaDestino)
	Destino := models.Destino{Nome:"Destino de Teste", Foto:"imagem.jpg", Preco:499}
	body, _ := json.Marshal(Destino)
	path := "/destinos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("PATCH", path, bytes.NewBuffer(body))
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t,http.StatusOK, resposta.Code)
	var DestinoMock models.Destino
	json.Unmarshal(resposta.Body.Bytes(), &DestinoMock)
	assert.Equal(t, "Destino de Teste", DestinoMock.Nome)
	assert.Equal(t, "imagem.jpg", DestinoMock.Foto)
	assert.Equal(t, 499., DestinoMock.Preco)
}

func TestCriaDestinoHandler(t *testing.T){
	database.ConectaComBancoDeDados()
	r := SetupDasRotasDeTeste()
	r.POST("/destinos", controllers.CriaDestino)
	Destino := models.Destino{Nome:"Destino de Teste Criação", Foto:"imagem.jpg", Preco:899}
	body, _ := json.Marshal(Destino)
	path := "/destinos"
	req, _ := http.NewRequest("POST", path, bytes.NewBuffer(body))
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t,http.StatusOK, resposta.Code)
	var DestinoMock models.Destino
	json.Unmarshal(resposta.Body.Bytes(), &DestinoMock)
	assert.NotEqual(t, 0, DestinoMock.ID)
	assert.Equal(t, "Destino de Teste Criação", DestinoMock.Nome)
	assert.Equal(t, "imagem.jpg", DestinoMock.Foto)
	assert.Equal(t, 899., DestinoMock.Preco)
	if DestinoMock.ID != 0{
		ID = int(DestinoMock.ID)
		defer DeletaDestinoMock()
	}
}

func TestDeletaDestinoPorIDHandler(t *testing.T){
	database.ConectaComBancoDeDados()
	CriaDestinoMock()
	r := SetupDasRotasDeTeste()
	r.DELETE("/destinos/:id", controllers.DeletaDestino)
	req, _ := http.NewRequest("DELETE", "/destinos/" + strconv.Itoa(ID), nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t,http.StatusOK, resposta.Code)
}


func CriaDestinoMock(){
	Destino := models.Destino{Nome:"Destino Mockado", Preco:499, Foto:"imagem.jpg"}
	database.DB.Create(&Destino)
	ID = int(Destino.ID)
}

func DeletaDestinoMock(){
	var Destino models.Destino
	database.DB.Delete(&Destino, ID)
}
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/textproto"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thormesfap/jornada-milhas/controllers"
	"github.com/thormesfap/jornada-milhas/database"
	"github.com/thormesfap/jornada-milhas/models"
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
	Destino := models.Destino{Nome:"Destino de Teste", Preco:499}
	body, _ := json.Marshal(Destino)
	path := "/destinos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("PATCH", path, bytes.NewBuffer(body))
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t,http.StatusOK, resposta.Code)
	var DestinoMock models.Destino
	json.Unmarshal(resposta.Body.Bytes(), &DestinoMock)
	assert.Equal(t, "Destino de Teste", DestinoMock.Nome)
	assert.Equal(t, 499., DestinoMock.Preco)
}

func TestCriaDestinoHandler(t *testing.T){
	database.ConectaComBancoDeDados()
	r := SetupDasRotasDeTeste()
	r.POST("/destinos", controllers.CriaDestino)
	Destino := models.Destino{Nome:"Destino de Teste Criação", Preco:899}
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
	assert.Equal(t, 899., DestinoMock.Preco)
	if DestinoMock.ID != 0{
		ID = int(DestinoMock.ID)
		defer DeletaDestinoMock()
	}
}

func TestAdicionaFotoAoDestino(t *testing.T){
	database.ConectaComBancoDeDados()
	CriaDestinoMock()
	defer DeletaDestinoMock()
	r := SetupDasRotasDeTeste()
	r.POST("/destinos/:id", controllers.AdicionaFotoAoDestino)
	body := new(bytes.Buffer)
	multipartWriter := multipart.NewWriter(body)
	//Create multipart header
    fileHeader := make(textproto.MIMEHeader)
    fileHeader.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, "imagem", "avatar.jpg"))
    fileHeader.Set("Content-Type", "text/plain")
    writer, err := multipartWriter.CreatePart(fileHeader)
    assert.Nil(t, err)
    //Copy file to file multipart writer
    file, err := os.Open("avatar.jpg")
    assert.Nil(t, err)
    io.Copy(writer, file)
    // close the writer before making the request
    multipartWriter.Close()
    req, _ := http.NewRequest(http.MethodPost, "/destinos/" + strconv.Itoa(ID), body)
    resp := httptest.NewRecorder()
    req.Header.Add("Content-Type", multipartWriter.FormDataContentType())
    r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
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
	Destino := models.Destino{Nome:"Destino Mockado", Preco:499}
	database.DB.Create(&Destino)
	ID = int(Destino.ID)
}

func DeletaDestinoMock(){
	var Destino models.Destino
	database.DB.Delete(&Destino, ID)
}
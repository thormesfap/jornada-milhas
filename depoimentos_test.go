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
	depoimento := models.Depoimento{Autor:"Autor de Teste para Edição", Depoimento:"Depoimento alterado com sucesso através do teste"}
	body, _ := json.Marshal(depoimento)
	path := "/depoimentos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("PATCH", path, bytes.NewBuffer(body))
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t,http.StatusOK, resposta.Code)
	var DepoimentoMock models.Depoimento
	json.Unmarshal(resposta.Body.Bytes(), &DepoimentoMock)
	assert.Equal(t, "Autor de Teste para Edição", DepoimentoMock.Autor)
	assert.Equal(t, "Depoimento alterado com sucesso através do teste", DepoimentoMock.Depoimento)
}

func TestCriaDepoimentoHandler(t *testing.T){
	database.ConectaComBancoDeDados()
	r := SetupDasRotasDeTeste()
	r.POST("/depoimentos", controllers.CriaDepoimento)
	depoimento := models.Depoimento{Autor:"Autor de Teste para Criação", Depoimento:"Depoimento criado com sucesso através do teste"}
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

func TestAdicionaFotoAoDepoimento(t *testing.T){
	database.ConectaComBancoDeDados()
	CriaDepoimentoMock()
	defer DeletaDepoimentoMock()
	r := SetupDasRotasDeTeste()
	r.POST("/depoimentos/:id", controllers.AdicionaFotoAoDepoimento)
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
    req, _ := http.NewRequest(http.MethodPost, "/depoimentos/" + strconv.Itoa(ID), body)
    resp := httptest.NewRecorder()
    req.Header.Add("Content-Type", multipartWriter.FormDataContentType())
    r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}


func CriaDepoimentoMock(){
	depoimento := models.Depoimento{Autor:"Autor de Teste", Depoimento:"Depoimento de Teste para criação"}
	database.DB.Create(&depoimento)
	ID = int(depoimento.ID)
}

func DeletaDepoimentoMock(){
	var depoimento models.Depoimento
	database.DB.Delete(&depoimento, ID)
}
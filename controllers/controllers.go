package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thormesfap/jornada-milhas/database"
	"github.com/thormesfap/jornada-milhas/models"
)

func TodosDepoimentos(w http.ResponseWriter, r *http.Request) {

	var d []models.Depoimento
	database.DB.Find(&d)
	json.NewEncoder(w).Encode(d)
}

func DepoimentosHome(w http.ResponseWriter, r *http.Request) {
	var d []models.Depoimento
	database.DB.Find(&d)
	json.NewEncoder(w).Encode(d)
}
func RetornaDepoimento(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var d models.Depoimento
	database.DB.First(&d, id)
	json.NewEncoder(w).Encode(d)
}
func CriaDepoimento(w http.ResponseWriter, r *http.Request) {
	var depoimento models.Depoimento
	json.NewDecoder(r.Body).Decode(&depoimento)
	database.DB.Create(&depoimento)
	json.NewEncoder(w).Encode(depoimento)
}
func DeletaDepoimento(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var p models.Depoimento
	database.DB.Delete(&p, id)
	json.NewEncoder(w).Encode(p)
}
func EditaDepoimento(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var d models.Depoimento
	database.DB.First(&d, id)
	json.NewDecoder(r.Body).Decode(&d)
	database.DB.Save(&d)
	json.NewEncoder(w).Encode(d)
}

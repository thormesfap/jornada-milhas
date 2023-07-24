package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"github.com/thormesfap/jornada-milhas/controllers"
	"github.com/thormesfap/jornada-milhas/middleware"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/api/depoimentos", controllers.TodosDepoimentos).Methods("Get")
	r.HandleFunc("/api/depoimentos-home", controllers.DepoimentosHome).Methods("Get")
	r.HandleFunc("/api/depoimentos", controllers.CriaDepoimento).Methods("Post")
	r.HandleFunc("/api/depoimentos/{id}", controllers.RetornaDepoimento).Methods("Get")
	r.HandleFunc("/api/depoimentos/{id}", controllers.DeletaDepoimento).Methods("Delete")
	r.HandleFunc("/api/depoimentos/{id}", controllers.EditaDepoimento).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}

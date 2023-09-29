package routes

import (
	"github.com/gorilla/mux"
	"github.com/josepnapitupulu/API_Kegiatan/controllers"
)

var RegisterKegiatansRoutes = func(router *mux.Router) {
	// router.HandleFunc("/", controllers.Index).Methods("GET")
	// router.HandleFunc("/jemaatBaru", controllers.Create).Methods("POST")
	router.HandleFunc("/kegiatan/", controllers.CreateKegiatan).Methods("POST")
	router.HandleFunc("/kegiatan/", controllers.GetKegiatan).Methods("GET")
	router.HandleFunc("/kegiatan/{kegiatanId}", controllers.GetKegiatanById).Methods("GET")
	router.HandleFunc("/kegiatan/{kegiatanId}", controllers.UpdateKegiatan).Methods("PUT")
	router.HandleFunc("/kegiatan/{kegiatanId}", controllers.DeleteKegiatan).Methods("DELETE")
}

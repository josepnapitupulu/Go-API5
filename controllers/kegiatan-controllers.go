package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	
	"github.com/gorilla/mux"
	"github.com/josepnapitupulu/API_Kegiatan/models"
	"github.com/josepnapitupulu/API_Kegiatan/utils"
)

var NewKegiatan models.Kegiatan
// var tmpl * template.Template
// func init(){
// 	tmpl = template.Must(template.ParseFiles("jemaat.html"))
// }

// func Index(w http.ResponseWriter, r *http.Request){
// 	temp, _ := template.ParseFiles("views/jemaat.html")
// 	temp.Execute(w, nil)
// }

// func Create(w http.ResponseWriter, r *http.Request){
// 	temp, _ := template.ParseFiles("views/createJemaat.html")
// 	temp.Execute(w, nil)
// }

func GetKegiatan(w http.ResponseWriter, r *http.Request) {
	newKegiatans := models.GetAllKegiatans()
	res, _ := json.Marshal(newKegiatans)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetKegiatanById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	kegiatanId := vars["kegiatanId"]
	IdKegiatan, err := strconv.ParseInt(kegiatanId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	kegiatanDetails, _ := models.GetKegiatanbyId(IdKegiatan)
	res, _ := json.Marshal(kegiatanDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateKegiatan(w http.ResponseWriter, r *http.Request) {
	CreateKegiatan := &models.Kegiatan{}
	utils.ParseBody(r, CreateKegiatan)
	b := CreateKegiatan.CreateKegiatan()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteKegiatan(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	kegiatanId := vars["kegiatanId"]
	ID, err := strconv.ParseInt(kegiatanId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	kegiatan := models.DeleteKegiatan(ID)
	res, _ := json.Marshal(kegiatan)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateKegiatan(w http.ResponseWriter, r *http.Request) {
    var updateKegiatan = &models.Kegiatan{}
    utils.ParseBody(r, updateKegiatan)
    vars := mux.Vars(r)
    kegiatanId := vars["kegiatanId"]
    ID, err := strconv.ParseInt(kegiatanId, 0, 0)
    if err != nil {
        fmt.Println("error while parsing")
    }
    kegiatanDetails, db := models.GetKegiatanbyId(ID)
    if updateKegiatan.NamaKegiatan != "" {
        kegiatanDetails.NamaKegiatan = updateKegiatan.NamaKegiatan
    }
    if updateKegiatan.JenisKegiatan != "" {
        kegiatanDetails.JenisKegiatan = updateKegiatan.JenisKegiatan
    }
    if updateKegiatan.JadwalKegiatan != "" {
        kegiatanDetails.JadwalKegiatan = updateKegiatan.JadwalKegiatan
    }
    if updateKegiatan.LokasiKegiatan != "" {
        kegiatanDetails.LokasiKegiatan = updateKegiatan.LokasiKegiatan
    }
    db.Save(&kegiatanDetails)
    res, _ := json.Marshal(kegiatanDetails)
    w.Header().Set("Content-Type", "pkglication/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}
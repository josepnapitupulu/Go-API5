package models

import (
	"github.com/jinzhu/gorm"
	"github.com/josepnapitupulu/API_Kegiatan/config"
	)

var db *gorm.DB

type Kegiatan struct {
	gorm.Model
	NamaKegiatan string `json:"nama_kegiatan"`
	JenisKegiatan string `json:"jenis_kegiatan"`
	JadwalKegiatan string `json:"jadwal_kegiatan"`
	LokasiKegiatan string `json:"lokasi_kegiatan"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Kegiatan{})
}

func (b *Kegiatan) CreateKegiatan() *Kegiatan {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllKegiatans() []Kegiatan {
	var Kegiatans []Kegiatan
	db.Find(&Kegiatans)
	return Kegiatans
}

func GetKegiatanbyId(id_kegiatan int64) (*Kegiatan, *gorm.DB) {
	var getKegiatan Kegiatan
	db := db.Where("ID=?", id_kegiatan).Find(&getKegiatan)
	return &getKegiatan, db
}

func DeleteKegiatan(id_kegiatan int64) Kegiatan {
	var kegiatan Kegiatan
	db.Where("ID=?", id_kegiatan).Delete(kegiatan)
	return kegiatan 
}
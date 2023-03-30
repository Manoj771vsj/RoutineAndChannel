package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/manoj771vsj/Routines/entity"
	"gorm.io/gorm"
)

type Ownerinterface interface {
	GetDetails(w http.ResponseWriter, r *http.Request)
	AddDetails(w http.ResponseWriter, r *http.Request)
}

type OwnerDatabase struct {
	DB *gorm.DB
}

func (d *OwnerDatabase) GetDetails(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Message{Msg: "Id is null"})
		return
	}

	var tag entity.Owner

	if result := d.DB.First(&tag, id); result.Error != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Message{Msg: "Entry with given id doesnot exist"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tag.OwnerData)
}

func (d *OwnerDatabase) AddDetails(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var tag1 entity.Owner
	var tag entity.OwnerData
	json.Unmarshal(reqBody, &tag)
	tag1.OwnerData = tag
	if result := d.DB.Create(&tag1); result.Error != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Message{Msg: "Id Already exists"})
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tag1)
}

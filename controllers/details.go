package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/manoj771vsj/Routines/service"

	"gorm.io/gorm"
)

type Detailinterface interface {
	ShowCarDetails(w http.ResponseWriter, r *http.Request)
	//GetCardetails(w http.ResponseWriter, r *http.Request)
}

type CarDetailsDatabase struct {
	DB *gorm.DB
}

func (d CarDetailsDatabase) ShowCarDetails(w http.ResponseWriter, r *http.Request) {
	var merge service.CarDetailsService = service.NewDetailsService()
	id := r.URL.Query().Get("id")
	res := merge.GetDetails(id)
	w.Header().Set("Content-Type", "application/json")
	if result := d.DB.Create(&res); result.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Message{Msg: "Id Already exists"})
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

//func (d CarDetailsDatabase) GetCardetails(w http.ResponseWriter, r *http.Request) {
// 	var merge service.CarDetailsService = service.NewDetailsService()
// 	id := r.URL.Query().Get("id")
// 	var post *entity.CarDetails = service.CarDetailsCache.Get(id)
// 	if post == nil {
// 		post = merge.FetchCarDetails(id)
// 		service.CarDetailsCache.Set(id, post)
// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(http.StatusOK)
// 		json.NewEncoder(w).Encode(post)
// 	} else {
// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(http.StatusOK)
// 		json.NewEncoder(w).Encode(post)
// 	}

// }

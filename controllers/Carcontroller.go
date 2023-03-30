package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/manoj771vsj/Routines/entity"
	"gorm.io/gorm"
)

type Carinterface interface {
	GetDetails(w http.ResponseWriter, r *http.Request)
	AddDetails(w http.ResponseWriter, r *http.Request)
}

type CarDatabase struct {
	DB *gorm.DB
}

type Message struct {
	Msg string
}

func (d *CarDatabase) GetDetails(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Message{Msg: "Id is null"})
		return
	}

	var tag entity.Car

	if result := d.DB.First(&tag, id); result.Error != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(result.Error)
		json.NewEncoder(w).Encode(Message{Msg: "Entry with given id doesnot exist"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tag.CarData)
}

func (d *CarDatabase) AddDetails(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var tag1 entity.Car
	var tag entity.CarData
	json.Unmarshal(reqBody, &tag)
	tag1.CarData = tag
	w.Header().Set("Content-Type", "application/json")
	if result := d.DB.Create(&tag1); result.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Message{Msg: "Id Already exists"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tag1)
}

// import (
// 	"encoding/json"
// 	"log"
// 	"net/http"

// 	"github.com/manoj771vsj/Routines/entity"
// 	"github.com/manoj771vsj/Routines/service"
// 	"gorm.io/gorm"
// )

// var (
// 	carDataChannel   = make(chan entity.CarData)
// 	OwnerDataChannel = make(chan entity.OwnerData)
// )

// type Info interface {
// 	GetCar(w http.ResponseWriter, r *http.Request)
// 	//CreateCarEntry(w http.ResponseWriter, r *http.Request)
// 	GetOwnerData(w http.ResponseWriter, r *http.Request)
// 	//CreateOwnerEntry(w http.ResponseWriter, r *http.Request)
// }

// type Database struct {
// 	DB *gorm.DB
// }

// type Message struct {
// 	Msg string
// }

// func (d Database) GetCar(w http.ResponseWriter, r *http.Request) {

// 	var tag entity.CarData
// 	result := d.DB.Find(&tag /*[]int{1,2}*/) //Return all entries since no condition was given
// 	if result.Error != nil {
// 		log.Fatal(result.Error)
// 	}
// 	w.Header().Add("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(tag)

// }

// func (d Database) GetOwnerData(w http.ResponseWriter, r *http.Request) {

// 	var tag entity.OwnerData
// 	result := d.DB.Find(&tag /*[]int{1,2}*/) //Return all entries since no condition was given
// 	if result.Error != nil {
// 		log.Fatal(result.Error)
// 	}
// 	w.Header().Add("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(tag)
// }

// func (d Database) GetCarDetailData(w http.ResponseWriter, r *http.Request) {
// 	poi := service.NewDetailsService()
// 	res := poi.GetDetails()
// 	// go d.GetCar(w, r)
// 	// go d.GetOwnerData(w, r)
// 	// car, _ := getCarData()
// 	// owner, _ := getOwnerData()
// 	// res := entity.CarDetails{
// 	// 	ID:        car.ID,
// 	// 	Brand:     car.Brand,
// 	// 	Model:     car.Model,
// 	// 	Year:      car.Year,
// 	// 	FirstName: owner.FirstName,
// 	// 	LastName:  owner.LastName,
// 	// 	Email:     owner.Email,
// 	// }
// 	if result := d.DB.Create(&res); result.Error != nil {
// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(http.StatusBadRequest)
// 		json.NewEncoder(w).Encode(Message{Msg: "Id Already exists"})
// 		return
// 	}
// 	w.Header().Add("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(res)

// }

// // func getCarData() (entity.CarData, error) {
// // 	var car entity.CarData
// // 	car = <-carDataChannel
// // 	return car, nil
// // }

// // func getOwnerData() (entity.OwnerData, error) {
// // 	var owner entity.OwnerData
// // 	owner = <-OwnerDataChannel
// // 	return owner, nil
// // }

package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/manoj771vsj/Routines/entity"
)

const (
	carServiceUrl = "http://localhost:3000/car"
)

type CarService interface {
	FetchData(string)
}

type fetchCarDataService struct{}

func NewCarService() CarService {
	return &fetchCarDataService{}
}

func (*fetchCarDataService) FetchData(id string) {
	client := http.Client{}
	fmt.Printf("Fetching the url %s", carServiceUrl)

	resp, _ := client.Get(carServiceUrl + fmt.Sprintf("?id=%s", id))
	fmt.Println(fmt.Sprintf("?id=%s", id))
	var car entity.CarData
	json.NewDecoder(resp.Body).Decode(&car)
	fmt.Println(car)
	carDataChannel <- car
	//fmt

}

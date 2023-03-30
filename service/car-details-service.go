package service

import (
	"fmt"

	"github.com/manoj771vsj/Routines/entity"
)

// var (
// 	CarDetailsCache cache.CarDetailsCache
// )

const (
	cardetailsServiceUrl = "http://localhost:3000/car"
)

var (
	carService       CarService   = NewCarService()
	ownerService     OwnerService = NewOwnerService()
	carDataChannel                = make(chan entity.CarData)
	OwnerDataChannel              = make(chan entity.OwnerData)
)

type CarDetailsService interface {
	GetDetails(string) entity.CarDetails
	//FetchCarDetails(id string) *entity.CarDetails
}

type service struct{}

func NewDetailsService( /*cache cache.CarDetailsCache*/ ) CarDetailsService {
	//CarDetailsCache = cache
	return &service{}
}

func (*service) GetDetails(id string) entity.CarDetails {
	go carService.FetchData(id)

	go ownerService.FetchData(id)

	car, _ := getCarData()
	owner, _ := getOwnerData()
	fmt.Print(car)
	return entity.CarDetails{
		ID:        car.ID,
		Brand:     car.Brand,
		Model:     car.Model,
		Year:      car.Year,
		FirstName: owner.FirstName,
		LastName:  owner.LastName,
		Email:     owner.Email,
	}

}

// func (*service) FetchCarDetails(id string) *entity.CarDetails {
// 	client := http.Client{}
// 	resp, _ := client.Get(cardetailsServiceUrl + fmt.Sprintf("?id=%s", id))
// 	fmt.Println(fmt.Sprintf("?id=%s", id))
// 	var cardetails entity.CarDetails
// 	json.NewDecoder(resp.Body).Decode(&cardetails)
// 	return &cardetails
// }

func getCarData() (entity.CarData, error) {
	r1 := <-carDataChannel
	return r1, nil
}

func getOwnerData() (entity.OwnerData, error) {
	r1 := <-OwnerDataChannel
	return r1, nil
}

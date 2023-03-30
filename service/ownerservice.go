package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/manoj771vsj/Routines/entity"
)

const (
	ownerServiceUrl = "http://localhost:3000/owner"
)

type OwnerService interface {
	FetchData(string)
}

type fetchOwnerDataService struct{}

func NewOwnerService() CarService {
	return &fetchOwnerDataService{}
}

func (*fetchOwnerDataService) FetchData(id string) {
	client := http.Client{}
	fmt.Printf("Fetching the url %s", ownerServiceUrl)

	resp, _ := client.Get(ownerServiceUrl + fmt.Sprintf("?id=%s", id))
	var owner entity.OwnerData
	json.NewDecoder(resp.Body).Decode(&owner)
	OwnerDataChannel <- owner
}

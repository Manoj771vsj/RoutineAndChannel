package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/manoj771vsj/Routines/DataBase"
	"github.com/manoj771vsj/Routines/controllers"
)

func Routes() {
	var Carrout controllers.Carinterface
	Carrout = &controllers.CarDatabase{
		DB: DataBase.ConnectToCarDB(),
	}

	var Ownerrout controllers.Ownerinterface
	Ownerrout = &controllers.OwnerDatabase{
		DB: DataBase.ConnectOwnerDB(),
	}

	var Mergerout controllers.Detailinterface
	Mergerout = &controllers.CarDetailsDatabase{
		DB: DataBase.ConnectCarDetailsDB(),
	}

	myrouter := mux.NewRouter()

	myrouter.HandleFunc("/car", Carrout.GetDetails).Methods("GET")
	myrouter.HandleFunc("/owner", Ownerrout.GetDetails).Methods("GET")
	myrouter.HandleFunc("/cardetails", Mergerout.ShowCarDetails).Methods("GET")
	//myrouter.HandleFunc("/cardetailsforcache", Mergerout.GetCardetails).Methods("GET")
	myrouter.HandleFunc("/addowner", Ownerrout.AddDetails).Methods("POST")
	myrouter.HandleFunc("/addcar", Carrout.AddDetails).Methods("POST")

	fmt.Println("api running")

	http.ListenAndServe(":3000", myrouter)

}

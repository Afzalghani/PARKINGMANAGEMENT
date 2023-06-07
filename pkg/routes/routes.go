package routes

import (
	"net/http"
	controller "parking/pkg/Controller"
)

func Routes() http.Handler {
	mux := http.NewServeMux();
	
	mux.HandleFunc("/",controller.Hello)
	mux.HandleFunc("/car",controller.CreateCarHandler)
	mux.HandleFunc("/cars",controller.GetCarsHandler)
	mux.HandleFunc("/payment",controller.CreatePayemnet)
	mux.HandleFunc("/payments",controller.GetPaymentDetails)
	mux.HandleFunc("/parker",controller.Createparker)
	mux.HandleFunc("/parkers",controller.GetParkerDetails)
    mux.HandleFunc("/slot", controller.CreateParkingSlot)
	mux.HandleFunc("/slots",controller.GetParkingSlotDetails);
	mux.HandleFunc("/parkedCar",controller.CreateParkedSpace)
	mux.HandleFunc("/parkedCars",controller.GetParkedSpaceDetails);
   
	return mux;
}
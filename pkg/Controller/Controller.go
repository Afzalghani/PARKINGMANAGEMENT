package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"parking/pkg/model"
)

func Hello(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("this is working condition"))
}

//Create the Parked Space
func CreateParkedSpace(w http.ResponseWriter, r *http.Request){
	var parkedSpace model.ParkedCar;

	//Decode the request that is comming 

	err := json.NewDecoder(r.Body).Decode(&parkedSpace);

	if err!= nil{
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError);
		return;
	}

	_, err = db.Exec("INSERT INTO parkedcar(parkedcarId, carId,slotId,parkerId,status,paymentId) VALUES(?,?,?,?,?,?)", &parkedSpace.ParkedCarID,&parkedSpace.CarID,&parkedSpace.SlotID, &parkedSpace.ParkerID,&parkedSpace.Status,&parkedSpace.PaymentID);


	if err!= nil{
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	   return;
	}

	
}


//Get the ParkedCareDetails with all the details form payment to space etc.
func GetParkedSpaceDetails(w http.ResponseWriter, r *http.Request){

	var parkedCars []model.ParkedCar;

	rows, err := db.Query("SELECT * FROM parkedcar");
	if err!= nil{
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return;
	}

	for rows.Next(){
		var  parkedSpace model.ParkedCar

		err = rows.Scan(&parkedSpace.ParkedCarID,&parkedSpace.CarID,&parkedSpace.SlotID, &parkedSpace.ParkerID,&parkedSpace.Status,&parkedSpace.PaymentID)
        
		if err!= nil{
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		parkedCars= append(parkedCars, parkedSpace)
	}

	_ = json.NewEncoder(w).Encode(parkedCars);



}

//Get the parkingSlot details
func GetParkingSlotDetails(w http.ResponseWriter, r *http.Request){
	var slots []model.ParkingSlot;

	rows, err := db.Query("SELECT * FROM parkingslot");

	if err!= nil{
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	for rows.Next(){
		var slot model.ParkingSlot;

		err := rows.Scan(&slot.SlotID, &slot.Size)

		if err!= nil{
		    log.Println(err)
		    w.WriteHeader(http.StatusInternalServerError)
		     return
		}

		slots = append(slots, slot)
	   
	}

	json.NewEncoder(w).Encode(slots);
	
     

}

//Create the ParkingSlot 
func CreateParkingSlot(w http.ResponseWriter, r *http.Request){

	var slot model.ParkingSlot;

	//Just decode that data which is comming in the form of json
	//And then marshal that in slot structure

	err := json.NewDecoder(r.Body).Decode(&slot)

	if err!=nil{
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Insert into database

	_, err = db.Exec("INSERT INTO parkingslot(slotId,size) VALUES(?,?)",&slot.SlotID, &slot.Size);

	if err!= nil{
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)



}



// Get Parker details
func GetParkerDetails(w http.ResponseWriter, r *http.Request){
	var parkers []model.Parker;

	rows, err :=  db.Query("SELECT * FROM parker");
  
	if err!=nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	for rows.Next(){
		var parker model.Parker;

		err := rows.Scan(&parker.ParkerID,&parker.Name)

		if err!=nil{
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return;
		}

		parkers = append(parkers, parker)

	}


	jsonParker, err := json.Marshal(parkers)

	if err!=nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonParker)


}


//Create the parker 
func Createparker(w http.ResponseWriter, r *http.Request) {
	var parker model.Parker
	err := json.NewDecoder(r.Body).Decode(&parker)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = db.Exec("INSERT INTO parker(parkerId,name) VALUES(?,?)",parker.ParkerID,parker.Name)

	if err!=nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return;
	}
	w.WriteHeader(http.StatusCreated)
}

func CreatePayemnet(w http.ResponseWriter, r *http.Request) {
	var payment model.Payment

	err := json.NewDecoder(r.Body).Decode(&payment)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Insert into database

	_, err = db.Exec("INSERT INTO  PAYMENT(paymentId, PaymentAmount, paymentDate,paymentLocation) VALUES (?, ?, ?,?)", payment.PaymentID, payment.PaymentAmount, payment.PaymentDate, payment.PaymentLocation)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

}

func GetPaymentDetails(w http.ResponseWriter, r *http.Request) {

	// Fetch all the payment details from the database

	rows, err := db.Query("SELECT * FROM payment")

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	//create slice of payment which will used to store the value

	var payments []model.Payment

	for rows.Next() {
		var payment model.Payment

		err := rows.Scan(&payment.PaymentID, &payment.PaymentAmount, &payment.PaymentDate, &payment.PaymentLocation)

		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		payments = append(payments, payment)
	}

	//Convert the payments sllice to json

	jsonBytes, err := json.Marshal(payments)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)

}
func GetCarsHandler(w http.ResponseWriter, r *http.Request) {
	// Fetch all cars from the database
	rows, err := db.Query("SELECT * FROM Car")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Create a slice to store the retrieved cars
	var cars []model.Car

	// Iterate over the result rows and populate the cars slice
	for rows.Next() {
		var car model.Car

		err := rows.Scan(&car.CarID, &car.Name, &car.Model, &car.Owner)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		cars = append(cars, car)
	}

	// Convert the cars slice to JSON and send it as the response
	jsonBytes, err := json.Marshal(cars)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}

func CreateCarHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request body into a Car struct
	var car model.Car
	err := json.NewDecoder(r.Body).Decode(&car)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Insert the car into the database
	_, err = db.Exec("INSERT INTO Car (carId, name, model,owner) VALUES (?, ?,?, ?)", car.CarID, car.Name, car.Model, car.Owner)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

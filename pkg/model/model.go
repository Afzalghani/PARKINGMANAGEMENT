package model

type Car struct {
	CarID int    `json:"carId"`
	Name  string `json:"name"`
	Model string `json:"model"`
	Owner string `json:"owner"`
}

type ParkingSlot struct {
	SlotID int    `json:"slotId"`
	Size   string `json:"size"`
}

type Parker struct {
	ParkerID int    `json:"parkerId"`
	Name     string `json:"name"`
}

type Payment struct {
	PaymentID       int     `json:"paymentId"`
	PaymentAmount   float64 `json:"paymentAmount"`
	PaymentDate     string  `json:"paymentDate"`
	PaymentLocation string  `json:"paymentLocation"`
}

type ParkedCar struct {
	ParkedCarID int    `json:"parkedcarId"`
	CarID       int    `json:"carId"`
	SlotID      int    `json:"slotId"`
	ParkerID    int    `json:"parkerId"`
	Status      string `json:"status"`
	PaymentID   int    `json:"paymentId"`
}

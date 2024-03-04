package controller

import "homework/2-oy/10-dars/rent_car/storage"

type Controller struct {
	Store storage.Store
}

func NewController(store storage.Store) Controller {
	return Controller{
		Store: store,
	}
}

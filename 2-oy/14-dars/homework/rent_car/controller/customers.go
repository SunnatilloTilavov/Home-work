package controller

import (
	"encoding/json"
	// "errors"
	"fmt"
	"homework/2-oy/13-dars/rent_car/models"
	"homework/2-oy/13-dars/rent_car/pkg/check"
	"net/http"

	"github.com/google/uuid"
)

func (c Controller) Customer(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		c.CreateCustomer(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		_, ok := values["id"]
		if !ok {
			c.GetAllCustomers(w, r)
		}
	case http.MethodPut:
		values := r.URL.Query()
		_, ok := values["id"]
		if ok {
			c.UpdateCustomer(w, r)
		}

	case http.MethodDelete:
		values := r.URL.Query()
		_, ok := values["id"]
		if ok {
			c.DeleteCustomer(w, r)
		}

	default:
		handleResponse(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

func (c Controller) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	customer := models.Customer{}

	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		errStr := fmt.Sprintf("error while decoding request body, err: %v\n", err)
		fmt.Println(errStr)
		handleResponse(w, http.StatusBadRequest, errStr)
		return
	}

	if  !check.ValidateEmail(customer.Gmail) {
		fmt.Println("error while validating gamil: ", customer.Gmail)
		handleResponse(w, http.StatusBadRequest,"error gmail address")
		return
	}

	if  !check.ValidatePhone(customer.Phone) {
		fmt.Println("error while validating phone: ", customer.Phone)
		handleResponse(w, http.StatusBadRequest,"error phone add")
		return
	}
	

	id, err := c.Store.Customer.Create(customer)
	if err != nil {
		fmt.Println("error while creating customer, err: ", err)
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, id)
}

func (c Controller) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	customer := models.Customer{}

	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		errStr := fmt.Sprintf("error while decoding request body, err: %v\n", err)
		fmt.Println(errStr)
		handleResponse(w, http.StatusBadRequest, errStr)
		return
	}

	if  !check.ValidateEmail(customer.Gmail) {
		fmt.Println("error while validating gmail: ", customer.Gmail)
		handleResponse(w, http.StatusBadRequest,"error gmail address")
		return
	}

	if  !check.ValidatePhone(customer.Phone) {
		fmt.Println("error while validating phone: ", customer.Phone)
		handleResponse(w, http.StatusBadRequest,"error phone address")
		return
	}

	customer.Id = r.URL.Query().Get("id")

	err := uuid.Validate(customer.Id)
	if err != nil {
		fmt.Println("error while validating, err: ", err)
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	id, err := c.Store.Customer.Update(customer)
	if err != nil {
		fmt.Println("error while creating customer, err: ", err)
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, id)
}

func (c Controller) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	var (
		values = r.URL.Query()
		search,last_name,phone string
	)
	if _, ok := values["search"]; ok {
		search = values["search"][0]
	}

	if _, ok := values["last_name"]; ok {
		last_name = values["last_name"][0]
	}

	if _, ok := values["phone"]; ok {
		phone = values["phone"][0]
	}


	customers, err := c.Store.Customer.GetAllCustomers(search,last_name,phone)
	if err != nil {
		fmt.Println("error while getting customers, err: ", err)
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(w, http.StatusOK, customers)
}

func (c Controller) GetById(w http.ResponseWriter, r *http.Request) {
	var (
		values = r.URL.Query()
		id string
	)
	if _, ok := values["id"]; ok {
		id = values["id"][0]
	}

	customers, err := c.Store.Customer.GetById(id)
	if err != nil {
		fmt.Println("error while getting customers, err: ", err)
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(w, http.StatusOK, customers)
}

func (c Controller) DeleteCustomer(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	fmt.Println("id: ", id)

	err := uuid.Validate(id)
	if err != nil {
		fmt.Println("error while validating id, err: ", err)
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err = c.Store.Customer.Delete(id)
	if err != nil {
		fmt.Println("error while deleting customer, err: ", err)
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, id)
}

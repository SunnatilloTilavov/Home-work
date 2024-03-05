package models

type Customer struct {
	Id          string  `json:"id"`
	First_name        string  `json:"first_name"`
	Last_name       string  `json:"last_name"`
	Gmail       string  `json:"gmail"`
	Phone     string  `json:"phone"`
	CreatedAt   string  `json:"createdAt"`
	UpdatedAt   string  `json:"updatedAt"`
}

type GetAllCustomersResponse struct {
	Customers  []Customer `json:"customer"`
	Count int64 `json:"count"`
}



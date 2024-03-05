package storage

import (
	"database/sql"
	"fmt"
	"homework/2-oy/13-dars/rent_car/models"
	"homework/2-oy/13-dars/rent_car/pkg"
	"github.com/google/uuid"
)

type customerRepo struct {
	db *sql.DB
}

func NewCustomer(db *sql.DB) customerRepo {
	return customerRepo{
		db: db,
	}
}

/*
create (body) id,err
update (body) id,err
delete (id) err
get (id) body,err
getAll (search) []body,count,err
*/

func (c *customerRepo) Create(customer models.Customer) (string, error) {

	id := uuid.New()

	query := ` INSERT INTO customers (
		id ,       
		first_name,
		last_name ,
		gmail ,    
		phone     )
		VALUES($1,$2,$3,$4,$5) `

	_, err := c.db.Exec(query,
		id.String(),
		customer.First_name, customer.Last_name,
		customer.Gmail, customer.Phone)
	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (c *customerRepo) Update(customer models.Customer) (string, error) {

	query := ` UPDATE customers set
	        first_name=$1,
	        last_name=$2,
	        gmail=$3,
	        phone=$4,
			updated_at=CURRENT_TIMESTAMP
		WHERE id = $5 AND deleted_at=0
	`

	_, err := c.db.Exec(query,
		customer.First_name, customer.Last_name,
		customer.Gmail, customer.Phone,customer.Id)

	if err != nil {
		return "", err
	}

	return customer.Id, nil
}

func (c customerRepo) GetAllCustomers(search string,last_name string,phone string) (models.GetAllCustomersResponse, error) {
	var (
		resp   = models.GetAllCustomersResponse{}
		filter = ""
	)

	if search != "" {
		filter += fmt.Sprintf(` and  first_name ILIKE  '%%%v%%' `, search)
	}
	if last_name !="" {
		filter += fmt.Sprintf(` and  last_name ILIKE  '%%%v%%' `, last_name)
	}
	 if phone!="" {
		filter += fmt.Sprintf(` and  phone ILIKE  '%%%v%%' `, phone)
	}



	// fmt.Println("filter: ", filter)

	rows, err := c.db.Query(`select 
				count(id) OVER(),
				id,        
				last_name ,
				first_name,
				gmail ,    
				phone, 
				created_at::date,
				updated_at
	  FROM customers WHERE deleted_at = 0 ` + filter + ``)

	if err != nil {
		return resp, err
	}
	for rows.Next() {
		var (
			customer      = models.Customer{}
			updateAt sql.NullString
		)

		if err := rows.Scan(
			&resp.Count,
			&customer.Id,
			&customer.Last_name,
			&customer.First_name,
			&customer.Gmail,
			&customer.Phone,
			&customer.CreatedAt,
			&updateAt); err != nil {
			return resp, err
		}

		customer.UpdatedAt = pkg.NullStringToString(updateAt)
		resp.Customers = append(resp.Customers, customer)
	}
	return resp, nil
}


func (c customerRepo) GetById(id string) (models.GetAllCustomersResponse, error) {
	var (
		resp   = models.GetAllCustomersResponse{}

	)

	rows, err := c.db.Query(`select 
				count(id) OVER(),
				id,        
				last_name ,
				first_name,
				gmail ,    
				phone, 
				created_at::date,
				updated_at
	  FROM customers WHERE deleted_at = 0 and id=$1 `,id)
	if err != nil {
		return resp, err
	}
	for rows.Next() {
		var (
			customer      = models.Customer{}
			updateAt sql.NullString
		)

		if err := rows.Scan(
			&resp.Count,
			&customer.Id,
			&customer.Last_name,
			&customer.First_name,
			&customer.Gmail,
			&customer.Phone,
			&customer.CreatedAt,
			&updateAt); err != nil {
			return resp, err
		}

		customer.UpdatedAt = pkg.NullStringToString(updateAt)
		resp.Customers = append(resp.Customers, customer)
	}
	return resp, nil
}


func (c *customerRepo) Delete(id string) error {

	query := ` UPDATE customers set
			deleted_at = date_part('epoch', CURRENT_TIMESTAMP)::int
		WHERE id = $1 AND deleted_at=0
	`

	_, err := c.db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}

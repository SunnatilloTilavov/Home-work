// func (i *Inventory) Update(c country.Country, name string, code int, id string) error {
// 	_, err := i.db.Exec(
// 		`UPDATE countries SET 
// 		    name=$2,
// 			code=$3,
// 			update_at=NOW()
// 			WHERE id = '$1'`, id, name, code)
// 	if err != nil {
// 		fmt.Println("error while updating country err: ", err)
// 		return err
// 	}

// 	return n
package main
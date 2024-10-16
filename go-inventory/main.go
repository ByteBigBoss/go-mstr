package main // PACKAGE DECLARATION [`main` is the entry point of program]
// WHEN `main` package is executed, the `main()` function is called

import (
	"fmt" // USE THIS PACKAGE TO FORMAT AND PRINT TEXT TO THE CONSOLE
)

// PRODUCT STRUCT REPRESENTS AN ITEM IN THE INVENTORY
//** `struct` IS A CUSTOM DATA TYPE THAT GROUPS TOGETHER VARIABLES OF DIFFERENT TYPES **//
type Product struct { 
	ID    int
	Name  string
	Price float64
	Stock int
}

// FUNCTION TO CREATE A NEW PRODUCT
func NewProduct(id int, name string, price float64, stock int) Product {
	return Product{
		ID:    id,
		Name:  name,
		Price: price,
		Stock: stock,
	}
}

// FUNCTION TO DISPLAY PRODUCT DETAILS
func (p Product) Display() { // THIS IS A `method bound` TO THE `struct`
	fmt.Printf(
		"ID: \033[32m%d\033[0m, Name: \033[33m%s\033[0m, Price: \033[35m%.2f\033[0m, Stock:\033[36m%d\033[0m\n",
		p.ID, p.Name, p.Price, p.Stock)
}

// FUNCTION TO LIST ALL PRODUCTS IN INVENTORY.
func ListProducts(products []Product){ 
	fmt.Println("\n--- Product List ---")
	for _, product := range products{ // ITERATING OVER A SLICE OF `Product` STRUCT
		product.Display()
	}
}

func main(){
	// CREATE A LIST TO STORE PRODUCTS.
	inventory := []Product{} // Slices = dunamic arrays | WE STORE OUR PRODUCTS IN A SLICE

	// ADD SOME PRODUCTS TO THE INVENTORY.
	inventory = append(inventory, NewProduct(1, "Laptop", 999.99, 10)) // APPENDING TO SLICE
	inventory = append(inventory, NewProduct(2, "Smartphone", 499.99, 20))
	inventory = append(inventory, NewProduct(3, "Tablet", 299.99, 15))

	// LIST ALL PRODUCTS IN THE INVENTORY.
	ListProducts(inventory)
}
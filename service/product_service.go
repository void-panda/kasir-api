package service

import "kasir-api/model"

var products = []model.Product{
	{ID: 1, Name: "Product 1", Price: 10000, Stock: 10},
	{ID: 2, Name: "Product 2", Price: 20000, Stock: 20},
	{ID: 3, Name: "Product 3", Price: 30000, Stock: 30},
}

func GetAllProducts() []model.Product {
	return products
}

func GetProductByID(id int) *model.Product {
	for i, product := range products {
		if product.ID == id {
			return &products[i]
		}
	}
	return nil
}

func CreateProduct(product model.Product) model.Product {
	product.ID = len(products) + 1
	products = append(products, product)
	return product
}

func UpdateProduct(id int, updatedProduct model.Product) *model.Product {
	for i := range products {
		if products[i].ID == id {
			if updatedProduct.Name != "" {
				products[i].Name = updatedProduct.Name
			}
			if updatedProduct.Price != 0 {
				products[i].Price = updatedProduct.Price
			}
			if updatedProduct.Stock != 0 {
				products[i].Stock = updatedProduct.Stock
			}
			return &products[i]
		}
	}
	return nil
}

func DeleteProduct(id int) bool {
	for i, p := range products {
		if p.ID == id {
			products = append(products[:i], products[i+1:]...)
			return true
		}
	}
	return false
}

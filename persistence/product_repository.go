package persistence

import (
	"Product-app/domain"
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

type IProductRepository interface {
	GetAllProducts() []domain.Product //* slice ile dönecek

}

type ProductRepository struct {
	dbPool *pgxpool.Pool
}


func (productRepository *ProductRepository) GetAllProducts() []domain.Product {
	productRows, err := productRepository.dbPool.Query(context.Background(), "SELECT * FROM products")
	if err != nil {
		log.Fatal(err)
	}

	products := []domain.Product{}
	var id int64
	var name string
	var price float32
	var discount float32
	var store string

	for productRows.Next(){
		productRows.Scan(&id, &name, &price, &discount, &store)
		products = append(products, domain.Product{
			Id: id, 
			Name: name, 
			Price: price, 
			Discount: discount, 
			Store: store,
		})
	}
	return products
}

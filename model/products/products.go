package model

import (
	"StoreManager/conn"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

// product structure -- how a product looks like
type Product struct {
	ID bson.ObjectId  `bson:"_id"`
	Name string `bson:"name" binding:"required"`
	Price int `bson:"price" binding:"required"`
	CreatedAt time.Time     `bson:"created_at"`
	UpdatedAt time.Time     `bson:"updated_at"`
}

type ProductId struct {
	ID bson.ObjectId `bson:"_id" json:"id" binding:"required"`
}

type ProductName struct {
	Name string `bson:"name" binding:"required" json:"name"`
}


// type for a list of products
type Products []Product

const ProductCollection = "product"

// utility function to connect to database and get connection
func ConnectDBHelper()(*mgo.Database, Product) {
	db := conn.GetMongoDB()
	product := Product{}
	return db, product
}

// fetches a product by id
func ProductById(id bson.ObjectId, productCollection string) (Product, error){
	db, product := ConnectDBHelper()
	err := db.C(productCollection).Find(bson.M{"_id": &id}).One(&product)
	return product, err
}

// fetches a product by name
func ProductByName(name string, productCollection string) ([]Product, error) {
	db := conn.GetMongoDB()
	product := make([]Product, 0)
	err := db.C(productCollection).Find(bson.M{"name": name}).All(&product)
	return product, err
}

// fetches a product by id
func (productId *ProductId) GetProductById(id bson.ObjectId)(Product, error)  {
	db, product := ConnectDBHelper()
	err := db.C(ProductCollection).Find(bson.M{"_id": &id}).One(&product)
	return product, err
}

// method on product name struct that fetches a product by name
func (productName *ProductName) GetProductByName(name string) (Product, error)  {
	db, product := ConnectDBHelper()
	err := db.C(ProductCollection).Find(bson.M{"name": &name}).One(&product)
	return product, err
}

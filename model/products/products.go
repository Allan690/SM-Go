package model

import (
	"StoreManager/conn"
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

// type for a list of products
type Products []Product


// fetches a product by id
func ProductById(id bson.ObjectId, productCollection string) (Product, error){
	db := conn.GetMongoDB()
	product := Product{}
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


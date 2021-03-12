package models

import (
	"StoreManager/conn"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Sale struct {
	ID bson.ObjectId `bson:"_id"`
	Product bson.ObjectId `bson:"productId" binding:"required"`
	User bson.ObjectId `bson:"userId" binding:"required" binding:"required"`
	CreatedAt time.Time  `bson:"created_at"`
	UpdatedAt time.Time  `bson:"updated_at"`
}

type Sales []Sale

func GetSaleInfo(salesId bson.ObjectId, salesCollection string)(Sale, error) {
	db := conn.GetMongoDB()
	sale := Sale{}
	err := db.C(salesCollection).FindId(&salesId).One(&sale)
	return sale, err
}

// fetches a sale by id
func GetOneSale(salesId bson.ObjectId, salesCollection string)([]bson.M, error){
	db := conn.GetMongoDB()
	var resp []bson.M
	pipe := db.C(salesCollection).Pipe([]bson.M{
		{"$match": bson.M{ "_id": salesId}},
		{ "$lookup": bson.M{
			"from": "product",
			"localField": "productId",
			"foreignField": "_id",
			"as": "product",
	}},
	{"$lookup": bson.M{
			"from": "user",
			"localField": "userId",
			"foreignField": "_id",
			"as": "user",
		}},
	})
	err := pipe.Iter().All(&resp)
	return resp, err
}

// fetches all sales currently
func GetAllSales(salesCollection string) ([]bson.M, error) {
	db := conn.GetMongoDB()
	var resp []bson.M
	pipe := db.C(salesCollection).Pipe([]bson.M{
		{ "$lookup": bson.M{
			"from": "product",
			"localField": "productId",
			"foreignField": "_id",
			"as": "product",
		}},
		{"$lookup": bson.M{
			"from": "user",
			"localField": "userId",
			"foreignField": "_id",
			"as": "user",
		}},
	})
	err := pipe.Iter().All(&resp)
	return resp, err
}

// get all sales belonging to a user
func GetSaleByUserId(userId bson.ObjectId, salesCollection string)([]bson.M, error) {
	db := conn.GetMongoDB()
	var resp []bson.M
	pipe := db.C(salesCollection).Pipe([]bson.M{
		{"$match": bson.M{ "userId": userId}},
		{ "$lookup": bson.M{
			"from": "product",
			"localField": "productId",
			"foreignField": "_id",
			"as": "product",
		}},
		{"$lookup": bson.M{
			"from": "user",
			"localField": "userId",
			"foreignField": "_id",
			"as": "user",
		}},
	})
	err := pipe.Iter().All(&resp)
	return resp, err
}

// get all sales by product id
func GetSalesByProductId(productId bson.ObjectId, salesCollection string)([]bson.M, error) {
	db := conn.GetMongoDB()
	var resp []bson.M
	pipe := db.C(salesCollection).Pipe([]bson.M{
		{"$match": bson.M{ "productId": productId}},
		{ "$lookup": bson.M{
			"from": "product",
			"localField": "productId",
			"foreignField": "_id",
			"as": "product",
		}},
		{"$lookup": bson.M{
			"from": "user",
			"localField": "userId",
			"foreignField": "_id",
			"as": "user",
		}},
	})
	err := pipe.Iter().All(&resp)
	return resp, err
}

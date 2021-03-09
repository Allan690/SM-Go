package models

import (
	"StoreManager/conn"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Sale struct {
	ID bson.ObjectId `bson:"_id"`
	Product bson.ObjectId `bson:"productId"`
	User bson.ObjectId `bson:"userId"`
	CreatedAt time.Time  `bson:"created_at"`
	UpdatedAt time.Time  `bson:"updated_at"`
}


type Sales []Sale

// gets a sale by id
func GetOneSale(salesId bson.ObjectId, salesCollection string)(Sales, error){
	matchStage := bson.M{
		"$match": bson.M{
			"_id": salesId,
		},
	}
	pipelineResult := make([]Sale, 0)
	pipeline := make([]bson.M, 0)
	lookupProductStage := bson.M{
		"$lookup": bson.M{
			"from": "product",
			"foreignField":"salesId",
			"localField":"productId",
			"as":"product",
		},
	}
	lookupUserStage := bson.M{ "$lookup": bson.M{
		"from": "user",
		"foreignField":"_id",
		"localField":"userId",
		"as":"user",
	}}

	pipeline = append(pipeline, matchStage, lookupProductStage, lookupUserStage)
	db := conn.GetMongoDB()
	err := db.C(salesCollection).Pipe(pipeline).All(&pipelineResult)
	return pipelineResult, err
}

// get all sales belonging to a user
func GetSalesByUserId(userId bson.ObjectId, salesCollection string)(Sale, error) {
	db := conn.GetMongoDB()
	sale := Sale{}
	err := db.C(salesCollection).Find(bson.M{"userId": &userId}).One(&sale)
	return sale, err
}

// get all sales by product id
func GetSalesByProductId(productId bson.ObjectId, salesCollection string)(Sale, error) {
	db := conn.GetMongoDB()
	sale := Sale{}
	err := db.C(salesCollection).Find(bson.M{"productId": &productId}).One(&sale)
	return sale, err
}


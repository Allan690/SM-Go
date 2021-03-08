package model

import (
	"gopkg.in/mgo.v2/bson"
)


type Product struct {
	ID bson.ObjectId  `bson:"_id"`
	Name string `bson:"name" binding:"required"`

}


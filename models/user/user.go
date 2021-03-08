package model

import (
	"StoreManager/conn"
	"gopkg.in/mgo.v2/bson"
	"time"
)

// user structure

type User struct {
	 ID  bson.ObjectId `bson:"_id"`
	 Name string `bson:"name"  binding:"required"`
	 Address   string        `bson:"address"  binding:"required"`
	 Age       int           `bson:"age"  binding:"required"`
	 CreatedAt time.Time     `bson:"created_at"`
	 UpdatedAt time.Time     `bson:"updated_at"`
}

//users list
type Users []User

func UserInfo(id bson.ObjectId, userCollection string) (User, error) {
	db := conn.GetMongoDB()
	user := User{}
	err := db.C(userCollection).Find(bson.M{"_id": &id}).One(&user)
	return user, err
}
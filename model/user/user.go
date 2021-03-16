package model

import (
	"StoreManager/conn"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

// user structure

type User struct {
	 ID  bson.ObjectId `bson:"_id" json:"id"`
	 Name string `bson:"name"  binding:"required"`
	 Email string `bson:"email" binding:"required"`
	 Password string `bson:"password" binding:"required"`
	 Address   string        `bson:"address"  binding:"required"`
	 Age       int           `bson:"age"  binding:"required"`
	 CreatedAt time.Time     `bson:"created_at"`
	 UpdatedAt time.Time     `bson:"updated_at"`
}

type LoginDetails struct {
	Email string `json:"email" binding:"required"`
	Password string  `json:"password" binding:"required"`
}

//users list
type Users []User

const UserCollection = "user"

func UserInfo(id bson.ObjectId) (User, error) {
	db := conn.GetMongoDB()
	user := User{}
	err := db.C(UserCollection).Find(bson.M{"_id": &id}).One(&user)
	return user, err
}

// creates a user
func (user *User) CreateUser(userDetails *User) (*User, error) {
	db := conn.GetMongoDB()
	bytes, err := bcrypt.GenerateFromPassword([]byte(userDetails.Password), 14)
	if err != nil {
		return nil, err
	}
	userDetails.Password = string(bytes)
	err = db.C(UserCollection).Insert(userDetails)
	userDetails.Password = ""
	return userDetails, err
}

// handles user login
func (user *LoginDetails) LoginUser() (*User, error) {
	db := conn.GetMongoDB()
	user_ := User{}
	err := db.C(UserCollection).Find(bson.M{"email": user.Email}).One(&user_)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user_.Password), []byte(user.Password))
	user_.Password = ""
	return &user_, err
}

// gets a user by email
func (user *User) GetUser(email string) (*User, error) {
	db := conn.GetMongoDB()
	err := db.C(UserCollection).Find(bson.M{"email": email}).One(&user)
	return user, err
}

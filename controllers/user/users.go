package user

import (
	"StoreManager/conn"
	user "StoreManager/models/user"
	"errors"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"time"
)

const UserCollection = "user"

var (
	errNotExist = errors.New("user does not exist")
	errInvalidId = errors.New("invalid ID")
	errInvalidBody = errors.New("invalid request body")
	errInsertionFailed = errors.New("error in the user insertion")
	errUpdateFailed  = errors.New("error in the user update")
	errDeletionFailed  = errors.New("error in the user deletion")
)

func CreateUser(c *gin.Context) {
	db := conn.GetMongoDB()
	user := user.User{}
	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInvalidBody.Error()})
		return
	}
	user.ID = bson.NewObjectId()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	err = db.C(UserCollection).Insert(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInsertionFailed.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "user": &user})
}

func UpdateUser(c *gin.Context) {
	db := conn.GetMongoDB()
	var id = bson.ObjectIdHex(c.Param("id"))
	existingUser, err := user.UserInfo(id, UserCollection)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInvalidId.Error()})
		return
	}
	err = c.Bind(&existingUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInvalidBody.Error()})
		return
	}
	existingUser.ID = id
	existingUser.UpdatedAt = time.Now()
	err = db.C(UserCollection).Update(bson.M{"_id": &id}, existingUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errUpdateFailed.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "user": &existingUser})
}

func GetAllUsers(c *gin.Context) {
	db := conn.GetMongoDB()
	users := user.Users{}
	err := db.C(UserCollection).Find(bson.M{}).All(&users)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errNotExist.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "users": &users})
}


func GetUser(c *gin.Context) {
	var id = bson.ObjectIdHex(c.Param("id")) // Get Param
	user_, err := user.UserInfo(id, UserCollection)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInvalidId.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "user": &user_})
}

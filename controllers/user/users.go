package user

import (
	"StoreManager/conn"
	user "StoreManager/model/user"
	"errors"
	"fmt"
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
	user_ := user.User{}
	err := c.Bind(&user_)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInvalidBody.Error()})
		return
	}
	user_.ID = bson.NewObjectId()
	user_.CreatedAt = time.Now()
	user_.UpdatedAt = time.Now()
	createdUser, err := user_.CreateUser(&user_)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInsertionFailed.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "user": &createdUser})
}

func UpdateUser(c *gin.Context) {
	db := conn.GetMongoDB()
	var id = bson.ObjectIdHex(c.Param("id"))
	existingUser, err := user.UserInfo(id)
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

func DeleteUser(c *gin.Context) {
	db := conn.GetMongoDB()
	var id = bson.ObjectIdHex(c.Param("id"))
	existingUser, err := user.UserInfo(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInvalidId.Error()})
		return
	}
	err = c.Bind(&existingUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInvalidBody.Error()})
		return
	}
	err = db.C(UserCollection).Remove(bson.M{"_id": &id})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errDeletionFailed.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "User deleted successfully"})
}


func GetUser(c *gin.Context) {
	var id = bson.ObjectIdHex(c.Param("id")) // Get Param
	user_, err := user.UserInfo(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInvalidId.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "user": &user_})
}

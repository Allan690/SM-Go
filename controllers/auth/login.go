package auth

import (
	user "StoreManager/model/user"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

func LoginController(c *gin.Context) {
	var payload user.LoginDetails
	err:= c.Bind(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}
	user_, err := payload.LoginUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}
	token, _ := user.JWTAuthService().GenerateToken(payload.Email, bson.ObjectId.Hex(user_.ID))
	c.JSON(http.StatusOK, gin.H{"status": "success", "token": token})
}

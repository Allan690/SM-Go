package sales

import (
	"StoreManager/conn"
	sale "StoreManager/model/sales"
	"errors"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strings"
	"time"
)

const SalesCollection = "sales"

// struct for the incoming sale creation request
type IncomingSaleCreateRequest struct {
	UserId string `json:"userId" binding:"required"`
	ProductId string `json:"productId" binding:"required"`
}

type IncomingSaleUpdateRequest struct {
	UserId string `json:"userId"`
	ProductId string `json:"productId"`
}

var (
	errNotExist = errors.New("sale does not exist")
	errInvalidId = errors.New("invalid ID")
	errInvalidBody = errors.New("invalid request body")
	errInsertionFailed = errors.New("error in the sale insertion")
	errUpdateFailed  = errors.New("error in the sale update")
)

// creates a sale
func CreateSale(c *gin.Context) {
	db := conn.GetMongoDB()
	var saleRequest IncomingSaleCreateRequest
	_sale := sale.Sale{}
	err := c.Bind(&saleRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInvalidBody.Error()})
		return
	}
	_sale.Product = bson.ObjectIdHex(saleRequest.ProductId)
	_sale.User = bson.ObjectIdHex(saleRequest.UserId)
	_sale.ID = bson.NewObjectId()
	_sale.CreatedAt = time.Now()
	_sale.UpdatedAt = time.Now()
	err = db.C(SalesCollection).Insert(_sale)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInsertionFailed.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "sale": &_sale})

}

// updates a sale by id
func UpdateSale(c *gin.Context) {
	db := conn.GetMongoDB()
	var saleRequest IncomingSaleUpdateRequest
	var id = bson.ObjectIdHex(c.Param("id"))
	existingSale, err := sale.GetSaleInfo(id, SalesCollection)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInvalidId.Error()})
		return
	}
	err = c.Bind(&saleRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInvalidBody.Error()})
		return
	}
	if len(strings.TrimSpace(saleRequest.ProductId)) > 0 {
		existingSale.Product = bson.ObjectIdHex(saleRequest.ProductId)
	}
	if len(strings.TrimSpace(saleRequest.UserId)) > 0 {
		existingSale.User = bson.ObjectIdHex(saleRequest.UserId)
	}
	existingSale.UpdatedAt = time.Now()
	err = db.C(SalesCollection).Update(bson.M{"_id": &id}, existingSale)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errUpdateFailed.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "sale": &existingSale})
}

func GetAllSales(c *gin.Context) {
	allSales, err := sale.GetAllSales(SalesCollection)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errNotExist.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "sales": &allSales})
}


// gets a specific sale
func GetSaleById(c *gin.Context) {
	var id = bson.ObjectIdHex(c.Param("id"))
	existingSale, err := sale.GetOneSale(id, SalesCollection)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errNotExist.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "sale": &existingSale})
}

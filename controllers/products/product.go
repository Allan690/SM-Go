package products

import (
	"StoreManager/conn"
	product "StoreManager/model/products"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strings"
	"time"
)


const ProductCollection = "product"

var (
	errNotExist = errors.New("product does not exist")
	errInvalidId = errors.New("invalid ID")
	errInvalidName = errors.New("invalid param name")
	errInvalidBody = errors.New("invalid request body")
	errInsertionFailed = errors.New("error in the product insertion")
	errUpdateFailed  = errors.New("error in the product update")
	errDeletionFailed  = errors.New("error in the product deletion")
)

// create a product in the db and return it
func CreateProduct(c *gin.Context) {
	db := conn.GetMongoDB()
	_product := product.Product{}
	err := c.Bind(&_product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInvalidBody.Error()})
		return
	}
	_product.ID = bson.NewObjectId()
	_product.CreatedAt = time.Now()
	_product.UpdatedAt = time.Now()
	err = db.C(ProductCollection).Insert(_product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInsertionFailed.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "product": &_product})
}

// update a product in the db
func UpdateProduct(c *gin.Context) {
	db := conn.GetMongoDB()
	var id = bson.ObjectIdHex(c.Param("id"))
	existingProduct, err := product.ProductById(id, ProductCollection)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInvalidId.Error()})
		return
	}
	err = c.Bind(&existingProduct)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInvalidBody.Error()})
		return
	}
	existingProduct.ID = id
	existingProduct.UpdatedAt = time.Now()
	err = db.C(ProductCollection).Update(bson.M{"_id": &id}, existingProduct)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errUpdateFailed.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "product": &existingProduct})
}


// fetches all products
func GetAllProducts(c *gin.Context) {
	db := conn.GetMongoDB()
	products := product.Products{}
	var q = c.Request.URL.Query()
	var name = q.Get("name")
	if len(strings.TrimSpace(name)) > 0 {
		fmt.Println("here")
		product_, err := product.ProductByName(name, ProductCollection)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInvalidName.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "success", "product": &product_})
	} else {
		err := db.C(ProductCollection).Find(bson.M{}).All(&products)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errNotExist.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "success", "products": &products})
	}
}


// get product by id
func GetProductById(c *gin.Context) {
	var id = bson.ObjectIdHex(c.Param("id")) // Get Param
	product_, err := product.ProductById(id, ProductCollection)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInvalidId.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "product": &product_})
}

// delete product
func DeleteProduct(c *gin.Context) {
	db := conn.GetMongoDB()
	var id = bson.ObjectIdHex(c.Param("id"))
	existingProduct, err := product.ProductById(id, ProductCollection)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInvalidId.Error()})
		return
	}
	err = c.Bind(&existingProduct)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInvalidBody.Error()})
		return
	}
	err = db.C(ProductCollection).Remove(bson.M{"_id": &id})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errDeletionFailed.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Product deleted successfully"})
}

package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type user struct {
	Email    string
	Password string
}

var client *mongo.Client
var coll *mongo.Collection

func mongoConnect() {
	var err error
	client, err = mongo.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
		return
	}

}
func main() {
	mongoConnect()
	r := gin.Default()
	r.POST("/userdetails", userDetailsHandler)
	r.Run()

}
func userDetailsHandler(c *gin.Context) {
	var newUser user
	if err := c.ShouldBindJSON(&newUser); err != nil {
		log.Println("userDetailsHandler ShouldBindJSON error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	coll = client.Database("Registration").Collection("signup")

	_, err := coll.InsertOne(context.TODO(), newUser)
	if err != nil {
		log.Println("userDetailsHandler MongoDB InsertOne error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save user details"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User details successfully received and saved", "email": newUser.Email, "password": newUser.Password})
}

package main

import (
    "net/http"
	"fmt"
	"encoding/json"

	"github.com/AbdulConsole/bagStoreApi/database"
    "github.com/gin-gonic/gin"
)

type Product struct {
	ID	int	`json:"id"`
	Name	string	`json:"name"`
	Price	float64	`json:"price"`
	Image	string	`json:"image"`
	Description	string	`json:"description"`
}

func main() {
	r := gin.Default()
	database.ConnectDatabase()
	
	r.POST("/products", postProduct)
	//r.GET("/products", getProducts)
	
	r.Run() // listen and serve on 0.0.0.0:8080
}

func postProduct(c *gin.Context) {
    body := Product {}

	data, err := c.GetRawData()
    if err != nil {
      c.AbortWithStatusJSON(400, "User is not defined")
      return
   }
   err = json.Unmarshal(data, &body)
   if err != nil {
      c.AbortWithStatusJSON(400, "Bad Input")
      return
   }
    //products = append(products, newProduct)
    //c.IndentedJSON(http.StatusCreated, newProduct)

	_, err = database.Db.Exec("INSERT INTO bags (name, price, image, description) VALUES($1,$2,$3,$4)", body.Name, body.Price, body.Image, body.Description);
	if err != nil {
		fmt.Println(err)
	    c.AbortWithStatusJSON(400, "Couldn't create the new user.")
	} else {
	    c.JSON(http.StatusOK, "User is successfully created.")
	}
    
}



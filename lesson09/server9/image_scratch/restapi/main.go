package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Token struct {
	ID         int
	Token      string `json:"token"`
	CreatedAtt string `json:"createdAtt"`
	ExpiredAt  string `json:"expiredAt"`
}

var idInner = 4

// postTokens adds an Token from JSON received in the request body.
func postTokens(c *gin.Context) {
	var newToken Token

	// Call BindJSON to bind the received JSON to
	// newToken.
	if err := c.BindJSON(&newToken); err != nil {
		return
	}
	newToken.ID = idInner
	// Add the new Token to the slice.
	Tokens = append(Tokens, newToken)
	c.IndentedJSON(http.StatusCreated, newToken)
	idInner++
}

// Token represents data about a record Token.

var Tokens = []Token{
	{ID: 1, Token: "Ali-Baba:Bagdad", CreatedAtt: "2021-08-01 9:57:23 pm", ExpiredAt: "2021-08-11 9:57:23 pm"},
	{ID: 2, Token: "Sherlock_Holms:London", CreatedAtt: "2021-08-01 10:1:18 pm", ExpiredAt: "2021-08-11 9:59:40 pm"},
	{ID: 3, Token: "Megre:Paris", CreatedAtt: "2021-08-01 10:1:18 pm", ExpiredAt: "2021-08-11 10:1:18 pm"},
}

func main() {
	router := gin.Default()
	router.GET("/tokens", getTokens)
	router.GET("/tokens/:id", getTokenByID)
	router.POST("/tokens", postTokens)

	router.Run("localhost:8082")
}

// getTokens responds with the list of all Tokens as JSON.
func getTokens(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Tokens)
}

// getTokenByID locates the Token whose ID value matches the id
// parameter sent by the client, then returns that Token as a response.
func getTokenByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of Tokens, looking for
	// an Token whose ID value matches the parameter.
	for _, a := range Tokens {
		ids, err := strconv.Atoi(id)
		if err != nil {
			log.Fatal(err)
			return
		}
		if a.ID == ids {

			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Token not found"})
}

package controllers

import (
	"net/http"

	"github.com/0xMarvell/watchlist/pkg/config"
	"github.com/0xMarvell/watchlist/pkg/models"
	"github.com/0xMarvell/watchlist/pkg/utils"
	"github.com/gin-gonic/gin"
)

// AddShow adds a new show to the user's watchlist
func AddShow(c *gin.Context) {
	// Get  data off request body and bind to payload struct
	var newShowPayload struct {
		Name      string `json:"name"`
		Category  string `json:"category"`
		Completed bool   `json:"completed"`
	}
	err := c.Bind(&newShowPayload)
	utils.CheckErr("c.Bind error: ", err)

	// Store details in DB
	show := models.Show{}
	result := config.DB.Create(&show)

	if (result.Error != nil) || (show.Name == "" || show.Category == "") {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "bad request: failed to add show to watchlist",
		})
		return
	}

	// Return JSON response to confirm operation success
	c.JSON(http.StatusCreated, gin.H{
		"success":      true,
		"show_details": show,
	})
}

func GetWatchlist(c *gin.Context) {

}

func GetShow(c *gin.Context) {

}

func UpdateShow(c *gin.Context) {

}

func DeleteShow(c *gin.Context) {

}

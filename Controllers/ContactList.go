package Controllers

import (
	"net/http"
        "github.com/nayaksamdeep/contact-list-service/Models"
	"github.com/gin-gonic/gin"
)

func GetContacts(c *gin.Context) {
	var contact []Models.Contact
	err := Models.GetAllContacts(&contact)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, contact)
	}
}

func CreateAContact(c *gin.Context) {
	var contact Models.Contact
	c.BindJSON(&contact)
	err := Models.CreateAContact(&contact)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, contact)
	}
}

func GetAContact(c *gin.Context) {
	id := c.Params.ByName("id")
	var contact Models.Contact
	err := Models.GetAContact(&contact, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, contact)
	}
}

func UpdateAContact(c *gin.Context) {
	var contact Models.Contact
	id := c.Params.ByName("id")
	err := Models.GetAContact(&contact, id)
	if err != nil {
		c.JSON(http.StatusNotFound, contact)
	}
	c.BindJSON(&contact)
	err = Models.UpdateAContact(&contact, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, contact)
	}
}

func DeleteAContact(c *gin.Context) {
	var contact Models.Contact
	id := c.Params.ByName("id")
	err := Models.DeleteAContact(&contact, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
	}
}

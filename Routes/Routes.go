package Routes

import (
	"../Controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("Contacts", Controllers.GetContacts)
		v1.POST("Contact", Controllers.CreateAContact)
		v1.GET("Contacts/:id", Controllers.GetAContact)
		v1.PUT("Contacts/:id", Controllers.UpdateAContact)
		v1.DELETE("Contacts/:id", Controllers.DeleteAContact)
	}

	return r
}

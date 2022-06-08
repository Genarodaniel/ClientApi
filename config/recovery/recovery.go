package recovery

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Filter(c *gin.Context, recovered interface{}) {
	if err, ok := recovered.(string); ok {
		fmt.Println("Recovered error: ", err)
	}
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Error"})
}

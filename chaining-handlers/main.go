//Golang doesn't have python-Django like decorators but here is 
//a small example of what you can do

package main 

//you'll need gin for this. Use go get -u github.com/gin-gonic/gin
import (
	"github.com/gin-gonic/gin"
	"net/http"
)	


func Handler(h gin.HandlerFunc, decors ...func(gin.HandlerFunc)gin.HandlerFunc) gin.HandlerFunc {
	for i := range decors {
		d := decors[len(decors) - 1 - i]  // iterate in reverse
		h = d(h)
	}
	return h
}

func isSuperUser (handlerFunc gin.HandlerFunc) gin.HandlerFunc {
	return func (context *gin.Context) {
		// test your condition here
		if true {
			handlerFunc(context)
		} else {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status":"you are not authorized"})
		}
	}
}


func main() {
	router := gin.Default()
	router.GET("/welcome", Handler(welcome, isSuperUser))
	router.Run()
}

func welcome(context *gin.Context) {
	context.JSON(200, gin.H{"status": "ok"})
}
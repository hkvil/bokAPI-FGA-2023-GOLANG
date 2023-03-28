package routers

import(
	"bookapi/controllers"
	"github.com/gin-gonic/gin"
)

func CarServer() *gin.Engine{
	router:=gin.Default()
	router.POST("/cars",controllers.CreateCar)
	router.PUT("/cars/:carID",controllers.UpdateCar)
	router.GET("/cars/:carID",controllers.GetCar)
	router.DELETE("/cars/:carID",controllers.GetCar)

	return router
}

func BookServer() *gin.Engine{
	router:=gin.Default()
	router.POST("/books",controllers.CreateBook)
	router.PUT("/books/:bookID",controllers.UpdateBook)
	router.GET("/books",controllers.GetAllBook)
	router.GET("/books/:bookID",controllers.GetBookByID)
	router.DELETE("/books/:bookID",controllers.DeleteBook)
	return router
}
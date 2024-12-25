package api

import (
	"github.com/gin-gonic/gin"

	"github.com/Melom01/go-boilerplate/pkg/api/handler"
	"github.com/Melom01/go-boilerplate/pkg/api/middleware"
	"github.com/Melom01/go-boilerplate/pkg/config"
)

type ServerHttp struct {
	engine *gin.Engine
}

func CreateServerHttp(firebaseService *config.FirebaseService, bookImpl *handler.BookHandler) *ServerHttp {
	engine := gin.New()
	engine.Use(gin.Logger())

	api := engine.Group("/api")
	{
		api.Use(middleware.ValidateToken(firebaseService))
		api.GET("v1/books", bookImpl.FindAll)
		api.GET("v1/books/:id", bookImpl.FindById)
		api.POST("v1/books", bookImpl.AddNewBook)
		api.DELETE("v1/books/:id", bookImpl.DeleteBookById)
	}

	return &ServerHttp{engine: engine}
}

func (sh *ServerHttp) Start(port string) {
	err := sh.engine.Run(":" + port)
	if err != nil {
		panic(err)
	}
}

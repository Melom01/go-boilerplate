package api

import "github.com/gin-gonic/gin"

type ServerHttp struct {
	engine *gin.Engine
}

func NewServerHttp() *ServerHttp {
	engine := gin.New()
	engine.Use(gin.Logger())

	return &ServerHttp{engine: engine}
}

func (sh *ServerHttp) Start(port string) {
	err := sh.engine.Run(":" + port)
	if err != nil {
		panic(err)
	}
}

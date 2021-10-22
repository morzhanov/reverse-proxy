package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type controller struct {
	r    *gin.Engine
	port string
}

type Controller interface {
	Listen() error
}

func (c *controller) handleHttpErr(ctx *gin.Context, err error) {
	ctx.String(http.StatusInternalServerError, err.Error())
	log.Printf("error in the REST handler: %s", err.Error())
}

func (c *controller) handleGet(ctx *gin.Context) {
	ctx.String(http.StatusOK, "handled get on PORT: %s", c.port)
}

func (c *controller) handlePost(ctx *gin.Context) {
	ctx.String(http.StatusOK, "handled post on PORT: %s", c.port)
}

func (c *controller) Listen() error {
	return c.r.Run(fmt.Sprintf(":%s", c.port))
}

func NewController(port string) Controller {
	r := gin.Default()
	c := &controller{r, port}
	r.GET("/", c.handleGet)
	r.POST("/", c.handleGet)
	return c
}

func main() {
	c := NewController(os.Getenv("PORT"))
	if err := c.Listen(); err != nil {
		log.Fatalf("error in server %s", err)
	}
}

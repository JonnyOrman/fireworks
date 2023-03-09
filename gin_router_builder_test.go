package fireworks

import (
	"fmt"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type RequestHandlerMock[T any] struct {
	mock.Mock
}

func (this RequestHandlerMock[T]) Handle(ginContext *gin.Context) {}

func TestAddGetBuild(t *testing.T) {
	requestHandler := new(RequestHandlerMock[interface{}])

	sut := NewGinRouterBuilder()

	sut.AddGet("/", requestHandler.Handle)

	result := sut.Build()

	routes := result.Routes()

	assert.Equal(t, len(routes), 1)

	route := routes[0]

	fmt.Println("ROUTE HANDLER: " + route.Handler)

	assert.Equal(t, "GET", route.Method)
	assert.Equal(t, "/", route.Path)
	assert.Equal(t, "github.com/jonnyorman/fireworks.RequestHandlerMock[...].Handle-fm", route.Handler)
}

func TestAddPostBuild(t *testing.T) {
	requestHandler := new(RequestHandlerMock[interface{}])

	sut := NewGinRouterBuilder()

	sut.AddPost("/", requestHandler.Handle)

	result := sut.Build()

	routes := result.Routes()

	assert.Equal(t, len(routes), 1)

	route := routes[0]

	fmt.Println("ROUTE HANDLER: " + route.Handler)

	assert.Equal(t, "POST", route.Method)
	assert.Equal(t, "/", route.Path)
	assert.Equal(t, "github.com/jonnyorman/fireworks.RequestHandlerMock[...].Handle-fm", route.Handler)
}

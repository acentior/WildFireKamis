package controller

import "github.com/gin-gonic/gin"

type Controller interface {
	Handler() gin.HandlerFunc
	Route() string
	Method() string
}

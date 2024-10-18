package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
)

func Home(c *gin.Context) {
	// Отправляем HTML-шаблон "base.html" в ответ на запрос
	fmt.Println("Rendering home.html")
	c.HTML(http.StatusOK, "home.html", gin.H{})
}
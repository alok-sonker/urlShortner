package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	COUNTER int64 = 1000000
	urlMap  map[int64]string
)

func healthCheck(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to short url service")
}
func getAllURL(c *gin.Context) {
	c.JSON(http.StatusOK, urlMap)
}
func getURL(c *gin.Context) {
	url := URLJSON{}
	c.BindJSON(&url)
	fmt.Println(url)
	urlStr := getMappedURL(url.URL)
	c.String(http.StatusOK, urlStr)
}
func createURL(c *gin.Context) {
	url := URLJSON{}
	err := c.BindJSON(&url)
	if err != nil {
		c.String(http.StatusOK, err.Error())
	}
	shortURL := createMapping(url)
	if shortURL != "" {
		c.String(http.StatusBadRequest, aliasOfURL+shortURL)
	}
}

func init() {
	urlMap = make(map[int64]string)
}
func createMapping(url URLJSON) string {
	urlMap[COUNTER] = url.URL
	urlStr := fmt.Sprintf("%v", COUNTER)
	COUNTER++
	return urlStr
}
func getMappedURL(url string) string {
	idStr := strings.Split(url, aliasOfURL)
	id, err := strconv.Atoi(idStr[1])
	if err != nil {
		return ""
	}
	if v, ok := urlMap[int64(id)]; ok {
		return v
	}
	return ""
}

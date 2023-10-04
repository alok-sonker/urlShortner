package main

import (
	redis "dep/v2/redis"
	"fmt"
	"net/http"
	parser "net/url"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	COUNTER     int64 = 1000000
	urlMap      map[int64]string
	metricsMap  map[string]int32
	redisClient redis.Redis
)

func init() {
	urlMap = make(map[int64]string)
	metricsMap = make(map[string]int32)
	redisClient = redis.NewRedisClient()
}

func healthCheck(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to short url service")
}
func getAllURL(c *gin.Context) {
	c.JSON(http.StatusOK, urlMap)
}
func getURL(c *gin.Context) {
	url := URLJSON{}
	err := c.BindJSON(&url)
	if err != nil {
		return
	}
	urlStr := getMappedURL(url.URL)
	c.String(http.StatusOK, urlStr)
}
func createURL(c *gin.Context) {
	url := URLJSON{}
	err := c.BindJSON(&url)
	if err != nil {
		c.String(http.StatusOK, err.Error())
	}
	shortURL := createMapping(url.URL)
	hostName, err := parseURL(url.URL)
	if err == nil && hostName != "" {
		if _, ok := metricsMap[hostName]; !ok {
			metricsMap[hostName] = 1
		} else {
			metricsMap[hostName]++
		}
	}

	if shortURL != "" {
		c.String(http.StatusBadRequest, aliasOfURL+shortURL)
	}
}

func indexHandler(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

type myForm struct {
	URL string `form:"url"`
}

func formHandler(c *gin.Context) {
	var fakeForm myForm
	err := c.Bind(&fakeForm)
	if err != nil {
		return
	}

	shortURL := createMapping(fakeForm.URL)
	hostName, err := parseURL(fakeForm.URL)
	if err == nil && hostName != "" {
		if _, ok := metricsMap[hostName]; !ok {
			metricsMap[hostName] = 1
		} else {
			metricsMap[hostName]++
		}
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"url": aliasOfURL + shortURL,
	})
	//c.JSON(200, gin.H{"url": fakeForm.URL})
}

func parseURL(url string) (string, error) {
	if !strings.Contains(url, "http") {
		url = "http://" + url
	}
	u, err := parser.Parse(url)
	if err != nil {
		return "", err
	}
	return u.Hostname(), nil
}
func getMetrics(c *gin.Context) {
	m := map[string]int32{}
	if len(metricsMap) > 3 {
		m = sortMapByValue(metricsMap)
	} else {
		m = metricsMap
	}
	c.JSON(http.StatusOK, m)
}

func createMapping(url string) string {
	urlMap[COUNTER] = url

	counterStr := fmt.Sprintf("%v", COUNTER)
	err := redisClient.SetValue(counterStr, url)
	if err != nil {
		return ""
	}
	COUNTER++
	return counterStr
}
func getMappedURL(url string) string {
	idStr := strings.Split(url, aliasOfURL)
	id, err := strconv.Atoi(idStr[1])
	if err != nil {
		return ""
	}
	//if v, ok := urlMap[int64(id)]; ok {
	//	return v
	//}
	if v, ok := redisClient.GetValue(fmt.Sprintf("%v", id)); ok {
		return v
	}
	return ""
}

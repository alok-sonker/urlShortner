package main

import (
	"fmt"
)

const (
	LEGTH_OF_URL = 7
)

// Create a basic URL shortner with map
var (
	COUNTER    int64 = 1000000
	urlMap     map[int64]string
	aliasOfURL string = "https://url.com/"
)

func main() {

	// URL to short
	URL := "http://aloksonker.me"
	// Intialise map
	urlMap = make(map[int64]string)
	urlMap[COUNTER] = URL
	fmt.Println(aliasOfURL + fmt.Sprintf("%v", COUNTER))
	COUNTER++
}

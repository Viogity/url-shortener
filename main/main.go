package main

import (
	"fmt"
	_ "fmt"
	_ "strings"
)

func main() {
	fmt.Print("Hello, Enter your URL: ")
	var longUrl string = "url"
	fmt.Scanln(longUrl)
	fmt.Println(longUrl)
}

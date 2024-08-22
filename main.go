package main

import "fmt"

// import (
// 	"net/http"

// 	"github.com/labstack/echo/v4"
// )

// func main() {
// 	e := echo.New()
// 	e.GET("/", func(c echo.Context) error {
// 		return c.String(http.StatusOK, "Hello, World!")
// 	})
// 	e.Logger.Fatal(e.Start(":1323"))
// }

func main() {
	//implicit type
	var a int = 10
	fmt.Println(a)
	//Dynamisc type
	b := 20
	fmt.Println(b)
	c := 2.2
	fmt.Println(c)
}

package router

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	// "github.com/gorilla/websocket"
)

func dumpHandler(c echo.Context, reqBody, resBody []byte) {
	output := fmt.Sprintf("%#v", c.Request().Header)

	fmt.Printf("Request Header: %v\n", output)
	fmt.Printf("Request Body: %v\n", string(reqBody))
	fmt.Printf("Response Body: %v\n", string(resBody))
}

func Init() *echo.Echo {

	e := echo.New()

	e.Use(middleware.BodyDump(dumpHandler))

	// hub := newHub()
	// go hub.run()
	e.GET("/ws", func (c echo.Context) error {
		// serveWs(hub, c.Response(), c.Request())
		// return nil
		return c.String(http.StatusOK, "Hello, World!")
	})

	return e
}

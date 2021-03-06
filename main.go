package main

import (
	"fmt"
	"net/http"
	"os"

	"goConfig/config"

	_ "goConfig/init"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		env := os.Getenv("configor_env")
		fmt.Println("env:", env, "conn", config.Config.DB.Conn)
		conn := config.Config.DB.Conn
		return c.String(http.StatusOK, conn)
	})
	//config.InitConfig()

	e.Start(":1113")
}

package main

import (
	"fmt"
	gocon "github.com/ahmetask/go-con"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func main() {
	e := echo.New()

	configServerPort := "8082"
	applicationPort := "8081"

	type CustomConfigInner struct {
		A string
	}

	type CustomConfig struct {
		C int
		D bool
		E CustomConfigInner
	}

	var MyConfig = CustomConfig{
		C: 1,
		D: true,
		E: CustomConfigInner{
			A: "go-con-client-2",
		},
	}

	go func() {
		for {
			fmt.Println("go-con-client-2")
			fmt.Println(fmt.Sprintf("My Config %v", MyConfig))
			time.Sleep(10 * time.Second)
		}
	}()

	go func() {
		// Pass Your Config variable address to Listener
		err := gocon.ListenConfigChange(configServerPort, &MyConfig, nil)
		panic(err)
	}()

	e.GET("/config", func(c echo.Context) error {
		return c.JSON(http.StatusOK, MyConfig)
	})

	e.Logger.Fatal(e.Start(":" + applicationPort))
}

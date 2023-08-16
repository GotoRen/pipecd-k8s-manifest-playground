package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/GotoRen/pipecd-k8s-manifest-playground/webhook/internal"
	"github.com/labstack/echo/v4"
)

const addrPort = ":8081"

func main() {
	e := echo.New()
	e.HideBanner = true

	e.POST("/api/hooks/services", internal.NewHooks) // Webhook endpoint -> http://localhost:8081/api/hooks/services

	err := e.Start(addrPort)
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}

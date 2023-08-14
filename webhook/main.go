package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/GotoRen/pipecd-k8s-manifest-playground/webhook/domain"
	"github.com/labstack/echo/v4"
)

const addrPort = ":8081"

func main() {
	e := echo.New()
	e.HideBanner = true

	e.POST("/api/hooks/services", newHooks) // Webhook endpoint -> http://localhost:8081/api/hooks/services

	err := e.Start(addrPort)
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}
}

// newHooks reads the contents of the PipeCD Webhook and returns HooksData.
func newHooks(ctx echo.Context) error {
	allHeaders := ctx.Request().Header
	for key, values := range allHeaders {
		for _, value := range values {
			log.Printf("Info - Header => %s: %s\n", key, value)
		}
	}

	webhookData := new(domain.HooksData)

	body, err := io.ReadAll(ctx.Request().Body)
	if err != nil {
		return err
	}
	defer ctx.Request().Body.Close()

	err = json.Unmarshal(body, &webhookData)
	if err != nil {
		return err
	}

	jsonData, err := json.MarshalIndent(webhookData, "", "  ")
	if err != nil {
		return err
	}

	fmt.Fprintln(os.Stdout, string(jsonData))

	return nil
}

package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/GotoRen/pipecd-k8s-manifest-playground/webhook/domain"
	"github.com/GotoRen/pipecd-k8s-manifest-playground/webhook/helper"
	"github.com/labstack/echo/v4"
)

// NewHooks reads the contents of the PipeCD Webhook and returns HooksData.
func NewHooks(ctx echo.Context) error {
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

	log.Println("[DEBUG] Deployment.Trigger.Commit.CreatedAt:", helper.EpochSeconds(webhookData.Metadata.Deployment.Trigger.Commit.CreatedAt).ConvertToISO8601())
	log.Println("[DEBUG] Deployment.Trigger.Timestamp", helper.EpochSeconds(webhookData.Metadata.Deployment.Trigger.Timestamp).ConvertToISO8601())
	log.Println("[DEBUG] Deployment.CreatedAt", helper.EpochSeconds(webhookData.Metadata.Deployment.CreatedAt).ConvertToISO8601())
	log.Println("[DEBUG] Deployment.UpdatedAt", helper.EpochSeconds(webhookData.Metadata.Deployment.UpdatedAt).ConvertToISO8601())

	return nil
}

package main

import (
	"fmt"
	"net/url"

	"github.com/GotoRen/pipecd-k8s-manifest-playground/grafana-client/internal"
	grafana "github.com/grafana/grafana-api-golang-client"
)

func main() {
	gclient, err := internal.NewGrafanaClient()
	if err != nil {
		fmt.Println(err)
	}

	// 全てのアノテーションを取得
	params := url.Values{} // 範囲指定する場合: https://github.com/grafana/grafana-api-golang-client/blob/master/annotation_test.go#L53-L56
	annotations, err := gclient.GetAnnotations(params)
	if err != nil {
		fmt.Println(err)
	}

	for _, annotation := range annotations {
		fmt.Println("ID:", annotation.ID)
		fmt.Println("AlertID:", annotation.AlertID)
		fmt.Println("DashboardID:", annotation.DashboardID)
		fmt.Println("DashboardUID:", annotation.DashboardUID)
		fmt.Println("PanelID:", annotation.PanelID)
		fmt.Println("UserID:", annotation.UserID)
		fmt.Println("UserName:", annotation.UserName)
		fmt.Println("NewState:", annotation.NewState)
		fmt.Println("PrevState:", annotation.PrevState)
		fmt.Println("Time:", annotation.Time)
		fmt.Println("TimeEnd:", annotation.TimeEnd)
		fmt.Println("Text:", annotation.Text)
		fmt.Println("Metric:", annotation.Metric)
		fmt.Println("RegionID:", annotation.RegionID)
		fmt.Println("Type:", annotation.Type)
		fmt.Println("Tags:", annotation.Tags)
		fmt.Println("IsRegion:", annotation.IsRegion)
	}

	// annotations 追加処理
	define := &grafana.Annotation{
		DashboardUID: "5Aiamcj4k",
		Time:         1689655800000,
		IsRegion:     false,
		Tags:         []string{"event:deployed"},
		Text:         "sample",
	}

	res, err := gclient.CreateAnnotations(define)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("[DEBUG] AnnotationID:", res)
}

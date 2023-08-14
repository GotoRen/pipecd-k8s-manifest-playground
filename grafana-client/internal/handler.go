package internal

import (
	"net/url"

	grafana "github.com/grafana/grafana-api-golang-client"
)

// GetAnnotations gets all Grafana annotations.
func (client *GrafanaClient) GetAnnotations(params url.Values) ([]*grafana.Annotation, error) {
	annotationList := []*grafana.Annotation{}

	annotations, err := client.Conn.Annotations(params)
	if err != nil {
		return nil, err
	}

	for _, v := range annotations {
		annotation := &grafana.Annotation{
			ID:           v.ID,
			AlertID:      v.AlertID,
			DashboardID:  v.DashboardID,
			DashboardUID: v.DashboardUID,
			PanelID:      v.PanelID,
			UserID:       v.UserID,
			UserName:     v.UserName,
			NewState:     v.NewState,
			PrevState:    v.PrevState,
			Time:         v.Time,
			TimeEnd:      v.TimeEnd,
			Text:         v.Text,
			Metric:       v.Metric,
			RegionID:     v.RegionID,
			Type:         v.Type,
			Tags:         v.Tags,
			IsRegion:     v.IsRegion,
		}

		annotationList = append(annotationList, annotation)
	}

	return annotationList, nil
}

// CreateAnnotations creates new Grafana annotations.
func (client *GrafanaClient) CreateAnnotations(define *grafana.Annotation) (int64, error) {
	res, err := client.Conn.NewAnnotation(define)
	if err != nil {
		return 1, err
	}

	return res, nil
}

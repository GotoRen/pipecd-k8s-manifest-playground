package internal

import (
	grafana "github.com/grafana/grafana-api-golang-client"
)

const (
	apiKey   = "glsa_zhl48pD4kccNbIHCia8TCHVllHpN36oc_dc5c0e32"
	endPoint = "http://10.10.10.85"
)

type GrafanaClient struct {
	Conn *grafana.Client
}

func NewGrafanaClient() (*GrafanaClient, error) {
	conn, err := grafana.New(endPoint, grafana.Config{APIKey: apiKey})
	if err != nil {
		return nil, err
	}

	return &GrafanaClient{Conn: conn}, nil
}

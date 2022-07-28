package alerting

import (
	"path"
	// cfg "github.com/prometheus/alertmanager/config"
)

const (
	GET    = "GET"
	POST   = "POST"
	DELETE = "DELETE"
	v2     = "/api/v2"
	v1     = "/api/v1"
)

// type ConfigMapData struct {
// }

// func UnmarshalConfigMap(p *Plugin) {
// 	// get config map data from k8s client
// 	// key := "alertmanager.yaml"

// 	// unmarshal config map data
// 	c := cfg.Config
// 	for _, r := range c.Receivers {

// 	}
// 	// return that data
// }

type AlertManagerAPI struct {
	endpoint string
	api      string
	route    string
	verb     string
}

func (a *AlertManagerAPI) construct() string {
	return path.Join(a.endpoint, a.api, a.route)
}

func (a *AlertManagerAPI) IsReady() bool {
	return false
}

func (a *AlertManagerAPI) IsHealthy() bool {
	return false
}

// WithHttpV2
//## OpenAPI reference
// https://github.com/prometheus/alertmanager/blob/main/api/v2/openapi.yaml
//
func WithHttpV2(verb string, endpoint string, route string) *AlertManagerAPI {
	return &AlertManagerAPI{
		endpoint: endpoint,
		api:      v2,
		route:    route,
		verb:     verb,
	}
}

// WithHttpV1
// ## Reference
// https://prometheus.io/docs/alerting/latest/clients/
func WithHttpV1(verb string, endpoint string, route string) *AlertManagerAPI {
	return &AlertManagerAPI{
		endpoint: endpoint,
		api:      v1,
		route:    route,
		verb:     verb,
	}
}

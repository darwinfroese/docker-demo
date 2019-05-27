package server

import (
	"log"
	"net/http"
)

const (
	healthCheckFailMessage = "Unable to perform health check"
	healthyMessage         = "HEALTHY"
	unhealthyMessage       = "UNHEALTHY"
	unavailableMessage     = "UNAVAILABLE"
)

func healthCheckFull() *healthInfo {
	health := healthInfo{
		WebStatus:   healthyMessage,
		LoginStatus: healthCheckService("http://login.docker.demo:8000/api/v1/health"),
		ShopStatus:  healthCheckService("http://shop.docker.demo:8000/api/v1/health"),
	}

	return &health
}

func healthCheckService(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return healthCheckFailMessage
	}

	switch resp.StatusCode {
	case http.StatusOK:
		return healthyMessage
	case http.StatusNotFound:
		return unavailableMessage
	default:
		return unhealthyMessage
	}
}

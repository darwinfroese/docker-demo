package server

import (
	"log"
	"net/http"
)

const (
	healthyMessage     = "HEALTHY"
	unhealthyMessage   = "UNHEALTHY"
	unavailableMessage = "UNAVAILABLE"
)

func healthCheckFull() *healthInfo {
	health := healthInfo{
		WebStatus:   healthyMessage,
		LoginStatus: healthCheckService("http://login.docker.demo/api/v1/health"),
		ShopStatus:  healthCheckService("http://shop.docker.demo/api/v1/health"),
	}

	return &health
}

func healthCheckService(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return unavailableMessage
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

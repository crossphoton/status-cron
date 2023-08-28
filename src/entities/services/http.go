package services

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"net/http"

	"github.com/crossphoton/status-cron/src/entities"
)

const HTTP entities.ServiceType = "http"

type HTTPService struct {
	Service entities.Service
}

func (service *HTTPService) Run() (res entities.Result) {
	res.ServiceId = service.Service.Id

	data := service.Service.Data
	status, ok := data["status"].(float64)
	if !ok || data["status"] == 0 {
		status = 200
	}
	method, ok := data["method"].(string)
	if !ok {
		method = "GET"
	}

	var body bytes.Buffer
	enc := gob.NewEncoder(&body)
	enc.Encode(data["body"])

	// Form http request
	req, err := http.NewRequest(method, service.Service.Url, &body)
	if err != nil {
		res.Reason = err.Error()
		return
	}

	// Add headers
	d, ok := data["headers"].(map[string]interface{})
	if ok {
		for k, v := range d {
			req.Header.Add(k, v.(string))
		}
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return entities.Result{
			ServiceId: service.Service.Id,
			Success:   false,
			Reason:    err.Error(),
		}
	}

	if resp.StatusCode != int(status) {
		res.Reason = fmt.Sprintf("Got status %d instead.", resp.StatusCode)
		return
	}

	res.Success = true
	return
}

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

const ALERT_ENDPOINT = "/alert"
const DEFAULT_PORT = "8080"

func main() {
	servicePort := getEnvVariable("ANB_PORT", DEFAULT_PORT)
	log.Infof("starting service on port %s", servicePort)
	http.HandleFunc(ALERT_ENDPOINT, HandleAlert)
	http.ListenAndServe(fmt.Sprintf(":%s", servicePort), nil)
}

func HandleAlert(responseWriter http.ResponseWriter, request *http.Request) {

	requestMethod := request.Method

	log.Infof("handle %s request on %s", requestMethod, ALERT_ENDPOINT)

	if requestMethod == "POST" {
		body, error := io.ReadAll(request.Body)
		if error != nil {
			panic(error)
		}
		log.Debugf("POST body: %s", string(body))
	} else {
		log.Warnf("HTTP method %s for endpoint %s not found", requestMethod, ALERT_ENDPOINT)
	}

	log.Info("handled request")
}

func getEnvVariable(key string, fallback string) string {
	envVariable := os.Getenv(key)
	if len(envVariable) == 0 {
		envVariable = fallback
	}
	return envVariable
}

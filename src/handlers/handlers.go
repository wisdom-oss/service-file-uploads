package handlers

import (
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"

	e "microservice/errors"
	"microservice/helpers"
	"microservice/vars"
)

func AuthorizationCheck(nextHandler http.Handler) http.Handler {
	return http.HandlerFunc(
		func(responseWriter http.ResponseWriter, request *http.Request) {
			logger := log.WithFields(
				log.Fields{
					"middleware": true,
					"title":      "AuthorizationCheck",
				},
			)
			logger.Debug("Checking the incoming request for authorization information set by the gateway")
			if request.URL.Path == "/ping" {
				nextHandler.ServeHTTP(responseWriter, request)
				return
			}
			// Get the scopes the requesting user has
			scopes := request.Header.Get("X-Authenticated-Scope")
			// Check if the string is empty
			if strings.TrimSpace(scopes) == "" {
				logger.Warning("Unauthorized request detected. The required header had no content or was not set")
				helpers.SendRequestError(e.UnauthorizedRequest, responseWriter)
				return
			}

			scopeList := strings.Split(scopes, ",")
			if !helpers.StringArrayContains(scopeList, vars.ScopeConfiguration.ScopeValue) {
				logger.Error("Request rejected. The user is missing the scope needed for accessing this service")
				helpers.SendRequestError(e.MissingScope, responseWriter)
				return
			}
			// Call the next handler which will continue handling the request
			nextHandler.ServeHTTP(responseWriter, request)
		},
	)
}

/*
PingHandler

This handler is used to test if the service is able to ping itself. This is done to run a healthcheck on the container
*/
func PingHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

// UploadHandler accepts the request body and stores it in a location
func UploadHandler(responseWriter http.ResponseWriter, request *http.Request) {
	// Create the files hash sum
	hashSum, err := helpers.CalculateFileHash(request.Body)
	if err != nil {
		helpers.SendRequestError(e.UnprocessableEntity, responseWriter)
	}

	// SQL Query: Query the existence of the hash sum in the file database
	sqlFileUploadedQuery := `SELECT EXISTS(SELECT id FROM files.uploads WHERE hash = $1)`
	uploadedFiles, err := vars.PostgresConnection.Query(sqlFileUploadedQuery, hashSum)
	var fileExists bool
	for uploadedFiles.Next() {
		scanError := uploadedFiles.Scan(&fileExists)
		if scanError != nil {
			log.WithError(scanError).Error("Unable to scan the result rows")
			helpers.SendRequestError(e.UnsupportedHTTPMethod, responseWriter) // FIXME: Change to internal server error
		}
	}
}

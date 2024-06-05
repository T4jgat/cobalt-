package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)

	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthCheckerHandler)

	router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)
	router.HandlerFunc(http.MethodPut, "/v1/users/activated", app.activateUserHandler)

	router.HandlerFunc(http.MethodPut, "/v1/users/update/:id", app.requirePermission("info:write", app.updateUserHandler))
	router.HandlerFunc(http.MethodDelete, "/v1/users/delete/:id", app.requirePermission("info:write", app.deleteUserHandler))
	//router.HandleFunc(http.MethodGet, app.getAllUsersHandler)
	router.HandlerFunc(http.MethodGet, "/v1/users", app.getAllUsersHandler)

	router.HandlerFunc(http.MethodPost, "/v1/tokens/authentication", app.createAuthenticationTokenHandler)

	return app.recoverPanic(app.rateLimit(app.authenticate(router)))
}

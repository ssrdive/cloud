package main

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	r := mux.NewRouter()
	r.Handle("/", http.HandlerFunc(app.home)).Methods("GET")
	r.HandleFunc("/authenticate", http.HandlerFunc(app.authenticate)).Methods("POST")
	r.Handle("/dropdown/{name}", app.validateToken(http.HandlerFunc(app.dropdownHandler))).Methods("GET")
	r.Handle("/dropdown/condition/{name}/{where}/{value}", app.validateToken(http.HandlerFunc(app.dropdownConditionHandler))).Methods("GET")

	r.Handle("/item/create", app.validateToken(http.HandlerFunc(app.createItem))).Methods("POST")
	r.Handle("/item/all", app.validateToken(http.HandlerFunc(app.allItems))).Methods("GET")
	r.Handle("/item/search", app.validateToken(http.HandlerFunc(app.itemSearch))).Methods("GET")
	r.Handle("/item/{id}", app.validateToken(http.HandlerFunc(app.itemDetails))).Methods("GET")
	r.Handle("/item/details/byid/{id}", app.validateToken(http.HandlerFunc(app.itemDetailsById))).Methods("GET")
	r.Handle("/item/update/byid", app.validateToken(http.HandlerFunc(app.updateItemById))).Methods("POST")

	r.Handle("/saleWatches/all", app.validateToken(http.HandlerFunc(app.saleWatches))).Methods("GET")
	r.Handle("/saleAll/all", app.validateToken(http.HandlerFunc(app.saleAll))).Methods("GET")
	r.Handle("/saleCloudIDSearch/search", app.validateToken(http.HandlerFunc(app.saleCloudIDSearch))).Methods("GET")
	r.Handle("/saleChassisSearch/search", app.validateToken(http.HandlerFunc(app.saleChassisSearch))).Methods("GET")
	r.Handle("/saleInfoSearch/search", app.validateToken(http.HandlerFunc(app.saleInfoSearch))).Methods("GET")
	r.Handle("/saleCloudIDInfor/details/{id}", app.validateToken(http.HandlerFunc(app.saleCloudIdInformation))).Methods("GET")
	r.Handle("/cloudIDComments/details/{id}", app.validateToken(http.HandlerFunc(app.SaleCloudIDComments))).Methods("GET")
	r.Handle("/daterangeall/search", app.validateToken(http.HandlerFunc(app.sysByRangeAll))).Methods("GET")

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	r.Handle("/businesspartner/create", app.validateToken(http.HandlerFunc(app.createBusinessPartner))).Methods("POST")

	r.Handle("/static/", http.StripPrefix("/static", fileServer))

	return standardMiddleware.Then(handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(r))
}

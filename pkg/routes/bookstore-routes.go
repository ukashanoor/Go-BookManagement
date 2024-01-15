package routes

import "github.com/gorilla/mux"

var RegisterBookStoreRoutes = func(r *mux.Router) {
	r.HandleFunc("/book/", controller.GetBook).Methods("GET")
	r.HandleFunc("/book/", controller.CreateBook).Methods("POST")
	r.HandleFunc("/book/{id}", controller.GetBookById).Methods("GET")
	r.HandleFunc("/book/{id}", controller.UpdateBook).Methods("PUT")
	r.HandleFunc("/book/{id}", controller.DeleteBook).Methods("DELETE")
}

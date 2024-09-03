package handler

func (d *bookHandler) Routes() {

	d.r.HandleFunc("/api/v1/books", AuthAdminMiddleware(d.CreateBook)).Methods("POST")
	d.r.HandleFunc("/api/v1/books", d.GetAllBooks).Methods("GET")
	d.r.HandleFunc("/api/v1/books/{id}", AuthAdminMiddleware(d.GetBookById)).Methods("GET")
	d.r.HandleFunc("/api/v1/books/{id}", AuthAdminMiddleware(d.UpdateBook)).Methods("PUT")
	d.r.HandleFunc("/api/v1/books/{id}", AuthAdminMiddleware(d.DeleteBook)).Methods("DELETE")
}

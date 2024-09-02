package handler

func (d *authorHandler) Routes() {

	d.r.HandleFunc("/api/v1/authors", AuthAdminMiddleware(d.CreateAuthor)).Methods("POST")
	d.r.HandleFunc("/api/v1/authors", d.GetAllAuthors).Methods("GET")
	d.r.HandleFunc("/api/v1/authors/{id}", AuthAdminMiddleware(d.GetAuthorById)).Methods("GET")
	d.r.HandleFunc("/api/v1/authors/{id}", AuthAdminMiddleware(d.UpdateAuthor)).Methods("PUT")
	d.r.HandleFunc("/api/v1/authors/{id}", AuthAdminMiddleware(d.DeleteAuthor)).Methods("DELETE")
}

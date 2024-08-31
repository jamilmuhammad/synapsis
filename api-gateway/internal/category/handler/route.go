package handler

func (d *categoryHandler) Routes() {

	d.r.HandleFunc("/api/v1/categories", AuthAdminMiddleware(d.CreateCategory)).Methods("POST")
	d.r.HandleFunc("/api/v1/categories", d.GetAllCategories).Methods("GET")
	d.r.HandleFunc("/api/v1/categories/{id}", AuthAdminMiddleware(d.GetCategoryById)).Methods("GET")
	d.r.HandleFunc("/api/v1/categories/{id}", AuthAdminMiddleware(d.UpdateCategory)).Methods("PUT")
	d.r.HandleFunc("/api/v1/categories/{id}", AuthAdminMiddleware(d.DeleteCategory)).Methods("DELETE")
}

package delivery

func (d *usersDelivery) Routes() {
	d.r.HandleFunc("/users", d.CreateArticle).Methods("POST")
	d.r.HandleFunc("/users", d.GetArticles).Methods("GET")
	d.r.HandleFunc("/users/{id}", d.GetArticle).Methods("GET")
	d.r.HandleFunc("/users/{id}", d.UpdateArticle).Methods("PUT")
	d.r.HandleFunc("/users/{id}", d.DeleteArticle).Methods("DELETE")
}

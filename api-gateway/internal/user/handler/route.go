package handler

func (d *userHandler) Routes() {

	d.r.HandleFunc("/api/v1/users", AuthAdminMiddleware(d.CreateUser)).Methods("POST")
	d.r.HandleFunc("/api/v1/users/auth", d.Login).Methods("POST")
	d.r.HandleFunc("/api/v1/users/auth/refresh-token", AuthAdminMiddleware(d.RefreshToken)).Methods("POST")
	d.r.HandleFunc("/api/v1/users", d.GetAllUsers).Methods("GET")
	d.r.HandleFunc("/api/v1/users/{id}", AuthAdminMiddleware(d.GetUserById)).Methods("GET")
	d.r.HandleFunc("/api/v1/users/{id}", AuthAdminMiddleware(d.UpdateUser)).Methods("PUT")
	d.r.HandleFunc("/api/v1/users/{id}", AuthAdminMiddleware(d.DeleteUser)).Methods("DELETE")
}

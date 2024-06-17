package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/register", app.RegistrationHandler)
	mux.HandleFunc("/api/v1/login", app.LoginHandler)
	mux.HandleFunc("/api/v1/doctor", app.CreateDoctorHandler)
	mux.HandleFunc("/api/v1/doctors", app.GetDoctorsByAgeHandler)
	mux.HandleFunc("/api/v1/doctor/:name", app.GetDoctorByNameHandler)
	mux.HandleFunc("/api/v1/doctor/update/:name", app.UpdateDoctorInfoHandler)
	mux.HandleFunc("/api/v1/doctor/delete/:name", app.DeleteDoctorHandler)
	return mux
}

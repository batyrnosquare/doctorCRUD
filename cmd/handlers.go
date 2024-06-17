package main

import (
	"doctorsFinal/internal"
	"encoding/json"
	"net/http"
)

func (app *application) RegistrationHandler(w http.ResponseWriter, r *http.Request) {

	var user internal.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = app.users.CreateUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("User created"))
}

func (app *application) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user internal.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}
	user, err = app.users.FindUser(user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	token, err := internal.GenerateToken(user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Authorization", token)
	w.Write([]byte("User logged in"))
}

func (app *application) CreateDoctorHandler(w http.ResponseWriter, r *http.Request) {
	var doctor internal.Doctor
	err := json.NewDecoder(r.Body).Decode(&doctor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = app.doctors.CreateDoctor(&doctor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Doctor created"))
}

func (app *application) GetDoctorsByAgeHandler(w http.ResponseWriter, r *http.Request) {
	sortDir := r.URL.Query().Get("sortDir")
	var doctors []internal.Doctor
	if sortDir == "asc" {
		doctors, err := app.doctors.GetDoctorsSortByAgeAsc()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(doctors)
		return
	}
	if sortDir == "desc" {
		doctors, err := app.doctors.GetDoctorsSortByAgeDesc()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(doctors)
		return
	}

	json.NewEncoder(w).Encode(doctors)
}

func (app *application) GetDoctorByNameHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	doctor, err := app.doctors.GetDoctorByName(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(doctor)
}

func (app *application) UpdateDoctorInfoHandler(w http.ResponseWriter, r *http.Request) {
	var doctor internal.Doctor
	err := json.NewDecoder(r.Body).Decode(&doctor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = app.doctors.UpdateDoctor(&doctor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Doctor info updated"))
}

func (app *application) DeleteDoctorHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	err := app.doctors.DeleteDoctor(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Doctor deleted"))
}

package controllers

import "net/http"

// CreateUser Create a new user
func CreateUser(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Creating user"))
}

// GetUsers Get all users
func GetUsers(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Getting all users"))
}

// GetUser Get user by id
func GetUser(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Getting user"))
}

// UpdateUser Update user by id
func UpdateUser(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Updating user"))
}
// DeleteUser remove user by id
func DeleteUser(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Deleting user"))
}

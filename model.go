package main

// Users models
type Users struct {
	ID   int    `form:"id" json:"id"`
	Name string `form:"name" json:"name"`
	Age  int    `form:"age" json:"age"`
}

// Response return value
type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Users
}

package main

import (
	"fmt"
	"net/http"
	"trafilea-prepare/controller"
	"trafilea-prepare/repository"
	"trafilea-prepare/service"
)

func main() {
	repo := repository.NewSubscriptionRepository()
	service := service.NewSubscriptionService(repo)
	controller := controller.NewSubscriptionController(service)

	http.HandleFunc("/subscriptions", controller.GetSubscriptions)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}

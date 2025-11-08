package main

import (
	"fmt"
	"net/http"
	"trafilea-prepare/infrastructure/repository"
	"trafilea-prepare/interfaces/http/handler"
	"trafilea-prepare/usecase/subscription"
)

func main() {
	// Dependency injection: outer layers depend on inner layers
	// Infrastructure layer
	repo := repository.NewSubscriptionRepository()

	// Use case layer
	subscriptionUseCase := subscription.NewSubscriptionUseCase(repo)

	// Interface layer
	subscriptionHandler := handler.NewSubscriptionHandler(subscriptionUseCase)

	// HTTP routing
	http.HandleFunc("/subscriptions", subscriptionHandler.GetSubscriptions)
	
	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}


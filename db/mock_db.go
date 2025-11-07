package db

import (
	"trafilea-prepare/model"
)

var Subscriptions = []model.Subscription{
	{ID: 1, UserID: 1, Service: "Netflix", Status: "active"},
	{ID: 2, UserID: 1, Service: "Spotify", Status: "canceled"},
	{ID: 3, UserID: 2, Service: "Amazon Prime", Status: "active"},
}

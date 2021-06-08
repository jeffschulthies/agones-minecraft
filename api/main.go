package main

import (
	"go.uber.org/zap"

	"agones-minecraft/config"
	"agones-minecraft/db"
	"agones-minecraft/log"
	"agones-minecraft/routers"
	"agones-minecraft/services/auth/jwt"
	"agones-minecraft/services/auth/sessions"
	"agones-minecraft/services/auth/twitch"
	"agones-minecraft/services/k8s"
	"agones-minecraft/services/k8s/agones"
	"agones-minecraft/services/validator"
)

func main() {
	// Load environment variables and .env config
	config.LoadConfig()
	// Sets global zap logger
	log.Init()
	// Initializes k8s cluster config
	k8s.InitConfig()
	// Initializes agones client and informer
	agones.Init()
	// Initializes cookie store
	sessions.Init()
	// Initializes database connections and migrates (in development)
	db.Init()
	// Initializes Twitch ODIC provider for login with Twitch
	twitch.Init()
	// Initializes JWT token store
	jwt.Init()
	// Initializes custom validators
	validator.InitV1()

	r := routers.NewRouter()

	port := config.GetPort()
	if err := r.Run(":" + port); err != nil {
		zap.L().Fatal("error starting server", zap.Error(err))
	}
}

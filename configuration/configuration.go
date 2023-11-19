package configuration

import (
	"go.dscinflux.xyz/types"
)

func getConfig() types.Config {
	return types.Config{
		ApiVersion: 1,
		Database: types.Database{
			Url: "mongodb+srv://admin:c27Yt7yvcFZnll54@dscinflux.46fj1ar.mongodb.net/",
		},
		Web: types.Web{
			Port:           "7070",
			ImageUploadKey: "d6b358da96fd1e63b3eadc17fbae7bdb",
			//ReturnUrl: "http://localhost:8080/v1/auth/callback"
			ReturnUrl:      "https://dscinflux.xyz",
		},
		Client: types.Client{
			Id:       "1148743937599746141",
			Secret:   "La1_WJKl0M9jGoePXrXhegYBK3O1Zyy_",
			Token:    "MTE0ODc0MzkzNzU5OTc0NjE0MQ.GM1vMN.FKPXKgXZZmxIU3Af38L-G6ijZ4A8LbADko_x-g",
			//Callback: "http://localhost:8008/v1/auth/callback"
			Callback: "https://dscinflux.xyz/v1/auth/callback",
		},
		Collection: "entities",
		//APIUrl: "http://localhost:8080/v1/auth/callback"
		APIUrl:     "https://dscinflux.xyz/v1/auth/callback",
	}
}

func GetConfig() types.Config {
	return getConfig()
}

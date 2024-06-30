package configuration

import (
	"os"

	"go.dscinflux.xyz/types"
)

func getConfig() types.Config {
	return types.Config{
		ApiVersion: 1,
		Database: types.Database{
			Url: os.Getenv("DATABASE_URL"),
		},
		Web: types.Web{
			Port:           "3000",
			ImageUploadKey: os.Getenv("ImageUploadKey"),
			//ReturnUrl: "http://localhost:8080/v1/auth/callback"
			ReturnUrl: "https://dscinflux.xyz",
		},
		Client: types.Client{
			Id:     "1215433663085285466",
			Secret: os.Getenv("CleintSecret"),
			Token:  os.Getenv("ClientToken"),
			//Callback: "http://localhost:8008/v1/auth/callback"
			Callback: "https://dscinflux.xyz/v1/auth/callback",
		},
		Collection: "entities",
		//APIUrl: "http://localhost:8080/v1/"
		APIUrl: "https://api.dscinflux.xyz",
	}
}

func GetConfig() types.Config {
	return getConfig()
}

package main

import (
	"fmt"

	"github.com/drlatt/go-config-demo/src/appsettings"
)

type AppSettings struct {
	Log struct {
		MinFilter string `envconfig:"optional"`
	}
	Cors struct {
		Origins []string `envconfig:"optional"`
	}
	SomeService struct {
		URL string `json:"Url" envconfig:"optional"`
	}
	SomePublisher struct {
		Env     string `envconfig:"optional"`
		Project string `envconfig:"optional"`
		Topic   string `envconfig:"optional"`
	}
	Google struct {
		Application struct {
			Credentials string `envconfig:"optional"`
		}
	}
	SomeXML struct {
		Storage struct {
			BucketID string `json:"BucketId" envconfig:"optional"`
		}
		AuthBasicToken string `envconfig:"optional"`
	} `json:"SomeXml"`
	DefaultProvider string `envconfig:"optional"`
}

func main() {
	// placeholder variable
	settings := AppSettings{}

	if err := appsettings.ReadFromFileAndEnv(&settings); err != nil {
		panic(err)
	}

	// do something with the settings
	fmt.Printf("%+v", settings)
}

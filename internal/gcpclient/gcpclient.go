package gcpclient

import (
	"context"
	"log"

	"cloud.google.com/go/storage"
	"github.com/spf13/viper"
)

func Client(config *viper.Viper) (*storage.BucketHandle, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Println("ERR::0X234")
		return nil, err
	}

	log.Println("Bucket Name: " + config.GetString("app.bucket.name"))

	return client.Bucket(config.GetString("app.bucket.name")), nil
}

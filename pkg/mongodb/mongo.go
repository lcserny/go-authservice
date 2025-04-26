package mongodb

import (
	"context"

	"github.com/lcserny/go-authservice/src/auth"
	"github.com/lcserny/go-authservice/src/config"
	"github.com/lcserny/go-authservice/src/db"
	"github.com/lcserny/go-authservice/src/logging"
	"github.com/lcserny/go-authservice/src/users"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func connect(url string) *mongo.Client {
	client, err := mongo.Connect(options.Client().ApplyURI(url))
	if err != nil {
		logging.Fatal(err.Error())
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		logging.Fatal(err.Error())
	}
	
	logging.Info("MongoDB connected")

	return client
}

func NewMongoRepositoryProvider(cfg *config.Config) db.RepositoryProvider {
	client := connect(cfg.Database.Url)
	return &mongoRepositoryProvider{
		client: client,
	}
}

type mongoRepositoryProvider struct {
	client *mongo.Client
}

func (rp *mongoRepositoryProvider) GetUserRepository() users.UserRepository {
	return nil
}

func (rp *mongoRepositoryProvider) GetAuthRepository() auth.AuthRepository {
	return nil
}

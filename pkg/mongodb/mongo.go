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

const (
	usersCollection = "users"
	authCollection = "auth"
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
		cfg: cfg,
		client: client,
	}
}

type mongoRepositoryProvider struct {
	cfg *config.Config
	client *mongo.Client
}

func (rp *mongoRepositoryProvider) GetUserRepository() users.UserRepository {
	collection := rp.client.Database(rp.cfg.Database.Database).Collection(usersCollection)
	return &mongoUserRepository{
		collection,
	}
}

func (rp *mongoRepositoryProvider) GetAuthRepository() auth.AuthRepository {
	collection := rp.client.Database(rp.cfg.Database.Database).Collection(authCollection)
	return &mongoAuthRepository{
		collection,
	}
}

type mongoUserRepository struct {
	collection *mongo.Collection
}

func (ur *mongoUserRepository) GetUserByID(ctx context.Context, id string) (*users.User, error) {
	var user users.User
	err := ur.collection.FindOne(ctx, map[string]any{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil	
}

type mongoAuthRepository struct {
	collection *mongo.Collection
}

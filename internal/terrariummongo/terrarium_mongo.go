package terrariummongo

import (
	"context"
	"errors"
	"fmt"

	"github.com/dylanrhysscott/terrarium/internal/terrariummongo/orgs"
	"github.com/dylanrhysscott/terrarium/internal/terrariummongo/vcs"
	"github.com/dylanrhysscott/terrarium/internal/terrariumvcs/github"

	"github.com/dylanrhysscott/terrarium/pkg/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TerrariumMongo implements Mongo support for Terrarium for all API's
type TerrariumMongo struct {
	Host     string
	User     string
	Password string
	Database string
	client   *mongo.Client
}

// Connect iniitialises a database connection to mongo
func (m *TerrariumMongo) Connect(ctx context.Context) error {
	connectionStr := fmt.Sprintf("mongodb://%s", m.Host)
	if m.User != "" && m.Password != "" {
		connectionStr = fmt.Sprintf("mongodb://%s:%s@%s", m.User, m.Password, m.Host)
	}
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionStr))
	if err != nil {
		return err
	}
	m.client = client
	return nil
}

// Organizations returns a Mongo compatible organization store which implements the OrganizationStore interface
func (m *TerrariumMongo) Organizations() types.OrganizationStore {
	return &orgs.OrganizationBackend{
		CollectionName: "organizations",
		Client:         m.client,
		Database:       m.Database,
	}
}

// VCS returns a Mongo compatible VCS store which implements the VCSStore interface
func (m *TerrariumMongo) VCS() types.VCSStore {
	return &vcs.VCSBackend{
		CollectionName: "vcs",
		Client:         m.client,
		Database:       m.Database,
	}
}

// Sources returns a Sources struct which implements the SourceStore interface
func (m *TerrariumMongo) GithubSources() types.SourceStore {
	return &github.GithubBackend{}
}

// New creates a TerrariumMongo driver
func New(host string, user string, password string, database string) (*TerrariumMongo, error) {
	if host == "" {
		return nil, errors.New("mongo host cannot be empty")
	}
	driver := &TerrariumMongo{
		Host:     host,
		User:     user,
		Password: password,
		Database: database,
	}
	err := driver.Connect(context.TODO())
	if err != nil {
		return nil, err
	}
	err = driver.Organizations().Init()
	if err != nil {
		return nil, err
	}
	return driver, nil
}

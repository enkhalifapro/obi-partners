package mongo

import (
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// Connector wraps mongodb data access functionalities
type Connector struct {
	client *mongo.Client
	dbName string
}

// NewConnector initiates a new mongodb connector
func NewConnector(dbURI string, dbName string) (*Connector, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(dbURI))
	if err != nil {
		return &Connector{}, errors.Wrapf(err, "conneting to %s error", dbURI)
	}

	return &Connector{
		client: client,
		dbName: dbName,
	}, nil
}

// CreateOne document, returns inserted id and error if any
func (c *Connector) CreateOne(collectionName string, document interface{}) (string, error) {
	coll := c.client.Database(c.dbName).Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := coll.InsertOne(ctx, document)
	if err != nil {
		return "", errors.Wrap(err, "inserting one document error")
	}
	return res.InsertedID.(string), nil
}

// CreateMany documents
func (c *Connector) CreateMany(collectionName string, documents []interface{}) error {
	coll := c.client.Database(c.dbName).Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_, err := coll.InsertMany(ctx, documents)
	if err != nil {
		return errors.Wrap(err, "inserting many documents error")
	}
	return nil
}

// FindOne document
func (c *Connector) FindOne(collectionName string, filter interface{}, res interface{}) error {
	coll := c.client.Database(c.dbName).Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := coll.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		return errors.Wrap(err, "finding one document error")
	}
	return nil
}

// FindMany documents
func (c *Connector) FindMany(collectionName string, filter interface{}, res []interface{}) error {
	coll := c.client.Database(c.dbName).Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := coll.Find(ctx, filter)
	if err != nil {
		return errors.Wrap(err, "finding documents error")
	}
	if err = cur.Decode(res); err != nil {
		return errors.Wrap(err, "decoding found documents error")
	}
	return nil
}

// UpdateOne document
func (c *Connector) UpdateOne(collectionName string, filter interface{}, update interface{}) error {
	coll := c.client.Database(c.dbName).Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.Wrap(err, "updating documents error")
	}
	return nil
}

// UpdateMany documents
func (c *Connector) UpdateMany(collectionName string, filter interface{}, update interface{}) error {
	coll := c.client.Database(c.dbName).Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_, err := coll.UpdateMany(ctx, filter, update)
	if err != nil {
		return errors.Wrap(err, "updating documents error")
	}
	return nil
}

// DeleteOne document
func (c *Connector) DeleteOne(collectionName string, filter interface{}) error {
	coll := c.client.Database(c.dbName).Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := coll.DeleteOne(ctx, filter)
	if err != nil {
		return errors.Wrap(err, "deleting document error")
	}
	return nil
}

// DeleteMany documents
func (c *Connector) DeleteMany(collectionName string, filter interface{}) error {
	coll := c.client.Database(c.dbName).Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_, err := coll.DeleteMany(ctx, filter)
	if err != nil {
		return errors.Wrap(err, "deleting documents error")
	}
	return nil
}


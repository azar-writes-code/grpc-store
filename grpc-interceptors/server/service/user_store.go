package service

import (
	"context"
	"fmt"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"

	"ingens/configs"
)

type UserStore interface {
	Save(user *User) error

	Find(username string) (*User, error)
}

type InMemoryUserStore struct {
	mutex  sync.RWMutex
	users  map[string]*User
	client *mongo.Client
}

func NewInMemoryUserStore(client *mongo.Client) *InMemoryUserStore {
	return &InMemoryUserStore{
		users:  make(map[string]*User),
		client: client,
	}
}

func (store *InMemoryUserStore) getCollName() *mongo.Collection {
	coll := configs.GetCollection(store.client, "users")
	return coll
}

func (store *InMemoryUserStore) Save(user *User) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.users[user.Username] != nil {
		return fmt.Errorf("username exists")
	}
	result, err := store.getCollName().InsertOne(context.Background(), user)
	if err != nil {
		return err
	}
	store.users[user.Username] = user.Clone()
	fmt.Println("--> user: ", result)
	return nil
}

func (store *InMemoryUserStore) Find(username string) (*User, error) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	user := store.users[username]

	if user == nil {
		return nil, nil
	}

	return user.Clone(), nil
}

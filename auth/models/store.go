package models

import (
	"errors"
	"sync"
	"time"
)

// UserStore provides in-memory storage for users
// In production, this would be replaced with a database
type UserStore struct {
	users  map[int]*User
	emails map[string]*User
	usernames map[string]*User
	nextID int
	mutex  sync.RWMutex
}

// NewUserStore creates a new user store
func NewUserStore() *UserStore {
	return &UserStore{
		users:     make(map[int]*User),
		emails:    make(map[string]*User),
		usernames: make(map[string]*User),
		nextID:    1,
	}
}

// CreateUser creates a new user
func (s *UserStore) CreateUser(username, email, hashedPassword string) (*User, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Check if username already exists
	if _, exists := s.usernames[username]; exists {
		return nil, errors.New("username already exists")
	}

	// Check if email already exists
	if _, exists := s.emails[email]; exists {
		return nil, errors.New("email already exists")
	}

	user := &User{
		ID:        s.nextID,
		Username:  username,
		Email:     email,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	s.users[user.ID] = user
	s.emails[email] = user
	s.usernames[username] = user
	s.nextID++

	return user, nil
}

// GetUserByUsername retrieves a user by username
func (s *UserStore) GetUserByUsername(username string) (*User, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	user, exists := s.usernames[username]
	if !exists {
		return nil, errors.New("user not found")
	}

	return user, nil
}

// GetUserByID retrieves a user by ID
func (s *UserStore) GetUserByID(id int) (*User, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	user, exists := s.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}

	return user, nil
}

// GetAllUsers returns all users (for admin purposes)
func (s *UserStore) GetAllUsers() []*User {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	users := make([]*User, 0, len(s.users))
	for _, user := range s.users {
		users = append(users, user)
	}

	return users
}
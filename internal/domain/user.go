package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email     string             `json:"email" bson:"email"`
	Username  string             `json:"username" bson:"username"`
	Password  string             `json:"-" bson:"password"`
	FirstName string             `json:"first_name" bson:"first_name"`
	LastName  string             `json:"last_name" bson:"last_name"`
	IsActive  bool               `json:"is_active" bson:"is_active"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

type CreateUserRequest struct {
	Email     string `json:"email" binding:"required,email"`
	Username  string `json:"username" binding:"required,min=3,max=50"`
	Password  string `json:"password" binding:"required,min=6"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
}

type UpdateUserRequest struct {
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
	IsActive  *bool   `json:"is_active,omitempty"`
}

type UserResponse struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRepository interface {
	Create(user *User) error
	GetByID(id string) (*User, error)
	GetByEmail(email string) (*User, error)
	GetByUsername(username string) (*User, error)
	Update(id string, user *User) error
	Delete(id string) error
	List(offset, limit int) ([]*User, int64, error)
}

type UserService interface {
	CreateUser(req *CreateUserRequest) (*UserResponse, error)
	GetUser(id string) (*UserResponse, error)
	UpdateUser(id string, req *UpdateUserRequest) (*UserResponse, error)
	DeleteUser(id string) error
	ListUsers(page, limit int) ([]*UserResponse, int64, error)
}
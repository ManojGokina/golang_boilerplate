package service

import (
	"errors"

	"backend/internal/domain"
	"backend/pkg/logger"

	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepo domain.UserRepository
	logger   logger.Logger
}

func NewUserService(userRepo domain.UserRepository, logger logger.Logger) domain.UserService {
	return &userService{
		userRepo: userRepo,
		logger:   logger,
	}
}

func (s *userService) CreateUser(req *domain.CreateUserRequest) (*domain.UserResponse, error) {
	// Check if user already exists
	if _, err := s.userRepo.GetByEmail(req.Email); err == nil {
		return nil, errors.New("user with this email already exists")
	}

	if _, err := s.userRepo.GetByUsername(req.Username); err == nil {
		return nil, errors.New("user with this username already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		s.logger.Error("Failed to hash password: ", err)
		return nil, errors.New("failed to process password")
	}

	user := &domain.User{
		Email:     req.Email,
		Username:  req.Username,
		Password:  string(hashedPassword),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		IsActive:  true,
	}

	if err := s.userRepo.Create(user); err != nil {
		s.logger.Error("Failed to create user: ", err)
		return nil, errors.New("failed to create user")
	}

	return s.toUserResponse(user), nil
}

func (s *userService) GetUser(id string) (*domain.UserResponse, error) {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("user not found")
		}
		s.logger.Error("Failed to get user: ", err)
		return nil, errors.New("failed to get user")
	}

	return s.toUserResponse(user), nil
}

func (s *userService) UpdateUser(id string, req *domain.UpdateUserRequest) (*domain.UserResponse, error) {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("user not found")
		}
		s.logger.Error("Failed to get user: ", err)
		return nil, errors.New("failed to get user")
	}

	// Update fields if provided
	if req.FirstName != nil {
		user.FirstName = *req.FirstName
	}
	if req.LastName != nil {
		user.LastName = *req.LastName
	}
	if req.IsActive != nil {
		user.IsActive = *req.IsActive
	}

	if err := s.userRepo.Update(id, user); err != nil {
		s.logger.Error("Failed to update user: ", err)
		return nil, errors.New("failed to update user")
	}

	return s.toUserResponse(user), nil
}

func (s *userService) DeleteUser(id string) error {
	if _, err := s.userRepo.GetByID(id); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return errors.New("user not found")
		}
		s.logger.Error("Failed to get user: ", err)
		return errors.New("failed to get user")
	}

	if err := s.userRepo.Delete(id); err != nil {
		s.logger.Error("Failed to delete user: ", err)
		return errors.New("failed to delete user")
	}

	return nil
}

func (s *userService) ListUsers(page, limit int) ([]*domain.UserResponse, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	offset := (page - 1) * limit
	users, total, err := s.userRepo.List(offset, limit)
	if err != nil {
		s.logger.Error("Failed to list users: ", err)
		return nil, 0, errors.New("failed to list users")
	}

	responses := make([]*domain.UserResponse, len(users))
	for i, user := range users {
		responses[i] = s.toUserResponse(user)
	}

	return responses, total, nil
}

func (s *userService) toUserResponse(user *domain.User) *domain.UserResponse {
	return &domain.UserResponse{
		ID:        user.ID.Hex(),
		Email:     user.Email,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		IsActive:  user.IsActive,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
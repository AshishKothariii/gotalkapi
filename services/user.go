package services

import (
	"context"

	"github.com/AshishKothariii/gotalkapi/models"
	"github.com/AshishKothariii/gotalkapi/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// UserService handles the business logic for user operations
type UserService interface {
    GetAllUsers(ctx context.Context) ([]*models.User, error)
    GetUserByID(ctx context.Context, id primitive.ObjectID) (*models.User, error)
    CreateUser(ctx context.Context, username, email, password string) (*models.User, error)
    CheckUserCredentials(ctx context.Context, username, password string) (*models.User, error)
    GetUserByUserName(ctx context.Context, username string) (*models.User, error)

    
}

type userService struct {
    repo repository.UserRepository
}

// NewUserService creates a new instance of UserService
func NewUserService(repo repository.UserRepository) UserService {
    return &userService{repo}
}

func (s *userService) GetAllUsers(ctx context.Context) ([]*models.User, error) {
    return s.repo.GetAllUsers(ctx)
}

func (s *userService) GetUserByID(ctx context.Context, id primitive.ObjectID) (*models.User, error) {
    return s.repo.GetUserByID(ctx, id)
}

func (s *userService) CreateUser(ctx context.Context, username, email, password string) (*models.User, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }
    user := &models.User{
        Username: username,
        Email:    email,
        Password: string(hashedPassword),
    }
    err = s.repo.CreateUser(ctx, user)
    return user, err
}

func (s *userService) CheckUserCredentials(ctx context.Context, username, password string) (*models.User, error) {
    user, err := s.repo.GetUserByUsername(ctx, username)
    if err != nil {
        return nil, err
    }
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
        return nil, err
    }
    return user, nil
}
func (s *userService) GetUserByUserName(ctx context.Context, username string) (*models.User, error) {
    return s.repo.GetUserByUsername(ctx, username)
}
package service

import (
	"context"
	"fmt"
	"github.com/Den4ik117/examly/config"
	"github.com/Den4ik117/examly/internal/model"
	"github.com/Den4ik117/examly/internal/repository"
	"github.com/Den4ik117/examly/internal/util"
	"github.com/golang-jwt/jwt/v5"
	"github.com/guregu/null/v5"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthService struct {
	repo repository.Auth
}

func NewAuthService(repo repository.Auth) *AuthService {
	return &AuthService{repo: repo}
}

type RegisterInput struct {
	FirstName            string `json:"first_name" validate:"required,max=255"`
	LastName             string `json:"last_name" validate:"required,max=255"`
	Email                string `json:"email" validate:"required,email,max=255"`
	Password             string `json:"password" validate:"required,eqfield=PasswordConfirmation"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required"`
}

type LoginInput struct {
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required"`
}

type UpdateUserInput struct {
	FirstName string `json:"first_name" validate:"required,max=255"`
	LastName  string `json:"last_name" validate:"required,max=255"`
	Email     string `json:"email" validate:"required,email,max=255"`
}

type Claims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

func (s *AuthService) CreateUser(u RegisterInput) (int, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	exists, err := s.repo.IsExistsUserByEmail(u.Email)
	if err != nil {
		return 0, nil
	}
	if exists {
		return 0, fmt.Errorf("user with email %s already exists", u.Email)
	}

	user := model.User{
		UUID:      util.GenerateUUID(),
		Email:     null.StringFrom(u.Email),
		FirstName: null.StringFrom(u.FirstName),
		LastName:  null.StringFrom(u.LastName),
		Role:      model.MemberRole,
		Password:  null.StringFrom(string(bytes)),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return s.repo.CreateUser(user)
}

func (s *AuthService) Login(input LoginInput) (string, error) {
	u, err := s.repo.GetUserByEmail(input.Email)
	if err != nil {
		return "", err
	}
	if u.ID == 0 {
		return "", fmt.Errorf("user with email %s not found", input.Email)
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password.String), []byte(input.Password))
	if err != nil {
		return "", fmt.Errorf("invalid password")
	}

	return generateTokenByUserID(u.ID)
}

func (s *AuthService) GetGuestToken() (string, error) {
	user := model.User{
		Role:      model.GuestRole,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	ctx := context.Background()
	context.WithValue(ctx, "fdsf", "fdsf")
	context.with

	id, err := s.repo.CreateUser(user)
	if err != nil {
		return "", err
	}

	return generateTokenByUserID(id)
}

func (s *AuthService) CheckAuth(t string) (*model.User, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(t, claims, func(token *jwt.Token) (any, error) {
		return []byte(config.Envs.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	user, err := s.repo.GetUserByID(claims.UserID)
	if err != nil {
		return nil, err
	}
	if user.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return &user, nil
}

func (s *AuthService) UpdateUser(user model.User, u *UpdateUserInput) (*model.User, error) {
	user.FirstName = null.StringFrom(u.FirstName)
	user.LastName = null.StringFrom(u.LastName)
	user.Email = null.StringFrom(u.Email)
	user.UpdatedAt = time.Now()

	err := s.repo.UpdateUser(&user)
	if err != nil {
		return nil, err
	}

	updated, err := s.repo.GetUserByID(user.ID)

	return &updated, err
}

func generateTokenByUserID(id int) (string, error) {
	claims := &Claims{
		UserID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 365)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Envs.SecretKey))
}

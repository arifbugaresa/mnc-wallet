package user

import (
	"encoding/json"
	"errors"
	"github.com/arifbugaresa/mnc-wallet/middlewares"
	"github.com/arifbugaresa/mnc-wallet/utils/common"
	"github.com/arifbugaresa/mnc-wallet/utils/response"
	"github.com/arifbugaresa/mnc-wallet/utils/session"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"strings"
	"time"
)

type Service interface {
	SignUp(ctx *gin.Context, dataBody SignUpRequest) (response SignUpResponse, err error)
	Login(ctx *gin.Context, dataBody LoginRequest) (response LoginResponse, err error)
	Logout(ctx *gin.Context) (err error)
	GetMyProfile(ctx *gin.Context) (response GetProfileResponse, err error)
	UpdateMyProfile(ctx *gin.Context, dataBody UpdateProfileRequest) (err error)
}

type UserService struct {
	repo  *UserRepository
	redis *redis.Client
}

func NewService(repo *UserRepository, redis *redis.Client) *UserService {
	return &UserService{
		repo:  repo,
		redis: redis,
	}
}

func (s *UserService) SignUp(ctx *gin.Context, dataBody SignUpRequest) (response SignUpResponse, err error) {
	signUpModel := SignUpModel{
		FirstName:    dataBody.FirstName,
		LastName:     dataBody.LastName,
		Pin:          dataBody.Pin,
		PhoneNumber:  dataBody.PhoneNumber,
		Address:      dataBody.Address,
		DefaultTable: common.DefaultTable{}.GetDefaultTableWithoutToken(ctx),
	}

	err = s.repo.SignUp(ctx, signUpModel)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			err = errors.New("phone number already registered")
		}
		return
	}

	user, err := s.repo.GetUserByPhone(ctx, LoginModel{PhoneNumber: dataBody.PhoneNumber})
	if err != nil {
		return
	}

	response = SignUpResponse{
		UserId:      user.UserId,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		PhoneNumber: user.PhoneNumber,
		Address:     user.Address,
	}

	return
}

func (s *UserService) Login(ctx *gin.Context, dataBody LoginRequest) (response LoginResponse, err error) {
	user, err := s.CheckUser(ctx, dataBody)
	if err != nil {
		return
	}

	// generate token
	claims := middlewares.Claims{}

	jwtToken, err := claims.GenerateJwtToken()
	if err != nil {
		return
	}

	err = s.SetRedisSession(ctx, user, jwtToken)
	if err != nil {
		return
	}

	response.Token = jwtToken

	return
}

func (s UserService) CheckUser(ctx *gin.Context, dataBody LoginRequest) (user LoginModel, err error) {
	user, err = s.repo.GetUserByPhone(ctx, LoginModel{PhoneNumber: dataBody.PhoneNumber})
	if err != nil {
		return
	}

	if user.PhoneNumber == "" {
		err = errors.New("user not found")
		return
	}

	// check pin
	if user.Pin != dataBody.Pin {
		err = errors.New("“Phone Number and PIN doesn’t match")
	}

	return
}

func (s *UserService) SetRedisSession(ctx *gin.Context, user LoginModel, jwtToken string) (err error) {
	redisSession := session.RedisData{
		UserId:      user.UserId,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		PhoneNumber: user.PhoneNumber,
		Address:     user.Address,
	}

	jsonBytes, err := json.Marshal(redisSession)
	if err != nil {
		return
	}

	s.redis.Set(ctx, jwtToken, string(jsonBytes), 1*time.Hour)

	return
}

func (s UserService) Logout(ctx *gin.Context) (err error) {
	token, err := middlewares.GetJwtTokenFromHeader(ctx)
	if err != nil {
		return
	}
	s.redis.Del(ctx, token)
	return
}

func (s UserService) GetMyProfile(ctx *gin.Context) (response GetProfileResponse, err error) {
	auth, err := middlewares.GetSession(ctx)
	if err != nil {
		return
	}

	authUser, err := s.repo.GetUserByPhone(ctx, LoginModel{PhoneNumber: auth.PhoneNumber})
	if err != nil {
		return
	}

	response = GetProfileResponse{
		UserId:      authUser.UserId,
		FirstName:   authUser.FirstName,
		LastName:    authUser.LastName,
		PhoneNumber: authUser.PhoneNumber,
		Address:     authUser.Address,
	}

	return
}

func (s UserService) UpdateMyProfile(ctx *gin.Context, dataBody UpdateProfileRequest) (err error) {
	auth, err := middlewares.GetSession(ctx)
	if err != nil {
		response.GenerateErrorResponse(ctx, err.Error())
	}

	data := UpdateProfileModel{
		UserId:       auth.UserId,
		FirstName:    dataBody.FirstName,
		LastName:     dataBody.LastName,
		Address:      dataBody.Address,
		DefaultTable: common.DefaultTable{}.GetDefaultTable(ctx),
	}

	err = s.repo.UpdateProfileById(ctx, data)
	if err != nil {
		return
	}

	return
}

package user

import (
	"errors"
	"github.com/arifbugaresa/mnc-wallet/utils/common"
)

type (
	SignUpRequest struct {
		FirstName   string `json:"first_name"`
		LastName    string `json:"last_name"`
		PhoneNumber string `json:"phone_number"`
		Address     string `json:"address"`
		Pin         int64  `json:"pin"`
	}

	SignUpModel struct {
		FirstName   string `db:"first_name"`
		LastName    string `db:"last_name"`
		PhoneNumber string `db:"phone_number"`
		Address     string `db:"address"`
		Pin         int64  `db:"pin"`
		common.DefaultTable
	}

	SignUpResponse struct {
		UserId      string `json:"user_id"`
		FirstName   string `json:"first_name"`
		LastName    string `json:"last_name"`
		PhoneNumber string `json:"phone_number"`
		Address     string `json:"address"`
		CreatedDate string `json:"created_date"`
	}
)

func (r SignUpRequest) ValidateSignUpRequest() error {
	if r.FirstName == "" {
		return errors.New("firstname is required")
	}

	if r.PhoneNumber == "" {
		return errors.New("phone number is required")
	}

	if r.Pin == 0 {
		return errors.New("pin is required")
	}

	return nil
}

type (
	LoginRequest struct {
		PhoneNumber string `json:"phone_number"`
		Pin         int64  `json:"pin"`
	}

	LoginResponse struct {
		Token string `json:"access_token"`
	}

	LoginModel struct {
		PhoneNumber string `db:"phone_number"`
		UserId      string `db:"user_id"`
		FirstName   string `db:"first_name"`
		LastName    string `db:"last_name"`
		Address     string `db:"address"`
		Pin         int64  `db:"pin"`
	}
)

func (r LoginRequest) ValidateLoginRequest() error {
	if r.PhoneNumber == "" {
		return errors.New("phone number is required")
	}

	if r.Pin == 0 {
		return errors.New("pin is required")
	}

	return nil
}

type (
	LogoutRequest struct {
		Token string `json:"access_token"`
	}
)

func (r LogoutRequest) ValidateLogoutRequest() (err error) {
	if r.Token == "" {
		return errors.New("token is required")
	}

	return
}

type (
	GetProfileResponse struct {
		UserId      string `json:"user_id"`
		FirstName   string `json:"first_name"`
		LastName    string `json:"last_name"`
		PhoneNumber string `json:"phone_number"`
		Address     string `json:"address"`
	}

	UpdateProfileRequest struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Address   string `json:"address"`
	}

	UpdateProfileModel struct {
		UserId    string `db:"user_id" goqu:"skipupdate"`
		FirstName string `db:"first_name"`
		LastName  string `db:"last_name"`
		Address   string `db:"address"`
		common.DefaultTable
	}
)

func (r UpdateProfileRequest) ValidateUpdateProfileRequest() (err error) {
	if r.FirstName == "" {
		return errors.New("first name is required")
	}

	return
}

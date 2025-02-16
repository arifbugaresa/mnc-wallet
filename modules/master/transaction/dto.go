package transaction

import (
	"errors"
	"github.com/arifbugaresa/mnc-wallet/utils/common"
)

type (
	TopUpRequest struct {
		Amount float64 `json:"amount"`
	}

	TopUpModel struct {
		UserId        string  `json:"user_id"`
		TopUpId       string  `db:"top_up_id"`
		Amount        float64 `db:"amount"`
		BalanceBefore float64 `db:"balance_before"`
		BalanceAfter  float64 `db:"balance_after"`
		common.DefaultTable
	}

	TopUpResponse struct {
		TopUpId       string  `json:"top_up_id"`
		Amount        float64 `json:"amount_top_up"`
		BalanceBefore float64 `json:"balance_before"`
		BalanceAfter  float64 `json:"balance_after"`
		CreatedDate   string  `json:"created_date"`
	}
)

func (r TopUpRequest) ValidateTopUpRequest() error {
	if r.Amount == 0 {
		return errors.New("amount is required")
	}

	return nil
}

type (
	TransferRequest struct {
		TargetUser string  `json:"target_user"`
		Amount     float64 `json:"amount"`
		Remarks    string  `json:"remark"`
	}

	TransferModel struct {
		UserId        string  `db:"user_id"`
		UserIdTarget  string  `db:"user_id_target"`
		Amount        float64 `db:"amount"`
		Remarks       string  `db:"remarks"`
		BalanceBefore float64 `db:"balance_before"`
		BalanceAfter  float64 `db:"balance_after"`
		common.DefaultTable
	}

	TransferResponse struct {
		TransferId    string  `json:"top_up_id"`
		Amount        float64 `json:"amount_top_up"`
		Remarks       string  `json:"remark"`
		BalanceBefore float64 `json:"balance_before"`
		BalanceAfter  float64 `json:"balance_after"`
		CreatedDate   string  `json:"created_date"`
	}
)

func (r TransferRequest) ValidateTransferRequest() error {
	if r.Amount == 0 {
		return errors.New("amount is required")
	}

	if r.Remarks == "" {
		return errors.New("remarks is required")
	}

	if r.TargetUser == "" {
		return errors.New("target user is required")
	}

	return nil
}

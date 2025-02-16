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

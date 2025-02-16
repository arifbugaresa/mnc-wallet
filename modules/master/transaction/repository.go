package transaction

import (
	"github.com/arifbugaresa/mnc-wallet/utils/common"
	"github.com/arifbugaresa/mnc-wallet/utils/constant/table"
	"github.com/doug-martin/goqu/v9"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	TopUp(ctx *gin.Context, model TopUpModel) (output TopUpModel, err error)
}

type UserRepository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) TopUp(ctx *gin.Context, model TopUpModel) (output TopUpModel, err error) {
	tx, err := r.db.BeginTxx(ctx, nil)
	defer tx.Rollback()

	dialect := goqu.NewTx("postgres", tx)

	// select for update
	sql, args, err := dialect.
		From(table.TM_USER_ACCOUNTS).
		Select(
			goqu.COALESCE(goqu.I("balance"), 0.00),
		).
		Where(
			goqu.Ex{"user_id": model.UserId},
		).
		ForUpdate(goqu.Wait).
		ToSQL()
	if err != nil {
		return
	}

	var balance float64
	err = tx.QueryRow(sql, args...).Scan(&balance)
	if err != nil {
		return
	}

	var balanceAfterTopUp = balance + model.Amount

	// update balance
	sql, args, err = dialect.Update(table.TM_USER_ACCOUNTS).
		Set(goqu.Record{"balance": balanceAfterTopUp}).
		Where(
			goqu.Ex{"user_id": model.UserId},
		).
		ToSQL()
	if err != nil {
		return
	}

	_, err = tx.Exec(sql, args...)
	if err != nil {
		return
	}

	// insert table detail top up
	sql, args, err = dialect.
		Insert(table.TR_USER_TOPUPS).
		Rows(
			goqu.Record{
				"amount":         model.Amount,
				"balance_before": balance,
				"balance_after":  balanceAfterTopUp,
				"created_at":     model.CreatedAt,
				"created_by":     model.UserId,
				"updated_at":     model.UpdatedAt,
				"updated_by":     model.UserId,
			},
		).
		Returning("top_up_id").
		ToSQL()
	if err != nil {
		return
	}

	var topUpId string
	err = tx.QueryRow(sql, args...).Scan(&topUpId)
	if err != nil {
		return
	}

	err = tx.Commit()
	if err != nil {
		return
	}

	output = TopUpModel{
		TopUpId:       topUpId,
		Amount:        model.Amount,
		BalanceBefore: balance,
		BalanceAfter:  balanceAfterTopUp,
		DefaultTable:  common.DefaultTable{}.GetDefaultTable(ctx),
	}

	return
}

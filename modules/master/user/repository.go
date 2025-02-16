package user

import (
	"github.com/arifbugaresa/mnc-wallet/utils/constant/table"
	"github.com/doug-martin/goqu/v9"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	SignUp(ctx *gin.Context, model SignUpModel) (err error)
	GetUserByPhone(ctx *gin.Context, req LoginModel) (record LoginModel, err error)
	UpdateProfileById(ctx *gin.Context, model UpdateProfileModel) (err error)
}

type UserRepository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) SignUp(ctx *gin.Context, model SignUpModel) (err error) {
	dialect := goqu.New("postgres", r.db)
	dataset := dialect.
		Insert(table.TM_USER_ACCOUNTS).
		Rows(model)

	_, err = dataset.Executor().ExecContext(ctx)
	if err != nil {
		return err
	}

	return
}

func (r *UserRepository) GetUserByPhone(ctx *gin.Context, req LoginModel) (record LoginModel, err error) {
	dialect := goqu.New("postgres", r.db)
	dataset := dialect.
		Select(
			goqu.I("a.user_id"),
			goqu.I("a.phone_number"),
			goqu.I("a.first_name"),
			goqu.I("a.last_name"),
			goqu.I("a.address"),
			goqu.I("a.pin"),
		).
		From(
			goqu.T(table.TM_USER_ACCOUNTS).As("a"),
		).
		Where(
			goqu.I("a.phone_number").Eq(req.PhoneNumber),
		)

	_, err = dataset.ScanStructContext(ctx, &record)
	if err != nil {
		return
	}

	return
}

func (r *UserRepository) UpdateProfileById(ctx *gin.Context, model UpdateProfileModel) (err error) {
	dialect := goqu.New("postgres", r.db)
	dataset := dialect.
		Update(table.TM_USER_ACCOUNTS).
		Set(model).
		Where(goqu.I("user_id").Eq(model.UserId))

	_, err = dataset.Executor().ExecContext(ctx)
	if err != nil {
		return err
	}

	return
}

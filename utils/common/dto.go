package common

import (
	"github.com/arifbugaresa/mnc-wallet/middlewares"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type DefaultTable struct {
	ID        int64  `db:"id" goqu:"skipinsert,skipupdate"`
	CreatedBy string `db:"created_by" goqu:"skipupdate"`
	CreatedAt string `db:"created_at" goqu:"skipupdate"`
	UpdatedBy string `db:"updated_by"`
	UpdatedAt string `db:"updated_at"`
}

func (d DefaultTable) GetDefaultTable(ctx *gin.Context) DefaultTable {
	var (
		timeNow = time.Now().Format("2006-01-02")
	)

	auth, _ := middlewares.GetSession(ctx)
	return DefaultTable{
		CreatedBy: auth.UserId,
		UpdatedBy: auth.UserId,
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
	}
}

func (d DefaultTable) GetDefaultTableWithoutToken(ctx *gin.Context) DefaultTable {
	var (
		timeNow = time.Now().Format("2006-01-02")
	)

	return DefaultTable{
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
	}
}

type DefaultListRequest struct {
	Page   int64
	Limit  int64
	Search Search
	Sort   Sort
}

type Sort struct {
	Field string
	Order string
}

type Search struct {
	Field string
	Value string
}

func (d DefaultListRequest) GetParamRequest(ctx *gin.Context) DefaultListRequest {
	page, _ := strconv.ParseInt(ctx.Query("page"), 10, 64)
	limit, _ := strconv.ParseInt(ctx.Query("limit"), 10, 64)

	return DefaultListRequest{
		Page:  page,
		Limit: limit,
		Sort: Sort{
			Field: ctx.Query("sort_field"),
			Order: ctx.Query("sort_order"),
		},
		Search: Search{
			Field: ctx.Query("search_field"),
			Value: ctx.Query("search_value"),
		},
	}
}

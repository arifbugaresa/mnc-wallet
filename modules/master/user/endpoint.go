package user

import (
	"github.com/arifbugaresa/mnc-wallet/middlewares"
	"github.com/arifbugaresa/mnc-wallet/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

func Initiator(router *gin.Engine, dbConnection *sqlx.DB, redisConnection *redis.Client) {
	var (
		userRepo = NewRepository(dbConnection)
		userSrv  = NewService(userRepo, redisConnection)
	)

	publicAPI := router.Group("/api/users")
	{
		publicAPI.POST("/register", func(c *gin.Context) {
			SignUpEndpoint(c, userSrv)
		})
		publicAPI.POST("/login", func(c *gin.Context) {
			LoginEndpoint(c, userSrv)
		})
	}

	protectedApi := router.Group("/api/users")
	protectedApi.Use(middlewares.JwtMiddleware())
	{
		protectedApi.POST("/logout", func(c *gin.Context) {
			LogoutEndpoint(c, userSrv)
		})
		protectedApi.GET("/profile", func(c *gin.Context) {
			GetProfileEndpoint(c, userSrv)
		})
		protectedApi.PUT("/profile", func(c *gin.Context) {
			UpdateProfileEndpoint(c, userSrv)
		})
	}
}

func SignUpEndpoint(ctx *gin.Context, userSrv Service) {
	var (
		dataBody SignUpRequest
	)

	if err := ctx.ShouldBindJSON(&dataBody); err != nil {
		response.GenerateErrorResponse(ctx, err.Error())
		return
	}

	err := dataBody.ValidateSignUpRequest()
	if err != nil {
		response.GenerateErrorResponse(ctx, err.Error())
		return
	}

	result, err := userSrv.SignUp(ctx, dataBody)
	if err != nil {
		response.GenerateErrorResponse(ctx, err.Error())
		return
	}

	response.GenerateSuccessResponseWithData(ctx, "awesome, you are a member now!", result)
}

func LoginEndpoint(ctx *gin.Context, userSrv Service) {
	var (
		dataBody LoginRequest
	)

	if err := ctx.ShouldBindJSON(&dataBody); err != nil {
		response.GenerateErrorResponse(ctx, err.Error())
		return
	}

	err := dataBody.ValidateLoginRequest()
	if err != nil {
		response.GenerateErrorResponse(ctx, err.Error())
		return
	}

	result, err := userSrv.Login(ctx, dataBody)
	if err != nil {
		response.GenerateErrorResponse(ctx, err.Error())
		return
	}

	response.GenerateSuccessResponseWithData(ctx, "login successful", result)
}

func LogoutEndpoint(ctx *gin.Context, userSrv Service) {
	err := userSrv.Logout(ctx)
	if err != nil {
		response.GenerateErrorResponse(ctx, err.Error())
		return
	}

	response.GenerateSuccessResponse(ctx, "you have successfully logged out")
}

func GetProfileEndpoint(ctx *gin.Context, userSrv Service) {
	record, err := userSrv.GetMyProfile(ctx)
	if err != nil {
		response.GenerateErrorResponse(ctx, err.Error())
		return
	}

	response.GenerateSuccessResponseWithData(ctx, "you have successfully get your profile", record)
}

func UpdateProfileEndpoint(ctx *gin.Context, userSrv Service) {
	var (
		dataBody UpdateProfileRequest
	)

	if err := ctx.ShouldBindJSON(&dataBody); err != nil {
		response.GenerateErrorResponse(ctx, err.Error())
		return
	}

	err := dataBody.ValidateUpdateProfileRequest()
	if err != nil {
		response.GenerateErrorResponse(ctx, err.Error())
		return
	}

	err = userSrv.UpdateMyProfile(ctx, dataBody)
	if err != nil {
		response.GenerateErrorResponse(ctx, err.Error())
		return
	}

	response.GenerateSuccessResponse(ctx, "successfully update your profile")
}

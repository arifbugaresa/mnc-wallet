package upload

import (
	"github.com/arifbugaresa/mnc-wallet/middlewares"
	"github.com/arifbugaresa/mnc-wallet/utils/common"
	"github.com/arifbugaresa/mnc-wallet/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

func Initiator(router *gin.Engine, dbConnection *sqlx.DB) {
	var (
		uploadRepo = NewRepository(dbConnection)
		uploadSrv  = NewService(uploadRepo)
	)

	router.Static("/get-file", "./"+viper.GetString("storage.upload.file"))

	apiProtected := router.Group("/api/uploads")
	apiProtected.Use(middlewares.JwtMiddleware())
	{
		apiProtected.POST("", func(c *gin.Context) {
			UploadFileEndpoint(c, uploadSrv)
		})
		apiProtected.GET("", func(c *gin.Context) {
			GetFileEndpoint(c, uploadSrv)
		})
	}
}

func UploadFileEndpoint(ctx *gin.Context, uploadSrv Service) {
	file, err := ctx.FormFile("file")
	if err != nil {
		response.GenerateErrorResponse(ctx, err.Error())
		return
	}

	moduleName := ctx.PostForm("module")
	if moduleName == "" {
		response.GenerateErrorResponse(ctx, "module name is required")
		return
	}

	record, err := uploadSrv.UploadFile(ctx, UploadFileRequest{
		ModuleName: moduleName,
		File:       file,
	})
	if err != nil {
		response.GenerateErrorResponse(ctx, err.Error())
		return
	}

	response.GenerateSuccessResponseWithData(ctx, "successfully upload file", record)
}

func GetFileEndpoint(ctx *gin.Context, uploadSrv Service) {
	filePath := ctx.Query("file_path")
	if filePath == "" {
		response.GenerateErrorResponse(ctx, "file path is required")
		return
	}

	data := GetFileResponse{
		FilePath:   filePath,
		PreviewURL: common.GetPreviewURL(filePath),
	}

	response.GenerateSuccessResponseWithData(ctx, "successfully get file", data)
}

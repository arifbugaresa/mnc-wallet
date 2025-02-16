#!/bin/bash

# Periksa apakah direktori tujuan diberikan sebagai argumen
if [ -z "$1" ]; then
  echo "Penggunaan: $0 <project_directory>/<target_directory>"
  exit 1
fi

# Direktori tujuan dari argumen pertama
PROJECT_DIR=$(dirname "$1")
TARGET_DIR="$1"

# Ekstrak nama modul dari nama folder tujuan dan kapitalisasi huruf pertama
MODULE_NAME=$(basename "$TARGET_DIR")
MODULE_NAME_CAPITALIZED=$(echo "$MODULE_NAME" | awk '{print toupper(substr($0,1,1)) tolower(substr($0,2))}')
MODULE_NAME_LOWERCASE=$(echo "$MODULE_NAME" | awk '{print tolower($0)}')

# Buat direktori proyek jika belum ada
if [ ! -d "$PROJECT_DIR" ]; then
  echo "Membuat direktori proyek $PROJECT_DIR..."
  mkdir -p "$PROJECT_DIR"
fi

# Buat direktori tujuan jika belum ada
if [ ! -d "$TARGET_DIR" ]; then
  echo "Membuat direktori $TARGET_DIR..."
  mkdir -p "$TARGET_DIR"
fi

# Membuat file dto.go
DTO_FILE="$TARGET_DIR/dto.go"
DTO_CONTENT="package $MODULE_NAME_LOWERCASE

type ${MODULE_NAME_CAPITALIZED}DTO struct {
}
"
if [ ! -f "$DTO_FILE" ]; then
  echo "Membuat file $DTO_FILE..."
  echo -e "$DTO_CONTENT" > "$DTO_FILE"
  echo "File $DTO_FILE berhasil dibuat."
else
  echo "File $DTO_FILE sudah ada, melewatkan..."
fi

# Membuat file repository.go
REPOSITORY_FILE="$TARGET_DIR/repository.go"
REPOSITORY_CONTENT="package $MODULE_NAME_LOWERCASE

import (
	\"github.com/jmoiron/sqlx\"
)

type Repository interface {
	Get()
}

type ${MODULE_NAME_CAPITALIZED}Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *${MODULE_NAME_CAPITALIZED}Repository {
	return &${MODULE_NAME_CAPITALIZED}Repository{db: db}
}

func (r *${MODULE_NAME_CAPITALIZED}Repository) Get() {}
"
if [ ! -f "$REPOSITORY_FILE" ]; then
  echo "Membuat file $REPOSITORY_FILE..."
  echo -e "$REPOSITORY_CONTENT" > "$REPOSITORY_FILE"
  echo "File $REPOSITORY_FILE berhasil dibuat."
else
  echo "File $REPOSITORY_FILE sudah ada, melewatkan..."
fi

# Membuat file service.go
SERVICE_FILE="$TARGET_DIR/service.go"
SERVICE_CONTENT="package $MODULE_NAME_LOWERCASE

type Service interface {
	GetService() (interface{}, error)
}

type ${MODULE_NAME_CAPITALIZED}Service struct {
	repo *${MODULE_NAME_CAPITALIZED}Repository
}

func NewService(repo *${MODULE_NAME_CAPITALIZED}Repository) *${MODULE_NAME_CAPITALIZED}Service {
	return &${MODULE_NAME_CAPITALIZED}Service{repo: repo}
}

func (s *${MODULE_NAME_CAPITALIZED}Service) GetService() (interface{}, error) {
	// Implementasi logic untuk GetService
	return nil, nil
}
"
if [ ! -f "$SERVICE_FILE" ]; then
  echo "Membuat file $SERVICE_FILE..."
  echo -e "$SERVICE_CONTENT" > "$SERVICE_FILE"
  echo "File $SERVICE_FILE berhasil dibuat."
else
  echo "File $SERVICE_FILE sudah ada, melewatkan..."
fi

# Membuat file endpoint.go
ENDPOINT_FILE="$TARGET_DIR/endpoint.go"
ENDPOINT_CONTENT="package $MODULE_NAME_LOWERCASE

import (
	\"github.com/gin-gonic/gin\"
	\"github.com/jmoiron/sqlx\"
)

func Initiator(router *gin.Engine, dbConnection *sqlx.DB) {
	var (
		${MODULE_NAME_LOWERCASE}Repo = NewRepository(dbConnection)
		${MODULE_NAME_LOWERCASE}Srv  = NewService(${MODULE_NAME_LOWERCASE}Repo)
	)

	api := router.Group(\"/api/$MODULE_NAME_LOWERCASE\")
	{
		api.GET(\"/get\", func(c *gin.Context) {
			Get(c, ${MODULE_NAME_LOWERCASE}Srv)
		})
	}
}

func Get(ctx *gin.Context, ${MODULE_NAME_LOWERCASE}Srv Service) {
	record, err := ${MODULE_NAME_LOWERCASE}Srv.GetService()
	if err != nil {
		response.GenerateErrorResponse(ctx, err.Error())
		return
	}

	response.GenerateSuccessResponseWithData(ctx, \"successfully retrieved\", record)
}
"
if [ ! -f "$ENDPOINT_FILE" ]; then
  echo "Membuat file $ENDPOINT_FILE..."
  echo -e "$ENDPOINT_CONTENT" > "$ENDPOINT_FILE"
  echo "File $ENDPOINT_FILE berhasil dibuat."
else
  echo "File $ENDPOINT_FILE sudah ada, melewatkan..."
fi

echo "Semua file berhasil dibuat di $TARGET_DIR."

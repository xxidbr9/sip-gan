package sql

import (
	"github.com/xxidbr9/sip-gan/backend/app/utils/shared"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewSqlDB(uri string) (*gorm.DB, error) {
	gormDB, err := gorm.Open(postgres.Open(uri), &gorm.Config{})

	if err != nil {
		return &gorm.DB{}, shared.NewInfrastructureError(err)
	}

	return gormDB, nil
}

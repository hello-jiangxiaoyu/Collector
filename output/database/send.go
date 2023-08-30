package database

import "gorm.io/gorm"

type GormOutput struct {
	db *gorm.DB
}

func NewGormOutput(dbType string, dsn string) (*GormOutput, error) {
	db, err := NewGormDB(dbType, dsn)
	if err != nil {
		return nil, err
	}
	return &GormOutput{db: db}, nil
}

func (g *GormOutput) SendLog(logs map[string]any) error {
	return nil
}

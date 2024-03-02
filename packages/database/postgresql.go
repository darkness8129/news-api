package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var _ Database = (*postgreSQLDatabase)(nil)

type postgreSQLDatabase struct {
	db *gorm.DB
}

type Options struct {
	User     string
	Password string
	Database string
	Host     string
}

func NewPostgreSQLDatabase(opt Options) (*postgreSQLDatabase, error) {
	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s",
		opt.User, opt.Password, opt.Database, opt.Host,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to DB: %w", err)
	}

	// needed for automatic creating IDs for new records
	err = db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`).Error
	if err != nil {
		return nil, fmt.Errorf("failed to create uuid-ossp extension: %w", err)
	}

	return &postgreSQLDatabase{db}, nil
}

func (s *postgreSQLDatabase) DB() interface{} {
	return s.db
}

func (p *postgreSQLDatabase) Close() error {
	if p.DB != nil {
		db, err := p.db.DB()
		if err != nil {
			return fmt.Errorf("failed to get db: %w", err)
		}

		err = db.Close()
		if err != nil {
			return fmt.Errorf("failed to close postgresql connection: %w", err)
		}
	}

	return nil
}

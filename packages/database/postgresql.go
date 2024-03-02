package database

import (
	"darkness8129/news-api/packages/logging"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var _ Database = (*postgreSQLDatabase)(nil)

type postgreSQLDatabase struct {
	db     *gorm.DB
	logger logging.Logger
}

type Options struct {
	User     string
	Password string
	Database string
	Host     string
	Logger   logging.Logger
}

func NewPostgreSQLDatabase(opt Options) (*postgreSQLDatabase, error) {
	logger := opt.Logger.Named("PostgreSQLDatabase")

	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s",
		opt.User, opt.Password, opt.Database, opt.Host,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		logger.Error("failed to connect to DB", "err", err)
		return nil, fmt.Errorf("failed to connect to DB: %w", err)
	}

	// needed for automatic creating IDs for new records
	err = db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`).Error
	if err != nil {
		logger.Error("failed to create uuid-ossp extension", "err", err)
		return nil, fmt.Errorf("failed to create uuid-ossp extension: %w", err)
	}

	return &postgreSQLDatabase{
		db:     db,
		logger: logger,
	}, nil
}

func (p *postgreSQLDatabase) DB() interface{} {
	return p.db
}

func (p *postgreSQLDatabase) Close() error {
	logger := p.logger.Named("Close")

	if p.DB != nil {
		db, err := p.db.DB()
		if err != nil {
			logger.Error("failed to get db", "err", err)
			return fmt.Errorf("failed to get db: %w", err)
		}

		err = db.Close()
		if err != nil {
			logger.Error("failed to close postgresql connection", "err", err)
			return fmt.Errorf("failed to close postgresql connection: %w", err)
		}
	}

	logger.Info("successfully closed connection to DB")
	return nil
}

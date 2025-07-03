package initializa

import (
	"fmt"
	"janx-admin/global"
	"time"

	model "janx-admin/app/model"

	"github.com/fatih/color"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DatabaseBuilder struct {
	dsn        string
	logLevel   logger.LogLevel
	config     *gorm.Config
	poolConfig *ConnectionPoolConfig
	db         *gorm.DB
	err        error
}

type ConnectionPoolConfig struct {
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

// 日志级别映射表
var logLevelMap = map[string]logger.LogLevel{
	"silent": logger.Silent,
	"error":  logger.Error,
	"warn":   logger.Warn,
	"info":   logger.Info,
}

func InitPostgres() {
	// 使用Builder模式和链式调用重构
	builder := NewDatabaseBuilder().
		BuildDSN().
		ConfigureLogger().
		SetPoolConfig().
		Connect().
		TestConnection().
		Complete()

	// 统一错误处理
	if builder.err != nil {
		panic(fmt.Errorf("database initialization failed: %w", builder.err))
	}

	dbAutoMigrate(builder.db)

	global.Db = builder.db
	fmt.Println("PostgreSQL database connected successfully")
}

// NewDatabaseBuilder 创建数据库构建器
func NewDatabaseBuilder() *DatabaseBuilder {
	return &DatabaseBuilder{
		config: &gorm.Config{},
	}
}

// BuildDSN 构建数据源名称
func (b *DatabaseBuilder) BuildDSN() *DatabaseBuilder {
	if b.err != nil {
		return b
	}

	dsnTemplate := "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable"
	b.dsn = fmt.Sprintf(dsnTemplate,
		global.Conf.Db.Host,
		global.Conf.Db.Hostname, // 修改为 Username
		global.Conf.Db.Password,
		global.Conf.Db.Database,
		global.Conf.Db.Port,
	)
	return b
}

// ConfigureLogger 配置日志级别
func (b *DatabaseBuilder) ConfigureLogger() *DatabaseBuilder {
	if b.err != nil {
		return b
	}

	determineLogLevel := func() logger.LogLevel {
		if !global.Conf.Db.LogStatus {
			return logger.Silent
		}
		if level, exists := logLevelMap[global.Conf.Db.LogLevel]; exists {
			return level
		}
		return logger.Info // 默认级别
	}

	b.logLevel = determineLogLevel()
	b.config.Logger = logger.Default.LogMode(b.logLevel)
	return b
}

// SetPoolConfig 设置连接池配置
func (b *DatabaseBuilder) SetPoolConfig() *DatabaseBuilder {
	if b.err != nil {
		return b
	}

	b.poolConfig = &ConnectionPoolConfig{
		MaxOpenConns:    int(global.Conf.Db.MaxOpenConns),
		MaxIdleConns:    int(global.Conf.Db.MaxIdleConns),
		ConnMaxLifetime: time.Duration(global.Conf.Db.ConnMaxLifetime) * time.Second,
	}
	return b
}

// Connect 建立数据库连接
func (b *DatabaseBuilder) Connect() *DatabaseBuilder {
	if b.err != nil {
		return b
	}

	b.db, b.err = gorm.Open(postgres.Open(b.dsn), b.config) // 修改为 postgres
	return b
}

// TestConnection 测试连接并应用连接池配置
func (b *DatabaseBuilder) TestConnection() *DatabaseBuilder {
	if b.err != nil {
		return b
	}

	configurePool := func() error {
		sqlDB, err := b.db.DB()
		if err != nil {
			return fmt.Errorf("failed to get underlying sql.DB: %w", err)
		}

		// 批量应用连接池配置
		sqlDB.SetMaxOpenConns(b.poolConfig.MaxOpenConns)
		sqlDB.SetMaxIdleConns(b.poolConfig.MaxIdleConns)
		sqlDB.SetConnMaxLifetime(b.poolConfig.ConnMaxLifetime)

		return sqlDB.Ping()
	}

	if err := configurePool(); err != nil {
		b.err = fmt.Errorf("connection pool configuration failed: %w", err)
	}

	return b
}

// Complete 完成构建过程
func (b *DatabaseBuilder) Complete() *DatabaseBuilder {
	return b
}

func dbAutoMigrate(db *gorm.DB) {
	if err := db.AutoMigrate(
		&model.User{},
		&model.Role{},
	); err != nil {
		red := color.New(color.FgRed).Add(color.Bold) // 创建红色加粗的格式
		red.Printf("failed to migrate database:\n")   // 打印红色错误信息并退出
		panic(err.Error())
	}
}

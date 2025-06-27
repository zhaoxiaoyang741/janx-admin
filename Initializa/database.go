package initializa

import (
	"fmt"
	"janx-admin/global"
	"time"

	"gorm.io/driver/mysql"
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

func InitMySQL() {
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

	global.Db = builder.db
	fmt.Println("MySQL database connected successfully")
}

// NewDatabaseBuilder 创建数据库构建器
func NewDatabaseBuilder() *DatabaseBuilder {
	return &DatabaseBuilder{
		config: &gorm.Config{},
	}
}

// BuildDSN 构建数据源名称 - 使用函数式方法简化字符串构建
func (b *DatabaseBuilder) BuildDSN() *DatabaseBuilder {
	if b.err != nil {
		return b
	}

	dsnTemplate := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	b.dsn = fmt.Sprintf(dsnTemplate,
		global.Conf.Db.Hostname,
		global.Conf.Db.Password,
		global.Conf.Db.Host,
		global.Conf.Db.Port,
		global.Conf.Db.Database,
	)
	return b
}

// ConfigureLogger 配置日志级别 - 使用map替代switch，支持函数式配置
func (b *DatabaseBuilder) ConfigureLogger() *DatabaseBuilder {
	if b.err != nil {
		return b
	}

	// 使用函数式方法确定日志级别
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

// SetPoolConfig 设置连接池配置 - 封装配置逻辑
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

	b.db, b.err = gorm.Open(mysql.Open(b.dsn), b.config)
	return b
}

// TestConnection 测试连接并应用连接池配置 - 合并相关操作
func (b *DatabaseBuilder) TestConnection() *DatabaseBuilder {
	if b.err != nil {
		return b
	}

	// 使用闭包处理连接池配置，简化错误处理
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

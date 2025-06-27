package initializa

import (
	"fmt"
	"janx-admin/global"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

func InitLogger() {
	// 确保日志目录存在
	createLogDir()

	// 创建自定义formatter
	customFormatter := &CustomFormatter{
		TimestampFormat: "2006/01/02 15:04:05",
	}

	// 创建新的logger实例
	logger := logrus.New()

	// 设置日志级别
	level := getLogLevel(global.Conf.Log.Level)
	logger.SetLevel(level)

	// 设置格式化器
	logger.SetFormatter(customFormatter)

	// 启用调用者信息记录
	logger.SetReportCaller(true)

	// 设置输出到控制台
	logger.SetOutput(os.Stdout)

	// 创建info日志writer
	infoWriter := &lumberjack.Logger{
		Filename:   getLogFileName(fmt.Sprintf("%s/info", global.Conf.Log.Path), "info"),
		MaxSize:    global.Conf.Log.MaxSize,
		MaxBackups: global.Conf.Log.MaxBackups,
		MaxAge:     global.Conf.Log.MaxAge,
		Compress:   global.Conf.Log.Compress,
	}

	// 创建error日志writer
	errorWriter := &lumberjack.Logger{
		Filename:   getLogFileName(fmt.Sprintf("%s/error", global.Conf.Log.Path), "error"),
		MaxSize:    global.Conf.Log.MaxSize,
		MaxBackups: global.Conf.Log.MaxBackups,
		MaxAge:     global.Conf.Log.MaxAge,
		Compress:   global.Conf.Log.Compress,
	}

	// 创建hook来分别处理不同级别的日志
	hook := &CustomHook{
		InfoWriter:  infoWriter,
		ErrorWriter: errorWriter,
		Formatter:   customFormatter,
	}

	// 添加钩子，将日志写入文件
	logger.AddHook(hook)

	// 将自定义logger设置为全局logger
	global.Logger = logger

}

// 创建日志目录
func createLogDir() {
	// 获取当前工作目录
	cwd, err := os.Getwd()
	if err != nil {
		panic(fmt.Sprintf("Failed to get current working directory: %v", err))
	}

	// 转换相对路径为绝对路径
	logDir := global.Conf.Log.Path
	if !filepath.IsAbs(logDir) {
		logDir = filepath.Join(cwd, logDir)
	}

	// 更新全局配置中的路径为绝对路径
	global.Conf.Log.Path = logDir

	// 创建主日志目录
	if err := os.MkdirAll(logDir, 0755); err != nil {
		panic(fmt.Sprintf("Failed to create main log directory %s: %v", logDir, err))
	}

	// 创建info和error子目录
	dirs := []string{
		fmt.Sprintf("%s/info", logDir),
		fmt.Sprintf("%s/error", logDir),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			panic(fmt.Sprintf("Failed to create log directory %s: %v", dir, err))
		}
	}

	fmt.Printf("Log directories created at: %s\n", logDir)
}

// 生成日志文件名（只按日期，不按时间）
func getLogFileName(basePath, logType string) string {
	today := time.Now().Format("2006-01-02")
	// 使用filepath.Join确保路径分隔符适合当前操作系统
	return filepath.Join(basePath, fmt.Sprintf("%s-%s.log", logType, today))
}

// 获取日志级别
func getLogLevel(level string) logrus.Level {
	switch level {
	case "debug":
		return logrus.DebugLevel
	case "info":
		return logrus.InfoLevel
	case "warn":
		return logrus.WarnLevel
	case "error":
		return logrus.ErrorLevel
	default:
		return logrus.InfoLevel
	}
}

// CustomHook 自定义hook，用于将不同级别的日志写入不同文件
type CustomHook struct {
	InfoWriter  *lumberjack.Logger
	ErrorWriter *lumberjack.Logger
	Formatter   logrus.Formatter
}

func (hook *CustomHook) Fire(entry *logrus.Entry) error {
	line, err := hook.Formatter.Format(entry)
	if err != nil {
		return err
	}

	switch entry.Level {
	case logrus.PanicLevel, logrus.FatalLevel, logrus.ErrorLevel:
		_, err = hook.ErrorWriter.Write(line)
	case logrus.WarnLevel, logrus.InfoLevel, logrus.DebugLevel:
		_, err = hook.InfoWriter.Write(line)
	}

	return err
}

func (hook *CustomHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// CustomFormatter 自定义日志格式化器
type CustomFormatter struct {
	TimestampFormat string
}

// Format 实现logrus.Formatter接口
func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// 格式化时间戳
	timestamp := entry.Time.Format(f.TimestampFormat)

	// 获取调用者信息
	var funcName, fileInfo string
	if entry.HasCaller() {
		// 从完整函数路径中提取函数名，去掉包名部分
		parts := strings.Split(entry.Caller.Function, ".")
		if len(parts) > 0 {
			funcName = parts[len(parts)-1] + "()"
		} else {
			funcName = entry.Caller.Function + "()"
		}
		fileInfo = fmt.Sprintf("%s:%d", filepath.Base(entry.Caller.File), entry.Caller.Line)
	}

	// 按照格式组织日志内容: 2006/01/02 15:04:05 [ERROR] main() main:20 Your log message here
	logMessage := fmt.Sprintf("%s [%s] %s %s %s\n",
		timestamp,
		strings.ToUpper(entry.Level.String()),
		funcName,
		fileInfo,
		entry.Message)

	return []byte(logMessage), nil
}

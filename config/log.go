package config

type Log struct {
	Level      string `json:"level" yaml:"level"` // 日志级别
	Path       string `json:"path" yaml:"path"`
	MaxSize    int    `json:"maxsize" yaml:"maxsize"`       // 单个文件最大大小(MB)
	MaxBackups int    `json:"maxbackups" yaml:"maxbackups"` // 保留旧文件数量
	MaxAge     int    `json:"maxage" yaml:"maxage"`         // 保留天数
	Compress   bool   `json:"compress" yaml:"compress"`     // 是否压缩
}

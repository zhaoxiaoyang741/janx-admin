package config

type System struct {
	Port          uint   `json:"port" yaml:"port"`
	Debug         bool   `json:"debug" yaml:"debug"`
	StaticPath    string `json:"staticpath" yaml:"staticpath"`
	UrlPathPrefix string `json:"url-path-prefix" yaml:"url-path-prefix"`
}

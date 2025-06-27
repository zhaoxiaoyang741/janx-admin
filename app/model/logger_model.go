package model

type SystemRouteLog struct {
	Url       string `json:"url"`
	Method    string `json:"method"`
	Ip        string `json:"ip"`
	UserAgent string `json:"user_agent"`
	Duration  int64  `json:"duration"`
}

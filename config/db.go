package config

type Db struct {
	Host            string `json:"host" yaml:"host"`
	Port            uint   `json:"port" yaml:"port"`
	Hostname        string `json:"hostname" yaml:"hostname"`
	Password        string `json:"password" yaml:"password"`
	Database        string `json:"database" yaml:"database"`
	MaxOpenConns    uint   `json:"maxopenconns" yaml:"maxopenconns"`
	MaxIdleConns    uint   `json:"maxidleconns" yaml:"maxidleconns"`
	ConnMaxLifetime uint   `json:"connmaxlifetime" yaml:"connmaxlifetime"`
	LogStatus       bool   `json:"logstatus" yaml:"logstatus"`
	LogLevel        string `json:"loglevel" yaml:"loglevel"`
}

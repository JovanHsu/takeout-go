package nacos

import (
	"github.com/spf13/viper"
	"log"
	"time"
)

type (
	NacosConfig struct {
		NacosClientConfig *NacosClientConfig   `mapstructure:"client"`
		NacosServerConfig []*NacosServerConfig `mapstructure:"server"`
	}

	NacosClientConfig struct {
		TimeoutMs            uint64                   `mapstructure:"timeout"`                 // 请求Nacos服务器超时，默认值为10000ms,单位毫秒
		BeatInterval         int64                    `mapstructure:"beat_interval"`           // 向服务器发送心跳的时间间隔，缺省值为5000ms
		NamespaceId          string                   `mapstructure:"namespace_id"`            // namespaceId。当命名空间是公共的时候，在这里填写空白字符串。
		AppName              string                   `mapstructure:"app_name"`                // 应用名
		Endpoint             string                   `mapstructure:"endpoint"`                // 获取Nacos服务器地址的端点
		RegionId             string                   `mapstructure:"region_id"`               // the regionId for kms
		AccessKey            string                   `mapstructure:"access_key"`              // the AccessKey for kms
		SecretKey            string                   `mapstructure:"secret_key"`              // the SecretKey for kms
		OpenKMS              bool                     `mapstructure:"open_KMS"`                // it's to open kms,default is false. https://help.aliyun.com/product/28933.html
		CacheDir             string                   `mapstructure:"cache_dir"`               // the directory for persist nacos service info,default value is current path
		UpdateThreadNum      int                      `mapstructure:"update_thread_num"`       // 更新nacos服务信息的gorutine数，默认值为20
		NotLoadCacheAtStart  bool                     `mapstructure:"not_load_cache_at_start"` // not to load persistent nacos service info in CacheDir at start time
		UpdateCacheWhenEmpty bool                     `mapstructure:"update_cache_when_empty"` // update cache when get empty service instance from server
		Username             string                   `mapstructure:"username"`                // the username for nacos auth
		Password             string                   `mapstructure:"password"`                // the password for nacos auth
		LogDir               string                   `mapstructure:"log_dir"`                 // the directory for log, default is current path
		LogLevel             string                   `mapstructure:"log_level"`               // the level of log, it's must be debug,info,warn,error, default value is info
		LogSamplingConfig    *ClientLogSamplingConfig `mapstructure:"log_sampling_config"`     // the sampling config of log
		ContextPath          string                   `mapstructure:"context_path"`            // the nacos server contextpath
		LogRollingConfig     *ClientLogRollingConfig  `mapstructure:"log_rolling_config"`      // the log rolling config
	}

	NacosServerConfig struct {
		Scheme      string `mapstructure:"scheme"`       //the nacos server scheme
		ContextPath string `mapstructure:"context_path"` //the nacos server contextpath
		IpAddr      string `mapstructure:"ip_addr"`      //the nacos server address
		Port        uint64 `mapstructure:"port"`         //the nacos server port
	}

	ClientLogSamplingConfig struct {
		Initial    int           `mapstructure:"initial"`    //the sampling initial of log
		Thereafter int           `mapstructure:"thereafter"` //the sampling thereafter of log
		Tick       time.Duration `mapstructure:"tick"`       //the sampling tick of log
	}

	ClientLogRollingConfig struct {
		// MaxSize is the maximum size in megabytes of the log file before it gets
		// rotated. It defaults to 100 megabytes.
		MaxSize int `mapstructure:"max_size"`

		// MaxAge is the maximum number of days to retain old log files based on the
		// timestamp encoded in their filename.  Note that a day is defined as 24
		// hours and may not exactly correspond to calendar days due to daylight
		// savings, leap seconds, etc. The default is not to remove old log files
		// based on age.
		MaxAge int `mapstructure:"max_age"`

		// MaxBackups is the maximum number of old log files to retain.  The default
		// is to retain all old log files (though MaxAge may still cause them to get
		// deleted.)
		MaxBackups int `mapstructure:"max_backups"`

		// LocalTime determines if the time used for formatting the timestamps in
		// backup files is the computer's local time.  The default is to use UTC
		// time.
		LocalTime bool `mapstructure:"local_time"`

		// Compress determines if the rotated log files should be compressed
		// using gzip. The default is not to perform compression.
		Compress bool `mapstructure:"compress"`
	}
)

func LoadConfig() *NacosConfig {
	viper.SetConfigName("nacos.config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./conf")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	var config *NacosConfig
	if err != nil {
		log.Fatal("读取Nacos配置文件失败.", err)
	}
	_ = viper.Unmarshal(config)
	return config
}

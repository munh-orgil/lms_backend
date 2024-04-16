package global

import (
	"github.com/spf13/viper"
)

type Config struct {
	// system
	Port        string `mapstructure:"PORT"`
	SystemCode  string `mapstructure:"SYSTEM_CODE"`
	SystemToken string `mapstructure:"SYSTEM_TOKEN"`
	// db
	DBHost            string `mapstructure:"DB_HOST"`
	DBUserName        string `mapstructure:"DB_USERNAME"`
	DBUserPassword    string `mapstructure:"DB_PASSWORD"`
	DBName            string `mapstructure:"DB_NAME"`
	DBPort            string `mapstructure:"DB_PORT"`
	DBSchema          string `mapstructure:"DB_SCHEMA"`
	DBTablePrefix     string `mapstructure:"DB_TABLE_PREFIX"`
	DBMaxIdleConn     int    `mapstructure:"DB_MAX_IDLE_CONN"`
	DBMaxOpenConn     int    `mapstructure:"DB_MAX_OPEN_CONN"`
	DBMaxConnLifetime int    `mapstructure:"DB_MAX_CONN_LIFETIME"`
	// jwt
	JwtSecret           string `mapstructure:"JWT_SECRET"`
	JwtSecretPubKeyPath string `mapstructure:"JWT_SECRET_PUBLIC_KEY_PATH"`
	JwtSecretPrvKeyPath string `mapstructure:"JWT_SECRET_PRIVATE_KEY_PATH"`
	// otp
	OtpWaitSecond   int    `mapstructure:"OTP_WAIT_SECOND"`
	GmailUsername   string `mapstructure:"GMAIL_USERNAME"`
	GmailPassword   string `mapstructure:"GMAIL_PASSWORD"`
	PathOtpTemplate string `mapstructure:"PATH_OTP_TEMPLATE"`
}

var Conf *Config

func LoadConfig(path string) (err error) {

	if Conf, err = load("app", "env", path, Conf); err != nil {
		return err
	}

	return nil
}

func load[T any](configName, configType, path string, responseType T) (T, error) {
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	viper.AddConfigPath(path)

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return responseType, err
	}

	var res T
	if err := viper.Unmarshal(&res); err != nil {
		return responseType, err
	}
	return res, nil
}

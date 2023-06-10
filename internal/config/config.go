package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func GetConf() {
	viper.AddConfigPath(".")
	viper.AddConfigPath("./..")
	viper.AddConfigPath("./../..")
	viper.SetConfigName("config")

	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Warningf("%v", err)
	}
}

func GetPort() string {
	return viper.GetString("port")
}

func RedisUserWorkerHost() string {
	return viper.GetString("redis.user_worker.host")
}

func RedisUserWorkerPassword() string {
	return viper.GetString("redis.user_worker.password")
}

func RedisUserWorkerDB() int {
	return viper.GetInt("redis.user_worker.db")
}

func RedisCompanyWorkerHost() string {
	return viper.GetString("redis.company_worker.host")
}

func RedisCompanyWorkerPassword() string {
	return viper.GetString("redis.company_worker.password")
}

func RedisCompanyWorkerDB() int {
	return viper.GetInt("redis.company_worker.db")
}

func RedisNotificationWorkerHost() string {
	return viper.GetString("redis.notification_worker.host")
}

func RedisNotificationWorkerPassword() string {
	return viper.GetString("redis.notification_worker.password")
}

func RedisNotificationWorkerDB() int {
	return viper.GetInt("redis.notification_worker.db")
}

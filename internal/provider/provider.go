package provider

import (
	"github.com/bagastri07/asyqnmon-multi-worker/internal/config"
	"github.com/bagastri07/asyqnmon-multi-worker/internal/constant"
	"github.com/hibiken/asynq"
	"github.com/hibiken/asynqmon"
)

type AsyqnmonProvider struct {
	UserWorker         *asynqmon.HTTPHandler
	CompanyWorker      *asynqmon.HTTPHandler
	NotificationWorker *asynqmon.HTTPHandler
}

func NewAsyqnmonProvider() *AsyqnmonProvider {
	userWorker := asynqmon.New(asynqmon.Options{
		RootPath: constant.UserWorkerURL,
		RedisConnOpt: asynq.RedisClientOpt{
			Addr:     config.RedisUserWorkerHost(),
			Password: config.RedisUserWorkerPassword(),
			DB:       config.RedisUserWorkerDB(),
		},
	})

	companyWorker := asynqmon.New(asynqmon.Options{
		RootPath: constant.CompanyWorkerURL,
		RedisConnOpt: asynq.RedisClientOpt{
			Addr:     config.RedisCompanyWorkerHost(),
			Password: config.RedisCompanyWorkerPassword(),
			DB:       config.RedisCompanyWorkerDB(),
		},
	})

	notificationWorker := asynqmon.New(asynqmon.Options{
		RootPath: constant.NotificationWorkerURL,
		RedisConnOpt: asynq.RedisClientOpt{
			Addr:     config.RedisNotificationWorkerHost(),
			Password: config.RedisNotificationWorkerPassword(),
			DB:       config.RedisNotificationWorkerDB(),
		},
	})

	return &AsyqnmonProvider{
		UserWorker:         userWorker,
		CompanyWorker:      companyWorker,
		NotificationWorker: notificationWorker,
	}
}

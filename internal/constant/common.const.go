package constant

import "github.com/bagastri07/asyqnmon-multi-worker/internal/model"

const (
	UserWorkerURL         = "/user-worker"
	CompanyWorkerURL      = "/company-worker"
	NotificationWorkerURL = "/notification-worker"
)

var (
	WorkerData = model.PageData{
		Title: "Asynqmon Monitoring Multi Worker",
		Items: WorkerList,
	}

	WorkerList = []model.MenuItem{
		{Name: "user-worker", URL: UserWorkerURL + "/"},
		{Name: "company-worker", URL: CompanyWorkerURL + "/"},
		{Name: "notification-worker", URL: NotificationWorkerURL + "/"},
	}
)

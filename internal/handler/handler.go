package handler

import (
	"net/http"

	"github.com/bagastri07/asyqnmon-multi-worker/internal/constant"
	"github.com/bagastri07/asyqnmon-multi-worker/internal/provider"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	asynqmonProv *provider.AsyqnmonProvider
}

func NewHandler(ap *provider.AsyqnmonProvider) *Handler {
	return &Handler{
		asynqmonProv: ap,
	}
}

func (h *Handler) RootHandler(c echo.Context) error {

	return c.Render(http.StatusOK, "index.html", constant.WorkerData)
}

func (h *Handler) UserWorkerHandler(c echo.Context) error {
	userWorkerHandler := h.asynqmonProv.UserWorker

	return echo.WrapHandler(userWorkerHandler)(c)
}

func (h *Handler) CompanyWorkerHandler(c echo.Context) error {
	companyWorkerHandler := h.asynqmonProv.CompanyWorker

	return echo.WrapHandler(companyWorkerHandler)(c)
}

func (h *Handler) NotificationWorkerHandler(c echo.Context) error {
	notificationWorkerHandler := h.asynqmonProv.NotificationWorker

	return echo.WrapHandler(notificationWorkerHandler)(c)
}

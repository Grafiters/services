package routes

import (
	"net/http"
	controllers "riskmanagement/controllers/pelaporan"
	"riskmanagement/lib"
	"riskmanagement/middlewares"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
	"golang.org/x/time/rate"
)

type PelaporanRoutes struct {
	logger              logger.Logger
	handler             lib.RequestHandler
	PelaporanController controllers.PelaporanController
	authMiddleware      middlewares.JWTAuthMiddleware
}

func (p PelaporanRoutes) Setup() {

	// request_per_second, err := strconv.Atoi(os.Getenv("REQUEST_PER_SECOND"))
	// if err != nil {
	// 	request_per_second = 1
	// }

	// burst_size, err := strconv.Atoi(os.Getenv("BURST_SIZE"))
	// if err != nil {
	// 	burst_size = 1
	// }

	p.logger.Zap.Info("Setting up routes")
	// apiLimiter := rate.NewLimiter(rate.Limit(request_per_second), burst_size)
	// apiLimiter := rate.NewLimiter(1, 1)
	api := p.handler.Gin.Group("api/v1/pelaporan").Use(p.authMiddleware.Handler())
	{
		// api.POST("/insert-draft", p.PelaporanController.InsertDraftSurat)
		api.POST("/draft-list", p.PelaporanController.GetDraftList)
		api.POST("/draft-detail", p.PelaporanController.GetDraftDetail)
		api.POST("/approval-list", p.PelaporanController.GetApprovalList)

		// api.POST("/create-draft", rateLimit(apiLimiter), p.PelaporanController.TambahDraftSurat)
		api.POST("/create-draft", p.PelaporanController.TambahDraftSurat)
		api.POST("/approve", p.PelaporanController.Approve)
		api.POST("/reject", p.PelaporanController.Reject)
		api.POST("/delete", p.PelaporanController.Delete)
	}
}

func rateLimit(limiter *rate.Limiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		if limiter.Allow() {
			c.Next()
		} else {
			// lib.ReturnToJson(c, 200, "400", "Request Terlalu Banyak", false)
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Terlalu Banyak Request. Silahkan Coba Lagi Nanti."})
		}
	}
}

func NewPelaporanRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	PelaporanController controllers.PelaporanController,
	authMiddleware middlewares.JWTAuthMiddleware,
) PelaporanRoutes {
	return PelaporanRoutes{
		logger:              logger,
		handler:             handler,
		PelaporanController: PelaporanController,
		authMiddleware:      authMiddleware,
	}
}

package middlewares

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"riskmanagement/lib"
	models "riskmanagement/models/audittrail"
	auditTrail "riskmanagement/services/audittrail"
	"strings"
	"text/template"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
	"go.uber.org/zap"
)

type AuditTrailMiddleware struct {
	handler    lib.RequestHandler
	logger     logger.Logger
	db         lib.Database
	auditTrail auditTrail.AuditTrailDefinition
}

func NewAuditTrailMiddleware(
	handler lib.RequestHandler,
	logger logger.Logger,
	db lib.Database,
	auditTrail auditTrail.AuditTrailDefinition,
) AuditTrailMiddleware {
	return AuditTrailMiddleware{
		handler:    handler,
		logger:     logger,
		db:         db,
		auditTrail: auditTrail,
	}
}

func (m AuditTrailMiddleware) Setup() {
	m.logger.Zap.Info("setting up audit trail middleware")
}

func (m AuditTrailMiddleware) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("masuk middleware")
		if strings.Contains(c.FullPath(), "store") {
			// ... logic audit trail kamu di sini
			var bodyBytes []byte
			ipaddress := template.HTMLEscapeString(c.ClientIP())

			if c.Request.Body != nil {
				bodyBytes, _ = io.ReadAll(c.Request.Body)
			}

			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

			// Ambil data request body
			var req map[string]interface{}
			if err := json.Unmarshal(bodyBytes, &req); err != nil {
				req = map[string]interface{}{}
			}

			aktifitas := "Store"
			if strings.Contains(c.FullPath(), "briefing") {
				aktifitas = "Briefing"
			} else if strings.Contains(c.FullPath(), "coaching") {
				aktifitas = "Coaching"
			} else if strings.Contains(c.FullPath(), "verifikasi") {
				aktifitas = "Verifikasi"
			}

			logPayload := models.AuditTrail{
				PN:          toString(req["pernr"]),
				NamaBrcUrc:  toString(req["maker_desc"]),
				REGION:      toString(req["REGION"]),
				RGDESC:      toString(req["RGDESC"]),
				MAINBR:      toString(req["MAINBR"]),
				MBDESC:      toString(req["MBDESC"]),
				BRANCH:      toString(req["BRANCH"]),
				BRDESC:      toString(req["BRDESC"]),
				NoPelaporan: toString(req["no_pelaporan"]),
				Aktifitas:   aktifitas,
				IpAddress:   ipaddress,
				Lokasi:      toString(req["lokasi"]),
			}

			fmt.Println("logPayload ==>", logPayload)
			go func(payload models.AuditTrail) {
				if _, err := m.auditTrail.Store(payload); err != nil {
					m.logger.Zap.Error("failed to store audit trail", zap.Error(err))
				}
			}(logPayload)
		}
		c.Next()
	}
}

func toString(val interface{}) string {
	if val == nil {
		return ""
	}
	if s, ok := val.(string); ok {
		return s
	}
	return ""
}

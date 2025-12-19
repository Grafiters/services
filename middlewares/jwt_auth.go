package middlewares

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	services "riskmanagement/services/auth"
	"strings"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

// JWTAuthMiddleware middleware for jwt authentication
type JWTAuthMiddleware struct {
	service services.JWTAuthService
	logger  logger.Logger
}

// NewJWTAuthMiddleware creates new jwt auth middleware
func NewJWTAuthMiddleware(
	logger logger.Logger,
	service services.JWTAuthService,
) JWTAuthMiddleware {
	return JWTAuthMiddleware{
		service: service,
		logger:  logger,
	}
}

type PernrData struct {
	Pernr string `"json:pernr"`
}

func cleanJSONBody(body []byte) []byte {
	// Replace NBSP (0xC2 0xA0) dengan spasi biasa
	body = bytes.ReplaceAll(body, []byte{0xC2, 0xA0}, []byte{0x20})
	// Kalau ada BOM (0xEF 0xBB 0xBF), hapus juga
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))
	return body
}

// Setup sets up jwt auth middleware
func (m JWTAuthMiddleware) Setup() {}

func (m JWTAuthMiddleware) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		contentType := c.ContentType()
		var pernr string

		switch contentType {
		case "application/json":
			body, _ := io.ReadAll(c.Request.Body)
			// simpan ulang body biar bisa dipakai lagi downstream
			c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

			// Debug print
			// fmt.Println("Body (raw bytes) =>", body)
			fmt.Println("Body (as string) =>", string(body))

			// Bersihkan hidden char sebelum Unmarshal
			cleanBody := cleanJSONBody(body)

			var data PernrData
			if err := json.Unmarshal(cleanBody, &data); err != nil {
				fmt.Println("Unmarshal error =>", err)
			} else {
				fmt.Println("bodyyy =======", data.Pernr)
				pernr = data.Pernr
			}

		case "multipart/form-data":
			pernr = c.PostForm("pernr")
		}

		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")

		if len(t) == 2 {
			authToken := t[1]
			authorized, err, claim := m.service.Authorize(authToken, pernr)
			// fmt.Println("T1 ====> ", t[1])
			fmt.Println("authorized", authorized)
			fmt.Println("claim", claim)

			if authorized {
				// fmt.Println("mashok 1")
				c.Set("pernr", pernr)
				c.Next()
				return
			}

			c.JSON(http.StatusUnauthorized, gin.H{
				"code":  401,
				"error": err.Error(),
			})

			m.logger.Zap.Error(err.Error())
			c.Abort()
			return
		}
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":  401,
			"error": "you are not authorized",
		})
		c.Abort()
	}
}

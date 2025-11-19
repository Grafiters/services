package middlewares

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"riskmanagement/constants"
	"riskmanagement/lib"
	"riskmanagement/lib/env"
	"text/template"

	"github.com/gin-gonic/gin"
	elastic "gitlab.com/golang-package-library/elasticsearch"
	"gitlab.com/golang-package-library/logger"
)

// CorsMiddleware middleware for cors
type LogActivityMiddleware struct {
	handler lib.RequestHandler
	logger  logger.Logger
	env     env.Env
	elastic elastic.Elasticsearch
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// NewCorsMiddleware creates new cors middleware
func NewLogActivityMiddleware(
	handler lib.RequestHandler,
	logger logger.Logger, env env.Env,
) LogActivityMiddleware {
	return LogActivityMiddleware{
		handler: handler,
		logger:  logger,
		env:     env,
	}
}

// Setup sets up cors middleware
func (m LogActivityMiddleware) Setup() {
	m.logger.Zap.Info("Setting up Log Activity")
	
	// m.handler.Gin.Use(m.DummyMiddleware)
	m.handler.Gin.Use(
		func(c *gin.Context) {
			client := m.elastic.Client
			m.logger.Zap.Info("incoming request")
			c.Set(constants.DBTransaction, client)

			// client := m.elastic.Client
			uri := template.HTMLEscapeString(c.Request.RequestURI)
			agent := template.HTMLEscapeString(c.Request.UserAgent())
			ipaddress := template.HTMLEscapeString(c.ClientIP())

			// uri := c.Request.RequestURI
			// agent := c.Request.UserAgent()
			// ipaddress := c.ClientIP()

			// get request body
			var bodyBytes []byte
			var err error

			if c.Request.Body != nil {
				bodyBytes, err = ioutil.ReadAll(c.Request.Body)
				if err != nil {
					m.logger.Zap.Info(err.Error())
				}
			}
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
			blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
			c.Writer = blw

			c.Next()

			if statusInList(c.Writer.Status(), []int{http.StatusOK, http.StatusCreated}) {
				m.logger.Zap.Info("response success")
				lib.CreateLogActivityToDB(client, uri, agent, ipaddress, string(bodyBytes), blw.body.String())
			}
		})
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (m LogActivityMiddleware) DummyMiddleware(c *gin.Context) {

	_ = func(ctx *gin.Context, reqBody, respBody []byte) {
		uri := template.HTMLEscapeString(ctx.Request.RequestURI)
		agent := template.HTMLEscapeString(ctx.Request.UserAgent())
		ipaddress := template.HTMLEscapeString(ctx.ClientIP())

		// uri := c.Request.RequestURI
		// agent := c.Request.UserAgent()
		// ipaddress := c.ClientIP()

		fmt.Println("Middleware", uri)
		lib.CreateLogActivityToDB(m.elastic.Client, uri, agent, ipaddress, string(reqBody), string(respBody))

	}
	// Pass on to the next-in-chain
	c.Next()
}

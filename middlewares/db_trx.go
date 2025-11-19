package middlewares

import (
	"net/http"

	"riskmanagement/constants"
	"riskmanagement/lib"

	"github.com/gin-gonic/gin"
	elastic "gitlab.com/golang-package-library/elasticsearch"
	"gitlab.com/golang-package-library/logger"
)

// DatabaseTrx middleware for transactions support for database
type DatabaseTrx struct {
	handler lib.RequestHandler
	logger  logger.Logger
	db      lib.Database
	elastic elastic.Elasticsearch
	// body    *bytes.Buffer
}

// statusInList function checks if context writer status is in provided list
func statusInList(status int, statusList []int) bool {
	for _, i := range statusList {
		if i == status {
			return true
		}
	}
	return false
}

// NewDatabaseTrx creates new database transactions middleware
func NewDatabaseTrx(
	handler lib.RequestHandler,
	logger logger.Logger,
	db lib.Database,
) DatabaseTrx {
	return DatabaseTrx{
		handler: handler,
		logger:  logger,
		db:      db,
	}
}

// Setup sets up database transaction middleware
func (m DatabaseTrx) Setup() {
	m.logger.Zap.Info("setting up database transaction middleware")

	m.handler.Gin.Use(func(c *gin.Context) {
		txHandle := m.db.DB.Begin()
		m.logger.Zap.Info("beginning database transaction")

		defer func() {
			if r := recover(); r != nil {
				txHandle.Rollback()
			}
		}()

		c.Set(constants.DBTransaction, txHandle)
		c.Next()
		// rollback transaction on server errors
		if c.Writer.Status() == http.StatusInternalServerError {
			m.logger.Zap.Info("rolling back transaction due to status code: 500")
			txHandle.Rollback()
		}

		// commit transaction on success status
		if statusInList(c.Writer.Status(), []int{http.StatusOK, http.StatusCreated}) {

			m.logger.Zap.Info("committing transactions")
			if err := txHandle.Commit().Error; err != nil {
				m.logger.Zap.Error("trx commit error: ", err)
			}
		}
	})
}

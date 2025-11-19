package product

import (
	"math"
	"riskmanagement/lib"
	models "riskmanagement/models/product"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type ProductDefinition interface {
	GetAll() (responses []models.ProductResponse, err error)
	GetAllWithPage(request *models.PageRequest) (responses []models.ProductResponse, totalRows int, totalData int, err error)
	GetOne(id int64) (responses models.ProductResponse, err error)
	Store(request *models.ProductRequest) (responses bool, err error)
	Update(request *models.ProductRequest) (responese bool, err error)
	Delete(id int64) (err error)
	WithTrx(trxHandle *gorm.DB) ProductRepository
	GetKodeProduct() (responses []models.KodeProduct, err error)
	GetProductByActivity(request *models.GetProductByActivityRequest) (responses []models.GetProductByActivityResponseNull, err error)
	GetProductBySegment(request *models.GetProductBySegmentRequest) (responses []models.GetProductBySegmentResponse, err error)
	SearchProduct(request *models.KeywordRequest) (responses []models.ProductResponsesNull, totalRows int, totalData int, err error)
}

type ProductRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewProductRepository(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
) ProductDefinition {
	return ProductRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// Delete implements ProductDefinition
func (product ProductRepository) Delete(id int64) (err error) {
	return product.db.DB.Where("id = ?", id).Delete(&models.ProductResponse{}).Error
}

// GetAll implements ProductDefinition
func (product ProductRepository) GetAll() (responses []models.ProductResponse, err error) {
	return responses, product.db.DB.Find(&responses).Error
}

// GetAll implements ProductDefinition
func (product ProductRepository) GetAllWithPage(request *models.PageRequest) (responses []models.ProductResponse, totalRows int, totalData int, err error) {
	var count int64
	db := product.db.DB.Table("product p").Count(&count).Limit(request.Limit).Offset(request.Offset)

	totalData = int(count)

	if totalData > 0 {
		totalRows = int(math.Ceil(float64(totalData) / float64(request.Limit)))
	}

	err = db.Scan(&responses).Error
	if err != nil {
		return nil, totalRows, totalData, err
	}

	return responses, totalRows, totalData, err

}

// GetOne implements ProductDefinition
func (product ProductRepository) GetOne(id int64) (responses models.ProductResponse, err error) {
	return responses, product.db.DB.Where("id = ?", id).Find(&responses).Error
}

// Store implements ProductDefinition
func (product ProductRepository) Store(request *models.ProductRequest) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")
	err = product.db.DB.Save(&models.ProductRequest{
		ID:            0,
		KodeProduct:   request.KodeProduct,
		Product:       request.Product,
		ActivityID:    request.ActivityID,
		SubActivityID: request.SubActivityID,
		LiniBisnisLv1: request.LiniBisnisLv1,
		LiniBisnisLv2: request.LiniBisnisLv2,
		LiniBisnisLv3: request.LiniBisnisLv3,
		CreatedAt:     &timeNow,
		Segment:       request.Segment,
	}).Error
	if err != nil {
		return false, err
	}

	return true, err
}

// Update implements ProductDefinition
func (product ProductRepository) Update(request *models.ProductRequest) (responese bool, err error) {
	timeNow := lib.GetTimeNow("timestime")
	err = product.db.DB.Save(&models.ProductRequest{
		ID:            request.ID,
		KodeProduct:   request.KodeProduct,
		Product:       request.Product,
		ActivityID:    request.ActivityID,
		SubActivityID: request.SubActivityID,
		LiniBisnisLv1: request.LiniBisnisLv1,
		LiniBisnisLv2: request.LiniBisnisLv2,
		LiniBisnisLv3: request.LiniBisnisLv3,
		CreatedAt:     request.CreatedAt,
		UpdatedAt:     &timeNow,
		Segment:       request.Segment,
	}).Error
	if err != nil {
		return false, err
	}

	return true, err
}

// WithTrx implements ProductDefinition
func (product ProductRepository) WithTrx(trxHandle *gorm.DB) ProductRepository {
	if trxHandle == nil {
		product.logger.Zap.Error("transaction Database not found in gin context")
		return product
	}

	product.db.DB = trxHandle
	return product
}

// GetKodeProduct implements ProductDefinition
func (product ProductRepository) GetKodeProduct() (responses []models.KodeProduct, err error) {
	query := `SELECT CONCAT("PROD.", RIGHT(CONCAT("0000",(count(*) + 1)), 4) ) 'kode_product' FROM product`

	product.logger.Zap.Info(query)
	rows, err := product.dbRaw.DB.Query(query)
	defer rows.Close()

	product.logger.Zap.Info("Rows ", rows)
	if err != nil {
		return responses, err
	}

	response := models.KodeProduct{}
	for rows.Next() {
		_ = rows.Scan(
			&response.KodeProduct,
		)

		responses = append(responses, response)
	}

	if err = rows.Err(); err != nil {
		return responses, err
	}

	return responses, err
}

func (product ProductRepository) GetProductByActivity(request *models.GetProductByActivityRequest) (responses []models.GetProductByActivityResponseNull, err error) {
	query := `SELECT id, kode_product as code, product as name 
			  FROM product
			  WHERE activity_id = ` + request.ActivityID

	product.logger.Zap.Info(query)
	rows, err := product.dbRaw.DB.Query(query)
	defer rows.Close()

	product.logger.Zap.Info("Rows ", rows)
	if err != nil {
		return responses, err
	}

	response := models.GetProductByActivityResponseNull{}
	for rows.Next() {
		_ = rows.Scan(
			&response.ID,
			&response.Code,
			&response.Name,
		)

		responses = append(responses, response)
	}

	if err = rows.Err(); err != nil {
		return responses, err
	}

	return responses, err
}

func (product ProductRepository) GetProductBySegment(request *models.GetProductBySegmentRequest) (responses []models.GetProductBySegmentResponse, err error) {
	// totalType := request.TotalType
	db := product.db.DB

	query := db.Table(`product p`).
		Select(`id, kode_product as code, product as name `)

		//Sorting
	if request.Segment != "" && request.Segment != "all" {
		query = query.Where(`p.segment`, request.Segment)
	} else {
		query = query.Where(`p.segment IN (SELECT segment FROM segment_realisasi_kredit WHERE is_active = 1)`)
	}

	query = query.
		Find(&responses)

	product.logger.Zap.Info("verifikasi-query-activity-unknown", query)

	if query.Error != nil {
		return responses, err
	}

	return responses, err
}

func (product ProductRepository) SearchProduct(request *models.KeywordRequest) (responses []models.ProductResponsesNull, totalRows int, totalData int, err error) {
	where := ``

	if request.ActivityID != "" {
		where += ` AND p.activity_id = ` + request.ActivityID
	}

	if request.Keyword != "" {
		where += ` AND p.product LIKE '%` + request.Keyword + `%'`
	}

	if request.Segment != "" {
		where += ` AND p.segment = '` + request.Segment + `'`
	}

	query := `SELECT p.id, p.product FROM product p WHERE p.kode_product IS NOT NULL ` + where + ` LIMIT ? OFFSET ?`

	product.logger.Zap.Info(query)
	rows, err := product.dbRaw.DB.Query(query, request.Limit, request.Offset)
	if err != nil {
		return responses, totalRows, totalData, err
	}

	product.logger.Zap.Info("rows ", rows)

	response := models.ProductResponsesNull{}
	for rows.Next() {
		err = rows.Scan(
			&response.ID,
			&response.Product,
		)
		if err != nil {
			return responses, totalRows, totalData, err
		}

		responses = append(responses, response)
	}

	if err = rows.Err(); err != nil {
		return responses, totalRows, totalData, err
	}

	paginationQuery := `SELECT count(p.product) FROM product p WHERE p.kode_product IS NOT NULL ` + where
	err = product.dbRaw.DB.QueryRow(paginationQuery).Scan(&totalData)
	if err != nil {
		return responses, totalRows, totalData, err
	}

	if totalData > 0 {
		totalRows = int(math.Ceil((float64(totalData)) / float64(request.Limit)))
	}
	return responses, totalRows, totalData, err
}

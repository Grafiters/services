package product

import (
	"database/sql"
	"fmt"
	"riskmanagement/lib"
	models "riskmanagement/models/product"
	repository "riskmanagement/repository/product"

	"gitlab.com/golang-package-library/logger"
)

type ProductDefinition interface {
	GetAll() (responses []models.ProductResponse, err error)
	GetOne(id int64) (responses models.ProductResponse, err error)
	Store(request *models.ProductRequest) (response bool, err error)
	Update(request *models.ProductRequest) (err error)
	Delete(id int64) (err error)
	GetKodeProduct() (responses []models.KodeProduct, err error)
	GetAllWithPage(request models.PageRequest) (responses []models.ProductResponse, pagination lib.Pagination, err error)
	GetProductByActivity(request models.GetProductByActivityRequest) (responses []models.GetProductByActivityResponseNull, pagination lib.Pagination, err error)
	GetProductBySegment(request models.GetProductBySegmentRequest) (responses []models.GetProductBySegmentResponse, pagination lib.Pagination, err error)
	SearchProduct(requests models.KeywordRequest) (responses []models.ProductResponse, pagination lib.Pagination, err error)
}

type ProductService struct {
	dbRaw      lib.Databases
	logger     logger.Logger
	repository repository.ProductDefinition
}

func NewProductService(
	dbRaw lib.Databases,
	logger logger.Logger,
	repository repository.ProductDefinition,
) ProductDefinition {
	return ProductService{
		dbRaw:      dbRaw,
		logger:     logger,
		repository: repository,
	}
}

func checkCount(rows *sql.Rows) (count int) {
	for rows.Next() {
		err := rows.Scan(&count)
		checkErr(err)
	}
	return count
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// Delete implements ProductDefinition
func (product ProductService) Delete(id int64) (err error) {
	return product.repository.Delete(id)
}

// GetAll implements ProductDefinition
func (product ProductService) GetAll() (responses []models.ProductResponse, err error) {
	return product.repository.GetAll()
}

// GetAllWithPage implements ProductDefinition
func (product ProductService) GetAllWithPage(request models.PageRequest) (responses []models.ProductResponse, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Limit = limit
	request.Page = page
	request.Offset = offset
	request.Order = order
	request.Sort = sort

	dataProduct, totalRows, totalData, err := product.repository.GetAllWithPage(&request)
	if err != nil {
		product.logger.Zap.Error(err)
		return responses, pagination, err
	}

	if totalData < 0 {
		product.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range dataProduct {
		responses = append(responses, models.ProductResponse{
			ID:            response.ID,
			KodeProduct:   response.KodeProduct,
			Product:       response.Product,
			ActivityID:    response.ActivityID,
			SubActivityID: response.SubActivityID,
			LiniBisnisLv1: response.LiniBisnisLv1,
			LiniBisnisLv2: response.LiniBisnisLv2,
			LiniBisnisLv3: response.LiniBisnisLv3,
			CreatedAt:     response.CreatedAt,
			UpdatedAt:     response.UpdatedAt,
			Segment:       response.Segment,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)
	return responses, pagination, err
}

// GetOne implements ProductDefinition
func (product ProductService) GetOne(id int64) (responses models.ProductResponse, err error) {
	return product.repository.GetOne(id)
}

// Store implements ProductDefinition
func (product ProductService) Store(request *models.ProductRequest) (response bool, err error) {
	rowsCheck, err := product.dbRaw.DB.Query("SELECT COUNT(*) FROM product WHERE product = ?", request.Product)

	checkErr(err)

	if checkCount(rowsCheck) < 1 {
		fmt.Println("service =", request)
		fmt.Println("Berhasil")
		status, err := product.repository.Store(request)
		if !status || err != nil {
			return false, err
		}

		return true, err
	}

	fmt.Println("Gagal")
	return false, err
}

// Update implements ProductDefinition
func (product ProductService) Update(request *models.ProductRequest) (err error) {
	status, err := product.repository.Update(request)
	if !status || err != nil {
		return err
	}

	return nil
}

// GetKodeActivity implements ProductDefinition
func (product ProductService) GetKodeProduct() (responses []models.KodeProduct, err error) {
	dataProduct, err := product.repository.GetKodeProduct()
	if err != nil {
		product.logger.Zap.Error(err)
		return responses, err
	}

	for _, response := range dataProduct {
		responses = append(responses, models.KodeProduct{
			KodeProduct: response.KodeProduct,
		})
	}

	return responses, err
}

func (product ProductService) GetProductByActivity(request models.GetProductByActivityRequest) (responses []models.GetProductByActivityResponseNull, pagination lib.Pagination, err error) {
	dataProduct, err := product.repository.GetProductByActivity(&request)

	responses = dataProduct

	return responses, pagination, err
}

func (product ProductService) GetProductBySegment(request models.GetProductBySegmentRequest) (responses []models.GetProductBySegmentResponse, pagination lib.Pagination, err error) {
	dataProduct, err := product.repository.GetProductBySegment(&request)

	responses = dataProduct

	return responses, pagination, err
}

func (product ProductService) SearchProduct(request models.KeywordRequest) (responses []models.ProductResponse, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort

	dataProduct, totalRows, totalData, err := product.repository.SearchProduct(&request)
	if err != nil {
		product.logger.Zap.Error(err)
		return responses, pagination, err
	}

	if totalData < 0 {
		product.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range dataProduct {
		responses = append(responses, models.ProductResponse{
			ID:      response.ID.Int64,
			Product: response.Product.String,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)
	return responses, pagination, err
}

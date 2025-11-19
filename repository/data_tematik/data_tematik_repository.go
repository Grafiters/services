package datatematik

import (
	"fmt"
	"regexp"
	"riskmanagement/lib"
	datatematik "riskmanagement/models/data_tematik"
	"strings"

	"gitlab.com/golang-package-library/logger"
)

type DataTematikDefinition interface {
	GetSampleDataTematik(request *datatematik.DataTematikRequest) (response datatematik.DataTematikResponse, totalData int64, err error)
	UpdateStatusDataSample(request *datatematik.RequestUpdate) (response bool, err error)
}

type DataTematikRepository struct {
	db     lib.Database
	logger logger.Logger
}

func NewDataTematikRepository(
	db lib.Database,
	logger logger.Logger,
) DataTematikDefinition {
	return DataTematikRepository{
		db:     db,
		logger: logger,
	}
}

func ConvertToSnakeCase(str string) string {
	// Replace all non-alphanumeric characters with underscores
	re := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	snake := re.ReplaceAllString(str, "_")

	// Convert to lowercase
	snake = strings.ToLower(snake)

	// Trim any leading or trailing underscores
	snake = strings.Trim(snake, "_")

	return snake
}

// GetSampleDataTematik implements DataTematikDefinition.
func (d DataTematikRepository) GetSampleDataTematik(request *datatematik.DataTematikRequest) (response datatematik.DataTematikResponse, totalData int64, err error) {
	if request.NamaTable == "" {
		return response, totalData, fmt.Errorf("Data Sample Repo not found !")
	}

	tableName := request.NamaTable + " lampiran"

	var columns []string
	err = d.db.DB.Table("information_schema.columns").
		Select("COLUMN_NAME").
		Where("table_name = ? AND table_schema = DATABASE()", request.NamaTable).
		Order("ORDINAL_POSITION").
		Pluck("COLUMN_NAME", &columns).Error

	if err != nil {
		fmt.Println("Error fetching column names:", err)
		return response, totalData, err
	}

	// fmt.Println(joinsColumns)
	selectedField := make([]string, len(columns))
	columnsList := make([]string, len(columns))
	for i, column := range columns {
		selectedField[i] = fmt.Sprintf("lampiran.`%s`", column)
		columnsList[i] = ConvertToSnakeCase(column)
	}

	joinsColumns := strings.Join(columnsList, ",")

	var rawData []map[string]interface{}
	query := d.db.DB.Table(tableName).
		Select(strings.Join(selectedField, ", ")).
		Joins(`LEFT JOIN tasklists_uker tu ON tu.tasklist_id = lampiran.tasklist_id AND tu.BRANCH = lampiran.BRANCH`). //AND tu.BRANCH = lampiran.BRANCH to be added
		// Joins(`LEFT JOIN tasklists_lampiran tl ON tl.tasklists_id = lampiran.tasklist_id`).
		Joins(`LEFT JOIN tasklists tl ON tl.id = lampiran.tasklist_id`).
		Where(`DATE(tl.created_at) = ?`, request.PeriodeData).
		Where(`tl.approval_status = 'Disetujui'`).
		Where(`tl.status = 'Aktif'`).
		Where(`tu.BRANCH = ?`, request.UnitKerja).
		Where(`lampiran.status NOT IN ('verify', 'draft')`).
		Group(`lampiran.id`).
		Order(`lampiran.id ASC`).
		Count(&totalData)

	err = query.Limit(int(request.Limit)).Offset(int(request.Offset)).Find(&rawData).Error
	if err != nil {
		fmt.Println("Error executing query:", err)
		return response, totalData, err
	}

	var columnsData []interface{}
	for _, row := range rawData {
		orderedRow := make(map[string]interface{})
		for _, column := range columns {
			orderedRow[ConvertToSnakeCase(column)] = row[column]
		}

		// fmt.Println(orderedRow)
		columnsData = append(columnsData, orderedRow)
	}

	response = datatematik.DataTematikResponse{
		Columns:     joinsColumns,
		ColumnsData: columnsData,
	}

	return response, totalData, nil
}

// UpdateStatusDataSample implements DataTematikDefinition.
func (d DataTematikRepository) UpdateStatusDataSample(request *datatematik.RequestUpdate) (response bool, err error) {
	fmt.Println("masuk repo", request)

	err = d.db.DB.Table(request.NamaTable).
		Where("id = ?", request.Id).
		Updates(&datatematik.RequestUpdate{
			Status:       request.Status,
			NoVerifikasi: request.NoVerifikasi,
		}).Error

	if err != nil {
		fmt.Println("Error updating data:", err)
		return false, err
	}

	return true, nil
}

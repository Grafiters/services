package pelaporan

import (
	"math"
	"riskmanagement/lib"
	models "riskmanagement/models/pelaporan"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type PelaporanDefinition interface {
	DraftList(requests models.DraftListRequest) (responses []models.DraftListResponse, totalRow int, totalData int, err error)
	ApprovalList(requests models.DraftListRequest) (responses []models.DraftListResponse, totalRow int, totalData int, err error)
	DraftDetail(id int64) (responses models.SuratDetail, err error)

	GenerateLaporan(requests models.GenerateLaporanRequest) (responses []models.GenerateLaporanResponse, total int, err error)
	StoreDraftPelaporan(requests *models.PelaporanDraft, tx *gorm.DB) (status bool, err error)
	Approve(requests *models.ApproveUpdate) (status bool, err error)
	Reject(requests *models.PenolakanCatatan, tx *gorm.DB) (status bool, err error)
	Delete(pelaporan *models.ApproveUpdate, rejected *models.PenolakanCatatan, tx *gorm.DB) (status bool, err error)

	GetPimpinanUker(branch string) (responses []models.PenerimaSuratResponse, err error)
	GetKodeUnik(branch string) (responses models.KodeResponse, jumlah int64, err error)
	GetNamaSigner(pernr string) (responses models.Signer, total int, err error)
}

type PelaporanRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewPelaporanRepository(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
) PelaporanDefinition {
	return PelaporanRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: 0,
	}
}

// DraftList implements PelaporanDefinition.
func (p PelaporanRepository) DraftList(requests models.DraftListRequest) (responses []models.DraftListResponse, totalRow int, totalData int, err error) {
	db := p.db.DB

	query := db.Table(`pelaporan_drafts pd`).
		Select(`pd.id, 
			pd.perihal "perihal",
			db.BRDESC "tujuan",
			pd.nomorSurat "nomor_surat",
			DATE_FORMAT(pd.executed,"%d %b %Y") "tanggal_surat",
			pd.responseStatus "response_status",
			CASE 
				WHEN pd.statusMCS = "1" THEN "Need Approve"
				WHEN pd.statusMCS = "2" AND pd.status = "Open" THEN "Proses Kirim"
				WHEN pd.statusMCS = "2" AND pd.status = "Executed" THEN "Berhasil Dikirim" 
				WHEN pd.statusMCS = "3" AND pd.status = "Rejected" THEN "Surat Ditolak"
				WHEN pd.statusMCS = "4" AND pd.status = "Fail" THEN "Gagal Dikirim"
				ELSE "Draft"
			END as "status_terakhir"`).
		Joins(`join dwh_branch db on pd.branchCodePenerima = LPAD(db.BRANCH,5,"0")`).
		Where(`idMaker = ?`, requests.PnMaker)

	if requests.Filter {
		if requests.JenisPencarian == "perihal" {
			query.Where(`pd.perihal like ?`, "%"+requests.Keyword+"%")
		} else if requests.JenisPencarian == "nomor_surat" {
			query.Where(`pd.nomorSurat like ?`, "%"+requests.Keyword+"%")
		} else if requests.JenisPencarian == "penerima" {
			query.Where(`db.BRDESC like ?`, "%"+requests.Keyword+"%")
		} else if requests.JenisPencarian == "tanggal" {
			query.Where(`date(pd.executed) = ?`, requests.Keyword)
		}
	}

	if requests.Status != "0" && requests.Status != "" && requests.Status != "all" {
		query.Where(`pd.statusMCS = ?`, requests.Status)
	}

	var count int64
	query.Count(&count)

	totalRow = int(count)

	if requests.Limit != 0 {
		query = query.Limit(requests.Limit)
	}

	if requests.Offset != 0 {
		query = query.Offset(requests.Offset)
	}

	query.Order("createdAt desc")

	err = query.Scan(&responses).Error

	result := float64(totalRow) / float64(requests.Limit)
	resultFinal := int(math.Ceil(result))

	if err != nil {
		p.logger.Zap.Error(err)
		return responses, 0, 0, err
	}

	return responses, resultFinal, totalRow, err
}

// ApprovalList implements PelaporanDefinition.
func (p PelaporanRepository) ApprovalList(requests models.DraftListRequest) (responses []models.DraftListResponse, totalRow int, totalData int, err error) {
	db := p.db.DB

	query := db.Table(`pelaporan_drafts pd`).
		Select(`pd.id, 
			pd.perihal "perihal",
			db.BRDESC "tujuan",
			pd.nomorSurat "nomor_surat",
			DATE_FORMAT(pd.executed,"%d %b %Y") "tanggal_surat",
			pd.responseStatus "response_status",
			CASE 
				WHEN pd.statusMCS = "1" THEN "Need Approve"
				WHEN pd.statusMCS = "2" AND pd.status = "Open" THEN "Proses Kirim"
				WHEN pd.statusMCS = "2" AND pd.status = "Executed" THEN "Berhasil Dikirim"  
				WHEN pd.statusMCS = "3" AND pd.status = "Rejected" THEN "Surat Ditolak"
				WHEN pd.statusMCS = "4" AND pd.status = "Fail" THEN "Gagal Dikirim"
				ELSE "Draft"
			END as "status_terakhir"`).
		Joins(`join dwh_branch db on pd.branchCodePenerima = LPAD(db.BRANCH,5,"0")`).
		Where(`pd.posisiApprover = ?`, requests.PnMaker)

	if requests.Filter {
		if requests.JenisPencarian == "perihal" {
			query.Where(`pd.perihal like ?`, "%"+requests.Keyword+"%")
		} else if requests.JenisPencarian == "nomor_surat" {
			query.Where(`pd.nomorSurat like ?`, "%"+requests.Keyword+"%")
		} else if requests.JenisPencarian == "penerima" {
			query.Where(`db.BRDESC like ?`, "%"+requests.Keyword+"%")
		} else if requests.JenisPencarian == "tanggal" {
			query.Where(`date(pd.executed) = ?`, requests.Keyword)
		}
		// query.Where(`pd.pnApprover like ? `, "%"+requests.PnMaker+"%")
	}
	// else {
	// 	query.Where(`pd.posisiApprover = ? `, requests.PnMaker)
	// }

	if requests.Status != "0" && requests.Status != "" && requests.Status != "all" {
		query.Where(`pd.statusMCS = ?`, requests.Status)
	}

	query.Where(`pd.statusMCS != '4'`)

	var count int64
	query.Count(&count)

	totalRow = int(count)

	if requests.Limit != 0 {
		query = query.Limit(requests.Limit)
	}

	if requests.Offset != 0 {
		query = query.Offset(requests.Offset)
	}

	query.Order("createdAt desc")

	err = query.Scan(&responses).Error

	result := float64(totalRow) / float64(requests.Limit)
	resultFinal := int(math.Ceil(result))

	if err != nil {
		p.logger.Zap.Error(err)
		return responses, 0, 0, err
	}

	return responses, resultFinal, totalRow, err
}

// DraftDetail implements PelaporanDefinition.
func (p PelaporanRepository) DraftDetail(id int64) (responses models.SuratDetail, err error) {
	db := p.db.DB

	query := db.Table(`pelaporan_drafts pd`).
		Select(`pd.id,
			pd.nomorSurat "nomor_surat",
			db.BRDESC "penerima",
			pd.kepadaYth "kepada_yth",
			pd.perihal "perihal",
			DATE_FORMAT(pd.executed,"%d %M %Y") "tanggal_surat",
			pd.isiSurat "isi_surat",
			pd.pnApprover "pn_approver",
			pd.posisiApprover "posisi_approver",
			pd.statusMCS "status_terakhir",
			pr.penolak "penolak",
	        pr.catatan "catatan",
			DATE_FORMAT(pr.tanggal_tolak,"%d %m %Y - %H:%i") "tanggal_tolak"
		`).
		Joins(`join dwh_branch db on pd.branchCodePenerima = LPAD(db.BRANCH,5,"0")`).
		Joins(`left join pelaporan_rejected pr on pd.id = pr.id_pelaporan`).
		Where(`pd.id = ? `, id)

	err = query.Find(&responses).Error

	if err != nil {
		p.logger.Zap.Error(err)
		return responses, err
	}

	return responses, err
}

// StoreDraftPelaporan implements PelaporanDefinition.
func (p PelaporanRepository) StoreDraftPelaporan(requests *models.PelaporanDraft, tx *gorm.DB) (status bool, err error) {
	// db := p.db.DB
	// err = db.Create(&requests).Error
	// if err != nil {
	// 	p.logger.Zap.Error(err)
	// 	return false, err
	// }
	// return true, err

	err = tx.Create(&requests).Error
	if err != nil {
		p.logger.Zap.Error(err)
		return false, err
	}
	return true, err
}

// Approve implements PelaporanDefinition.
func (p PelaporanRepository) Approve(requests *models.ApproveUpdate) (status bool, err error) {
	db := p.db.DB

	err = db.Save(&requests).Error
	if err != nil {
		p.logger.Zap.Error(err)
		return false, err
	}
	return true, err
}

func (p PelaporanRepository) Reject(requests *models.PenolakanCatatan, tx *gorm.DB) (status bool, err error) {
	err = tx.Create(&requests).Error
	if err != nil {
		p.logger.Zap.Error(err)
		return false, err
	}
	return true, err
}

func (p PelaporanRepository) Delete(pelaporan *models.ApproveUpdate, rejected *models.PenolakanCatatan, tx *gorm.DB) (status bool, err error) {
	err = tx.Delete(&pelaporan).Error
	if err != nil {
		p.logger.Zap.Error(err)
		return false, err
	}

	err = tx.Where("id_pelaporan", rejected.IDPelaporan).Delete(&rejected).Error
	if err != nil {
		p.logger.Zap.Error(err)
		return false, err
	}

	return true, err
}

// GenerateLaporan implements PelaporanDefinition.
func (p PelaporanRepository) GenerateLaporan(requests models.GenerateLaporanRequest) (responses []models.GenerateLaporanResponse, total int, err error) {
	db := p.db.DB

	query := db.Table(`tasklists t`).
		Select(`t.id, 
			a.name "aktifitas", 
			t.product_name "product",
			ri.risk_issue "risk_event", 
			ROW_NUMBER() over(order by count(t.risk_issue_id) desc) "prioritas"`).
		Joins(`join tasklists_uker tu on t.id = tu.tasklist_id `).
		Joins(`join activity a on t.activity_id = a.id`).
		Joins(`join risk_issue ri on t.risk_issue_id = ri.id `).
		Where(`t.status = "Aktif"`).
		Where(`t.approval_status = "Disetujui"`).
		Where(`t.rap = 1`).
		Where(`t.created_at BETWEEN ? AND ?`, requests.StartDate, requests.EndDate).
		Where(`tu.BRANCH = ?`, requests.Branch).
		Where(`tu.MAINBR = ? `, requests.Mainbr).
		Group(`t.risk_issue_id`).Group(`t.activity_id`).Limit(5)

	var count int64
	query.Count(&count)

	total = int(count)
	err = query.Scan(&responses).Error
	if err != nil {
		p.logger.Zap.Error(err)
		return responses, 0, err
	}

	return responses, total, nil
}

// GetPimpinanUker implements PelaporanDefinition.
func (p PelaporanRepository) GetPimpinanUker(branch string) (responses []models.PenerimaSuratResponse, err error) {
	db := p.db.DB

	query := db.Table(`pa0001_eof`).
		Select(`PERNR "pernr", SNAME "sname", ORGEH "orgeh"`).
		Where(`BRANCH = ? `, branch).
		Where(`HILFM in ('014','019','057')`)

	err = query.Scan(&responses).Error

	if err != nil {
		p.logger.Zap.Error(err)
		return responses, err
	}

	return responses, err
}

// GetKodeUnik
func (p PelaporanRepository) GetKodeUnik(branch string) (responses models.KodeResponse, jumlah int64, err error) {
	db := p.db.DB

	query := db.Table(`pelaporan_drafts`).
		Select(`kodeUnikSurat "kode"`).
		Where(`branchCodePenerima = ?`, branch).
		Where(`DATE(createdAt) = CURDATE()`).Order("id desc").Limit(1)

	err = query.Scan(&responses).Error

	query.Count(&jumlah)

	if err != nil {
		p.logger.Zap.Error(err)
		return responses, 0, err
	}

	return responses, jumlah, err
}

func (p PelaporanRepository) GetNamaSigner(pernr string) (responses models.Signer, total int, err error) {
	db := p.db.DB

	query := db.Table(`pa0001_eof pe `).
		Select(`PERNR "pn_signer", 
				SNAME "nama_signer",
				WERKS_TX "tempat",
				HTEXT "jabatan"`).
		Where(`PERNR = ? `, pernr)

	var count int64
	query.Count(&count)

	total = int(count)

	err = query.Find(&responses).Error
	if err != nil {
		p.logger.Zap.Error(err)
		return responses, 0, err
	}

	return responses, total, err
}

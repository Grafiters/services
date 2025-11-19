package pelaporan

import (
	"errors"
	"strconv"
	"strings"

	"gitlab.com/golang-package-library/logger"

	"riskmanagement/lib"
	models "riskmanagement/models/pelaporan"
	repository "riskmanagement/repository/pelaporan"
)

type PelaporanServicesDefinition interface {
	Generate(request models.DraftSuratRequest) (status bool, err error)
	GenerateIsiSuratRap(request models.SuratDraftRequest) (responses models.SuratDraftResponse, err error)

	GetDraftList(request models.DraftListRequest) (responses []models.DraftListResponse, pagination lib.Pagination, err error)
	GetApprovalList(request models.DraftListRequest) (responses []models.DraftListResponse, pagination lib.Pagination, err error)
	GetDraftDetail(id int64) (responses models.SuratDetailResponse, err error)

	Approve(request models.ApprovalRequest) (status bool, err error)
	Reject(request models.PenolakanCatatan) (status bool, err error)
	Delete(request models.ApprovalRequest) (status bool, err error)
}

type PelaporanServices struct {
	db         lib.Database
	logger     logger.Logger
	repository repository.PelaporanDefinition
}

func NewPelaporanServices(
	db lib.Database,
	logger logger.Logger,
	repository repository.PelaporanDefinition,
) PelaporanServicesDefinition {
	return PelaporanServices{
		db:         db,
		logger:     logger,
		repository: repository,
	}
}

// GetDraftList implements PelaporanServicesDefinition.
func (p PelaporanServices) GetDraftList(request models.DraftListRequest) (responses []models.DraftListResponse, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort

	data, totalRow, totalData, err := p.repository.DraftList(request)
	if err != nil {
		p.logger.Zap.Error(err)
		return responses, pagination, err
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRow, totalData)

	return data, pagination, nil
}

// GetApprovalList implements PelaporanServicesDefinition.
func (p PelaporanServices) GetApprovalList(request models.DraftListRequest) (responses []models.DraftListResponse, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort

	data, totalRow, totalData, err := p.repository.ApprovalList(request)
	if err != nil {
		p.logger.Zap.Error(err)
		return responses, pagination, err
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRow, totalData)

	return data, pagination, nil
}

// GetDraftDetail implements PelaporanServicesDefinition.
func (p PelaporanServices) GetDraftDetail(id int64) (responses models.SuratDetailResponse, err error) {

	dataSurat, err := p.repository.DraftDetail(id)
	if err != nil {
		p.logger.Zap.Error(err)
		return responses, err
	}

	_signer := strings.Split(dataSurat.PnApprover, "~")

	var Signer []models.Signer
	for _, value := range _signer {
		dataSigner, _, err := p.repository.GetNamaSigner(value)
		if err != nil {
			p.logger.Zap.Error(err)
			return responses, err
		}

		Signer = append(Signer, models.Signer{
			NamaSigner: dataSigner.NamaSigner,
			PnSigner:   dataSigner.PnSigner,
			Tempat:     dataSigner.Tempat,
			Jabatan:    dataSigner.Jabatan,
		})
	}

	responses = models.SuratDetailResponse{
		ID:             dataSurat.ID,
		NomorSurat:     dataSurat.NomorSurat,
		KepadaYth:      "Kepala/Pimpinan " + dataSurat.Penerima,
		Pengirim:       Signer[0].Tempat,
		TanggalSurat:   dataSurat.TanggalSurat,
		Perihal:        dataSurat.Perihal,
		IsiSurat:       dataSurat.IsiSurat,
		Signer:         Signer,
		PosisiApprover: dataSurat.PosisiApprover,
		StatusTerakhir: dataSurat.StatusTerakhir,
		Penolak:        dataSurat.Penolak,
		Catatan:        dataSurat.Catatan,
		TanggalTolak:   dataSurat.TanggalTolak,
	}

	return responses, err
}

// Generate implements PelaporanServicesDefinition.
func (p PelaporanServices) Generate(request models.DraftSuratRequest) (status bool, err error) {
	// timeNow := lib.GetTimeNow("timestime")
	// hari := lib.GetTimeNow("date2")

	tx := p.db.DB.Begin()

	timeNow := lib.GetTimeNow("timestime")
	hari := lib.GetTimeNow("date2")

	StatusApprover := "Signer"
	PnApprover := request.Signer[0].PnSigner

	//cek signer
	_, total, err := p.repository.GetNamaSigner(PnApprover)
	if err != nil {
		tx.Rollback()
		p.logger.Zap.Error(err)
		return false, err
	}

	if total == 0 {
		err1 := errors.New("Signer tidak ditemukan, silahkan cek kembali")
		tx.Rollback()
		p.logger.Zap.Error(err)
		return false, err1
	}

	//jika 2 signer
	if len(request.Signer) > 1 {
		StatusApprover = "Signer~Signer"
		PnApprover = request.Signer[0].PnSigner + "~" + request.Signer[1].PnSigner

		for _, value := range request.Signer {
			//cek signer
			_, total, err := p.repository.GetNamaSigner(value.PnSigner)
			if err != nil {
				tx.Rollback()
				p.logger.Zap.Error(err)
				return false, err
			}

			if total == 0 {
				err1 := errors.New("Signer tidak ditemukan, silahkan cek kembali.")
				tx.Rollback()
				p.logger.Zap.Error(err)
				return false, err1
			}
		}
	}

	if len(request.Uker) != 0 {

		for _, value := range request.Uker {

			var branch string = strings.Join([]string{strings.Repeat("0", 5-len(value.BRANCH)), value.BRANCH}, "")

			// penerima, err := p.repository.GetPimpinanUker(branch)
			// if err != nil {
			// 	p.logger.Zap.Error(err)
			// 	return false, err
			// }

			responseKode, jumlah, err := p.repository.GetKodeUnik(branch)
			if err != nil {
				tx.Rollback()
				p.logger.Zap.Error(err)
				return false, err
			}

			var kode_unik string
			if jumlah == 0 {
				nomor := "1"
				var nomor_urut string = strings.Join([]string{strings.Repeat("0", 5-len(nomor)), nomor}, "")
				kode_unik = "rap." + hari + "." + branch + "." + nomor_urut
			} else {
				nomor := strings.Split(responseKode.Kode, ".")

				nomor_terakhir, err := strconv.Atoi(nomor[3])
				// nomor_terakhir := nomor[3] + 1
				if err != nil {
					p.logger.Zap.Error(err)
				}

				nomor_terakhir = nomor_terakhir + 1
				// p.logger.Zap.Info(nomor_terakhir)
				var nomor_urut string = strings.Join([]string{strings.Repeat("0", 5-len(strconv.Itoa(nomor_terakhir))), strconv.Itoa(nomor_terakhir)}, "")
				kode_unik = "rap." + hari + "." + branch + "." + nomor_urut
			}

			// nomor := "2"
			// var nomor_urut string = strings.Join([]string{strings.Repeat("0", 5-len(nomor)), nomor}, "")
			// kode_unik := "rap." + hari + "." + branch + "." + nomor_urut

			//generate isi

			tahun := strconv.Itoa(int(request.Tahun))
			requestIsiSurat := models.SuratDraftRequest{
				Kanca:     value.MBDESC,
				KodeKanca: value.MAINBR,
				Uker:      value.BRDESC,
				KodeUker:  value.BRANCH,
				Semester:  request.Semester,
				Tahun:     tahun,
			}

			// perihal := "Laporan Risk Assesment Plan Semester " + request.Semester + " Tahun " + tahun
			isi_surat, err := p.GenerateIsiSuratRap(requestIsiSurat)
			if err != nil {
				tx.Rollback()
				p.logger.Zap.Error(err)
				return false, err
			}

			// if len(penerima) != 0 {
			// responses = append(responses, models.PelaporanDraft{
			// 	IdTemplate:         "3",
			// 	JenisSurat:         "RAP",
			// 	KodeUnikSurat:      kode_unik,
			// 	BranchCodePenerima: branch,
			// 	OrgehPenerima:      penerima[0].Orgeh,
			// 	BranchCodeTindasan: "",
			// 	OrgehTindasan:      "",
			// 	PnTindasan:         "",
			// 	PnPenerima:         penerima[0].Pernr,
			// 	KodeSurat:          "",
			// 	Kerahasiaan:        request.Kerahasiaan,
			// 	Kesegeraan:         request.Kerahasiaan,
			// 	KepadaYth:          "",
			// 	Perihal:            request.Perihal,
			// 	Semester:           request.Semester,
			// 	Tahun:              request.Tahun,
			// 	IsiSurat:           isi_surat.IsiSurat,
			// 	IdMaker:            request.PnMaker,
			// 	PnApprover:         PnApprover,
			// 	StatusApprover:     StatusApprover,
			// 	PosisiApprover:     request.Signer[0].PnSigner,
			// 	StatusMCS:          "0",
			// 	CreatedAt:          &timeNow,
			// })

			drafSurat := models.PelaporanDraft{
				IdTemplate:          "3",
				JenisSurat:          "RAP",
				KodeUnikSurat:       kode_unik,
				BranchCodePenerima:  branch,
				OrgehPenerima:       "",
				BranchCodeTindasan:  "",
				OrgehTindasan:       "",
				PnTindasan:          "",
				PnPenerima:          "",
				KodeSurat:           "RO-" + value.REGION + "/RMC",
				Kerahasiaan:         request.Kerahasiaan,
				Kesegeraan:          request.Kepentingan,
				KepadaYth:           "",
				Perihal:             request.Perihal,
				Semester:            request.Semester,
				Tahun:               tahun,
				IsiSurat:            isi_surat.IsiSurat,
				IdMaker:             request.PnMaker,
				PnApprover:          PnApprover,
				StatusApprover:      StatusApprover,
				PosisiApprover:      request.Signer[0].PnSigner,
				StatusMCS:           "1",
				CreatedAt:           &timeNow,
				Status:              "Approval",
				SuratKeluarApprover: "y",
			}

			_, err = p.repository.StoreDraftPelaporan(&drafSurat, tx)

			if err != nil {
				tx.Rollback()
				p.logger.Zap.Error(err)
				return false, err
			}

			// return responseDraft, err
			// } else {
			// 	err := errors.New("penerima tidak ditemukan")
			// 	return false, err
			// }

		}
	} else {
		tx.Rollback()
		err := errors.New("uker tidak ditemukan")
		return false, err
	}

	// responseDraft, err := p.repository.StoreDraftPelaporan(response)

	tx.Commit()
	return true, err
}

// GenerateIsiSurat
func (p PelaporanServices) GenerateIsiSuratRap(request models.SuratDraftRequest) (responses models.SuratDraftResponse, err error) {
	paragraf1 := `<p style="margin-right:12px; margin-left:6px; text-align:justify">
			<span style="font-size:12.0pt;line-height:106%;color:black;font-family:&quot;Times New Roman&quot;,serif">
			Berdasarkan hasil identifikasi risiko yang dilakukan oleh BRC/URC ` + request.Kanca + ` (` + request.KodeKanca + `) 
			pada Semester ` + request.Semester + ` Tahun ` + request.Tahun + `, 
			bersama ini kami sampaikan Risk Assesment Plan ` + request.Uker + ` (` + request.KodeUker + `) Semester ` + request.Semester + ` 
			Tahun ` + request.Tahun + ` sebagai berikut:
		</span>
	</p>
	</br>
	`

	paragraf2 := `<table class="MsoTableGrid" style="border-collapse:collapse; border:none">
	<tbody>
		<tr>
			<td style="border-bottom:1px solid black; width:37px; padding:0cm 7px 0cm 7px; border-top:1px solid black; border-right:1px solid black; border-left:1px solid black" valign="top">
			<p style="text-align:justify;font-size:12pt;line-height:normal;font-family:&quot;Times New Roman&quot;,serif">
				<b>No</b>
			</p>
			</td>
			<td style="border-bottom:1px solid black; width:85px; padding:0cm 7px 0cm 7px; border-top:1px solid black; border-right:1px solid black; border-left:none" valign="top">
			<p style="text-align:justify;font-size:12pt;line-height:normal;font-family:&quot;Times New Roman&quot;,serif">
				Aktifitas
			</p>
			</td>
			<td style="border-bottom:1px solid black; width:123px; padding:0cm 7px 0cm 7px; border-top:1px solid black; border-right:1px solid black; border-left:none" valign="top">
			<p style="text-align:justify;font-size:12pt;line-height:normal;font-family:&quot;Times New Roman&quot;,serif">
				Produk
			</p>
			</td>
			<td style="border-bottom:1px solid black; width:312px; padding:0cm 7px 0cm 7px; border-top:1px solid black; border-right:1px solid black; border-left:none" valign="top">
			<p style="text-align:center;font-size:12pt;line-height:normal;font-family:&quot;Times New Roman&quot;,serif">
				Risk Event
			</p>
			</td>
			<td style="border-bottom:1px solid black; width:66px; padding:0cm 7px 0cm 7px; border-top:1px solid black; border-right:1px solid black; border-left:none" valign="top">
			<p style="text-align:justify;font-size:12pt;line-height:normal;font-family:&quot;Times New Roman&quot;,serif">
				Prioritas
			</p>
			</td>
		</tr>`

	var StartDate, EndDate string

	if request.Semester == "1" {
		StartDate = request.Tahun + "-01-01"
		EndDate = request.Tahun + "-06-30"
	} else {
		StartDate = request.Tahun + "-07-01"
		EndDate = request.Tahun + "-12-31"
	}

	requestLaporan := models.GenerateLaporanRequest{
		Branch:    request.KodeUker,
		Mainbr:    request.KodeKanca,
		StartDate: StartDate,
		EndDate:   EndDate,
	}

	data_table, total, err := p.repository.GenerateLaporan(requestLaporan)
	if err != nil {
		p.logger.Zap.Error(err)
		return responses, err
	}

	if total == 0 {
		err1 := errors.New("Data tasklist pada " + request.Uker + " semester " + request.Semester + " " + request.Tahun + " tidak ditemukan")
		return responses, err1
	}

	var isi []string
	var isi_tabel string
	var nomor int64
	nomor = 1
	for _, item := range data_table {
		isi = append(isi, `<tr>
							<td style="border-bottom:1px solid black; width:37px; padding:0cm 7px 0cm 7px; border-top:none; border-right:1px solid black; border-left:1px solid black" valign="top">
								<p align="center" style="text-align:justify;font-size:11pt;line-height:normal;font-family:&quot;Times New Roman&quot;,serif">
									<b>
										`+strconv.Itoa(int(nomor))+`
									</b>
								</p>
							</td>
							<td style="border-bottom:1px solid black; width:85px; padding:0cm 7px 0cm 7px; border-top:none; border-right:1px solid black; border-left:none" valign="top">
								<p align="center" style="text-align:justify;font-size:11pt;line-height:normal;font-family:&quot;Times New Roman&quot;,serif">
									
												`+item.Aktifitas+`
											
								</p>
							</td>
							<td style="border-bottom:1px solid black; width:123px; padding:0cm 7px 0cm 7px; border-top:none; border-right:1px solid black; border-left:none" valign="top">
								<p align="center" style="text-align:justify;font-size:11pt;line-height:normal;font-family:&quot;Times New Roman&quot;,serif">
									
													`+item.Product+`
												
								</p>
							</td>
							<td style="border-bottom:1px solid black; width:312px; padding:0cm 7px 0cm 7px; border-top:none; border-right:1px solid black; border-left:none" valign="top">
								<p align="center" style="text-align:justify;font-size:11pt;line-height:normal;font-family:&quot;Times New Roman&quot;,serif">
									
													`+item.RiskEvent+`
											
								</p>
							</td>
							<td style="border-bottom:1px solid black; width:66px; padding:0cm 7px 0cm 7px; border-top:none; border-right:1px solid black; border-left:none" valign="top">
								<p align="center" style="text-align:justify;font-size:11pt;line-height:normal;font-family:&quot;Times New Roman&quot;,serif">
									
													`+item.Prioritas+`
													
								</p>
							</td>
						</tr>`)
		// isi_tabel = isi_tabel + isi_tabel
		nomor++
	}

	isi_tabel = strings.Join(isi, "")
	paragraf2 = paragraf2 + isi_tabel + `
		</tbody>
	</table>`

	paragraf3 := `</br>
					<p style="margin-right:12px; margin-left:6px; text-align:justify;font-size:12.0pt;line-height:106%;color:black;font-family:&quot;Times New Roman&quot;,serif">
							Demikian kami sampaikan, atas perhatian dan kerjasamanya kamu ucapkan terima kasih.
					</p>`

	isi_surat := paragraf1 + paragraf2 + paragraf3

	hasil_generate := models.SuratDraftResponse{
		Perihal:  request.Perihal,
		IsiSurat: isi_surat,
	}

	return hasil_generate, err
}

// Approve
func (p PelaporanServices) Approve(request models.ApprovalRequest) (status bool, err error) {
	dataSurat, err := p.repository.DraftDetail(request.ID)
	if err != nil {
		p.logger.Zap.Error(err)
		return false, err
	}

	_signer := strings.Split(dataSurat.PnApprover, "~")

	var Signer []models.Signer
	for _, value := range _signer {
		dataSigner, _, err := p.repository.GetNamaSigner(value)
		if err != nil {
			p.logger.Zap.Error(err)
			return false, err
		}

		Signer = append(Signer, models.Signer{
			NamaSigner: dataSigner.NamaSigner,
			PnSigner:   dataSigner.PnSigner,
			Tempat:     dataSigner.Tempat,
			Jabatan:    dataSigner.Jabatan,
		})
	}

	if dataSurat.PosisiApprover != request.PnApproval {
		_error := errors.New("PN Approval Tidak Sama")
		return false, _error
	}

	pnApprover := request.PnApproval
	statusMcs := "2"
	// responseStatus := "Sedang dikirim"
	_status := "Open"

	if len(Signer) > 1 {
		if request.PnApproval == Signer[0].PnSigner {
			pnApprover = Signer[1].PnSigner
			statusMcs = "1"
			_status = "Approval"
		}
	}

	req := models.ApproveUpdate{
		ID:             request.ID,
		PosisiApprover: pnApprover,
		StatusMCS:      statusMcs,
		// ResponseStatus: responseStatus,
		Status: _status,
	}

	status, err = p.repository.Approve(&req)
	if err != nil {
		p.logger.Zap.Error(err)
		return false, err
	}

	return status, err
}

func (p PelaporanServices) Reject(request models.PenolakanCatatan) (status bool, err error) {
	// timeNow := lib.GetTimeNow("timestime")

	tx := p.db.DB.Begin()
	timeNow := lib.GetTimeNow("timestime")

	dataSurat, err := p.repository.DraftDetail(request.IDPelaporan)
	if err != nil {
		p.logger.Zap.Error(err)
		return false, err
	}

	_signer := strings.Split(request.Penolak, " - ")

	if dataSurat.PosisiApprover != _signer[0] {
		tx.Rollback()
		_error := errors.New("PN Penolak Tidak Sama")
		return false, _error
	}

	req := models.ApproveUpdate{
		ID:        request.IDPelaporan,
		StatusMCS: "3",
		Status:    "Rejected",
	}

	status, err = p.repository.Approve(&req)
	if err != nil {
		tx.Rollback()
		p.logger.Zap.Error(err)
		return false, err
	}

	catatan := models.PenolakanCatatan{
		IDPelaporan:  request.IDPelaporan,
		Penolak:      request.Penolak,
		Catatan:      request.Catatan,
		TanggalTolak: &timeNow,
	}
	_, err = p.repository.Reject(&catatan, tx)

	if err != nil {
		tx.Rollback()
		p.logger.Zap.Error(err)
		return false, err
	}

	tx.Commit()
	return status, err
}

func (p PelaporanServices) Delete(request models.ApprovalRequest) (status bool, err error) {
	tx := p.db.DB.Begin()

	pelaporan := models.ApproveUpdate{
		ID: request.ID,
	}

	penolakan := models.PenolakanCatatan{
		IDPelaporan: request.ID,
	}

	status, err = p.repository.Delete(&pelaporan, &penolakan, tx)

	if err != nil {
		p.logger.Zap.Error(err)
		return false, err
	}

	tx.Commit()
	return status, err
}

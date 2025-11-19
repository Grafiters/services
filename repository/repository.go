package repository

import (
	activity "riskmanagement/repository/activity"
	admin_setting "riskmanagement/repository/admin_setting"
	aplikasi "riskmanagement/repository/aplikasi"
	audittrail "riskmanagement/repository/audittrail"
	briefing "riskmanagement/repository/briefing"
	coaching "riskmanagement/repository/coaching"
	"riskmanagement/repository/common"
	datatematik "riskmanagement/repository/data_tematik"
	download "riskmanagement/repository/download"
	eventtype1 "riskmanagement/repository/eventtypelv1"
	eventtype2 "riskmanagement/repository/eventtypelv2"
	eventtype3 "riskmanagement/repository/eventtypelv3"
	files "riskmanagement/repository/files"
	incident "riskmanagement/repository/incident"
	JenisTask "riskmanagement/repository/jenistask"
	laporan "riskmanagement/repository/laporan"
	linibisnis1 "riskmanagement/repository/linibisnislv1"
	linibisnis2 "riskmanagement/repository/linibisnislv2"
	linibisnis3 "riskmanagement/repository/linibisnislv3"
	linkcage "riskmanagement/repository/linkcage"
	majorproses "riskmanagement/repository/majorproses"
	managementUser "riskmanagement/repository/managementuser"
	materi "riskmanagement/repository/materi"
	megaproses "riskmanagement/repository/megaproses"
	menu "riskmanagement/repository/menu"
	monitoring "riskmanagement/repository/monitoring"
	mstKriteria "riskmanagement/repository/mstkriteria"
	mstrole "riskmanagement/repository/mstrole"
	msUker "riskmanagement/repository/msuker"
	"riskmanagement/repository/notifikasi"
	pekerja "riskmanagement/repository/pekerja"
	pelaporan "riskmanagement/repository/pelaporan"
	penyebabkejadianlv3 "riskmanagement/repository/penyebabkejadianlv3"
	pgsUser "riskmanagement/repository/pgsuser"
	product "riskmanagement/repository/product"
	questionner "riskmanagement/repository/questionner"
	riskcontrol "riskmanagement/repository/riskcontrol"
	riskIndicator "riskmanagement/repository/riskindicator"
	riskIssue "riskmanagement/repository/riskissue"
	riskType "riskmanagement/repository/risktype"
	subactivity "riskmanagement/repository/subactivity"
	subIncident "riskmanagement/repository/subincident"
	submajorproses "riskmanagement/repository/submajorproses"
	taskassignment "riskmanagement/repository/taskassignment"
	tasklists "riskmanagement/repository/tasklists"
	ukerkololaan "riskmanagement/repository/ukerkelolaan"
	unitKerja "riskmanagement/repository/unitkerja"
	uploaddata "riskmanagement/repository/uploaddata"
	user "riskmanagement/repository/user"
	verifikasi "riskmanagement/repository/verifikasi"
	verifikasirealisasi "riskmanagement/repository/verifikasirealisasi"
	verifikasiReportRealisasi "riskmanagement/repository/verifikasireportrealisasi"

	notificationGlobal "riskmanagement/repository/notification_global"
	organisasi "riskmanagement/repository/organisasi"

	"go.uber.org/fx"
)

// Module exports dependency
var Module = fx.Options(
	fx.Provide(user.NewUserRepository),
	fx.Provide(activity.NewActivityRepository),
	fx.Provide(admin_setting.NewAdminSettingRepository),
	fx.Provide(subactivity.NewSubActivityRepository),
	fx.Provide(product.NewProductRepository),
	fx.Provide(riskIssue.NewRiskIssueRepository),
	fx.Provide(riskIndicator.NewRiskIndicatorRepository),
	fx.Provide(riskIndicator.NewLampiranIndicatorRepository),
	fx.Provide(riskIndicator.NewMapThresholdRepository),
	fx.Provide(riskIndicator.NewMapRiskIssueRepository),
	fx.Provide(incident.NewIncidentRepository),
	fx.Provide(subIncident.NewSubIncidentRepository),
	fx.Provide(riskType.NewRiskTypeRepository),
	fx.Provide(tasklists.NewTasklistsRepository),
	fx.Provide(tasklists.NewTasklistsActivityRepository),
	fx.Provide(tasklists.NewTasklistsLampiranRepository),
	fx.Provide(tasklists.NewTasklistsDataAnomaliKRIDRepository),
	fx.Provide(tasklists.NewTasklistsUkerRepository),
	fx.Provide(tasklists.NewTasklistsDoneHistoryRepository),
	fx.Provide(tasklists.NewTasklistsDailyRepository),
	fx.Provide(unitKerja.NewUnitKerjaRepository),
	fx.Provide(briefing.NewBriefingRepository),
	fx.Provide(briefing.NewBriefingMateriRepository),
	fx.Provide(briefing.NewBriefingMapPesertaRepository),
	fx.Provide(files.NewFilesRepository),
	fx.Provide(materi.NewMateriRepository),
	fx.Provide(coaching.NewCoachingRepository),
	fx.Provide(coaching.NewCoachingActivityRepository),
	fx.Provide(coaching.NewCoachingMapPesertaRepository),
	fx.Provide(verifikasi.NewVerfikasiRepository),
	fx.Provide(verifikasi.NewVerifikasiAnomaliRepository),
	fx.Provide(verifikasi.NewVerifikasiPICRepository),
	fx.Provide(verifikasi.NewVerfikasiFilesRepository),
	fx.Provide(verifikasi.NewVerifikasiRiskControlRepository),
	fx.Provide(verifikasi.NewVerifikasiAnomaliDataKRIDRepository),
	fx.Provide(verifikasi.NewVerifikatiQuestionnerRepository),
	fx.Provide(verifikasi.NewVerifikasiPenyababKejadianRepository),
	fx.Provide(verifikasi.NewVerifikasiUsulanPerbaikanRepository),
	fx.Provide(verifikasi.NewVerifikasiDataTematikRepository),
	fx.Provide(riskcontrol.NewRiskControlRepository),
	fx.Provide(aplikasi.NewAplikasiRepository),
	fx.Provide(linibisnis1.NewLiniBisnisLv1Repository),
	fx.Provide(linibisnis2.NewLiniBisnisLv2Repository),
	fx.Provide(linibisnis3.NewLiniBisnisLv3Repository),
	fx.Provide(eventtype1.NewEventTypeLv1Repository),
	fx.Provide(eventtype2.NewEventTypeLv2Repository),
	fx.Provide(eventtype3.NewEventTypeLv3Repository),
	fx.Provide(megaproses.NewMegaprosesRepository),
	fx.Provide(majorproses.NewMajorProsesRepository),
	fx.Provide(submajorproses.NewSubMajorProsesRepository),
	fx.Provide(penyebabkejadianlv3.NewPenyebabKejadianLv3Repository),
	fx.Provide(riskIssue.NewMapAktifitasRepository),
	fx.Provide(riskIssue.NewMapEventRepository),
	fx.Provide(riskIssue.NewMapKejadianRepository),
	fx.Provide(riskIssue.NewMapLiniBisnisRepository),
	fx.Provide(riskIssue.NewMapProductRepository),
	fx.Provide(riskIssue.NewMapProsesRepository),
	fx.Provide(riskIssue.NewMapControlRepository),
	fx.Provide(riskIssue.NewMapIndicatorRepository),
	fx.Provide(msUker.NewMsUkerRepository),
	fx.Provide(mstKriteria.NewMstKriteriaRepository),
	fx.Provide(pgsUser.NewPgsUserRepository),
	fx.Provide(pgsUser.NewPgsUserApprovalRepository),
	fx.Provide(managementUser.NewManagementUserRepository),
	fx.Provide(managementUser.NewMapMenuRepository),
	fx.Provide(mstrole.NewMstRoleRepository),
	fx.Provide(mstrole.NewMstRoleMenuRepository),
	fx.Provide(download.NewDownloadRepository),
	fx.Provide(ukerkololaan.NewUkerKelolaanRepository),
	fx.Provide(audittrail.NewAuditTrailRepository),
	fx.Provide(questionner.NewQuestionnerRepository),
	fx.Provide(laporan.NewLaporanRepository),
	fx.Provide(common.NewCommonRepository),
	fx.Provide(notifikasi.NewNotifiksaiRepository),
	fx.Provide(monitoring.NewMonitoringRepository),
	fx.Provide(pelaporan.NewPelaporanRepository),
	fx.Provide(uploaddata.NewUploadDataRepository),
	fx.Provide(datatematik.NewDataTematikRepository),
	fx.Provide(verifikasirealisasi.NewVerfikasiRealisasiRepository),
	fx.Provide(verifikasiReportRealisasi.NewVerfikasiReportRealisasiRepository),
	fx.Provide(taskassignment.NewTaskAssignmentsRepository),
	fx.Provide(JenisTask.NewJenisTaskRepository),
	fx.Provide(pekerja.NewPekerjaRepository),
	fx.Provide(menu.NewMenuRepository),
	fx.Provide(linkcage.NewLinkcageRepository),
	fx.Provide(linkcage.NewLinkcageImageRepository),

	// notification
	fx.Provide(notificationGlobal.NewNotificationGlobalRepository),
	fx.Provide(organisasi.NewOrganisasiRepository),
)

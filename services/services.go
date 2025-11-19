package services

import (
	activity "riskmanagement/services/activity"
	admin_setting "riskmanagement/services/admin_setting"
	aplikasi "riskmanagement/services/aplikasi"
	audittrail "riskmanagement/services/audittrail"
	auth "riskmanagement/services/auth"
	briefing "riskmanagement/services/briefing"
	coaching "riskmanagement/services/coaching"
	"riskmanagement/services/common"
	datatematik "riskmanagement/services/data_tematik"
	download "riskmanagement/services/download"
	eventtype1 "riskmanagement/services/eventtypelv1"
	eventtype2 "riskmanagement/services/eventtypelv2"
	eventtype3 "riskmanagement/services/eventtypelv3"
	filemanager "riskmanagement/services/filemanager"
	getclientip "riskmanagement/services/getclientip"
	incident "riskmanagement/services/incident"
	jenisTask "riskmanagement/services/jenistask"
	krid "riskmanagement/services/krid"
	laporan "riskmanagement/services/laporan"
	linibisnis1 "riskmanagement/services/linibisnislv1"
	linibisnis2 "riskmanagement/services/linibisnislv2"
	linibisnis3 "riskmanagement/services/linibisnislv3"
	linkcage "riskmanagement/services/linkcage"
	majorproses "riskmanagement/services/majorproses"
	managementUser "riskmanagement/services/managementuser"
	materi "riskmanagement/services/materi"
	mcs "riskmanagement/services/mcs"
	megaproses "riskmanagement/services/megaproses"
	menu "riskmanagement/services/menu"
	monitoring "riskmanagement/services/monitoring"
	mstKriteria "riskmanagement/services/mstkriteria"
	mstrole "riskmanagement/services/mstrole"
	msUker "riskmanagement/services/msuker"
	"riskmanagement/services/notifikasi"
	pekerja "riskmanagement/services/pekerja"
	pelaporan "riskmanagement/services/pelaporan"
	penyebabKejadianLv3 "riskmanagement/services/penyebabkejadianlv3"
	pgsUser "riskmanagement/services/pgsuser"
	product "riskmanagement/services/product"
	questionner "riskmanagement/services/questionner"
	riskControl "riskmanagement/services/riskcontrol"
	riskIndicator "riskmanagement/services/riskindicator"
	riskIssue "riskmanagement/services/riskissue"
	riskType "riskmanagement/services/risktype"
	subactivity "riskmanagement/services/subactivity"
	subincident "riskmanagement/services/subincident"
	submajorproses "riskmanagement/services/submajorproses"
	taskassignment "riskmanagement/services/taskassignment"
	tasklists "riskmanagement/services/tasklists"
	ukerkelolaan "riskmanagement/services/ukerkelolaan"
	unitKerja "riskmanagement/services/unitkerja"
	uploaddata "riskmanagement/services/uploaddata"
	user "riskmanagement/services/user"
	verifikasi "riskmanagement/services/verifikasi"
	verifikasiRealisasi "riskmanagement/services/verifikasirealisasi"
	verifikasiReportRealisasi "riskmanagement/services/verifikasireportrealisasi"

	// notification global
	notificationGlobal "riskmanagement/services/notification_global"

	// organisasi
	organisasiService "riskmanagement/services/organisasi"

	// microservice
	realisasiService "riskmanagement/services/realisasiservice"

	mq "riskmanagement/services/mq"

	"go.uber.org/fx"
)

// Module exports services present
var Module = fx.Options(
	fx.Provide(user.NewUserService),
	fx.Provide(auth.NewJWTAuthService),
	fx.Provide(activity.NewActivityService),
	fx.Provide(admin_setting.NewAdminSettingService),
	fx.Provide(subactivity.NewSubActivityService),
	fx.Provide(product.NewProductService),
	fx.Provide(riskIssue.NewRiskIssueService),
	fx.Provide(riskIndicator.NewRiskIndicatorService),
	fx.Provide(incident.NewIncidentService),
	fx.Provide(subincident.NewSubIncidentService),
	fx.Provide(riskType.NewRiskTypeService),
	fx.Provide(tasklists.NewTasklistsService),
	fx.Provide(unitKerja.NewUnitKerjaService),
	fx.Provide(briefing.NewBriefingService),
	fx.Provide(materi.NewMateriService),
	fx.Provide(coaching.NewCoachingService),
	fx.Provide(verifikasi.NewVerifikasiService),
	fx.Provide(riskControl.NewRiskControService),
	fx.Provide(aplikasi.NewAplikasiService),
	fx.Provide(mcs.NewMcsService),
	fx.Provide(filemanager.NewFileManagerService),
	fx.Provide(linibisnis1.NewLiniBisnisLv1Service),
	fx.Provide(linibisnis2.NewLiniBisnisLv2Service),
	fx.Provide(linibisnis3.NewLiniBisnisLv3Service),
	fx.Provide(eventtype1.NewEventTypeLv1Service),
	fx.Provide(eventtype2.NewEventTypeLv2Service),
	fx.Provide(eventtype3.NewEventTypeLv3Service),
	fx.Provide(megaproses.NewMegaProsesService),
	fx.Provide(majorproses.NewMajorProsesService),
	fx.Provide(submajorproses.NewSubMajorProsesService),
	fx.Provide(penyebabKejadianLv3.NewPenyebabKejadianLv3Service),
	fx.Provide(msUker.NewMsUkerService),
	fx.Provide(mstKriteria.NewMstKriteriaService),
	fx.Provide(pgsUser.NewPgsUserService),
	fx.Provide(managementUser.NewManagementUserService),
	fx.Provide(mstrole.NewMstRoleService),
	fx.Provide(krid.NewKridService),
	fx.Provide(download.NewDownloadService),
	fx.Provide(ukerkelolaan.NewUkerKelolaanService),
	fx.Provide(getclientip.NewClientIPService),
	fx.Provide(audittrail.NewAuditTrailService),
	fx.Provide(questionner.NewQuestionnerService),
	fx.Provide(laporan.NewLaporanServices),
	fx.Provide(common.NewCommonService),
	fx.Provide(notifikasi.NewNotifikasiServices),
	fx.Provide(monitoring.NewMontioringServices),
	fx.Provide(pelaporan.NewPelaporanServices),
	fx.Provide(uploaddata.NewUploadDataService),
	fx.Provide(datatematik.NewDataTematikService),
	fx.Provide(verifikasiRealisasi.NewVerifikasiRealisasiService),
	fx.Provide(verifikasiReportRealisasi.NewVerifikasiReportRealisasiService),
	fx.Provide(taskassignment.NewTaskAssignmentService),
	fx.Provide(jenisTask.NewJenisTaskService),
	fx.Provide(pekerja.NewPekerjaService),
	fx.Provide(menu.NewMenuService),

	// notification global
	fx.Provide(notificationGlobal.NewNotificationGlobalService),

	// organisasi
	fx.Provide(organisasiService.NewOrganisasiService),
	// microservice
	fx.Provide(realisasiService.NewRealisasiService),
	fx.Provide(linkcage.NewLinkcageService),

	fx.Provide(mq.NewMQService),
)

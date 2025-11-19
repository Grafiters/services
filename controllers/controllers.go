package controllers

import (
	activity "riskmanagement/controllers/activity"
	admin_setting "riskmanagement/controllers/admin_setting"
	aplikasi "riskmanagement/controllers/aplikasi"
	audittrail "riskmanagement/controllers/audittrail"
	auth "riskmanagement/controllers/auth"
	briefing "riskmanagement/controllers/briefing"
	coaching "riskmanagement/controllers/coaching"
	common "riskmanagement/controllers/common"
	datatematik "riskmanagement/controllers/data_tematik"
	download "riskmanagement/controllers/download"
	eventtype1 "riskmanagement/controllers/eventtypelv1"
	eventtype2 "riskmanagement/controllers/eventtypelv2"
	eventtype3 "riskmanagement/controllers/eventtypelv3"
	filemanager "riskmanagement/controllers/filemanager"
	getclientip "riskmanagement/controllers/getclientip"
	incident "riskmanagement/controllers/incident"
	jenisTask "riskmanagement/controllers/jenistask"
	krid "riskmanagement/controllers/krid"
	laporan "riskmanagement/controllers/laporan"
	linibisnis1 "riskmanagement/controllers/linibisnislv1"
	linibisnis2 "riskmanagement/controllers/linibisnislv2"
	linibisnis3 "riskmanagement/controllers/linibisnislv3"
	majorproses "riskmanagement/controllers/majorproses"
	managementUser "riskmanagement/controllers/managementuser"
	materi "riskmanagement/controllers/materi"
	mcs "riskmanagement/controllers/mcs"
	megaproses "riskmanagement/controllers/megaproses"
	menu "riskmanagement/controllers/menu"
	monitoring "riskmanagement/controllers/monitoring"
	mstKriteria "riskmanagement/controllers/mstkriteria"
	mstrole "riskmanagement/controllers/mstrole"
	msUker "riskmanagement/controllers/msuker"
	notifikasi "riskmanagement/controllers/notifikasi"
	pekerja "riskmanagement/controllers/pekerja"
	"riskmanagement/controllers/pelaporan"
	penyebabkejadianlv3 "riskmanagement/controllers/penyebabkejadianlv3"
	pgsUser "riskmanagement/controllers/pgsuser"
	product "riskmanagement/controllers/product"
	questionner "riskmanagement/controllers/questionner"
	riskControl "riskmanagement/controllers/riskcontrol"
	riskIndicator "riskmanagement/controllers/riskindicator"
	riskIssue "riskmanagement/controllers/riskissue"
	riskType "riskmanagement/controllers/risktype"
	subactivity "riskmanagement/controllers/subactivity"
	subIncident "riskmanagement/controllers/subincident"
	submajorproses "riskmanagement/controllers/submajorproses"
	taskassignment "riskmanagement/controllers/taskassignment"
	tasklists "riskmanagement/controllers/tasklists"
	ukerkelolaan "riskmanagement/controllers/ukerkelolaan"
	unitKerja "riskmanagement/controllers/unitkerja"
	uploaddata "riskmanagement/controllers/uploaddata"
	user "riskmanagement/controllers/user"
	verifikasi "riskmanagement/controllers/verifikasi"
	verifikasiRealisasi "riskmanagement/controllers/verifikasirealisasi"
	verifikasiReportRealisasi "riskmanagement/controllers/verifikasireportrealisasi"

	// notification
	notifikasiGlobal "riskmanagement/controllers/notification_global"

	// organisasi
	organisasi "riskmanagement/controllers/organisasi"

	// microservice
	"riskmanagement/controllers/realisasicontroller"

	linkcage "riskmanagement/controllers/linkcage"

	mq "riskmanagement/controllers/mq"

	"go.uber.org/fx"
)

// Module exported for initializing application
var Module = fx.Options(
	fx.Provide(user.NewUserController),
	fx.Provide(auth.NewJWTAuthController),
	fx.Provide(activity.NewActivityController),
	fx.Provide(admin_setting.NewAdminSettingController),
	fx.Provide(subactivity.NewSubActivityController),
	fx.Provide(product.NewProductController),
	fx.Provide(riskIssue.NewRiskIssueController),
	fx.Provide(riskIndicator.NewRiskIndicatorController),
	fx.Provide(incident.NewIncidentController),
	fx.Provide(subIncident.NewSubIncidentController),
	fx.Provide(riskType.NewRiskTypeController),
	fx.Provide(tasklists.NewTasklistsController),
	fx.Provide(unitKerja.NewUnitKerjaController),
	fx.Provide(briefing.NewBriefingController),
	fx.Provide(materi.NewMateriController),
	fx.Provide(coaching.NewCoachingController),
	fx.Provide(verifikasi.NewVerifikasiController),
	fx.Provide(riskControl.NewRiskControlController),
	fx.Provide(aplikasi.NewAplikasiController),
	fx.Provide(mcs.NewMcsController),
	fx.Provide(filemanager.NewFileManagerController),
	fx.Provide(linibisnis1.NewLiniBisnisLV1Controller),
	fx.Provide(linibisnis2.NewLiniBisnisLV2Controller),
	fx.Provide(linibisnis3.NewLiniBisnisLV3Controller),
	fx.Provide(eventtype1.NewEventTypeLV1Controller),
	fx.Provide(eventtype2.NewEventTypeLV2Controller),
	fx.Provide(eventtype3.NewEventTypeLV3Controller),
	fx.Provide(megaproses.NewMegaProsesController),
	fx.Provide(majorproses.NewMajorProsesController),
	fx.Provide(submajorproses.NewSubMajorProsesController),
	fx.Provide(penyebabkejadianlv3.NewPenyebabKejadianLv3Controller),
	fx.Provide(msUker.NewMsUkerController),
	fx.Provide(mstKriteria.NewMstKriteriaController),
	fx.Provide(pgsUser.NewPgsUserController),
	fx.Provide(managementUser.NewManagementUserController),
	fx.Provide(mstrole.NewMstRoleController),
	fx.Provide(krid.NewKridController),
	fx.Provide(download.NewDownloadController),
	fx.Provide(ukerkelolaan.NewUkerKelolaanController),
	fx.Provide(getclientip.NewClientIPController),
	fx.Provide(audittrail.NewAuditTrailController),
	fx.Provide(questionner.NewQuestionnerController),
	fx.Provide(laporan.NewLaporanController),
	fx.Provide(common.NewCommonController),
	fx.Provide(notifikasi.NewNotifikasiController),
	fx.Provide(monitoring.NewMonitoringController),
	fx.Provide(pelaporan.NewPelaporanController),
	fx.Provide(uploaddata.NewUploadDataController),
	fx.Provide(datatematik.NewDataTematikController),
	fx.Provide(verifikasiRealisasi.NewVerifikasiRealisasiController),
	fx.Provide(verifikasiReportRealisasi.NewVerifikasiReportRealisasiController),
	fx.Provide(taskassignment.NewTaskAssignmentsController),
	fx.Provide(jenisTask.NewJenisTaskController),
	fx.Provide(pekerja.NewPekerjaController),
	fx.Provide(menu.NewMenuController),

	// notification
	fx.Provide(notifikasiGlobal.NewNotificationGlobalController),

	// organisasi
	fx.Provide(organisasi.NewOrganisasiController),

	// microservice
	fx.Provide(realisasicontroller.NewRealisasiController),

	fx.Provide(mq.NewMQController),
	fx.Provide(linkcage.NewLinkcageController),
)

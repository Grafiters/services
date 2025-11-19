package models

import (
	"riskmanagement/lib"

	ManagementUser "riskmanagement/models/managementuser"
)

type PgsUserRequest struct {
	ID            int64                    `json:"id"`
	PN            string                   `json:"pn"`
	NamaPekerja   string                   `json:"nama_pekerja"`
	UnitKerja     string                   `json:"unit_kerja"`
	REGION        string                   `json:"REGION"`
	RGDESC        string                   `json:"RGDESC"`
	RGNAME        string                   `json:"RGNAME"`
	MAINBR        string                   `json:"MAINBR"`
	MBDESC        string                   `json:"MBDESC"`
	MBNAME        string                   `json:"MBNAME"`
	BRANCH        string                   `json:"BRANCH"`
	BRDESC        string                   `json:"BRDESC"`
	BRNAME        string                   `json:"BRNAME"`
	JabatanPgs    string                   `json:"jabatan_pgs"`
	PeriodeAwal   string                   `json:"periode_awal"`
	PeriodeAkhir  string                   `json:"periode_akhir"`
	MakerID       string                   `json:"maker_id"`
	MakerDesc     string                   `json:"maker_desc"`
	MakerDate     *string                  `json:"maker_date"`
	LastMakerID   string                   `json:"last_maker_id"`
	LastMakerDesc string                   `json:"last_maker_desc"`
	LastMakerDate *string                  `json:"last_maker_date"`
	Status        string                   `json:"status"`
	Action        string                   `json:"action"`
	Approval      []PgsUserApprovalRequest `json:"approval"`
	CreatedAt     *string                  `json:"created_at"`
	UpdatedAt     *string                  `json:"updated_at"`
}

type Paginate struct {
	Order  string `json:"order"`
	Sort   string `json:"sort"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Page   int    `json:"page"`
	Penr   string `json:"pernr"`
}

type UpdateDelete struct {
	ID         int64   `json:"id"`
	DeleteFlag bool    `json:"delete_flag"`
	UpdatedAt  *string `json:"updated_at"`
}

type PgsUserResponses struct {
	ID            int64   `json:"id"`
	PN            string  `json:"pn"`
	NamaPekerja   string  `json:"nama_pekerja"`
	UnitKerja     string  `json:"unit_kerja"`
	REGION        string  `json:"REGION"`
	RGDESC        string  `json:"RGDESC"`
	RGNAME        string  `json:"RGNAME"`
	MAINBR        string  `json:"MAINBR"`
	MBDESC        string  `json:"MBDESC"`
	MBNAME        string  `json:"MBNAME"`
	BRANCH        string  `json:"BRANCH"`
	BRDESC        string  `json:"BRDESC"`
	BRNAME        string  `json:"BRNAME"`
	JabatanPgs    string  `json:"jabatan_pgs"`
	PeriodeAwal   string  `json:"periode_awal"`
	PeriodeAkhir  string  `json:"periode_akhir"`
	MakerID       string  `json:"maker_id"`
	MakerDesc     string  `json:"maker_desc"`
	MakerDate     *string `json:"maker_date"`
	LastMakerID   string  `json:"last_maker_id"`
	LastMakerDesc string  `json:"last_maker_desc"`
	LastMakerDate *string `json:"last_maker_date"`
	Status        string  `json:"status"`
	Action        string  `json:"action"`
	CreatedAt     *string `json:"created_at"`
	UpdatedAt     *string `json:"updated_at"`
}

type PgsUserResponseOne struct {
	ID            int64                     `json:"id"`
	PN            string                    `json:"pn"`
	NamaPekerja   string                    `json:"nama_pekerja"`
	UnitKerja     string                    `json:"unit_kerja"`
	REGION        string                    `json:"REGION"`
	RGDESC        string                    `json:"RGDESC"`
	RGNAME        string                    `json:"RGNAME"`
	MAINBR        string                    `json:"MAINBR"`
	MBDESC        string                    `json:"MBDESC"`
	MBNAME        string                    `json:"MBNAME"`
	BRANCH        string                    `json:"BRANCH"`
	BRDESC        string                    `json:"BRDESC"`
	BRNAME        string                    `json:"BRNAME"`
	JabatanPgs    string                    `json:"jabatan_pgs"`
	PeriodeAwal   string                    `json:"periode_awal"`
	PeriodeAkhir  string                    `json:"periode_akhir"`
	MakerID       string                    `json:"maker_id"`
	MakerDesc     string                    `json:"maker_desc"`
	MakerDate     *string                   `json:"maker_date"`
	LastMakerID   string                    `json:"last_maker_id"`
	LastMakerDesc string                    `json:"last_maker_desc"`
	LastMakerDate *string                   `json:"last_maker_date"`
	Status        string                    `json:"status"`
	Action        string                    `json:"action"`
	Approval      []PgsUserApprovalResponse `json:"approval"`
	CreatedAt     *string                   `json:"created_at"`
	UpdatedAt     *string                   `json:"updated_at"`
}

type ApprovalRequest struct {
	PN string `json:"pn"`
}

type PgsApprovalResponse struct {
	ID             int64  `json:"id"`
	PN             string `json:"pn"`
	NamaPekerja    string `json:"nama_pekerja"`
	UnitKerja      string `json:"unit_kerja"`
	JabatanPgs     string `json:"jabatan_pgs"`
	IDApproval     int64  `json:"id_approval"`
	Approval       string `json:"approval"`
	ApprovalDate   string `json:"approval_date"`
	ApprovalStatus bool   `json:"approval_status"`
}

type PgsApprovalResponseNull struct {
	ID             lib.NullInt64  `json:"id"`
	PN             lib.NullString `json:"pn"`
	NamaPekerja    lib.NullString `json:"nama_pekerja"`
	UnitKerja      lib.NullString `json:"unit_kerja"`
	JabatanPgs     lib.NullString `json:"jabatan_pgs"`
	IDApproval     lib.NullInt64  `json:"id_approval"`
	Approval       lib.NullString `json:"approval"`
	ApprovalDate   lib.NullString `json:"approval_date"`
	ApprovalStatus lib.NullBool   `json:"approval_status"`
}

type PgsUserRequestMaintainance struct {
	ID            int64   `json:"id"`
	PN            string  `json:"pn"`
	NamaPekerja   string  `json:"nama_pekerja"`
	REGION        string  `json:"REGION"`
	RGDESC        string  `json:"RGDESC"`
	RGNAME        string  `json:"RGNAME"`
	MAINBR        string  `json:"MAINBR"`
	MBDESC        string  `json:"MBDESC"`
	MBNAME        string  `json:"MBNAME"`
	BRANCH        string  `json:"BRANCH"`
	BRDESC        string  `json:"BRDESC"`
	BRNAME        string  `json:"BRNAME"`
	JabatanPgs    string  `json:"jabatan_pgs"`
	PeriodeAwal   string  `json:"periode_awal"`
	PeriodeAkhir  string  `json:"periode_akhir"`
	LastMakerID   string  `json:"last_maker_id"`
	LastMakerDesc string  `json:"last_maker_desc"`
	LastMakerDate *string `json:"last_maker_date"`
	Status        string  `json:"status"`
	Action        string  `json:"action"`
	UpdatedAt     *string `json:"updated_at"`
}

type PgsUserRequestUpdate struct {
	ID            int64                    `json:"id"`
	PN            string                   `json:"pn"`
	NamaPekerja   string                   `json:"nama_pekerja"`
	UnitKerja     string                   `json:"unit_kerja"`
	REGION        string                   `json:"REGION"`
	RGDESC        string                   `json:"RGDESC"`
	RGNAME        string                   `json:"RGNAME"`
	MAINBR        string                   `json:"MAINBR"`
	MBDESC        string                   `json:"MBDESC"`
	MBNAME        string                   `json:"MBNAME"`
	BRANCH        string                   `json:"BRANCH"`
	BRDESC        string                   `json:"BRDESC"`
	BRNAME        string                   `json:"BRNAME"`
	JabatanPgs    string                   `json:"jabatan_pgs"`
	PeriodeAwal   string                   `json:"periode_awal"`
	PeriodeAkhir  string                   `json:"periode_akhir"`
	LastMakerID   string                   `json:"last_maker_id"`
	LastMakerDesc string                   `json:"last_maker_desc"`
	LastMakerDate *string                  `json:"last_maker_date"`
	Action        string                   `json:"action"`
	Approval      []PgsUserApprovalRequest `json:"approval"`
	UpdatedAt     *string                  `json:"updated_at"`
}

type PgsUpdateRequest struct {
	ID            int64   `json:"id"`
	LastMakerID   string  `json:"last_maker_id"`
	LastMakerDesc string  `json:"last_maker_desc"`
	LastMakerDate *string `json:"last_maker_date"`
	Status        string  `json:"status"`
	Action        string  `json:"action"`
}

type PgsUpdateApproval struct {
	ID            int64            `json:"id"`
	LastMakerID   string           `json:"last_maker_id"`
	LastMakerDesc string           `json:"last_maker_desc"`
	LastMakerDate *string          `json:"last_maker_date"`
	Status        string           `json:"status"`
	Action        string           `json:"action"`
	Approval      []ApprovalUpdate `json:"approval"`
}

type LoginRequest struct {
	Pernr    string `json:"pernr"`
	Password string `json:"password"`
}

type Login struct {
	ClientID     string `json:"clientid"`
	ClientSecret string `json:"clientsecret"`
}

type UserSession struct {
	PERNR        string `json:"PERNR"`
	NIP          string `json:"NIP"`
	SNAME        string `json:"SNAME"`
	CORP_TITLE   string `json:"CORP_TITLE"`
	JGPG         string `json:"JGPG"`
	WERKS        string `json:"WERKS"`
	BTRTL        string `json:"BTRTL"`
	KOSTL        string `json:"KOSTL"`
	ORGEH        string `json:"ORGEH"`
	ORGEH_PGS    string `json:"ORGEH_PGS"`
	TIPE_UKER    string `json:"TIPE_UKER"`
	STELL        string `json:"STELL"`
	WERKS_TX     string `json:"WERKS_TX"`
	BTRTL_TX     string `json:"BTRTL_TX"`
	KOSTL_TX     string `json:"KOSTL_TX"`
	ORGEH_TX     string `json:"ORGEH_TX"`
	ORGEH_PGS_TX string `json:"ORGEH_PGS_TX"`
	PLANS_PGS    string `json:"PLANS_PGS"`
	PLANS_PGS_TX string `json:"PLANS_PGS_TX"`
	STELL_TX     string `json:"STELL_TX"`
	PLANS_TX     string `json:"PLANS_TX"`
	BRANCH       string `json:"BRANCH"`
	HILFM        string `json:"HILFM"`
	HILFM_PGS    string `json:"HILFM_PGS"`
	HTEXT        string `json:"HTEXT"`
	HTEXT_PGS    string `json:"HTEXT_PGS"`
	ADD_AREA     string `json:"ADD_AREA"`
	LAST_SYNC    string `json:"LAST_SYNC"`
	REGION       string `json:"REGION"`
	MAINBR       string `json:"MAINBR"`
	BRC          bool   `json:"BRC"`
	RGDESC       string `json:"RGDESC"`
	MBDESC       string `json:"MBDESC"`
	BRDESC       string `json:"BRDESC"`
}

type LoginResponseWithToken struct {
	Token           string                                  `json:"Tsoken`
	User            UserSession                             `json:"User"`
	UKER_BINAAN     []UnitKerja                             `json:UKER_BINAAN`
	ROLE_MENU       []ManagementUser.MenuResponse           `json:"ROLE_MENU"`
	ADDITIONAL_MENU []ManagementUser.AdditionalMenuResponse `json:"additional_menu"`
	ParameterBrc    string                                  `json:"parameter_brc"`
	LogoutUrl       string                                  `json:"logout_url"`
	// Menu  []ManagementUser.MenuResponse `json:"Menu"`
}

type UserSessionIncognito struct {
	PERNR        string `json:"PERNR"`
	NIP          string `json:"NIP"`
	SNAME        string `json:"SNAME"`
	CORP_TITLE   string `json:"CORP_TITLE"`
	JGPG         string `json:"JGPG"`
	WERKS        string `json:"WERKS"`
	BTRTL        string `json:"BTRTL"`
	KOSTL        string `json:"KOSTL"`
	ORGEH        string `json:"ORGEH"`
	ORGEH_PGS    string `json:"ORGEH_PGS"`
	TIPE_UKER    string `json:"TIPE_UKER"`
	STELL        string `json:"STELL"`
	WERKS_TX     string `json:"WERKS_TX"`
	BTRTL_TX     string `json:"BTRTL_TX"`
	KOSTL_TX     string `json:"KOSTL_TX"`
	ORGEH_TX     string `json:"ORGEH_TX"`
	ORGEH_PGS_TX string `json:"ORGEH_PGS_TX"`
	PLANS_PGS    string `json:"PLANS_PGS"`
	PLANS_PGS_TX string `json:"PLANS_PGS_TX"`
	STELL_TX     string `json:"STELL_TX"`
	PLANS_TX     string `json:"PLANS_TX"`
	BRANCH       string `json:"BRANCH"`
	HILFM        string `json:"HILFM"`
	HILFM_PGS    string `json:"HILFM_PGS"`
	HTEXT        string `json:"HTEXT"`
	HTEXT_PGS    string `json:"HTEXT_PGS"`
	ADD_AREA     string `json:"ADD_AREA"`
	STATUS       string `json:"STATUS"`
	REGION       string `json:"REGION"`
	MAINBR       string `json:"MAINBR"`
	BRC          bool   `json:"BRC"`
	RGDESC       string `json:"RGDESC"`
	MBDESC       string `json:"MBDESC"`
	BRDESC       string `json:"BRDESC"`
}

type LoginIncognitoResponseWithToken struct {
	Token           string                                  `json:"Token"`
	User            UserSessionIncognito                    `json:"User"`
	UKER_BINAAN     []UnitKerja                             `json:UKER_BINAAN`
	ROLE_MENU       []ManagementUser.MenuResponse           `json:"ROLE_MENU"`
	ADDITIONAL_MENU []ManagementUser.AdditionalMenuResponse `json:"additional_menu"`
	ParameterBrc    string                                  `json:"parameter_brc"`
	// Menu  []ManagementUser.MenuResponse `json:"Menu"`
}

type UserResponseLocal struct {
	PERNR      string `json:"PERNR"`
	WERKS      string `json:"WERKS"`
	BTRTL      string `json:"BTRTL"`
	KOSTL      string `json:"KOSTL"`
	ORGEH      string `json:"ORGEH"`
	ORGEHPGS   string `json:"ORGEH_PGS"`
	STELL      string `json:"STELL"`
	SNAME      string `json:"SNAME"`
	WERKSTX    string `json:"WERKS_TX"`
	BTRTLTX    string `json:"BTRTL_TX"`
	KOSTLTX    string `json:"KOSTL_TX"`
	ORGEHTX    string `json:"ORGEH_TX"`
	ORGEHPGSTX string `json:"ORGEH_PGS_TX"`
	STELLTX    string `json:"STELL_TX"`
	BRANCH     string `json:"BRANCH"`
	TIPEUKER   string `json:"TIPE_UKER"`
	HILFM      string `json:"HILFM"`
	HILFMPGS   string `json:"HILFM_PGS"`
	HTEXT      string `json:"HTEXT"`
	HTEXTPGS   string `json:"HTEXT_PGS"`
	CORPTITLE  string `json:"CORP_TITLE"`
}

type RequestPn struct {
	PERNR string `json:"pn"`
}

// end of

type UkerKelolaanRequest struct {
	PERNR    string `json:"PERNR"`
	TIPEUKER string `json:"TIPE_UKER"`
	BRANCH   string `json:"BRANCH"`
	HILFM    string `json:"HILFM"`
}

type UnitKerja struct {
	REGION string
	RGDESC string
	MAINBR string
	MBDESC string
	BRANCH string
	BRDESC string
}

func (p PgsUserRequestMaintainance) TableName() string {
	return "pgs_user"
}

func (p PgsUpdateRequest) TableName() string {
	return "pgs_user"
}

func (p PgsUserRequestUpdate) TableName() string {
	return "pgs_user"
}

func (p PgsUserRequest) TableName() string {
	return "pgs_user"
}

func (p PgsUserResponses) TableName() string {
	return "pgs_user"
}

func (p UpdateDelete) TableName() string {
	return "pgs_user"
}

// BrilianApps
type TokenRequest struct {
	Key  string `json:"key"`
	User string `json:"user"`
}

type LoginByToken struct {
	Token string `json:"token"`
}

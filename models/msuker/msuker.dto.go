package models

import "riskmanagement/lib"

type MsUkerRequest struct {
	SRCSYSID string `json:"SRCSYS_ID"`
	BRUNIT   string `json:"BRUNIT"`
	REGION   string `json:"REGION"`
	RGDESC   string `json:"RGDESC"`
	RGNAME   string `json:"RGNAME"`
	MAINBR   int64  `json:"MAINBR"`
	MBDESC   string `json:"MBDESC"`
	MBNAME   string `json:"MBNAME"`
	SUBBR    int64  `json:"SUBBR"`
	SBDESC   string `json:"SBDESC"`
	SBNAME   string `json:"SBNAME"`
	BRANCH   string `json:"BRANCH"`
	BRDESC   string `json:"BRDESC"`
	BRNAME   string `json:"BRNAME"`
	BIBR     string `json:"BIBR"`
}

type MsUkerResponse struct {
	SRCSYSID string `json:"SRCSYS_ID"`
	BRUNIT   string `json:"BRUNIT"`
	REGION   string `json:"REGION"`
	RGDESC   string `json:"RGDESC"`
	RGNAME   string `json:"RGNAME"`
	MAINBR   int64  `json:"MAINBR"`
	MBDESC   string `json:"MBDESC"`
	MBNAME   string `json:"MBNAME"`
	SUBBR    int64  `json:"SUBBR"`
	SBDESC   string `json:"SBDESC"`
	SBNAME   string `json:"SBNAME"`
	BRANCH   string `json:"BRANCH"`
	BRDESC   string `json:"BRDESC"`
	BRNAME   string `json:"BRNAME"`
	BIBR     string `json:"BIBR"`
}

type MsUkerResponseNull struct {
	SRCSYSID lib.NullString `json:"SRCSYS_ID"`
	BRUNIT   lib.NullString `json:"BRUNIT"`
	REGION   lib.NullString `json:"REGION"`
	RGDESC   lib.NullString `json:"RGDESC"`
	RGNAME   lib.NullString `json:"RGNAME"`
	MAINBR   lib.NullInt64  `json:"MAINBR"`
	MBDESC   lib.NullString `json:"MBDESC"`
	MBNAME   lib.NullString `json:"MBNAME"`
	SUBBR    lib.NullInt64  `json:"SUBBR"`
	SBDESC   lib.NullString `json:"SBDESC"`
	SBNAME   lib.NullString `json:"SBNAME"`
	BRANCH   lib.NullString `json:"BRANCH"`
	BRDESC   lib.NullString `json:"BRDESC"`
	BRNAME   lib.NullString `json:"BRNAME"`
	BIBR     lib.NullString `json:"BIBR"`
}

type MsUkerResponseGotOne struct {
	SRCSYSID string `json:"SRCSYS_ID"`
	BRUNIT   string `json:"BRUNIT"`
	REGION   string `json:"REGION"`
	RGDESC   string `json:"RGDESC"`
	RGNAME   string `json:"RGNAME"`
	MAINBR   int64  `json:"MAINBR"`
	MBDESC   string `json:"MBDESC"`
	MBNAME   string `json:"MBNAME"`
	SUBBR    int64  `json:"SUBBR"`
	SBDESC   string `json:"SBDESC"`
	SBNAME   string `json:"SBNAME"`
	BRANCH   string `json:"BRANCH"`
	BRDESC   string `json:"BRDESC"`
	BRNAME   string `json:"BRNAME"`
	BIBR     string `json:"BIBR"`
}

type BranchCodeInduk struct {
	BRANCH string `json:"BRANCH"`
	KOSTL  string `json:"KOSTL"`
	WERKS  string `json:"WERKS"`
}

type Region struct {
	REGION string `json:"REGION"`
}

type KeywordRequest struct {
	TipeUker     string `json:"tipe_uker"`
	BTRTL        string `json:"BTRTL"`
	BRANCH       string `json:"BRANCH"`
	KOSTL        string `json:"KOSTL"`
	Order        string `json:"order"`
	Sort         string `json:"sort"`
	Offset       int    `json:"offset"`
	Limit        int    `json:"limit"`
	Page         int    `json:"page"`
	Keyword      string `json:"keyword"`
	ParameterBrc string `json:"parameter_brc"`
}

type KeyRequest struct {
	Order   string `json:"order"`
	Sort    string `json:"sort"`
	Offset  int    `json:"offset"`
	Limit   int    `json:"limit"`
	Page    int    `json:"page"`
	PN      string `json:"pn"`
	Keyword string `json:"keyword"`
}

type SearchResponse struct {
	SRCSYSID string `json:"SRCSYS_ID"`
	BRUNIT   string `json:"BRUNIT"`
	REGION   string `json:"REGION"`
	RGDESC   string `json:"RGDESC"`
	RGNAME   string `json:"RGNAME"`
	MAINBR   int64  `json:"MAINBR"`
	MBDESC   string `json:"MBDESC"`
	MBNAME   string `json:"MBNAME"`
	SUBBR    int64  `json:"SUBBR"`
	SBDESC   string `json:"SBDESC"`
	SBNAME   string `json:"SBNAME"`
	BRANCH   string `json:"BRANCH"`
	BRDESC   string `json:"BRDESC"`
	BRNAME   string `json:"BRNAME"`
	BIBR     string `json:"BIBR"`
}

type MsPesertaNull struct {
	PERNR   lib.NullString `json:"PERNR"`
	SNAME   lib.NullString `json:"SNAME"`
	STEELTX lib.NullString `json:"STEELTX"`
}

type MsPeserta struct {
	PERNR   string `json:"PERNR"`
	SNAME   string `json:"SNAME"`
	STEELTX string `json:"STEELTX"`
}

type MsPelaku struct {
	PERNR   string `json:"PERNR"`
	SNAME   string `json:"SNAME"`
	STEELTX string `json:"STEELTX"`
	// BRANCH  string `json:"BRANCH"`
}

type MsPekerjaResponse struct {
	PERNR   string `json:"PERNR"`
	SNAME   string `json:"SNAME"`
	STEELTX string `json:"STEELTX"`
	BRANCH  string `json:"BRANCH"`
}

type Jabatan struct {
	HILFM   string `json:"HILFM"`
	HTEXT   string `json:"HTEXT"`
	STELLTX string `json:"STEELTX"`
}

type ListPeserta struct {
	PERNR       string `json:"pernr"`
	NamaPeserta string `json:"nama_peserta"`
	SteelTx     string `json:"steel_tx"`
	BRANCH      string `json:"BRANCH"`
}

type ListPesertaNull struct {
	PERNR       lib.NullString `json:"pernr"`
	NamaPeserta lib.NullString `json:"nama_peserta"`
	SteelTx     lib.NullString `json:"steel_tx"`
	BRANCH      lib.NullString `json:"BRANCH"`
}

type ListJabatanRequest struct {
	BRANCH string `json:"BRANCH"`
	KOSTL  string `json:"KOSTL"`
	WERKS  string `json:"WERKS"`
}

type ListJabatanResponse struct {
	HILFM  string `json:"HILFM"`
	HTEXT  string `json:"HTEXT"`
	Jumlah int    `json:"jumlah"`
}

type BranchByHilfmRequest struct {
	BRANCH string `json:"BRANCH"`
	HILFM  string `json:"HILFM"`
	KOSTL  string `json:"KOSTL"`
	WERKS  string `json:"WERKS"`
}

func (p MsUkerRequest) ParseRequest() MsUker {
	return MsUker{
		SRCSYSID: p.SRCSYSID,
		BRUNIT:   p.BRUNIT,
		REGION:   p.REGION,
		RGDESC:   p.RGDESC,
		RGNAME:   p.RGNAME,
		MAINBR:   p.MAINBR,
		MBDESC:   p.MBDESC,
		MBNAME:   p.MBNAME,
		SUBBR:    p.SUBBR,
		SBDESC:   p.SBDESC,
		SBNAME:   p.SBNAME,
		BRANCH:   p.BRANCH,
		BRDESC:   p.BRDESC,
		BRNAME:   p.BRNAME,
		BIBR:     p.BIBR,
	}
}

func (p MsUkerResponse) ParseResponse() MsUker {
	return MsUker{
		SRCSYSID: p.SRCSYSID,
		BRUNIT:   p.BRUNIT,
		REGION:   p.REGION,
		RGDESC:   p.RGDESC,
		RGNAME:   p.RGNAME,
		MAINBR:   p.MAINBR,
		MBDESC:   p.MBDESC,
		MBNAME:   p.MBNAME,
		SUBBR:    p.SUBBR,
		SBDESC:   p.SBDESC,
		SBNAME:   p.SBNAME,
		BRANCH:   p.BRANCH,
		BRDESC:   p.BRDESC,
		BRNAME:   p.BRNAME,
		BIBR:     p.BIBR,
	}
}

func (mu MsUkerRequest) TableName() string {
	return "dwh_branch"
}

func (mu MsUkerResponse) TableName() string {
	return "dwh_branch"
}

type SearchPNByRegionReq struct {
	REGION       string `json:"region"`
	Keyword      string `json:"keyword"`
	Limit        int64  `json:"limit"`
	ParameterBrc string `json:"parameter_brc"`
}

type SearchPNByRegionRes struct {
	PERNR string `json:"pernr"`
	SNAME string `json:"name"`
}

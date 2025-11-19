package jenistask

type JenisTask struct {
	ID        int    `json:"id"`
	JenisTask string `json:"jenis_task"`
	Kegiatan  string `json:"kegiatan"`
	Period    string `json:"period"`
	Range     string `json:"range"`
	Upload    string `json:"upload"`
}

type TaskRequest struct {
	TipeUker string `json:"tipe_uker"`
	Hilfm    string `json:"hilfm"`
	Kostl    string `json:"kostl"`
	Jgpg     string `json:"jgpg"`
	Stell    string `json:"stell"`
}

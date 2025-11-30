package dto

type BusinessProcessMap struct {
	ID                 string               `json:"id"`
	Code               string               `json:"code"`
	Level              string               `json:"level"`
	Name               string               `json:"name"`
	BusinessProcessMap []BusinessProcessMap `json:"business_processmap"`
}

type HierarchyPagination struct {
	List       []BusinessProcessMap `json:"list"`
	Pagination Pagination           `json:"pagination"`
}

type BusinessProcessNode struct {
	ActivityID           string
	ActivityCode         string
	ActivityName         string
	SubProcessID         string
	SubProcessCode       string
	SubProcessName       string
	ProcessID            string
	ProcessCode          string
	ProcessName          string
	SubBusinessCycleID   string
	SubBusinessCycleCode string
	SubBusinessCycleName string
	BusinessCycleID      string
	BusinessCycleCode    string
	BusinessCycleName    string
}

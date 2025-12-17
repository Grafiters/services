package dto

type MappingRiskEventBusinesProcessRelation struct {
	ID                   string `json:"id"`
	RiskEventID          string `json:"risk_event_id"`
	BusinessCode         string `json:"business_code"`
	BusinessCycle        string `json:"business_cycle"`
	BusinessCycleName    string `json:"business_cycle_name"`
	SubBusinessCode      string `json:"sub_business_code"`
	SubBusinessCycle     string `json:"sub_business_cycle"`
	SubBusinessCycleName string `json:"sub_business_cycle_name"`
	ActivityID           string `json:"activity_id"`
	ActivityCode         string `json:"activity_code"`
	ActivityName         string `json:"activity_name"`
	ProcessCode          string `json:"process_code"`
	ProcessID            string `json:"process_id"`
	ProcessName          string `json:"process_name"`
	SubProcessCode       string `json:"sub_process_code"`
	SubProcessID         string `json:"sub_process_id"`
	SubProcessName       string `json:"sub_process_nane"`
}

type MappingRiskEventBusinesProcessRelationRespnse struct {
	List       []MappingRiskEventBusinesProcessRelation `json:"list"`
	Pagination Pagination                               `json:"pagination"`
}

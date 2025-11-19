package models

type MapEventRequest struct {
	ID           int64  `json:"id"`
	IDRiskIssue  int64  `json:"id_risk_issue"`
	EventTypeLv1 string `json:"event_type_lv1"`
	EventTypeLv2 string `json:"event_type_lv2"`
	EventTypeLv3 string `json:"event_type_lv3"`
}

type MapEventResponse struct {
	ID           int64  `json:"id"`
	IDRiskIssue  int64  `json:"id_risk_issue"`
	EventTypeLv1 string `json:"event_type_lv1"`
	EventTypeLv2 string `json:"event_type_lv2"`
	EventTypeLv3 string `json:"event_type_lv3"`
}

type MapEventResponseFinal struct {
	ID               int64  `json:"id"`
	IDRiskIssue      int64  `json:"id_risk_issue"`
	EventTypeLv1     string `json:"event_type_lv1"`
	EventTypeLv1Desc string `json:"event_type_lv1_desc"`
	EventTypeLv2     string `json:"event_type_lv2"`
	EventTypeLv2Desc string `json:"event_type_lv2_desc"`
	EventTypeLv3     string `json:"event_type_lv3"`
	EventTypeLv3Desc string `json:"event_type_lv3_desc"`
}

func (p MapEventRequest) ParseRequest() MapEvent {
	return MapEvent{
		ID:           p.ID,
		IDRiskIssue:  p.IDRiskIssue,
		EventTypeLv1: p.EventTypeLv1,
		EventTypeLv2: p.EventTypeLv2,
		EventTypeLv3: p.EventTypeLv3,
	}
}

func (p MapEventRequest) ParseResponse() MapEvent {
	return MapEvent{
		ID:           p.ID,
		IDRiskIssue:  p.IDRiskIssue,
		EventTypeLv1: p.EventTypeLv1,
		EventTypeLv2: p.EventTypeLv2,
		EventTypeLv3: p.EventTypeLv3,
	}
}

func (me MapEventRequest) TableName() string {
	return "risk_issue_map_event"
}

func (me MapEventResponse) TableName() string {
	return "risk_issue_map_event"
}

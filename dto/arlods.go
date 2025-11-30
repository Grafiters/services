package dto

type Response[T any] struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       T      `json:"data"`
	Errors     *any   `json:"errors"`
}

type Pagination struct {
	Total int64 `json:"total"`
	Page  int   `json:"page"`
	Limit int   `json:"limit"`
	Next  int   `json:"next"`
	Prev  int   `json:"prev"`
}

type ControlAttributeFiler struct {
	CodeIDs []string `form:"code_ids"`
	Code    string   `form:"code"`
	Status  string   `form:"status"`
}

type DtoRiskControlAttributeResponse struct {
	List       []DtoAttribute `json:"list"`
	Pagination Pagination     `json:"pagination"`
}

type DtoAttribute struct {
	ID        string `json:"id"`
	CodeID    string `json:"code_id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

type MappingRiskEvent struct {
	ID            string `json:"id"`
	RiskEventID   string `json:"risk_event_id"`
	EventTypeLvl1 string `json:"event_type_lvl1"`
	EventTypeLvl2 string `json:"event_type_lvl2"`
	EventTypeLvl3 string `json:"event_type_lvl3"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}
type MappingCauseRiskEvent struct {
	ID             string `json:"id"`
	RiskEventID    string `json:"risk_event_id"`
	Incident       string `json:"incident"`
	SubIncident    string `json:"sub_incident"`
	SubSubIncident string `json:"sub_sub_incident"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

type MappingProductRiskEvent struct {
	ID          string `json:"id"`
	RiskEventID string `json:"risk_event_id"`
	ProductID   string `json:"product_id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type MappingRiskEventBusinesProcess struct {
	ID               string `json:"id"`
	RiskEventID      string `json:"risk_event_id"`
	BusinessCycle    string `json:"business_cycle"`
	SubBusinessCycle string `json:"sub_business_cycle"`
	ActivityID       string `json:"activity_id"`
	ProcessID        string `json:"process_id"`
	SubProcessID     string `json:"sub_process_id"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
}

type RiskEventControl struct {
	ID            string `json:"id"`
	RiskEventID   string `json:"risk_event_id"`
	RiskControlID string `json:"risk_control_id"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

type RiskEventIndicator struct {
	ID              string `json:"id"`
	RiskEventID     string `json:"risk_event_id"`
	RiskIndicatorID string `json:"risk_indicator_id"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}

type BulkRiskEventDetail struct {
	MappingEvent   []MappingRiskEvent               `json:"mapping_event"`
	MappingCause   []MappingCauseRiskEvent          `json:"mapping_cause"`
	MappingProduct []MappingProductRiskEvent        `json:"mapping_product"`
	MappingProcess []MappingRiskEventBusinesProcess `json:"mapping_process"`
	Controls       []RiskEventControl               `json:"controls"`
	Indicators     []RiskEventIndicator             `json:"indicators"`
}

type DetailMappingEvent struct {
	RiskEventID   string              `json:"risk_event_id"`
	MappingDetail BulkRiskEventDetail `json:"mapping_detail"`
}

type RiskEventControlMutateInput struct {
	RiskEventID string   `json:"risk_event_id"`
	TypeEvent   string   `json:"type_event" validate:"required,oneof='event' 'indicator'"`
	RiskID      []string `json:"risk_id" validate:"required"`
}

type MappingLVLRequest struct {
	RiskEventID string `json:"risk_event_id"`
	Lvl1        string `json:"lvl1"`
	Lvl2        string `json:"lvl2"`
	Lvl3        string `json:"lvl3"`
}

type MappingProductRiskEventRequest struct {
	RiskEventID string `json:"risk_event_id"`
	ProductID   string `json:"product_id"`
}

type MapingBusinessProcessRequest struct {
	RiskEventID      string `json:"risk_event_id"`
	ActivityID       string `json:"activity_id"`
	BusinessCycle    string `json:"business_cycle"`
	SubBusinessCycle string `json:"sub_business_cycle"`
	Process          string `json:"process"`
	SubProcess       string `json:"sub_process"`
}

type MappingRiskEventRequest struct {
	MappingRiskEvent        []MappingLVLRequest              `json:"mapping_risk_events"`
	MappingProductRiskEvent []MappingProductRiskEventRequest `json:"mapping_product_risk_events"`
	MappingCauseRiskEvent   []MappingLVLRequest              `json:"mapping_cause_risk_events"`
	MappingBusinessProcess  []MappingRiskEventBusinesProcess `json:"mapping_business_process"`
	MappingIndicatorControl []RiskEventControlMutateInput    `json:"mapping_indicator_control"`
}

type BulkMappingRiskEventRequest struct {
	Data []MappingRiskEventRequest `json:"data"`
}
type RiskEventBulkGetMapping struct {
	EventID []string `form:"event_id"`
}

type HttpResResponse struct {
	StatusCode int      `json:"status_code"`
	Message    string   `json:"message"`
	Data       any      `json:"data"`
	Errors     []string `json:"errors"`
}

type RowBPMapping struct {
	BusinessCycleID    []string `json:"business_cycle_id"`
	SubBusinessCycleID []string `json:"sub_business_cycle_id"`
	ProcessID          []string `json:"process_id"`
	SubProcessID       []string `json:"sub_process_id"`
	ActivityID         []string `json:"activity_id"`
}

type MappingEvent struct {
	EventLV1       []string `json:"event_lv_1"`
	EventLv2       []string `json:"event_lv_2"`
	EventLv3       []string `json:"event_lv_3"`
	Incident       []string `json:"indident"`
	SubIncident    []string `json:"sub_indident"`
	SubSubIncident []string `json:"sub_sub_indident"`
	ProductIDs     []string `json:"productIDs"`
	Control        []string `json:"control"`
	Indicator      []string `json:"indicator"`
}

type MappingEventSet struct {
	EventLV1       map[string]bool
	EventLv2       map[string]bool
	EventLv3       map[string]bool
	Incident       map[string]bool
	SubIncident    map[string]bool
	SubSubIncident map[string]bool
	ProductIDs     map[string]bool
	Control        map[string]bool
	Indicator      map[string]bool
}

func NewMappingEventSet(m MappingEvent) MappingEventSet {
	return MappingEventSet{
		EventLV1:       makeSet(m.EventLV1),
		EventLv2:       makeSet(m.EventLv2),
		EventLv3:       makeSet(m.EventLv3),
		Incident:       makeSet(m.Incident),
		SubIncident:    makeSet(m.SubIncident),
		SubSubIncident: makeSet(m.SubSubIncident),
		ProductIDs:     makeSet(m.ProductIDs),
		Control:        makeSet(m.Control),
		Indicator:      makeSet(m.Indicator),
	}
}

func makeSet(arr []string) map[string]bool {
	m := make(map[string]bool, len(arr))
	for _, v := range arr {
		m[v] = true
	}
	return m
}

func (m MappingEventSet) Contains(val string) bool {
	return m.EventLV1[val] ||
		m.EventLv2[val] ||
		m.EventLv3[val] ||
		m.Incident[val] ||
		m.SubIncident[val] ||
		m.SubSubIncident[val] ||
		m.ProductIDs[val] ||
		m.Control[val] ||
		m.Indicator[val]
}

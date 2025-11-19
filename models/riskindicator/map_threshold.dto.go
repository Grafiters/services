package models

type MapThresholdRequest struct {
	ID          int64   `json:"id"`
	IDIndicator int64   `json:"id_indicator"`
	KCK_1_MIN   float64 `json:"kck_1_min"`
	KCK_2_MIN   float64 `json:"kck_2_min"`
	KCK_3_MIN   float64 `json:"kck_3_min"`
	KCK_4_MIN   float64 `json:"kck_4_min"`
	KCK_5_MIN   float64 `json:"kck_5_min"`
	KC_1_MIN    float64 `json:"kc_1_min"`
	KC_2_MIN    float64 `json:"kc_2_min"`
	KC_3_MIN    float64 `json:"kc_3_min"`
	KC_4_MIN    float64 `json:"kc_4_min"`
	KC_5_MIN    float64 `json:"kc_5_min"`
	KCP_1_MIN   float64 `json:"kcp_1_min"`
	KCP_2_MIN   float64 `json:"kcp_2_min"`
	KCP_3_MIN   float64 `json:"kcp_3_min"`
	KCP_4_MIN   float64 `json:"kcp_4_min"`
	KCP_5_MIN   float64 `json:"kcp_5_min"`
	UN_1_MIN    float64 `json:"un_1_min"`
	UN_2_MIN    float64 `json:"un_2_min"`
	UN_3_MIN    float64 `json:"un_3_min"`
	UN_4_MIN    float64 `json:"un_4_min"`
	UN_5_MIN    float64 `json:"un_5_min"`
	KK_1_MIN    float64 `json:"kk_1_min"`
	KK_2_MIN    float64 `json:"kk_2_min"`
	KK_3_MIN    float64 `json:"kk_3_min"`
	KK_4_MIN    float64 `json:"kk_4_min"`
	KK_5_MIN    float64 `json:"kk_5_min"`
	KCK_1_MAX   float64 `json:"kck_1_max"`
	KCK_2_MAX   float64 `json:"kck_2_max"`
	KCK_3_MAX   float64 `json:"kck_3_max"`
	KCK_4_MAX   float64 `json:"kck_4_max"`
	KCK_5_MAX   float64 `json:"kck_5_max"`
	KC_1_MAX    float64 `json:"kc_1_max"`
	KC_2_MAX    float64 `json:"kc_2_max"`
	KC_3_MAX    float64 `json:"kc_3_max"`
	KC_4_MAX    float64 `json:"kc_4_max"`
	KC_5_MAX    float64 `json:"kc_5_max"`
	KCP_1_MAX   float64 `json:"kcp_1_max"`
	KCP_2_MAX   float64 `json:"kcp_2_max"`
	KCP_3_MAX   float64 `json:"kcp_3_max"`
	KCP_4_MAX   float64 `json:"kcp_4_max"`
	KCP_5_MAX   float64 `json:"kcp_5_max"`
	UN_1_MAX    float64 `json:"un_1_max"`
	UN_2_MAX    float64 `json:"un_2_max"`
	UN_3_MAX    float64 `json:"un_3_max"`
	UN_4_MAX    float64 `json:"un_4_max"`
	UN_5_MAX    float64 `json:"un_5_max"`
	KK_1_MAX    float64 `json:"kk_1_max"`
	KK_2_MAX    float64 `json:"kk_2_max"`
	KK_3_MAX    float64 `json:"kk_3_max"`
	KK_4_MAX    float64 `json:"kk_4_max"`
	KK_5_MAX    float64 `json:"kk_5_max"`
}

type MapThresholdResponse struct {
	ID          int64   `json:"id"`
	IDIndicator int64   `json:"id_indicator"`
	KCK_1_MIN   float64 `json:"kck_1_min"`
	KCK_2_MIN   float64 `json:"kck_2_min"`
	KCK_3_MIN   float64 `json:"kck_3_min"`
	KCK_4_MIN   float64 `json:"kck_4_min"`
	KCK_5_MIN   float64 `json:"kck_5_min"`
	KC_1_MIN    float64 `json:"kc_1_min"`
	KC_2_MIN    float64 `json:"kc_2_min"`
	KC_3_MIN    float64 `json:"kc_3_min"`
	KC_4_MIN    float64 `json:"kc_4_min"`
	KC_5_MIN    float64 `json:"kc_5_min"`
	KCP_1_MIN   float64 `json:"kcp_1_min"`
	KCP_2_MIN   float64 `json:"kcp_2_min"`
	KCP_3_MIN   float64 `json:"kcp_3_min"`
	KCP_4_MIN   float64 `json:"kcp_4_min"`
	KCP_5_MIN   float64 `json:"kcp_5_min"`
	UN_1_MIN    float64 `json:"un_1_min"`
	UN_2_MIN    float64 `json:"un_2_min"`
	UN_3_MIN    float64 `json:"un_3_min"`
	UN_4_MIN    float64 `json:"un_4_min"`
	UN_5_MIN    float64 `json:"un_5_min"`
	KK_1_MIN    float64 `json:"kk_1_min"`
	KK_2_MIN    float64 `json:"kk_2_min"`
	KK_3_MIN    float64 `json:"kk_3_min"`
	KK_4_MIN    float64 `json:"kk_4_min"`
	KK_5_MIN    float64 `json:"kk_5_min"`
	KCK_1_MAX   float64 `json:"kck_1_max"`
	KCK_2_MAX   float64 `json:"kck_2_max"`
	KCK_3_MAX   float64 `json:"kck_3_max"`
	KCK_4_MAX   float64 `json:"kck_4_max"`
	KCK_5_MAX   float64 `json:"kck_5_max"`
	KC_1_MAX    float64 `json:"kc_1_max"`
	KC_2_MAX    float64 `json:"kc_2_max"`
	KC_3_MAX    float64 `json:"kc_3_max"`
	KC_4_MAX    float64 `json:"kc_4_max"`
	KC_5_MAX    float64 `json:"kc_5_max"`
	KCP_1_MAX   float64 `json:"kcp_1_max"`
	KCP_2_MAX   float64 `json:"kcp_2_max"`
	KCP_3_MAX   float64 `json:"kcp_3_max"`
	KCP_4_MAX   float64 `json:"kcp_4_max"`
	KCP_5_MAX   float64 `json:"kcp_5_max"`
	UN_1_MAX    float64 `json:"un_1_max"`
	UN_2_MAX    float64 `json:"un_2_max"`
	UN_3_MAX    float64 `json:"un_3_max"`
	UN_4_MAX    float64 `json:"un_4_max"`
	UN_5_MAX    float64 `json:"un_5_max"`
	KK_1_MAX    float64 `json:"kk_1_max"`
	KK_2_MAX    float64 `json:"kk_2_max"`
	KK_3_MAX    float64 `json:"kk_3_max"`
	KK_4_MAX    float64 `json:"kk_4_max"`
	KK_5_MAX    float64 `json:"kk_5_max"`
}

func (mt MapThresholdRequest) TableName() string {
	return "risk_indicator_map_threshold"
}

func (mt MapThresholdResponse) TableName() string {
	return "risk_indicator_map_threshold"
}

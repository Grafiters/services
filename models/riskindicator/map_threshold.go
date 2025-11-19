package models

type MapThreshold struct {
	ID          int64
	IDIndicator int64
	KCK_1_MIN   float64
	KCK_2_MIN   float64
	KCK_3_MIN   float64
	KCK_4_MIN   float64
	KCK_5_MIN   float64
	KC_1_MIN    float64
	KC_2_MIN    float64
	KC_3_MIN    float64
	KC_4_MIN    float64
	KC_5_MIN    float64
	KCP_1_MIN   float64
	KCP_2_MIN   float64
	KCP_3_MIN   float64
	KCP_4_MIN   float64
	KCP_5_MIN   float64
	UN_1_MIN    float64
	UN_2_MIN    float64
	UN_3_MIN    float64
	UN_4_MIN    float64
	UN_5_MIN    float64
	KK_1_MIN    float64
	KK_2_MIN    float64
	KK_3_MIN    float64
	KK_4_MIN    float64
	KK_5_MIN    float64
	KCK_1_MAX   float64
	KCK_2_MAX   float64
	KCK_3_MAX   float64
	KCK_4_MAX   float64
	KCK_5_MAX   float64
	KC_1_MAX    float64
	KC_2_MAX    float64
	KC_3_MAX    float64
	KC_4_MAX    float64
	KC_5_MAX    float64
	KCP_1_MAX   float64
	KCP_2_MAX   float64
	KCP_3_MAX   float64
	KCP_4_MAX   float64
	KCP_5_MAX   float64
	UN_1_MAX    float64
	UN_2_MAX    float64
	UN_3_MAX    float64
	UN_4_MAX    float64
	UN_5_MAX    float64
	KK_1_MAX    float64
	KK_2_MAX    float64
	KK_3_MAX    float64
	KK_4_MAX    float64
	KK_5_MAX    float64
}

func (mt MapThreshold) TableName() string {
	return "risk_indicator_map_threshold"
}

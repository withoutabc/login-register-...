package model

type Message struct {
	MID     int64  `json:"mid"`
	SendUID int64  `json:"send_uid"`
	RecUID  int64  `json:"rec_uid"`
	Detail  string `json:"detail"`
}

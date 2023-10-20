package model

type TrafficLight struct {
	ID            uint   `json:"id" gorm:"primaryKey"`
	InfoShow      string `json:"info_show"`
	TimeRed       int    `json:"time_red"`
	TimeGreen     int    `json:"time_green"`
	TimeYellow    int    `json:"time_yellow"`
	TimeEmergency int    `json:"time_emergency"`
	IsEmergency   bool   `json:"is_emergency"`
}


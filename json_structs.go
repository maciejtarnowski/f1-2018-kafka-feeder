package f1_2018_kafka_feeder

type CarTelemetryJson struct {
	PlayerId int  `json:"player_id"`
	TimeStamp int64 `json:"timestamp"`
	EventId  int  `json:"event_id"`
	Speed    int  `json:"speed"`
	Throttle int  `json:"throttle"`
	Steer    int  `json:"steer"`
	Brake    int  `json:"brake"`
	Gear     int  `json:"gear"`
	Drs      bool `json:"drs"`
}

type CarMotionJson struct {
	PlayerId  int     `json:"player_id"`
	TimeStamp int64   `json:"timestamp"`
	EventId   int     `json:"event_id"`
	PosX      float32 `json:"pos_x"`
	PosY      float32 `json:"pos_y"`
	PosZ      float32 `json:"pos_z"`
}

type LapDataJson struct {
	PlayerId         int     `json:"player_id"`
	TimeStamp        int64   `json:"timestamp"`
	EventId          int     `json:"event_id"`
	CurrentLapNumber int     `json:"current_lap_number"`
	LastLapTime      float64 `json:"last_lap_time"`
}
package f1_2018_kafka_feeder

type CarTelemetryJson struct {
	Speed    int  `json:"speed"`
	Throttle int  `json:"throttle"`
	Steer    int  `json:"steer"`
	Brake    int  `json:"brake"`
	Gear     int  `json:"gear"`
	Drs      bool `json:"drs"`
}
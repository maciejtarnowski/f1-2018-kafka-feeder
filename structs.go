package f1_2018_kafka_feeder

/*
Source of structs: https://forums.codemasters.com/topic/30601-f1-2018-udp-specification/
 */

/*
struct PacketHeader
{
	uint16    m_packetFormat;         // 2018
	uint8     m_packetVersion;        // Version of this packet type, all start from 1
	uint8     m_packetId;             // Identifier for the packet type, see below
	uint64    m_sessionUID;           // Unique identifier for the session
	float     m_sessionTime;          // Session timestamp
	uint      m_frameIdentifier;      // Identifier for the frame the data was retrieved on
	uint8     m_playerCarIndex;       // Index of player's car in the array
};
*/
type PacketHeader struct {
	PacketFormat uint16
	PacketVersion uint8
	PacketId uint8
	SessionUID uint64
	SessionTime float32
	FrameIdentifier uint32
	PlayerCarIndex uint8
}

/*
struct CarTelemetryData
{
	uint16    m_speed;                      // Speed of car in kilometres per hour
	uint8     m_throttle;                   // Amount of throttle applied (0 to 100)
	int8      m_steer;                      // Steering (-100 (full lock left) to 100 (full lock right))
	uint8     m_brake;                      // Amount of brake applied (0 to 100)
	uint8     m_clutch;                     // Amount of clutch applied (0 to 100)
	int8      m_gear;                       // Gear selected (1-8, N=0, R=-1)
	uint16    m_engineRPM;                  // Engine RPM
	uint8     m_drs;                        // 0 = off, 1 = on
	uint8     m_revLightsPercent;           // Rev lights indicator (percentage)
	uint16    m_brakesTemperature[4];       // Brakes temperature (celsius)
	uint16    m_tyresSurfaceTemperature[4]; // Tyres surface temperature (celsius)
	uint16    m_tyresInnerTemperature[4];   // Tyres inner temperature (celsius)
	uint16    m_engineTemperature;          // Engine temperature (celsius)
	float     m_tyresPressure[4];           // Tyres pressure (PSI)
};
*/
type CarTelemetryData struct {
	Speed uint16
	Throttle uint8
	Steer int8
	Brake uint8
	Clutch uint8
	Gear int8
	EngineRPM uint16
	DRS uint8
	RevLightsPercent uint8
	BrakesTemp [4]uint16
	TyresSurfaceTemp [4]uint16
	TyresInnerTemp [4]uint16
	EngineTemp uint16
	TyresPressure [4]float32
}

/*
struct PacketCarTelemetryData
{
	PacketHeader        m_header;                // Header
	CarTelemetryData    m_carTelemetryData[20];
	uint32              m_buttonStatus;         // Bit flags specifying which buttons are being pressed currently - see appendices
};
*/
type PacketCarTelemetryData struct {
	Header PacketHeader
	Telemetry [20]CarTelemetryData
	ButtonStatus uint32
}


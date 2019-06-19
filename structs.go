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

/*
struct CarMotionData
{
	float         m_worldPositionX;           // World space X position
	float         m_worldPositionY;           // World space Y position
	float         m_worldPositionZ;           // World space Z position
	float         m_worldVelocityX;           // Velocity in world space X
	float         m_worldVelocityY;           // Velocity in world space Y
	float         m_worldVelocityZ;           // Velocity in world space Z
	int16         m_worldForwardDirX;         // World space forward X direction (normalised)
	int16         m_worldForwardDirY;         // World space forward Y direction (normalised)
	int16         m_worldForwardDirZ;         // World space forward Z direction (normalised)
	int16         m_worldRightDirX;           // World space right X direction (normalised)
	int16         m_worldRightDirY;           // World space right Y direction (normalised)
	int16         m_worldRightDirZ;           // World space right Z direction (normalised)
	float         m_gForceLateral;            // Lateral G-Force component
	float         m_gForceLongitudinal;       // Longitudinal G-Force component
	float         m_gForceVertical;           // Vertical G-Force component
	float         m_yaw;                      // Yaw angle in radians
	float         m_pitch;                    // Pitch angle in radians
	float         m_roll;                     // Roll angle in radians
};
 */
type CarMotionData struct {
	WorldPositionX float32
	WorldPositionY float32
	WorldPositionZ float32

	WorldVelocityX float32
	WorldVelocityY float32
	WorldVelocityZ float32

	WorldForwardDirX int16
	WorldForwardDirY int16
	WorldForwardDirZ int16

	WorldRightDirX int16
	WorldRightDirY int16
	WorldRightDirZ int16

	GForceLateral float32
	GForceLongitudinal float32
	GForceVertical float32

	Yaw float32
	Pitch float32
	Roll float32
}

/*
struct PacketMotionData
{
	PacketHeader    m_header;               // Header

	CarMotionData   m_carMotionData[20];    // Data for all cars on track

	// Extra player car ONLY data
	float         m_suspensionPosition[4];       // Note: All wheel arrays have the following order:
	float         m_suspensionVelocity[4];       // RL, RR, FL, FR
	float         m_suspensionAcceleration[4];   // RL, RR, FL, FR
	float         m_wheelSpeed[4];               // Speed of each wheel
	float         m_wheelSlip[4];                // Slip ratio for each wheel
	float         m_localVelocityX;              // Velocity in local space
	float         m_localVelocityY;              // Velocity in local space
	float         m_localVelocityZ;              // Velocity in local space
	float         m_angularVelocityX;            // Angular velocity x-component
	float         m_angularVelocityY;            // Angular velocity y-component
	float         m_angularVelocityZ;            // Angular velocity z-component
	float         m_angularAccelerationX;        // Angular velocity x-component
	float         m_angularAccelerationY;        // Angular velocity y-component
	float         m_angularAccelerationZ;        // Angular velocity z-component
	float         m_frontWheelsAngle;            // Current front wheels angle in radians
};
 */
type PacketMotionData struct {
	Header PacketHeader
	MotionData [20]CarMotionData

	SuspensionPosition [4]float32
	SuspensionVelocity [4]float32
	SuspensionAcceleration [4]float32
	WheelSpeed [4]float32
	WheelSlip [4]float32
	LocalVelocityX float32
	LocalVelocityY float32
	LocalVelocityZ float32
	AngularVelocityX float32
	AngularVelocityY float32
	AngularVelocityZ float32
	AngularAccelerationX float32
	AngularAccelerationY float32
	AngularAccelerationZ float32
	FrontWheelsAngle float32
}
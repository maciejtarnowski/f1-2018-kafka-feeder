-- KSQL streams for efficient telemetry visualisation

CREATE STREAM raw_car_motion_data (
    player_id int,
    event_id int,
    timestamp bigint,
    pos_x double,
    pos_y double,
    pos_z double
) WITH (kafka_topic = 'raw_car_motion_data', value_format = 'JSON', timestamp = 'timestamp');

CREATE STREAM raw_car_telemetry_data (
    player_id int,
    event_id int,
    timestamp bigint,
    speed int,
    throttle int,
    steer int,
    brake int,
    gear int,
    drs boolean
) WITH (kafka_topic = 'raw_car_telemetry_data', value_format = 'JSON', timestamp = 'timestamp');

CREATE STREAM raw_lap_data (
    player_id int,
    event_id int,
    timestamp bigint,
    current_lap_number int,
    last_lap_time double
) WITH (kafka_topic = 'raw_lap_data', value_format = 'JSON', timestamp = 'timestamp');

CREATE STREAM raw_lap_data_rekeyed WITH (partitions = 4) AS SELECT * FROM raw_lap_data PARTITION BY player_id;

CREATE STREAM enriched_car_motion WITH (kafka_topic = 'enriched_car_motion', timestamp = 'timestamp') AS
SELECT
    rcmd.timestamp AS timestamp,
    rcmd.player_id,
    rcmd.pos_x,
    rcmd.pos_y,
    rcmd.pos_z,
    rctd.speed,
    rctd.throttle,
    rctd.brake,
    rctd.steer,
    rctd.gear,
    rctd.drs
FROM
    raw_car_motion_data rcmd
INNER JOIN
    raw_car_telemetry_data rctd
WITHIN 10 MILLISECONDS
ON rctd.player_id = rcmd.player_id;

CREATE STREAM enriched_car_motion_lap WITH (kafka_topic = 'enriched_car_motion_lap', timestamp = 'timestamp') AS
SELECT
    ecm.timestamp AS timestamp,
    ecm.rcmd_player_id,
    ecm.pos_x,
    ecm.pos_y,
    ecm.pos_z,
    ecm.speed,
    ecm.throttle,
    ecm.brake,
    ecm.steer,
    ecm.gear,
    ecm.drs,
    rld.current_lap_number AS current_lap_number
FROM
    enriched_car_motion ecm
INNER JOIN
    raw_lap_data_rekeyed rld
WITHIN 10 MILLISECONDS
ON
    rld.player_id = ecm.rcmd_player_id
PARTITION BY current_lap_number;
use crate::lib::macros;
use crate::lib::util as utils;

/// Convert an Angle (degrees, minutes, and seconds) to Decimal Degrees
pub fn angle_to_decimal_degrees(degrees: f64, minutes: f64, seconds: f64) -> f64 {
    let a = seconds.abs() / 60.0;
    let b = (minutes.abs() + a) / 60.0;
    let c = degrees.abs() + b;
    let d = if degrees < 0.0 || minutes < 0.0 || seconds < 0.0 {
        -c
    } else {
        c
    };

    return d;
}

/// Convert Decimal Degrees to an Angle (degrees, minutes, and seconds)
///
/// ## Returns
/// degrees, minutes, seconds
pub fn decimal_degrees_to_angle(decimal_degrees: f64) -> (f64, f64, f64) {
    let unsigned_decimal = decimal_degrees.abs();
    let total_seconds = unsigned_decimal * 3600.0;
    let seconds_2_dp = utils::round_f64(total_seconds % 60.0, 2);
    let corrected_seconds = if seconds_2_dp == 60.0 {
        0.0
    } else {
        seconds_2_dp
    };
    let corrected_remainder = if seconds_2_dp == 60.0 {
        total_seconds + 60.0
    } else {
        total_seconds
    };
    let minutes = (corrected_remainder / 60.0).floor() % 60.0;
    let unsigned_degrees = (corrected_remainder / 3600.0).floor();
    let signed_degrees = if decimal_degrees < 0.0 {
        -1.0 * unsigned_degrees
    } else {
        unsigned_degrees
    };

    return (signed_degrees, minutes, corrected_seconds.floor());
}

/// Convert Right Ascension to Hour Angle
pub fn right_ascension_to_hour_angle(
    ra_hours: f64,
    ra_minutes: f64,
    ra_seconds: f64,
    lct_hours: f64,
    lct_minutes: f64,
    lct_seconds: f64,
    is_daylight_saving: bool,
    zone_correction: i32,
    local_day: f64,
    local_month: u32,
    local_year: u32,
    geographical_longitude: f64,
) -> (f64, f64, f64) {
    let daylight_saving = if is_daylight_saving == true { 1 } else { 0 };

    let hour_angle = macros::ra_ha(
        ra_hours,
        ra_minutes,
        ra_seconds,
        lct_hours,
        lct_minutes,
        lct_seconds,
        daylight_saving,
        zone_correction,
        local_day,
        local_month,
        local_year,
        geographical_longitude,
    );

    let hour_angle_hours = macros::dh_hour(hour_angle);
    let hour_angle_minutes = macros::dh_min(hour_angle);
    let hour_angle_seconds = macros::dh_sec(hour_angle);

    return (
        hour_angle_hours as f64,
        hour_angle_minutes as f64,
        hour_angle_seconds,
    );
}

/// Convert Hour Angle to Right Ascension
pub fn hour_angle_to_right_ascension(
    hour_angle_hours: f64,
    hour_angle_minutes: f64,
    hour_angle_seconds: f64,
    lct_hours: f64,
    lct_minutes: f64,
    lct_seconds: f64,
    is_daylight_saving: bool,
    zone_correction: i32,
    local_day: f64,
    local_month: u32,
    local_year: u32,
    geographical_longitude: f64,
) -> (f64, f64, f64) {
    let daylight_saving = if is_daylight_saving == true { 1 } else { 0 };

    let right_ascension = macros::ha_ra(
        hour_angle_hours,
        hour_angle_minutes,
        hour_angle_seconds,
        lct_hours,
        lct_minutes,
        lct_seconds as f64,
        daylight_saving,
        zone_correction,
        local_day as f64,
        local_month,
        local_year,
        geographical_longitude,
    );

    let right_ascension_hours = macros::dh_hour(right_ascension);
    let right_ascension_minutes = macros::dh_min(right_ascension);
    let right_ascension_seconds = macros::dh_sec(right_ascension);

    return (
        right_ascension_hours as f64,
        right_ascension_minutes as f64,
        right_ascension_seconds,
    );
}

/// Convert Equatorial Coordinates to Horizon Coordinates
pub fn equatorial_coordinates_to_horizon_coordinates(
    hour_angle_hours: f64,
    hour_angle_minutes: f64,
    hour_angle_seconds: f64,
    declination_degrees: f64,
    declination_minutes: f64,
    declination_seconds: f64,
    geographical_latitude: f64,
) -> (f64, f64, f64, f64, f64, f64) {
    let azimuth_in_decimal_degrees = macros::eq_az(
        hour_angle_hours,
        hour_angle_minutes,
        hour_angle_seconds,
        declination_degrees,
        declination_minutes,
        declination_seconds,
        geographical_latitude,
    );

    let altitude_in_decimal_degrees = macros::eq_alt(
        hour_angle_hours,
        hour_angle_minutes,
        hour_angle_seconds,
        declination_degrees,
        declination_minutes,
        declination_seconds,
        geographical_latitude,
    );

    let azimuth_degrees = macros::dd_deg(azimuth_in_decimal_degrees);
    let azimuth_minutes = macros::dd_min(azimuth_in_decimal_degrees);
    let azimuth_seconds = macros::dd_sec(azimuth_in_decimal_degrees);

    let altitude_degrees = macros::dd_deg(altitude_in_decimal_degrees);
    let altitude_minutes = macros::dd_min(altitude_in_decimal_degrees);
    let altitude_seconds = macros::dd_sec(altitude_in_decimal_degrees);

    return (
        azimuth_degrees,
        azimuth_minutes,
        azimuth_seconds,
        altitude_degrees,
        altitude_minutes,
        altitude_seconds,
    );
}

/// Convert Horizon Coordinates to Equatorial Coordinates
pub fn horizon_coordinates_to_equatorial_coordinates(
    azimuth_degrees: f64,
    azimuth_minutes: f64,
    azimuth_seconds: f64,
    altitude_degrees: f64,
    altitude_minutes: f64,
    altitude_seconds: f64,
    geographical_latitude: f64,
) -> (f64, f64, f64, f64, f64, f64) {
    let hour_angle_in_decimal_degrees = macros::hor_ha(
        azimuth_degrees,
        azimuth_minutes,
        azimuth_seconds,
        altitude_degrees,
        altitude_minutes,
        altitude_seconds,
        geographical_latitude,
    );

    let declination_in_decimal_degrees = macros::hor_dec(
        azimuth_degrees,
        azimuth_minutes,
        azimuth_seconds,
        altitude_degrees,
        altitude_minutes,
        altitude_seconds,
        geographical_latitude,
    );

    let hour_angle_hours = macros::dh_hour(hour_angle_in_decimal_degrees);
    let hour_angle_minutes = macros::dh_min(hour_angle_in_decimal_degrees);
    let hour_angle_seconds = macros::dh_sec(hour_angle_in_decimal_degrees);

    let declination_degrees = macros::dd_deg(declination_in_decimal_degrees);
    let declination_minutes = macros::dd_min(declination_in_decimal_degrees);
    let declination_seconds = macros::dd_sec(declination_in_decimal_degrees);

    return (
        hour_angle_hours as f64,
        hour_angle_minutes as f64,
        hour_angle_seconds,
        declination_degrees,
        declination_minutes,
        declination_seconds,
    );
}

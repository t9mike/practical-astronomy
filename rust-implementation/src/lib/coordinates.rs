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

/// Calculate Mean Obliquity of the Ecliptic for a Greenwich Date
pub fn mean_obliquity_of_the_ecliptic(
    greenwich_day: f64,
    greenwich_month: u32,
    greenwich_year: u32,
) -> f64 {
    let jd = macros::cd_jd(greenwich_day, greenwich_month, greenwich_year);
    let mjd = jd - 2451545.0;
    let t = mjd / 36525.0;
    let de1 = t * (46.815 + t * (0.0006 - (t * 0.00181)));
    let de2 = de1 / 3600.0;

    return 23.439292 - de2;
}

/// Convert Ecliptic Coordinates to Equatorial Coordinates
pub fn ecliptic_coordinate_to_equatorial_coordinate(
    ecliptic_longitude_degrees: f64,
    ecliptic_longitude_minutes: f64,
    ecliptic_longitude_seconds: f64,
    ecliptic_latitude_degrees: f64,
    ecliptic_latitude_minutes: f64,
    ecliptic_latitude_seconds: f64,
    greenwich_day: f64,
    greenwich_month: u32,
    greenwich_year: u32,
) -> (f64, f64, f64, f64, f64, f64) {
    let eclon_deg = macros::dms_dd(
        ecliptic_longitude_degrees,
        ecliptic_longitude_minutes,
        ecliptic_longitude_seconds,
    );
    let eclat_deg = macros::dms_dd(
        ecliptic_latitude_degrees,
        ecliptic_latitude_minutes,
        ecliptic_latitude_seconds,
    );
    let eclon_rad = eclon_deg.to_radians();
    let eclat_rad = eclat_deg.to_radians();
    let obliq_deg = macros::obliq(greenwich_day, greenwich_month, greenwich_year);
    let obliq_rad = obliq_deg.to_radians();
    let sin_dec =
        eclat_rad.sin() * obliq_rad.cos() + eclat_rad.cos() * obliq_rad.sin() * eclon_rad.sin();
    let dec_rad = sin_dec.asin();
    let dec_deg = macros::degrees(dec_rad);
    let y = eclon_rad.sin() * obliq_rad.cos() - eclat_rad.tan() * obliq_rad.sin();
    let x = eclon_rad.cos();
    let ra_rad = y.atan2(x);
    let ra_deg1 = macros::degrees(ra_rad);
    let ra_deg2 = ra_deg1 - 360.0 * (ra_deg1 / 360.0).floor();
    let ra_hours = macros::dd_dh(ra_deg2);

    let out_ra_hours = macros::dh_hour(ra_hours);
    let out_ra_minutes = macros::dh_min(ra_hours);
    let out_ra_seconds = macros::dh_sec(ra_hours);
    let out_dec_degrees = macros::dd_deg(dec_deg);
    let out_dec_minutes = macros::dd_min(dec_deg);
    let out_dec_seconds = macros::dd_sec(dec_deg);

    return (
        out_ra_hours as f64,
        out_ra_minutes as f64,
        out_ra_seconds,
        out_dec_degrees,
        out_dec_minutes,
        out_dec_seconds,
    );
}

/// Convert Equatorial Coordinates to Ecliptic Coordinates
pub fn equatorial_coordinate_to_ecliptic_coordinate(
    ra_hours: f64,
    ra_minutes: f64,
    ra_seconds: f64,
    dec_degrees: f64,
    dec_minutes: f64,
    dec_seconds: f64,
    gw_day: f64,
    gw_month: u32,
    gw_year: u32,
) -> (f64, f64, f64, f64, f64, f64) {
    let ra_deg = macros::dh_dd(macros::hms_dh(ra_hours, ra_minutes, ra_seconds));
    let dec_deg = macros::dms_dd(dec_degrees, dec_minutes, dec_seconds);
    let ra_rad = ra_deg.to_radians();
    let dec_rad = dec_deg.to_radians();
    let obliq_deg = macros::obliq(gw_day, gw_month, gw_year);
    let obliq_rad = obliq_deg.to_radians();
    let sin_ecl_lat =
        dec_rad.sin() * obliq_rad.cos() - dec_rad.cos() * obliq_rad.sin() * ra_rad.sin();
    let ecl_lat_rad = sin_ecl_lat.asin();
    let ecl_lat_deg = macros::degrees(ecl_lat_rad);
    let y = ra_rad.sin() * obliq_rad.cos() + dec_rad.tan() * obliq_rad.sin();
    let x = ra_rad.cos();
    let ecl_long_rad = y.atan2(x);
    let ecl_long_deg1 = macros::degrees(ecl_long_rad);
    let ecl_long_deg2 = ecl_long_deg1 - 360.0 * (ecl_long_deg1 / 360.0).floor();

    let out_ecl_long_deg = macros::dd_deg(ecl_long_deg2);
    let out_ecl_long_min = macros::dd_min(ecl_long_deg2);
    let out_ecl_long_sec = macros::dd_sec(ecl_long_deg2);
    let out_ecl_lat_deg = macros::dd_deg(ecl_lat_deg);
    let out_ecl_lat_min = macros::dd_min(ecl_lat_deg);
    let out_ecl_lat_sec = macros::dd_sec(ecl_lat_deg);

    return (
        out_ecl_long_deg,
        out_ecl_long_min,
        out_ecl_long_sec,
        out_ecl_lat_deg,
        out_ecl_lat_min,
        out_ecl_lat_sec,
    );
}

/// Convert Equatorial Coordinates to Galactic Coordinates
pub fn equatorial_coordinate_to_galactic_coordinate(
    ra_hours: f64,
    ra_minutes: f64,
    ra_seconds: f64,
    dec_degrees: f64,
    dec_minutes: f64,
    dec_seconds: f64,
) -> (f64, f64, f64, f64, f64, f64) {
    let ra_deg = macros::dh_dd(macros::hms_dh(ra_hours, ra_minutes, ra_seconds));
    let dec_deg = macros::dms_dd(dec_degrees, dec_minutes, dec_seconds);
    let ra_rad = ra_deg.to_radians();
    let dec_rad = dec_deg.to_radians();
    let sin_b = dec_rad.cos()
        * (27.4 as f64).to_radians().cos()
        * (ra_rad - (192.25 as f64).to_radians()).cos()
        + dec_rad.sin() * (27.4 as f64).to_radians().sin();
    let b_radians = sin_b.asin();
    let b_deg = macros::degrees(b_radians);
    let y = dec_rad.sin() - sin_b * (27.4 as f64).to_radians().sin();
    let x = dec_rad.cos()
        * (ra_rad - (192.25 as f64).to_radians()).sin()
        * (27.4 as f64).to_radians().cos();
    let long_deg1 = macros::degrees(y.atan2(x)) + 33.0;
    let long_deg2 = long_deg1 - 360.0 * (long_deg1 / 360.0).floor();

    let gal_long_deg = macros::dd_deg(long_deg2);
    let gal_long_min = macros::dd_min(long_deg2);
    let gal_long_sec = macros::dd_sec(long_deg2);
    let gal_lat_deg = macros::dd_deg(b_deg);
    let gal_lat_min = macros::dd_min(b_deg);
    let gal_lat_sec = macros::dd_sec(b_deg);

    return (
        gal_long_deg,
        gal_long_min,
        gal_long_sec,
        gal_lat_deg,
        gal_lat_min,
        gal_lat_sec,
    );
}

/// Convert Galactic Coordinates to Equatorial Coordinates
pub fn galactic_coordinate_to_equatorial_coordinate(
    gal_long_deg: f64,
    gal_long_min: f64,
    gal_long_sec: f64,
    gal_lat_deg: f64,
    gal_lat_min: f64,
    gal_lat_sec: f64,
) -> (f64, f64, f64, f64, f64, f64) {
    let glong_deg = macros::dms_dd(gal_long_deg, gal_long_min, gal_long_sec);
    let glat_deg = macros::dms_dd(gal_lat_deg, gal_lat_min, gal_lat_sec);
    let glong_rad = glong_deg.to_radians();
    let glat_rad = glat_deg.to_radians();
    let sin_dec = glat_rad.cos()
        * (27.4 as f64).to_radians().cos()
        * (glong_rad - (33 as f64).to_radians()).sin()
        + glat_rad.sin() * (27.4 as f64).to_radians().sin();
    let dec_radians = sin_dec.asin();
    let dec_deg = macros::degrees(dec_radians);
    let y = glat_rad.cos() * (glong_rad - (33 as f64).to_radians()).cos();
    let x = glat_rad.sin() * ((27.4 as f64).to_radians()).cos()
        - (glat_rad).cos()
            * ((27.4 as f64).to_radians()).sin()
            * (glong_rad - (33 as f64).to_radians()).sin();

    let ra_deg1 = macros::degrees(y.atan2(x)) + 192.25;
    let ra_deg2 = ra_deg1 - 360.0 * (ra_deg1 / 360.0).floor();
    let ra_hours1 = macros::dd_dh(ra_deg2);

    let ra_hours = macros::dh_hour(ra_hours1);
    let ra_minutes = macros::dh_min(ra_hours1);
    let ra_seconds = macros::dh_sec(ra_hours1);
    let dec_degrees = macros::dd_deg(dec_deg);
    let dec_minutes = macros::dd_min(dec_deg);
    let dec_seconds = macros::dd_sec(dec_deg);

    return (
        ra_hours as f64,
        ra_minutes as f64,
        ra_seconds,
        dec_degrees,
        dec_minutes,
        dec_seconds,
    );
}

/// Calculate the angle between two celestial objects
pub fn angle_between_two_objects(
    ra_long_1_hour_deg: f64,
    ra_long_1_min: f64,
    ra_long_1_sec: f64,
    dec_lat_1_deg: f64,
    dec_lat_1_min: f64,
    dec_lat_1_sec: f64,
    ra_long_2_hour_deg: f64,
    ra_long_2_min: f64,
    ra_long_2_sec: f64,
    dec_lat_2_deg: f64,
    dec_lat_2_min: f64,
    dec_lat_2_sec: f64,
    hour_or_degree: String,
) -> (f64, f64, f64) {
    let ra_long_1_decimal = if hour_or_degree == "H" {
        macros::hms_dh(ra_long_1_hour_deg, ra_long_1_min, ra_long_1_sec)
    } else {
        macros::dms_dd(ra_long_1_hour_deg, ra_long_1_min, ra_long_1_sec)
    };
    let ra_long_1_deg = if hour_or_degree == "H" {
        macros::dh_dd(ra_long_1_decimal)
    } else {
        ra_long_1_decimal
    };
    let ra_long_1_rad = ra_long_1_deg.to_radians();
    let dec_lat_1_deg1 = macros::dms_dd(dec_lat_1_deg, dec_lat_1_min, dec_lat_1_sec);
    let dec_lat_1_rad = dec_lat_1_deg1.to_radians();
    let ra_long_2_decimal = if hour_or_degree == "H" {
        macros::hms_dh(ra_long_2_hour_deg, ra_long_2_min, ra_long_2_sec)
    } else {
        macros::dms_dd(ra_long_2_hour_deg, ra_long_2_min, ra_long_2_sec)
    };
    let ra_long_2_deg = if hour_or_degree == "H" {
        macros::dh_dd(ra_long_2_decimal)
    } else {
        ra_long_2_decimal
    };
    let ra_long_2_rad = ra_long_2_deg.to_radians();
    let dec_lat_2_deg1 = macros::dms_dd(dec_lat_2_deg, dec_lat_2_min, dec_lat_2_sec);
    let dec_lat_2_rad = dec_lat_2_deg1.to_radians();

    let cos_d = dec_lat_1_rad.sin() * dec_lat_2_rad.sin()
        + dec_lat_1_rad.cos() * dec_lat_2_rad.cos() * (ra_long_1_rad - ra_long_2_rad).cos();
    let d_rad = cos_d.acos();
    let d_deg = macros::degrees(d_rad);

    let angle_deg = macros::dd_deg(d_deg);
    let angle_min = macros::dd_min(d_deg);
    let angle_sec = macros::dd_sec(d_deg);

    return (angle_deg, angle_min, angle_sec);
}

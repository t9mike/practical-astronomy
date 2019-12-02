use crate::lib::macros;
use crate::lib::util;

/// Calculate approximate position of the Moon.
///
/// ## Arguments
/// * `lct_hour` -- Local civil time, in hours.
/// * `lct_min` -- Local civil time, in minutes.
/// * `lct_sec` -- Local civil time, in seconds.
/// * `is_daylight_saving` -- Is daylight savings in effect?
/// * `zone_correction_hours` -- Time zone correction, in hours.
/// * `local_date_day` -- Local date, day part.
/// * `local_date_month` -- Local date, month part.
/// * `local_date_year` -- Local date, year part.
///
/// ## Returns
/// * `moon_ra_hour` -- Right ascension of Moon (hour part)
/// * `moon_ra_min` -- Right ascension of Moon (minutes part)
/// * `moon_ra_sec` -- Right ascension of Moon (seconds part)
/// * `moon_dec_deg` -- Declination of Moon (degrees part)
/// * `moon_dec_min` -- Declination of Moon (minutes part)
/// * `moon_dec_sec` -- Declination of Moon (seconds part)
pub fn approximate_position_of_moon(
    lct_hour: f64,
    lct_min: f64,
    lct_sec: f64,
    is_daylight_saving: bool,
    zone_correction_hours: i32,
    local_date_day: f64,
    local_date_month: u32,
    local_date_year: u32,
) -> (f64, f64, f64, f64, f64, f64) {
    let daylight_saving = if is_daylight_saving == true { 1 } else { 0 };

    let l0 = 91.9293359879052;
    let p0 = 130.143076320618;
    let n0 = 291.682546643194;
    let i: f64 = 5.145396;

    let gdate_day = macros::lct_gday(
        lct_hour,
        lct_min,
        lct_sec,
        daylight_saving,
        zone_correction_hours,
        local_date_day,
        local_date_month,
        local_date_year,
    );
    let gdate_month = macros::lct_gmonth(
        lct_hour,
        lct_min,
        lct_sec,
        daylight_saving,
        zone_correction_hours,
        local_date_day,
        local_date_month,
        local_date_year,
    );
    let gdate_year = macros::lct_gyear(
        lct_hour,
        lct_min,
        lct_sec,
        daylight_saving,
        zone_correction_hours,
        local_date_day,
        local_date_month,
        local_date_year,
    );

    let ut_hours = macros::lct_ut(
        lct_hour,
        lct_min,
        lct_sec,
        daylight_saving,
        zone_correction_hours,
        local_date_day,
        local_date_month,
        local_date_year,
    );
    let d_days = macros::cd_jd(gdate_day, gdate_month, gdate_year) - macros::cd_jd(0.0, 1, 2010)
        + ut_hours / 24.0;
    let sun_long_deg = macros::sun_long(
        lct_hour,
        lct_min,
        lct_sec,
        daylight_saving,
        zone_correction_hours,
        local_date_day,
        local_date_month,
        local_date_year,
    );
    let sun_mean_anomaly_rad = macros::sun_mean_anomaly(
        lct_hour,
        lct_min,
        lct_sec,
        daylight_saving,
        zone_correction_hours,
        local_date_day,
        local_date_month,
        local_date_year,
    );
    let lm_deg = macros::unwind_deg(13.1763966 * d_days + l0);
    let mm_deg = macros::unwind_deg(lm_deg - 0.1114041 * d_days - p0);
    let n_deg = macros::unwind_deg(n0 - (0.0529539 * d_days));
    let ev_deg = 1.2739 * ((2.0 * (lm_deg - sun_long_deg) - mm_deg).to_radians()).sin();
    let ae_deg = 0.1858 * (sun_mean_anomaly_rad).sin();
    let a3_deg = 0.37 * (sun_mean_anomaly_rad).sin();
    let mmd_deg = mm_deg + ev_deg - ae_deg - a3_deg;
    let ec_deg = 6.2886 * mmd_deg.to_radians().sin();
    let a4_deg = 0.214 * (2.0 * (mmd_deg).to_radians()).sin();
    let ld_deg = lm_deg + ev_deg + ec_deg - ae_deg + a4_deg;
    let v_deg = 0.6583 * (2.0 * (ld_deg - sun_long_deg).to_radians()).sin();
    let ldd_deg = ld_deg + v_deg;
    let nd_deg = n_deg - 0.16 * (sun_mean_anomaly_rad).sin();
    let y = ((ldd_deg - nd_deg).to_radians()).sin() * i.to_radians().cos();
    let x = (ldd_deg - nd_deg).to_radians().cos();

    let moon_long_deg = macros::unwind_deg(macros::degrees(y.atan2(x)) + nd_deg);
    let moon_lat_deg =
        macros::degrees(((ldd_deg - nd_deg).to_radians().sin() * i.to_radians().sin()).asin());
    let moon_ra_hours1 = macros::dd_dh(macros::ec_ra(
        moon_long_deg,
        0.0,
        0.0,
        moon_lat_deg,
        0.0,
        0.0,
        gdate_day,
        gdate_month,
        gdate_year,
    ));
    let moon_dec_deg1 = macros::ec_dec(
        moon_long_deg,
        0.0,
        0.0,
        moon_lat_deg,
        0.0,
        0.0,
        gdate_day,
        gdate_month,
        gdate_year,
    );

    let moon_ra_hour = macros::dh_hour(moon_ra_hours1);
    let moon_ra_min = macros::dh_min(moon_ra_hours1);
    let moon_ra_sec = macros::dh_sec(moon_ra_hours1);
    let moon_dec_deg = macros::dd_deg(moon_dec_deg1);
    let moon_dec_min = macros::dd_min(moon_dec_deg1);
    let moon_dec_sec = macros::dd_sec(moon_dec_deg1);

    return (
        moon_ra_hour as f64,
        moon_ra_min as f64,
        moon_ra_sec,
        moon_dec_deg,
        moon_dec_min,
        moon_dec_sec,
    );
}

/// Calculate approximate position of the Moon.
///
/// ## Arguments
/// * `lct_hour` -- Local civil time, in hours.
/// * `lct_min` -- Local civil time, in minutes.
/// * `lct_sec` -- Local civil time, in seconds.
/// * `is_daylight_saving` -- Is daylight savings in effect?
/// * `zone_correction_hours` -- Time zone correction, in hours.
/// * `local_date_day` -- Local date, day part.
/// * `local_date_month` -- Local date, month part.
/// * `local_date_year` -- Local date, year part.
///
/// ## Returns
/// * `moon_ra_hour` -- Right ascension of Moon (hour part)
/// * `moon_ra_min` -- Right ascension of Moon (minutes part)
/// * `moon_ra_sec` -- Right ascension of Moon (seconds part)
/// * `moon_dec_deg` -- Declination of Moon (degrees part)
/// * `moon_dec_min` -- Declination of Moon (minutes part)
/// * `moon_dec_sec` -- Declination of Moon (seconds part)
/// * `earth_moon_dist_km` -- Distance from Earth to Moon (km)
/// * `moon_hor_parallax_deg` -- Horizontal parallax of Moon (degrees)
pub fn precise_position_of_moon(
    lct_hour: f64,
    lct_min: f64,
    lct_sec: f64,
    is_daylight_saving: bool,
    zone_correction_hours: i32,
    local_date_day: f64,
    local_date_month: u32,
    local_date_year: u32,
) -> (f64, f64, f64, f64, f64, f64, f64, f64) {
    let daylight_saving = if is_daylight_saving == true { 1 } else { 0 };

    let gdate_day = macros::lct_gday(
        lct_hour,
        lct_min,
        lct_sec,
        daylight_saving,
        zone_correction_hours,
        local_date_day,
        local_date_month,
        local_date_year,
    );
    let gdate_month = macros::lct_gmonth(
        lct_hour,
        lct_min,
        lct_sec,
        daylight_saving,
        zone_correction_hours,
        local_date_day,
        local_date_month,
        local_date_year,
    );
    let gdate_year = macros::lct_gyear(
        lct_hour,
        lct_min,
        lct_sec,
        daylight_saving,
        zone_correction_hours,
        local_date_day,
        local_date_month,
        local_date_year,
    );

    let _ut_hours = macros::lct_ut(
        lct_hour,
        lct_min,
        lct_sec,
        daylight_saving,
        zone_correction_hours,
        local_date_day,
        local_date_month,
        local_date_year,
    );

    let (moon_ecliptic_longitude_deg, moon_ecliptic_latitude_deg, moon_horizontal_parallax_deg) =
        macros::moon_long_lat_hp(
            lct_hour,
            lct_min,
            lct_sec,
            daylight_saving,
            zone_correction_hours,
            local_date_day,
            local_date_month,
            local_date_year,
        );

    let nutation_in_longitude_deg = macros::nutat_long(gdate_day, gdate_month, gdate_year);
    let corrected_long_deg = moon_ecliptic_longitude_deg + nutation_in_longitude_deg;
    let earth_moon_distance_km = 6378.14 / moon_horizontal_parallax_deg.to_radians().sin();
    let moon_ra_hours_1 = macros::dd_dh(macros::ec_ra(
        corrected_long_deg,
        0.0,
        0.0,
        moon_ecliptic_latitude_deg,
        0.0,
        0.0,
        gdate_day,
        gdate_month,
        gdate_year,
    ));
    let moon_dec_deg1 = macros::ec_dec(
        corrected_long_deg,
        0.0,
        0.0,
        moon_ecliptic_latitude_deg,
        0.0,
        0.0,
        gdate_day,
        gdate_month,
        gdate_year,
    );

    let moon_ra_hour = macros::dh_hour(moon_ra_hours_1);
    let moon_ra_min = macros::dh_min(moon_ra_hours_1);
    let moon_ra_sec = macros::dh_sec(moon_ra_hours_1);
    let moon_dec_deg = macros::dd_deg(moon_dec_deg1);
    let moon_dec_min = macros::dd_min(moon_dec_deg1);
    let moon_dec_sec = macros::dd_sec(moon_dec_deg1);
    let earth_moon_dist_km = util::round_f64(earth_moon_distance_km, 0);
    let moon_hor_parallax_deg = util::round_f64(moon_horizontal_parallax_deg, 6);

    return (
        moon_ra_hour as f64,
        moon_ra_min as f64,
        moon_ra_sec,
        moon_dec_deg,
        moon_dec_min,
        moon_dec_sec,
        earth_moon_dist_km,
        moon_hor_parallax_deg,
    );
}

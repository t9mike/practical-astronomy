use crate::lib::macros;
use crate::lib::util as utils;

/// Calculate approximate position of the sun for a local date and time.
///
/// ## Arguments
/// * `lct_hours` -- Local civil time, in hours.
/// * `lct_minutes` -- Local civil time, in minutes.
/// * `lct_seconds` -- Local civil time, in seconds.
/// * `local_day` -- Local date, day part.
/// * `local_month` -- Local date, month part.
/// * `local_year` -- Local date, year part.
/// * `is_daylight_saving` -- Is daylight savings in effect?
/// * `zone_correction` -- Time zone correction, in hours.
///
/// ## Returns
/// * `sun_ra_hour` -- Right Ascension of Sun, hour part
/// * `sun_ra_min` -- Right Ascension of Sun, minutes part
/// * `sun_ra_sec` -- Right Ascension of Sun, seconds part
/// * `sun_dec_deg` -- Declination of Sun, degrees part
/// * `sun_dec_min` -- Declination of Sun, minutes part
/// * `sun_dec_sec` -- Declination of Sun, seconds part
pub fn approximate_position_of_sun(
    lct_hours: f64,
    lct_minutes: f64,
    lct_seconds: f64,
    local_day: f64,
    local_month: u32,
    local_year: u32,
    is_daylight_saving: bool,
    zone_correction: i32,
) -> (f64, f64, f64, f64, f64, f64) {
    let daylight_saving = if is_daylight_saving == true { 1 } else { 0 };

    let greenwich_date_day = macros::lct_gday(
        lct_hours,
        lct_minutes,
        lct_seconds,
        daylight_saving,
        zone_correction,
        local_day,
        local_month,
        local_year,
    );
    let greenwich_date_month = macros::lct_gmonth(
        lct_hours,
        lct_minutes,
        lct_seconds,
        daylight_saving,
        zone_correction,
        local_day,
        local_month,
        local_year,
    );
    let greenwich_date_year = macros::lct_gyear(
        lct_hours,
        lct_minutes,
        lct_seconds,
        daylight_saving,
        zone_correction,
        local_day,
        local_month,
        local_year,
    );
    let ut_hours = macros::lct_ut(
        lct_hours,
        lct_minutes,
        lct_seconds,
        daylight_saving,
        zone_correction,
        local_day,
        local_month,
        local_year,
    );
    let ut_days = ut_hours / 24.0;
    let jd_days = macros::cd_jd(
        greenwich_date_day,
        greenwich_date_month,
        greenwich_date_year,
    ) + ut_days;
    let d_days = jd_days - macros::cd_jd(0 as f64, 1, 2010);
    let n_deg = 360.0 * d_days / 365.242191;
    let m_deg1 =
        n_deg + macros::sun_e_long(0 as f64, 1, 2010) - macros::sun_peri(0 as f64, 1, 2010);
    let m_deg2 = m_deg1 - 360.0 * (m_deg1 / 360.0).floor();
    let e_c_deg = 360.0 * macros::sun_ecc(0 as f64, 1, 2010) * m_deg2.to_radians().sin()
        / std::f64::consts::PI;
    let l_s_deg1 = n_deg + e_c_deg + macros::sun_e_long(0 as f64, 1, 2010);
    let l_s_deg2 = l_s_deg1 - 360.0 * (l_s_deg1 / 360.0).floor();
    let ra_deg = macros::ec_ra(
        l_s_deg2,
        0 as f64,
        0 as f64,
        0 as f64,
        0 as f64,
        0 as f64,
        greenwich_date_day,
        greenwich_date_month,
        greenwich_date_year,
    );
    let ra_hours = macros::dd_dh(ra_deg);
    let dec_deg = macros::ec_dec(
        l_s_deg2,
        0 as f64,
        0 as f64,
        0 as f64,
        0 as f64,
        0 as f64,
        greenwich_date_day,
        greenwich_date_month,
        greenwich_date_year,
    );

    let sun_ra_hour = macros::dh_hour(ra_hours);
    let sun_ra_min = macros::dh_min(ra_hours);
    let sun_ra_sec = macros::dh_sec(ra_hours);
    let sun_dec_deg = macros::dd_deg(dec_deg);
    let sun_dec_min = macros::dd_min(dec_deg);
    let sun_dec_sec = macros::dd_sec(dec_deg);

    return (
        sun_ra_hour as f64,
        sun_ra_min as f64,
        sun_ra_sec,
        sun_dec_deg,
        sun_dec_min,
        sun_dec_sec,
    );
}

/// Calculate precise position of the sun for a local date and time.
///
/// ## Arguments
/// * `lct_hours` -- Local civil time, in hours.
/// * `lct_minutes` -- Local civil time, in minutes.
/// * `lct_seconds` -- Local civil time, in seconds.
/// * `local_day` -- Local date, day part.
/// * `local_month` -- Local date, month part.
/// * `local_year` -- Local date, year part.
/// * `is_daylight_saving` -- Is daylight savings in effect?
/// * `zone_correction` -- Time zone correction, in hours.
///
/// ## Returns
/// * `sun_ra_hour` -- Right Ascension of Sun, hour part
/// * `sun_ra_min` -- Right Ascension of Sun, minutes part
/// * `sun_ra_sec` -- Right Ascension of Sun, seconds part
/// * `sun_dec_deg` -- Declination of Sun, degrees part
/// * `sun_dec_min` -- Declination of Sun, minutes part
/// * `sun_dec_sec` -- Declination of Sun, seconds part
pub fn precise_position_of_sun(
    lct_hours: f64,
    lct_minutes: f64,
    lct_seconds: f64,
    local_day: f64,
    local_month: u32,
    local_year: u32,
    is_daylight_saving: bool,
    zone_correction: i32,
) -> (f64, f64, f64, f64, f64, f64) {
    let daylight_saving = if is_daylight_saving == true { 1 } else { 0 };

    let g_day = macros::lct_gday(
        lct_hours,
        lct_minutes,
        lct_seconds,
        daylight_saving,
        zone_correction,
        local_day,
        local_month,
        local_year,
    );
    let g_month = macros::lct_gmonth(
        lct_hours,
        lct_minutes,
        lct_seconds,
        daylight_saving,
        zone_correction,
        local_day,
        local_month,
        local_year,
    );
    let g_year = macros::lct_gyear(
        lct_hours,
        lct_minutes,
        lct_seconds,
        daylight_saving,
        zone_correction,
        local_day,
        local_month,
        local_year,
    );
    let sun_ecliptic_longitude_deg = macros::sun_long(
        lct_hours,
        lct_minutes,
        lct_seconds,
        daylight_saving,
        zone_correction,
        local_day,
        local_month,
        local_year,
    );
    let ra_deg = macros::ec_ra(
        sun_ecliptic_longitude_deg,
        0.0,
        0.0,
        0.0,
        0.0,
        0.0,
        g_day,
        g_month,
        g_year,
    );
    let ra_hours = macros::dd_dh(ra_deg);
    let dec_deg = macros::ec_dec(
        sun_ecliptic_longitude_deg,
        0.0,
        0.0,
        0.0,
        0.0,
        0.0,
        g_day,
        g_month,
        g_year,
    );

    let sun_ra_hour = macros::dh_hour(ra_hours);
    let sun_ra_min = macros::dh_min(ra_hours);
    let sun_ra_sec = macros::dh_sec(ra_hours);
    let sun_dec_deg = macros::dd_deg(dec_deg);
    let sun_dec_min = macros::dd_min(dec_deg);
    let sun_dec_sec = macros::dd_sec(dec_deg);

    return (
        sun_ra_hour as f64,
        sun_ra_min as f64,
        sun_ra_sec,
        sun_dec_deg,
        sun_dec_min,
        sun_dec_sec,
    );
}

/// Calculate distance to the Sun (in km), and angular size.
///
/// ## Arguments
/// * `lct_hours` -- Local civil time, in hours.
/// * `lct_minutes` -- Local civil time, in minutes.
/// * `lct_seconds` -- Local civil time, in seconds.
/// * `local_day` -- Local date, day part.
/// * `local_month` -- Local date, month part.
/// * `local_year` -- Local date, year part.
/// * `is_daylight_saving` -- Is daylight savings in effect?
/// * `zone_correction` -- Time zone correction, in hours.
///
/// ## Returns
/// * `sun_dist_km` -- Sun's distance, in kilometers
/// * `sun_ang_size_deg` -- Sun's angular size (degrees part)
/// * `sun_ang_size_min` -- Sun's angular size (minutes part)
/// * `sun_ang_size_sec` -- Sun's angular size (seconds part)
pub fn sun_distance_and_angular_size(
    lct_hours: f64,
    lct_minutes: f64,
    lct_seconds: f64,
    local_day: f64,
    local_month: u32,
    local_year: u32,
    is_daylight_saving: bool,
    zone_correction: i32,
) -> (f64, f64, f64, f64) {
    let daylight_saving = if is_daylight_saving == true { 1 } else { 0 };

    let g_day = macros::lct_gday(
        lct_hours,
        lct_minutes,
        lct_seconds,
        daylight_saving,
        zone_correction,
        local_day,
        local_month,
        local_year,
    );
    let g_month = macros::lct_gmonth(
        lct_hours,
        lct_minutes,
        lct_seconds,
        daylight_saving,
        zone_correction,
        local_day,
        local_month,
        local_year,
    );
    let g_year = macros::lct_gyear(
        lct_hours,
        lct_minutes,
        lct_seconds,
        daylight_saving,
        zone_correction,
        local_day,
        local_month,
        local_year,
    );
    let true_anomaly_deg = macros::sun_true_anomaly(
        lct_hours,
        lct_minutes,
        lct_seconds,
        daylight_saving,
        zone_correction,
        local_day,
        local_month,
        local_year,
    );
    let true_anomaly_rad = true_anomaly_deg.to_radians();
    let eccentricity = macros::sun_ecc(g_day, g_month, g_year);
    let f = (1.0 + eccentricity * true_anomaly_rad.cos()) / (1.0 - eccentricity * eccentricity);
    let r_km = 149598500.0 / f;
    let theta_deg = f * 0.533128;

    let sun_dist_km = utils::round_f64(r_km, 0);
    let sun_ang_size_deg = macros::dd_deg(theta_deg);
    let sun_ang_size_min = macros::dd_min(theta_deg);
    let sun_ang_size_sec = macros::dd_sec(theta_deg);

    return (
        sun_dist_km,
        sun_ang_size_deg,
        sun_ang_size_min,
        sun_ang_size_sec,
    );
}

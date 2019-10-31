use crate::lib::util as utils;

/// Convert a Civil Time (hours,minutes,seconds) to Decimal Hours
///
/// Original macro name: HMSDH
pub fn hms_dh(hours: f64, minutes: f64, seconds: f64) -> f64 {
    let f_hours = hours as f64;
    let f_minutes = minutes as f64;
    let f_seconds = seconds as f64;

    let a = f_seconds.abs() / 60.0;
    let b = (f_minutes.abs() + a) / 60.0;
    let c = f_hours.abs() + b;

    return if f_hours < 0.0 || f_minutes < 0.0 || f_seconds < 0.0 {
        -c
    } else {
        c
    };
}

/// Return the hour part of a Decimal Hours
///
/// Original macro name: DHHour
pub fn dh_hour(decimal_hours: f64) -> u32 {
    let a = decimal_hours.abs();
    let b = a * 3600.0;
    let c = utils::round_f64(b - 60.0 * (b / 60.0).floor(), 2);
    // let d = if c == 60.0 { 0.0 } else { c };
    let e = if c == 60.0 { b + 60.0 } else { b };

    return if decimal_hours < 0.0 {
        -(e / 3600.0).floor() as u32
    } else {
        (e / 3600.0).floor() as u32
    };
}

/// Return the minutes part of a Decimal Hours
///
/// Original macro name: DHMin
pub fn dh_min(decimal_hours: f64) -> u32 {
    let a = decimal_hours.abs();
    let b = a * 3600.0;
    let c = utils::round_f64(b - 60.0 * (b / 60.0).floor(), 2);
    let e = if c == 60.0 { b + 60.0 } else { b };

    return ((e / 60.0).floor() % 60.0) as u32;
}

/// Return the seconds part of a Decimal Hours
///
/// Original macro name: DHSec
pub fn dh_sec(decimal_hours: f64) -> f64 {
    let a = decimal_hours.abs();
    let b = a * 3600.0;
    let c = utils::round_f64(b - 60.0 * (b / 60.0).floor(), 2);
    let d = if c == 60.0 { 0.0 } else { c };

    return d;
}

/// Convert a Greenwich Date/Civil Date (day,month,year) to Julian Date
///
/// Original macro name: CDJD
pub fn cd_jd(day: f64, month: u32, year: u32) -> f64 {
    let f_day = day as f64;
    let f_month = month as f64;
    let f_year = year as f64;

    let y = if f_month < 3.0 { f_year - 1.0 } else { f_year };
    let m = if f_month < 3.0 {
        f_month + 12.0
    } else {
        f_month
    };

    let b: f64;

    if f_year > 1582.0 {
        let a = (y / 100.0).floor();
        b = 2.0 - a + (a / 4.0).floor();
    } else {
        if f_year == 1582.0 && f_month > 10.0 {
            let a = (y / 100.0).floor();
            b = 2.0 - a + (a / 4.0).floor();
        } else {
            if f_year == 1582.0 && f_month == 10.0 && f_day >= 15.0 {
                let a = (y / 100.0).floor();
                b = 2.0 - a + (a / 4.0).floor();
            } else {
                b = 0.0;
            }
        }
    }

    let c = if y < 0.0 {
        ((365.25 * y) - 0.75).floor()
    } else {
        (365.25 * y).floor()
    };

    let d = (30.6001 * (m + 1.0)).floor();

    return b + c + d + f_day + 1720994.5;
}

/// Returns the day part of a Julian Date
///
/// Original macro name: JDCDay
pub fn jdc_day(julian_date: f64) -> f64 {
    let i = (julian_date + 0.5).floor();
    let f = julian_date + 0.5 - i;
    let a = ((i - 1867216.25) / 36524.25).floor();
    let b = if i > 2299160.0 {
        i + 1.0 + a - (a / 4.0).floor()
    } else {
        i
    };
    let c = b + 1524.0;
    let d = ((c - 122.1) / 365.25).floor();
    let e = (365.25 * d).floor();
    let g = ((c - e) / 30.6001).floor();

    return c - e + f - (30.6001 * g).floor();
}

/// Returns the month part of a Julian Date
///
/// Original macro name: JDCMonth
pub fn jdc_month(julian_date: f64) -> u32 {
    let i = (julian_date + 0.5).floor();
    let _f = julian_date + 0.5 - i;
    let a = ((i - 1867216.25) / 36524.25).floor();
    let b = if i > 2299160.0 {
        i + 1.0 + a - (a / 4.0).floor()
    } else {
        i
    };
    let c = b + 1524.0;
    let d = ((c - 122.1) / 365.25).floor();
    let e = (365.25 * d).floor();
    let g = ((c - e) / 30.6001).floor();

    let return_value = if g < 13.5 { g - 1.0 } else { g - 13.0 };

    return return_value as u32;
}

/// Returns the year part of a Julian Date
///
/// Original macro name: JDCYear
pub fn jdc_year(julian_date: f64) -> u32 {
    let i = (julian_date + 0.5).floor();
    let _f = julian_date + 0.5 - i;
    let a = ((i - 1867216.25) / 36524.25).floor();
    let b = if i > 2299160.0 {
        i + 1.0 + a - (a / 4.0).floor()
    } else {
        i
    };
    let c = b + 1524.0;
    let d = ((c - 122.1) / 365.25).floor();
    let e = (365.25 * d).floor();
    let g = ((c - e) / 30.6001).floor();
    let h = if g < 13.5 { g - 1.0 } else { g - 13.0 };

    let return_value = if h > 2.5 { d - 4716.0 } else { d - 4715.0 };

    return return_value as u32;
}

/// Convert a Julian Date to Day-of-Week (e.g., Sunday)
///
/// Original macro name: FDOW
pub fn f_dow(julian_date: f64) -> String {
    let j = (julian_date - 0.5).floor() + 0.5;
    let n = (j + 1.5) % 7.0;

    let return_value: &str;
    match n as u32 {
        0 => return_value = "Sunday",
        1 => return_value = "Monday",
        2 => return_value = "Tuesday",
        3 => return_value = "Wednesday",
        4 => return_value = "Thursday",
        5 => return_value = "Friday",
        6 => return_value = "Saturday",
        _ => return_value = "Unknown",
    }

    return return_value.to_string();
}

/// Convert Right Ascension to Hour Angle
///
/// Original macro name: RAHA
pub fn ra_ha(
    ra_hours: f64,
    ra_minutes: f64,
    ra_seconds: f64,
    lct_hours: f64,
    lct_minutes: f64,
    lct_seconds: f64,
    daylight_saving: i32,
    zone_correction: i32,
    local_day: f64,
    local_month: u32,
    local_year: u32,
    geographical_longitude: f64,
) -> f64 {
    let a = lct_ut(
        lct_hours,
        lct_minutes,
        lct_seconds,
        daylight_saving,
        zone_correction,
        local_day,
        local_month,
        local_year,
    );
    let b = lct_gday(
        lct_hours,
        lct_minutes,
        lct_seconds,
        daylight_saving,
        zone_correction,
        local_day,
        local_month,
        local_year,
    );
    let c = lct_gmonth(
        lct_hours,
        lct_minutes,
        lct_seconds,
        daylight_saving,
        zone_correction,
        local_day,
        local_month,
        local_year,
    );
    let d = lct_gyear(
        lct_hours,
        lct_minutes,
        lct_seconds,
        daylight_saving,
        zone_correction,
        local_day,
        local_month,
        local_year,
    );
    let e = ut_gst(a, 0.0, 0.0, b, c, d);
    let f = gst_lst(e, 0.0, 0.0, geographical_longitude);
    let g = hms_dh(ra_hours, ra_minutes, ra_seconds);
    let h = f - g;

    return if h < 0.0 { 24.0 + h } else { h };
}

/// Convert Hour Angle to Right Ascension
///
/// Original macro name: HARA
pub fn ha_ra(
    hour_angle_hours: f64,
    hour_angle_minutes: f64,
    hour_angle_seconds: f64,
    lct_hours: f64,
    lct_minutes: f64,
    lct_seconds: f64,
    daylight_saving: i32,
    zone_correction: i32,
    local_day: f64,
    local_month: u32,
    local_year: u32,
    geographical_longitude: f64,
) -> f64 {
    let a = lct_ut(
        lct_hours,
        lct_minutes,
        lct_seconds,
        daylight_saving,
        zone_correction,
        local_day,
        local_month,
        local_year,
    );
    let b = lct_gday(
        lct_hours,
        lct_minutes,
        lct_seconds,
        daylight_saving,
        zone_correction,
        local_day,
        local_month,
        local_year,
    );
    let c = lct_gmonth(
        lct_hours,
        lct_minutes,
        lct_seconds,
        daylight_saving,
        zone_correction,
        local_day,
        local_month,
        local_year,
    );
    let d = lct_gyear(
        lct_hours,
        lct_minutes,
        lct_seconds,
        daylight_saving,
        zone_correction,
        local_day,
        local_month,
        local_year,
    );
    let e = ut_gst(a, 0.0, 0.0, b, c, d);
    let f = gst_lst(e, 0.0, 0.0, geographical_longitude);
    let g = hms_dh(hour_angle_hours, hour_angle_minutes, hour_angle_seconds);
    let h = f - g;

    return if h < 0.0 { 24.0 + h } else { h };
}

/// Convert Local Civil Time to Universal Time
///
/// Original macro name: LctUT
pub fn lct_ut(
    lct_hours: f64,
    lct_minutes: f64,
    lct_seconds: f64,
    daylight_saving: i32,
    zone_correction: i32,
    local_day: f64,
    local_month: u32,
    local_year: u32,
) -> f64 {
    let a = hms_dh(lct_hours, lct_minutes, lct_seconds as f64);
    let b = a - daylight_saving as f64 - zone_correction as f64;
    let c = local_day as f64 + (b / 24.0);
    let d = cd_jd(c, local_month, local_year);
    let e = jdc_day(d);
    let e1 = e.floor();

    return 24.0 * (e - e1);
}

/// Determine Greenwich Day for Local Time
///
/// Original macro name: LctGDay
pub fn lct_gday(
    lct_hours: f64,
    lct_minutes: f64,
    lct_seconds: f64,
    daylight_saving: i32,
    zone_correction: i32,
    local_day: f64,
    local_month: u32,
    local_year: u32,
) -> f64 {
    let a = hms_dh(lct_hours, lct_minutes, lct_seconds as f64);
    let b = a - daylight_saving as f64 - zone_correction as f64;
    let c = local_day as f64 + (b / 24.0);
    let d = cd_jd(c, local_month, local_year);
    let e = jdc_day(d);

    return e.floor();
}

/// Determine Greenwich Month for Local Time
///
/// Original macro name: LctGMonth
pub fn lct_gmonth(
    lct_hours: f64,
    lct_minutes: f64,
    lct_seconds: f64,
    daylight_saving: i32,
    zone_correction: i32,
    local_day: f64,
    local_month: u32,
    local_year: u32,
) -> u32 {
    let a = hms_dh(lct_hours, lct_minutes, lct_seconds as f64);
    let b = a - daylight_saving as f64 - zone_correction as f64;
    let c = local_day as f64 + (b / 24.0);
    let d = cd_jd(c, local_month, local_year);

    return jdc_month(d);
}

/// Determine Greenwich Year for Local Time
///
/// Original macro name: LctGYear
pub fn lct_gyear(
    lct_hours: f64,
    lct_minutes: f64,
    lct_seconds: f64,
    daylight_saving: i32,
    zone_correction: i32,
    local_day: f64,
    local_month: u32,
    local_year: u32,
) -> u32 {
    let a = hms_dh(lct_hours, lct_minutes, lct_seconds as f64);
    let b = a - daylight_saving as f64 - zone_correction as f64;
    let c = local_day as f64 + (b / 24.0);
    let d = cd_jd(c, local_month, local_year);

    return jdc_year(d);
}

/// Convert Universal Time to Greenwich Sidereal Time
///
/// Original macro name: UTGST
pub fn ut_gst(
    u_hours: f64,
    u_minutes: f64,
    u_seconds: f64,
    greenwich_day: f64,
    greenwich_month: u32,
    greenwich_year: u32,
) -> f64 {
    let a = cd_jd(greenwich_day as f64, greenwich_month, greenwich_year);
    let b = a - 2451545.0;
    let c = b / 36525.0;
    let d = 6.697374558 + (2400.051336 * c) + (0.000025862 * c * c);
    let e = d - (24.0 * (d / 24.0).floor());
    let f = hms_dh(u_hours, u_minutes, u_seconds);
    let g = f * 1.002737909;
    let h = e + g;
    return h - (24.0 * (h / 24.0).floor());
}

/// Convert Greenwich Sidereal Time to Local Sidereal Time
///
/// Original macro name: GSTLST
pub fn gst_lst(
    greenwich_hours: f64,
    greenwich_minutes: f64,
    greenwich_seconds: f64,
    geographical_longitude: f64,
) -> f64 {
    let a = hms_dh(greenwich_hours, greenwich_minutes, greenwich_seconds);
    let b = geographical_longitude / 15.0;
    let c = a + b;

    return c - (24.0 * (c / 24.0).floor());
}

/// Convert Equatorial Coordinates to Azimuth (in decimal degrees)
///
/// Original macro name: EQAz
pub fn eq_az(
    hour_angle_hours: f64,
    hour_angle_minutes: f64,
    hour_angle_seconds: f64,
    declination_degrees: f64,
    declination_minutes: f64,
    declination_seconds: f64,
    geographical_latitude: f64,
) -> f64 {
    let a = hms_dh(hour_angle_hours, hour_angle_minutes, hour_angle_seconds);
    let b = a * 15.0;
    let c = b.to_radians();
    let d = dms_dd(
        declination_degrees,
        declination_minutes,
        declination_seconds,
    );
    let e = d.to_radians();
    let f = geographical_latitude.to_radians();
    let g = e.sin() * f.sin() + e.cos() * f.cos() * c.cos();
    let h = -e.cos() * f.cos() * c.sin();
    let i = e.sin() - (f.sin() * g);
    let j = degrees(h.atan2(i));

    return j - 360.0 * (j / 360.0).floor();
}

/// Convert Degrees Minutes Seconds to Decimal Degrees
///
/// Original macro name: DMSDD
pub fn dms_dd(degrees: f64, minutes: f64, seconds: f64) -> f64 {
    let a = seconds.abs() / 60.0;
    let b = (minutes.abs() + a) / 60.0;
    let c = degrees.abs() + b;

    return if degrees < 0.0 || minutes < 0.0 || seconds < 0.0 {
        -c
    } else {
        c
    };
}

/// Convert W to Degrees
///
/// Original macro name: Degrees
pub fn degrees(w: f64) -> f64 {
    return w * 57.29577951;
}

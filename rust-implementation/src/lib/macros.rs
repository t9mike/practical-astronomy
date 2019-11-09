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

/// Convert Equatorial Coordinates to Altitude (in decimal degrees)
///
/// Original macro name: EQAlt
pub fn eq_alt(
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

    return degrees(g.asin());
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

/// Return Degrees part of Decimal Degrees
///
/// Original macro name: DDDeg
pub fn dd_deg(decimal_degrees: f64) -> f64 {
    let a = decimal_degrees.abs();
    let b = a * 3600.0;
    let c = utils::round_f64(b - 60.0 * (b / 60.0).floor(), 2);
    let _d = if c == 60.0 { 0.0 } else { c };
    let e = if c == 60.0 { 60.0 } else { b };

    return if decimal_degrees < 0.0 {
        -(e / 3600.0).floor()
    } else {
        (e / 3600.0).floor()
    };
}

/// Return Minutes part of Decimal Degrees
///
/// Original macro name: DDMin
pub fn dd_min(decimal_degrees: f64) -> f64 {
    let a = decimal_degrees.abs();
    let b = a * 3600.0;
    let c = utils::round_f64(b - 60.0 * (b / 60.0).floor(), 2);
    let _d = if c == 60.0 { 0.0 } else { c };
    let e = if c == 60.0 { b + 60.0 } else { b };

    return (e / 60.0).floor() % 60.0;
}

/// Return Seconds part of Decimal Degrees
///
/// Original macro name: DDSec
pub fn dd_sec(decimal_degrees: f64) -> f64 {
    let a = decimal_degrees.abs();
    let b = a * 3600.0;
    let c = utils::round_f64(b - 60.0 * (b / 60.0).floor(), 2);
    let d = if c == 60.0 { 0.0 } else { c };

    return d;
}

/// Convert Decimal Degrees to Degree-Hours
///
/// Original macro name: DDDH
pub fn dd_dh(decimal_degrees: f64) -> f64 {
    return decimal_degrees / 15.0;
}

/// Convert Degree-Hours to Decimal Degrees
///
/// Original macro name: DHDD
pub fn dh_dd(degree_hours: f64) -> f64 {
    return degree_hours * 15.0;
}

/// Convert Horizon Coordinates to Declination (in decimal degrees)
///
/// Original macro name: HORDec
pub fn hor_dec(
    azimuth_degrees: f64,
    azimuth_minutes: f64,
    azimuth_seconds: f64,
    altitude_degrees: f64,
    altitude_minutes: f64,
    altitude_seconds: f64,
    geographical_latitude: f64,
) -> f64 {
    let a = dms_dd(azimuth_degrees, azimuth_minutes, azimuth_seconds);
    let b = dms_dd(altitude_degrees, altitude_minutes, altitude_seconds);
    let c = a.to_radians();
    let d = b.to_radians();
    let e = geographical_latitude.to_radians();
    let f = d.sin() * e.sin() + d.cos() * e.cos() * c.cos();

    return degrees(f.asin());
}

/// Convert Horizon Coordinates to Hour Angle (in decimal degrees)
///
/// Original macro name: HORHa
pub fn hor_ha(
    azimuth_degrees: f64,
    azimuth_minutes: f64,
    azimuth_seconds: f64,
    altitude_degrees: f64,
    altitude_minutes: f64,
    altitude_seconds: f64,
    geographical_latitude: f64,
) -> f64 {
    let a = dms_dd(azimuth_degrees, azimuth_minutes, azimuth_seconds);
    let b = dms_dd(altitude_degrees, altitude_minutes, altitude_seconds);
    let c = a.to_radians();
    let d = b.to_radians();
    let e = geographical_latitude.to_radians();
    let f = d.sin() * e.sin() + d.cos() * e.cos() * c.cos();
    let g = -d.cos() * e.cos() * c.sin();
    let h = d.sin() - e.sin() * f;
    let i = dd_dh(degrees(g.atan2(h)));

    return i - 24.0 * (i / 24.0).floor();
}

/// Nutation of Obliquity
///
/// Original macro name: NutatObl
pub fn nutat_obl(greenwich_day: f64, greenwich_month: u32, greenwich_year: u32) -> f64 {
    let dj = cd_jd(greenwich_day, greenwich_month, greenwich_year) - 2415020.0;
    let t = dj / 36525.0;
    let t2 = t * t;

    let a = 100.0021358 * t;
    let b = 360.0 * (a - a.floor());

    let l1 = 279.6967 + 0.000303 * t2 + b;
    let l2 = 2.0 * l1.to_radians();

    let a = 1336.855231 * t;
    let b = 360.0 * (a - a.floor());

    let d1 = 270.4342 - 0.001133 * t2 + b;
    let d2 = 2.0 * d1.to_radians();

    let a = 99.99736056 * t;
    let b = 360.0 * (a - a.floor());

    let m1 = (358.4758 - 0.00015 * t2 + b).to_radians();
    //M1 = math.radians(M1)

    let a = 1325.552359 * t;
    let b = 360.0 * (a - a.floor());

    let m2 = (296.1046 + 0.009192 * t2 + b).to_radians();
    // M2 = math.radians(M2)

    let a = 5.372616667 * t;
    let b = 360.0 * (a - a.floor());

    let n1 = (259.1833 + 0.002078 * t2 - b).to_radians();
    //	N1 = math.radians(N1)

    let n2 = 2.0 * n1;

    let ddo = ((9.21 + 0.00091 * t) * n1.cos())
        + ((0.5522 - 0.00029 * t) * l2.cos() - 0.0904 * n2.cos())
        + (0.0884 * d2.cos() + 0.0216 * (l2 + m1).cos())
        + (0.0183 * (d2 - n1).cos() + 0.0113 * (d2 + m2).cos())
        - (0.0093 * (l2 - m1).cos() - 0.0066 * (l2 - n1).cos());

    return ddo / 3600.0;
}

/// Obliquity of the Ecliptic for a Greenwich Date
///
/// Original macro name: Obliq
pub fn obliq(greenwich_day: f64, greenwich_month: u32, greenwich_year: u32) -> f64 {
    let a = cd_jd(greenwich_day, greenwich_month, greenwich_year);
    let b = a - 2415020.0;
    let c = (b / 36525.0) - 1.0;
    let d = c * (46.815 + c * (0.0006 - (c * 0.00181)));
    let e = d / 3600.0;

    return 23.43929167 - e + nutat_obl(greenwich_day, greenwich_month, greenwich_year);
}

/// Convert Local Sidereal Time to Greenwich Sidereal Time
///
/// Original macro name: LSTGST
pub fn lst_gst(local_hours: f64, local_minutes: f64, local_seconds: f64, longitude: f64) -> f64 {
    let a = hms_dh(local_hours, local_minutes, local_seconds);
    let b = longitude / 15.0;
    let c = a - b;
    return c - (24.0 * (c / 24.0).floor());
}

/// Convert Greenwich Sidereal Time to Universal Time
///
/// Original macro name: GSTUT
pub fn gst_ut(
    greenwich_sidereal_hours: f64,
    greenwich_sidereal_minutes: f64,
    greenwich_sidereal_seconds: f64,
    greenwich_day: f64,
    greenwich_month: u32,
    greenwich_year: u32,
) -> f64 {
    let a = cd_jd(greenwich_day, greenwich_month, greenwich_year);
    let b = a - 2451545.0;
    let c = b / 36525.0;
    let d = 6.697374558 + (2400.051336 * c) + (0.000025862 * c * c);
    let e = d - (24.0 * (d / 24.0).floor());
    let f = hms_dh(
        greenwich_sidereal_hours,
        greenwich_sidereal_minutes,
        greenwich_sidereal_seconds,
    );
    let g = f - e;
    let h = g - (24.0 * (g / 24.0).floor());
    return h * 0.9972695663;
}

/// Calculate Sun's ecliptic longitude
///
/// Original macro name: SunLong
pub fn sun_long(lch: f64, lcm: f64, lcs: f64, ds: i32, zc: i32, ld: f64, lm: u32, ly: u32) -> f64 {
    let aa = lct_gday(lch, lcm, lcs, ds, zc, ld, lm, ly);
    let bb = lct_gmonth(lch, lcm, lcs, ds, zc, ld, lm, ly);
    let cc = lct_gyear(lch, lcm, lcs, ds, zc, ld, lm, ly);
    let ut = lct_ut(lch, lcm, lcs, ds, zc, ld, lm, ly);
    let dj = cd_jd(aa, bb, cc) - 2415020.0;
    let t = (dj / 36525.0) + (ut / 876600.0);
    let t2 = t * t;
    let a = 100.0021359 * t;
    let b = 360.0 * (a - (a).floor());

    let l = 279.69668 + 0.0003025 * t2 + b;
    let a = 99.99736042 * t;
    let b = 360.0 * (a - a.floor());

    let m1 = 358.47583 - (0.00015 + 0.0000033 * t) * t2 + b;
    let ec = 0.01675104 - 0.0000418 * t - 0.000000126 * t2;

    let am = m1.to_radians();
    let at = true_anomaly(am, ec);
    let _ae = eccentric_anomaly(am, ec);

    let a = 62.55209472 * t;
    let b = 360.0 * (a - (a).floor());

    let a1 = (153.23 + b).to_radians();
    let a = 125.1041894 * t;
    let _b = 360.0 * (a - (a).floor());

    let b1 = (216.57 + t).to_radians();
    let a = 91.56766028 * t;
    let b = 360.0 * (a - (a).floor());

    let c1 = (312.69 + b).to_radians();
    let a = 1236.853095 * t;
    let b = 360.0 * (a - (a).floor());

    let d1 = (350.74 - 0.00144 * t2 + b).to_radians();
    let e1 = (231.19 + 20.2 * t).to_radians();
    let a = 183.1353208 * t;
    let b = 360.0 * (a - (a).floor());
    let h1 = (353.4 + b).to_radians();

    let d2 = 0.00134 * a1.cos() + 0.00154 * b1.cos() + 0.002 * c1.cos();
    let d2 = d2 + 0.00179 * d1.sin() + 0.00178 * e1.sin();
    let d3 = 0.00000543 * a1.sin() + 0.00001575 * b1.sin();
    let d3 = d3 + 0.00001627 * c1.sin() + 0.00003076 * d1.cos();
    let _d3 = d3 + 0.00000927 * h1.sin();

    let sr = at + (l - m1 + d2).to_radians();
    let tp = 6.283185308;

    let sr = sr - tp * (sr / tp).floor();
    return degrees(sr);
}

/// Solve Kepler's equation, and return value of the true anomaly in radians
///
/// Original macro name: TrueAnomaly
pub fn true_anomaly(am: f64, ec: f64) -> f64 {
    let tp = 6.283185308;
    let m = am - tp * (am / tp).floor();
    let mut ae = m;

    while 1 == 1 {
        let d = ae - (ec * (ae).sin()) - m;
        if d.abs() < 0.000001 {
            break;
        }
        let d = d / (1.0 - (ec * (ae).cos()));
        ae = ae - d;
    }

    let a = ((1.0 + ec) / (1.0 - ec)).sqrt() * (ae / 2.0).tan();
    let at = 2.0 * a.atan();

    return at;
}

/// Solve Kepler's equation, and return value of the eccentric anomaly in radians
///
/// Original macro name: EccentricAnomaly
pub fn eccentric_anomaly(am: f64, ec: f64) -> f64 {
    let tp = 6.283185308;
    let m = am - tp * (am / tp).floor();
    let mut ae = m;

    while 1 == 1 {
        let d = ae - (ec * (ae).sin()) - m;

        if d.abs() < 0.000001 {
            break;
        }

        let d = d / (1.0 - (ec * (ae).cos()));
        ae = ae - d;
    }

    return ae;
}

/// Calculate effects of refraction
///
/// Original macro name: Refract
pub fn refract(y2: f64, sw: String, pr: f64, tr: f64) -> f64 {
    let y = y2.to_radians();

    let d = if &sw[..1].to_string().to_lowercase() == "t" {
        -1.0
    } else {
        1.0
    };

    if d == -1.0 {
        let y3 = y;
        let y1 = y;
        let mut r1 = 0.0;

        while 1 == 1 {
            let y = y1 + r1;
            let _q = y;
            let rf = refract_l3035(pr, tr, y, d);
            if y < -0.087 {
                return 0.0;
            }
            let r2 = rf;

            if (r2 == 0.0) || ((r2 - r1).abs() < 0.000001) {
                let q = y3;
                return degrees(q + rf);
            }

            r1 = r2;
        }
    }

    let rf = refract_l3035(pr, tr, y, d);

    if y < -0.087 {
        return 0.0;
    }

    let q = y;

    return degrees(q + rf);
}

/// Helper function for refract
pub fn refract_l3035(pr: f64, tr: f64, y: f64, d: f64) -> f64 {
    if y < 0.2617994 {
        if y < -0.087 {
            return 0.0;
        }

        let yd = degrees(y);
        let a = ((0.00002 * yd + 0.0196) * yd + 0.1594) * pr;
        let b = (273.0 + tr) * ((0.0845 * yd + 0.505) * yd + 1.0);

        return (-(a / b) * d).to_radians();
    }

    return -d * 0.00007888888 * pr / ((273.0 + tr) * (y).tan());
}

/// Calculate corrected hour angle in decimal hours
///
/// Original macro name: ParallaxHA
pub fn parallax_ha(
    hh: f64,
    hm: f64,
    hs: f64,
    dd: f64,
    dm: f64,
    ds: f64,
    sw: String,
    gp: f64,
    ht: f64,
    hp: f64,
) -> f64 {
    let a = gp.to_radians();
    let c1 = a.cos();
    let s1 = a.sin();

    let u = (0.996647 * s1 / c1).atan();
    let c2 = u.cos();
    let s2 = u.sin();
    let b = ht / 6378160.0;

    let rs = (0.996647 * s2) + (b * s1);

    let rc = c2 + (b * c1);
    let tp = 6.283185308;

    let rp = 1.0 / hp.to_radians().sin();

    let x = (dh_dd(hms_dh(hh, hm, hs))).to_radians();
    let x1 = x;
    let y = (dms_dd(dd, dm, ds)).to_radians();
    let y1 = y;

    let d = if &sw[..1].to_string().to_lowercase() == "t" {
        1.0
    } else {
        -1.0
    };

    if d == 1.0 {
        let (p, _q) = parallax_ha_l2870(x, y, rc, rp, rs, tp);
        return dd_dh(degrees(p));
    }

    let mut p1 = 0.0;
    let mut q1 = 0.0;
    let mut x_loop = x;
    let mut y_loop = y;
    while 1 == 1 {
        let (p, q) = parallax_ha_l2870(x_loop, y_loop, rc, rp, rs, tp);
        let p2 = p - x_loop;
        let q2 = q - y_loop;

        let aa = (p2 - p1).abs();
        let bb = (q2 - q1).abs();

        if (aa < 0.000001) && (bb < 0.000001) {
            let p = x1 - p2;
            let _q = y1 - q2;
            let _x_loop = x1;
            let _y_loop = y1;

            return dd_dh(degrees(p));
        }
        x_loop = x1 - p2;
        y_loop = y1 - q2;
        p1 = p2;
        q1 = q2;
    }

    return dd_dh(degrees(0.0));
}

/// Helper function for parallax_ha
pub fn parallax_ha_l2870(x: f64, y: f64, rc: f64, rp: f64, rs: f64, tp: f64) -> (f64, f64) {
    let cx = x.cos();
    let sy = y.sin();
    let cy = y.cos();

    let aa = (rc * x.sin()) / ((rp * cy) - (rc * cx));

    let dx = aa.atan();
    let p = x + dx;
    let cp = p.cos();

    let p = p - tp * (p / tp).floor();
    let q = (cp * (rp * sy - rs) / (rp * cy * cx - rc)).atan();

    return (p, q);
}

/// Calculate corrected declination in decimal degrees
///
/// Original macro name: ParallaxDec
/// HH,HM,HS,DD,DM,DS,SW,GP,HT,HP
pub fn parallax_dec(
    hh: f64,
    hm: f64,
    hs: f64,
    dd: f64,
    dm: f64,
    ds: f64,
    sw: String,
    gp: f64,
    ht: f64,
    hp: f64,
) -> f64 {
    let a = gp.to_radians();
    let c1 = a.cos();
    let s1 = a.sin();

    let u = (0.996647 * s1 / c1).atan();

    let c2 = u.cos();
    let s2 = u.sin();
    let b = ht / 6378160.0;
    let rs = (0.996647 * s2) + (b * s1);

    let rc = c2 + (b * c1);
    let tp = 6.283185308;

    let rp = 1.0 / hp.to_radians().sin();

    let x = (dh_dd(hms_dh(hh, hm, hs))).to_radians();
    let x1 = x;

    let y = (dms_dd(dd, dm, ds)).to_radians();
    let y1 = y;
    let d = if &sw[..1].to_string().to_lowercase() == "t" {
        1.0
    } else {
        -1.0
    };

    if d == 1.0 {
        let (_p, q) = parallax_dec_l2870(x, y, rc, rp, rs, tp);
        return degrees(q);
    }

    let mut p1 = 0.0;
    let mut q1 = 0.0;

    let mut x_loop = x;
    let mut y_loop = y;
    while 1 == 1 {
        let (p, q) = parallax_dec_l2870(x_loop, y_loop, rc, rp, rs, tp);
        let p2 = p - x_loop;
        let q2 = q - y_loop;
        let aa = (p2 - p1).abs();
        let _bb = (q2 - q1).abs();
        if (aa < 0.000001) && (b < 0.000001) {
            let _p = x1 - p2;
            let q = y1 - q2;
            let _x_loop = x1;
            let _y_loop = y1;
            return degrees(q);
        }
        x_loop = x1 - p2;
        y_loop = y1 - q2;
        p1 = p2;
        q1 = q2;
    }

    return degrees(0.0);
}

/// Helper function for parallax_dec
pub fn parallax_dec_l2870(x: f64, y: f64, rc: f64, rp: f64, rs: f64, tp: f64) -> (f64, f64) {
    let cx = x.cos();
    let sy = y.sin();
    let cy = y.cos();

    let aa = (rc * x.sin()) / ((rp * cy) - (rc * cx));
    let dx = aa.atan();
    let p = x + dx;
    let cp = p.cos();

    let p = p - tp * (p / tp).floor();
    let q = (cp * (rp * sy - rs) / (rp * cy * cx - rc)).atan();

    return (p, q);
}

/// Calculate Sun's angular diameter in decimal degrees
///
/// Original macro name: SunDia
pub fn sun_dia(lch: f64, lcm: f64, lcs: f64, ds: i32, zc: i32, ld: f64, lm: u32, ly: u32) -> f64 {
    let a = sun_dist(lch, lcm, lcs, ds, zc, ld, lm, ly);

    return 0.533128 / a;
}

/// Calculate Sun's distance from the Earth in astronomical units
///
/// Original macro name: SunDist
pub fn sun_dist(lch: f64, lcm: f64, lcs: f64, ds: i32, zc: i32, ld: f64, lm: u32, ly: u32) -> f64 {
    let aa = lct_gday(lch, lcm, lcs, ds, zc, ld, lm, ly);
    let bb = lct_gmonth(lch, lcm, lcs, ds, zc, ld, lm, ly);
    let cc = lct_gyear(lch, lcm, lcs, ds, zc, ld, lm, ly);
    let ut = lct_ut(lch, lcm, lcs, ds, zc, ld, lm, ly);
    let dj = cd_jd(aa, bb, cc) - 2415020.0;

    let t = (dj / 36525.0) + (ut / 876600.0);
    let t2 = t * t;

    let a = 100.0021359 * t;
    let b = 360.0 * (a - a.floor());
    let _l = 279.69668 + 0.0003025 * t2 + b;
    let a = 99.99736042 * t;
    let b = 360.0 * (a - (a).floor());
    let m1 = 358.47583 - (0.00015 + 0.0000033 * t) * t2 + b;
    let ec = 0.01675104 - 0.0000418 * t - 0.000000126 * t2;

    let am = m1.to_radians();
    let _at = true_anomaly(am, ec);
    let ae = eccentric_anomaly(am, ec);

    let a = 62.55209472 * t;
    let b = 360.0 * (a - a.floor());
    let a1 = (153.23 + b).to_radians();
    let a = 125.1041894 * t;
    let b = 360.0 * (a - a.floor());
    let b1 = (216.57 + b).to_radians();
    let a = 91.56766028 * t;
    let b = 360.0 * (a - a.floor());
    let c1 = (312.69 + b).to_radians();
    let a = 1236.853095 * t;
    let b = 360.0 * (a - a.floor());
    let d1 = (350.74 - 0.00144 * t2 + b).to_radians();
    let e1 = (231.19 + 20.2 * t).to_radians();
    let a = 183.1353208 * t;
    let b = 360.0 * (a - a.floor());
    let h1 = (353.4 + b).to_radians();

    let _d2 = (0.00134 * a1.cos() + 0.00154 * b1.cos() + 0.002 * c1.cos())
        + (0.00179 * d1.sin() + 0.00178 * e1.sin());
    let d3 = (0.00000543 * a1.sin() + 0.00001575 * b1.sin())
        + (0.00001627 * c1.sin() + 0.00003076 * d1.cos())
        + (0.00000927 * h1.sin());

    return 1.0000002 * (1.0 - ec * ae.cos()) + d3;
}

/// Calculate geocentric ecliptic longitude for the Moon
///
/// Original macro name: MoonLong
pub fn moon_long(lh: f64, lm: f64, ls: f64, ds: i32, zc: i32, dy: f64, mn: u32, yr: u32) -> f64 {
    let ut = lct_ut(lh, lm, ls, ds, zc, dy, mn, yr);
    let gd = lct_gday(lh, lm, ls, ds, zc, dy, mn, yr);
    let gm = lct_gmonth(lh, lm, ls, ds, zc, dy, mn, yr);
    let gy = lct_gyear(lh, lm, ls, ds, zc, dy, mn, yr);
    let t = ((cd_jd(gd, gm, gy) - 2415020.0) / 36525.0) + (ut / 876600.0);
    let t2 = t * t;

    let m1 = 27.32158213;
    let m2 = 365.2596407;
    let m3 = 27.55455094;
    let m4 = 29.53058868;
    let m5 = 27.21222039;
    let m6 = 6798.363307;
    let q = cd_jd(gd, gm, gy) - 2415020.0 + (ut / 24.0);
    let m1 = q / m1;
    let m2 = q / m2;
    let m3 = q / m3;
    let m4 = q / m4;
    let m5 = q / m5;
    let m6 = q / m6;
    let m1 = 360.0 * (m1 - (m1).floor());
    let m2 = 360.0 * (m2 - (m2).floor());
    let m3 = 360.0 * (m3 - (m3).floor());
    let m4 = 360.0 * (m4 - (m4).floor());
    let m5 = 360.0 * (m5 - (m5).floor());
    let m6 = 360.0 * (m6 - (m6).floor());

    let ml = 270.434164 + m1 - (0.001133 - 0.0000019 * t) * t2;
    let ms = 358.475833 + m2 - (0.00015 + 0.0000033 * t) * t2;
    let md = 296.104608 + m3 + (0.009192 + 0.0000144 * t) * t2;
    let me1 = 350.737486 + m4 - (0.001436 - 0.0000019 * t) * t2;
    let mf = 11.250889 + m5 - (0.003211 + 0.0000003 * t) * t2;
    let na = 259.183275 - m6 + (0.002078 + 0.0000022 * t) * t2;
    let a = (51.2 + 20.2 * t).to_radians();
    let s1 = a.sin();
    let s2 = ((na).to_radians()).sin();
    let b = 346.56 + (132.87 - 0.0091731 * t) * t;
    let s3 = 0.003964 * ((b).to_radians()).sin();
    let c = (na + 275.05 - 2.3 * t).to_radians();
    let s4 = c.sin();
    let ml = ml + 0.000233 * s1 + s3 + 0.001964 * s2;
    let ms = ms - 0.001778 * s1;
    let md = md + 0.000817 * s1 + s3 + 0.002541 * s2;
    let mf = mf + s3 - 0.024691 * s2 - 0.004328 * s4;
    let me1 = me1 + 0.002011 * s1 + s3 + 0.001964 * s2;
    let e = 1.0 - (0.002495 + 0.00000752 * t) * t;
    let e2 = e * e;
    let ml = (ml).to_radians();
    let ms = ms.to_radians();
    let _na = na.to_radians();
    let me1 = me1.to_radians();
    let mf = mf.to_radians();
    let md = md.to_radians();

    let l = 6.28875 * (md).sin() + 1.274018 * (2.0 * me1 - md).sin();
    let l = l + 0.658309 * (2.0 * me1).sin() + 0.213616 * (2.0 * md).sin();
    let l = l - e * 0.185596 * (ms).sin() - 0.114336 * (2.0 * mf).sin();
    let l = l + 0.058793 * (2.0 * (me1 - md)).sin();
    let l = l + 0.057212 * e * (2.0 * me1 - ms - md).sin() + 0.05332 * (2.0 * me1 + md).sin();
    let l = l + 0.045874 * e * (2.0 * me1 - ms).sin() + 0.041024 * e * (md - ms).sin();
    let l = l - 0.034718 * (me1).sin() - e * 0.030465 * (ms + md).sin();
    let l = l + 0.015326 * (2.0 * (me1 - mf)).sin() - 0.012528 * (2.0 * mf + md).sin();
    let l = l - 0.01098 * (2.0 * mf - md).sin() + 0.010674 * (4.0 * me1 - md).sin();
    let l = l + 0.010034 * (3.0 * md).sin() + 0.008548 * (4.0 * me1 - 2.0 * md).sin();
    let l = l - e * 0.00791 * (ms - md + 2.0 * me1).sin() - e * 0.006783 * (2.0 * me1 + ms).sin();
    let l = l + 0.005162 * (md - me1).sin() + e * 0.005 * (ms + me1).sin();
    let l = l + 0.003862 * (4.0 * me1).sin() + e * 0.004049 * (md - ms + 2.0 * me1).sin();
    let l = l + 0.003996 * (2.0 * (md + me1)).sin() + 0.003665 * (2.0 * me1 - 3.0 * md).sin();
    let l = l + e * 0.002695 * (2.0 * md - ms).sin() + 0.002602 * (md - 2.0 * (mf + me1)).sin();
    let l = l + e * 0.002396 * (2.0 * (me1 - md) - ms).sin() - 0.002349 * (md + me1).sin();
    let l = l + e2 * 0.002249 * (2.0 * (me1 - ms)).sin() - e * 0.002125 * (2.0 * md + ms).sin();
    let l = l - e2 * 0.002079 * (2.0 * ms).sin() + e2 * 0.002059 * (2.0 * (me1 - ms) - md).sin();
    let l = l - 0.001773 * (md + 2.0 * (me1 - mf)).sin() - 0.001595 * (2.0 * (mf + me1)).sin();
    let l = l + e * 0.00122 * (4.0 * me1 - ms - md).sin() - 0.00111 * (2.0 * (md + mf)).sin();
    let l = l + 0.000892 * (md - 3.0 * me1).sin() - e * 0.000811 * (ms + md + 2.0 * me1).sin();
    let l = l + e * 0.000761 * (4.0 * me1 - ms - 2.0 * md).sin();
    let l = l + e2 * 0.000704 * (md - 2.0 * (ms + me1)).sin();
    let l = l + e * 0.000693 * (ms - 2.0 * (md - me1)).sin();
    let l = l + e * 0.000598 * (2.0 * (me1 - mf) - ms).sin();
    let l = l + 0.00055 * (md + 4.0 * me1).sin() + 0.000538 * (4.0 * md).sin();
    let l = l + e * 0.000521 * (4.0 * me1 - ms).sin() + 0.000486 * (2.0 * md - me1).sin();
    let l = l + e2 * 0.000717 * (md - 2.0 * ms).sin();
    let mm = unwind(ml + l.to_radians());

    return degrees(mm);
}

/// Calculate geocentric ecliptic latitude for the Moon
///
/// Original macro name: MoonLat
pub fn moon_lat(lh: f64, lm: f64, ls: f64, ds: i32, zc: i32, dy: f64, mn: u32, yr: u32) -> f64 {
    let ut = lct_ut(lh, lm, ls, ds, zc, dy, mn, yr);
    let gd = lct_gday(lh, lm, ls, ds, zc, dy, mn, yr);
    let gm = lct_gmonth(lh, lm, ls, ds, zc, dy, mn, yr);
    let gy = lct_gyear(lh, lm, ls, ds, zc, dy, mn, yr);
    let t = ((cd_jd(gd, gm, gy) - 2415020.0) / 36525.0) + (ut / 876600.0);
    let t2 = t * t;

    let m1 = 27.32158213;
    let m2 = 365.2596407;
    let m3 = 27.55455094;
    let m4 = 29.53058868;
    let m5 = 27.21222039;
    let m6 = 6798.363307;
    let q = cd_jd(gd, gm, gy) - 2415020.0 + (ut / 24.0);
    let m1 = q / m1;
    let m2 = q / m2;
    let m3 = q / m3;
    let m4 = q / m4;
    let m5 = q / m5;
    let m6 = q / m6;
    let m1 = 360.0 * (m1 - (m1).floor());
    let m2 = 360.0 * (m2 - (m2).floor());
    let m3 = 360.0 * (m3 - (m3).floor());
    let m4 = 360.0 * (m4 - (m4).floor());
    let m5 = 360.0 * (m5 - (m5).floor());
    let m6 = 360.0 * (m6 - (m6).floor());

    let ml = 270.434164 + m1 - (0.001133 - 0.0000019 * t) * t2;
    let ms = 358.475833 + m2 - (0.00015 + 0.0000033 * t) * t2;
    let md = 296.104608 + m3 + (0.009192 + 0.0000144 * t) * t2;
    let me1 = 350.737486 + m4 - (0.001436 - 0.0000019 * t) * t2;
    let mf = 11.250889 + m5 - (0.003211 + 0.0000003 * t) * t2;
    let na = 259.183275 - m6 + (0.002078 + 0.0000022 * t) * t2;
    let a = (51.2 + 20.2 * t).to_radians();
    let s1 = (a).sin();
    let s2 = na.to_radians().sin();
    let b = 346.56 + (132.87 - 0.0091731 * t) * t;
    let s3 = 0.003964 * b.to_radians().sin();
    let c = (na + 275.05 - 2.3 * t).to_radians();
    let s4 = (c).sin();
    let ml = ml + 0.000233 * s1 + s3 + 0.001964 * s2;
    let ms = ms - 0.001778 * s1;
    let md = md + 0.000817 * s1 + s3 + 0.002541 * s2;
    let mf = mf + s3 - 0.024691 * s2 - 0.004328 * s4;
    let me1 = me1 + 0.002011 * s1 + s3 + 0.001964 * s2;
    let e = 1.0 - (0.002495 + 0.00000752 * t) * t;
    let e2 = e * e;
    let _ml = (ml).to_radians();
    let ms = (ms).to_radians();
    let na = (na).to_radians();
    let me1 = (me1).to_radians();
    let mf = (mf).to_radians();
    let md = (md).to_radians();

    let g = 5.128189 * (mf).sin() + 0.280606 * (md + mf).sin();
    let g = g + 0.277693 * (md - mf).sin() + 0.173238 * (2.0 * me1 - mf).sin();
    let g = g + 0.055413 * (2.0 * me1 + mf - md).sin() + 0.046272 * (2.0 * me1 - mf - md).sin();
    let g = g + 0.032573 * (2.0 * me1 + mf).sin() + 0.017198 * (2.0 * md + mf).sin();
    let g = g + 0.009267 * (2.0 * me1 + md - mf).sin() + 0.008823 * (2.0 * md - mf).sin();
    let g =
        g + e * 0.008247 * (2.0 * me1 - ms - mf).sin() + 0.004323 * (2.0 * (me1 - md) - mf).sin();
    let g = g + 0.0042 * (2.0 * me1 + mf + md).sin() + e * 0.003372 * (mf - ms - 2.0 * me1).sin();
    let g = g + e * 0.002472 * (2.0 * me1 + mf - ms - md).sin();
    let g = g + e * 0.002222 * (2.0 * me1 + mf - ms).sin();
    let g = g + e * 0.002072 * (2.0 * me1 - mf - ms - md).sin();
    let g = g + e * 0.001877 * (mf - ms + md).sin() + 0.001828 * (4.0 * me1 - mf - md).sin();
    let g = g - e * 0.001803 * (mf + ms).sin() - 0.00175 * (3.0 * mf).sin();
    let g = g + e * 0.00157 * (md - ms - mf).sin() - 0.001487 * (mf + me1).sin();
    let g = g - e * 0.001481 * (mf + ms + md).sin() + e * 0.001417 * (mf - ms - md).sin();
    let g = g + e * 0.00135 * (mf - ms).sin() + 0.00133 * (mf - me1).sin();
    let g = g + 0.001106 * (mf + 3.0 * md).sin() + 0.00102 * (4.0 * me1 - mf).sin();
    let g = g + 0.000833 * (mf + 4.0 * me1 - md).sin() + 0.000781 * (md - 3.0 * mf).sin();
    let g =
        g + 0.00067 * (mf + 4.0 * me1 - 2.0 * md).sin() + 0.000606 * (2.0 * me1 - 3.0 * mf).sin();
    let g = g + 0.000597 * (2.0 * (me1 + md) - mf).sin();
    let g = g
        + e * 0.000492 * (2.0 * me1 + md - ms - mf).sin()
        + 0.00045 * (2.0 * (md - me1) - mf).sin();
    let g = g + 0.000439 * (3.0 * md - mf).sin() + 0.000423 * (mf + 2.0 * (me1 + md)).sin();
    let g = g + 0.000422 * (2.0 * me1 - mf - 3.0 * md).sin()
        - e * 0.000367 * (ms + mf + 2.0 * me1 - md).sin();
    let g = g - e * 0.000353 * (ms + mf + 2.0 * me1).sin() + 0.000331 * (mf + 4.0 * me1).sin();
    let g = g + e * 0.000317 * (2.0 * me1 + mf - ms + md).sin();
    let g = g + e2 * 0.000306 * (2.0 * (me1 - ms) - mf).sin() - 0.000283 * (md + 3.0 * mf).sin();
    let w1 = 0.0004664 * (na).cos();
    let w2 = 0.0000754 * (c).cos();
    let bm = (g).to_radians() * (1.0 - w1 - w2);

    return degrees(bm);
}

/// Calculate horizontal parallax for the Moon
///
/// Original macro name: MoonHP
pub fn moon_hp(lh: f64, lm: f64, ls: f64, ds: i32, zc: i32, dy: f64, mn: u32, yr: u32) -> f64 {
    let ut = lct_ut(lh, lm, ls, ds, zc, dy, mn, yr);
    let gd = lct_gday(lh, lm, ls, ds, zc, dy, mn, yr);
    let gm = lct_gmonth(lh, lm, ls, ds, zc, dy, mn, yr);
    let gy = lct_gyear(lh, lm, ls, ds, zc, dy, mn, yr);
    let t = ((cd_jd(gd, gm, gy) - 2415020.0) / 36525.0) + (ut / 876600.0);
    let t2 = t * t;

    let m1 = 27.32158213;
    let m2 = 365.2596407;
    let m3 = 27.55455094;
    let m4 = 29.53058868;
    let m5 = 27.21222039;
    let m6 = 6798.363307;
    let q = cd_jd(gd, gm, gy) - 2415020.0 + (ut / 24.0);
    let m1 = q / m1;
    let m2 = q / m2;
    let m3 = q / m3;
    let m4 = q / m4;
    let m5 = q / m5;
    let m6 = q / m6;
    let m1 = 360.0 * (m1 - (m1).floor());
    let m2 = 360.0 * (m2 - (m2).floor());
    let m3 = 360.0 * (m3 - (m3).floor());
    let m4 = 360.0 * (m4 - (m4).floor());
    let m5 = 360.0 * (m5 - (m5).floor());
    let m6 = 360.0 * (m6 - (m6).floor());

    let ml = 270.434164 + m1 - (0.001133 - 0.0000019 * t) * t2;
    let ms = 358.475833 + m2 - (0.00015 + 0.0000033 * t) * t2;
    let md = 296.104608 + m3 + (0.009192 + 0.0000144 * t) * t2;
    let me1 = 350.737486 + m4 - (0.001436 - 0.0000019 * t) * t2;
    let mf = 11.250889 + m5 - (0.003211 + 0.0000003 * t) * t2;
    let na = 259.183275 - m6 + (0.002078 + 0.0000022 * t) * t2;
    let a = (51.2 + 20.2 * t).to_radians();
    let s1 = a.sin();
    let s2 = na.to_radians().sin();
    let b = 346.56 + (132.87 - 0.0091731 * t) * t;
    let s3 = 0.003964 * b.to_radians().sin();
    let c = (na + 275.05 - 2.3 * t).to_radians();
    let s4 = c.sin();
    let ml = ml + 0.000233 * s1 + s3 + 0.001964 * s2;
    let ms = ms - 0.001778 * s1;
    let md = md + 0.000817 * s1 + s3 + 0.002541 * s2;
    let mf = mf + s3 - 0.024691 * s2 - 0.004328 * s4;
    let me1 = me1 + 0.002011 * s1 + s3 + 0.001964 * s2;
    let e = 1.0 - (0.002495 + 0.00000752 * t) * t;
    let e2 = e * e;
    let _ml = (ml).to_radians();
    let ms = (ms).to_radians();
    let _na = (na).to_radians();
    let me1 = (me1).to_radians();
    let mf = (mf).to_radians();
    let md = (md).to_radians();

    let pm = 0.950724 + 0.051818 * (md).cos() + 0.009531 * (2.0 * me1 - md).cos();
    let pm = pm + 0.007843 * (2.0 * me1).cos() + 0.002824 * (2.0 * md).cos();
    let pm = pm + 0.000857 * (2.0 * me1 + md).cos() + e * 0.000533 * (2.0 * me1 - ms).cos();
    let pm = pm + e * 0.000401 * (2.0 * me1 - md - ms).cos();
    let pm = pm + e * 0.00032 * (md - ms).cos() - 0.000271 * (me1).cos();
    let pm = pm - e * 0.000264 * (ms + md).cos() - 0.000198 * (2.0 * mf - md).cos();
    let pm = pm + 0.000173 * (3.0 * md).cos() + 0.000167 * (4.0 * me1 - md).cos();
    let pm = pm - e * 0.000111 * (ms).cos() + 0.000103 * (4.0 * me1 - 2.0 * md).cos();
    let pm = pm - 0.000084 * (2.0 * md - 2.0 * me1).cos() - e * 0.000083 * (2.0 * me1 + ms).cos();
    let pm = pm + 0.000079 * (2.0 * me1 + 2.0 * md).cos() + 0.000072 * (4.0 * me1).cos();
    let pm = pm + e * 0.000064 * (2.0 * me1 - ms + md).cos()
        - e * 0.000063 * (2.0 * me1 + ms - md).cos();
    let pm = pm + e * 0.000041 * (ms + me1).cos() + e * 0.000035 * (2.0 * md - ms).cos();
    let pm = pm - 0.000033 * (3.0 * md - 2.0 * me1).cos() - 0.00003 * (md + me1).cos();
    let pm = pm - 0.000029 * (2.0 * (mf - me1)).cos() - e * 0.000029 * (2.0 * md + ms).cos();
    let pm =
        pm + e2 * 0.000026 * (2.0 * (me1 - ms)).cos() - 0.000023 * (2.0 * (mf - me1) + md).cos();
    let pm = pm + e * 0.000019 * (4.0 * me1 - ms - md).cos();

    return pm;
}

/// Convert angle in radians to equivalent angle in degrees.
///
/// Original macro name: Unwind
pub fn unwind(w: f64) -> f64 {
    return w - 6.283185308 * (w / 6.283185308).floor();
}

/// Mean ecliptic longitude of the Sun at the epoch
///
/// Original macro name: SunElong
pub fn sun_e_long(gd: f64, gm: u32, gy: u32) -> f64 {
    let t = (cd_jd(gd, gm, gy) - 2415020.0) / 36525.0;
    let t2 = t * t;
    let x = 279.6966778 + 36000.76892 * t + 0.0003025 * t2;

    return x - 360.0 * (x / 360.0).floor();
}

/// Longitude of the Sun at perigee
///
/// Original macro name: SunPeri
pub fn sun_peri(gd: f64, gm: u32, gy: u32) -> f64 {
    let t = (cd_jd(gd, gm, gy) - 2415020.0) / 36525.0;
    let t2 = t * t;
    let x = 281.2208444 + 1.719175 * t + 0.000452778 * t2;

    return x - 360.0 * (x / 360.0).floor();
}

/// Eccentricity of the Sun-Earth orbit
///
/// Original macro name: SunEcc
pub fn sun_ecc(gd: f64, gm: u32, gy: u32) -> f64 {
    let t = (cd_jd(gd, gm, gy) - 2415020.0) / 36525.0;
    let t2 = t * t;

    return 0.01675104 - 0.0000418 * t - 0.000000126 * t2;
}

/// Ecliptic - Declination (degrees)
///
/// Original macro name: ECDec
pub fn ec_dec(
    eld: f64,
    elm: f64,
    els: f64,
    bd: f64,
    bm: f64,
    bs: f64,
    gd: f64,
    gm: u32,
    gy: u32,
) -> f64 {
    let a = (dms_dd(eld, elm, els)).to_radians();
    let b = (dms_dd(bd, bm, bs)).to_radians();
    let c = (obliq(gd, gm, gy)).to_radians();
    let d = b.sin() * c.cos() + b.cos() * c.sin() * a.sin();
    return degrees(d.asin());
}

/// Ecliptic - Right Ascension (degrees)
///
/// Original macro name: ECRA
pub fn ec_ra(
    eld: f64,
    elm: f64,
    els: f64,
    bd: f64,
    bm: f64,
    bs: f64,
    gd: f64,
    gm: u32,
    gy: u32,
) -> f64 {
    let a = (dms_dd(eld, elm, els)).to_radians();
    let b = (dms_dd(bd, bm, bs)).to_radians();
    let c = (obliq(gd, gm, gy)).to_radians();
    let d = a.sin() * c.cos() - b.tan() * c.sin();
    let e = a.cos();
    let f = degrees(d.atan2(e));

    return f - 360.0 * (f / 360.0).floor();
}

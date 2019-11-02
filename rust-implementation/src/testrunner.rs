use crate::tests::coordinates as CST;
use crate::tests::datetime as DTT;

/// Run all functional tests.
pub fn run_tests() {
    run_datetime_tests();

    run_coordinate_tests();
}

pub fn run_datetime_tests() {
    DTT::test_easter(4, 20, 2003);
    DTT::test_day_numbers();

    let mut test_civil_time = DTT::TestCivilTimeScaffold {
        civil_hours: 18.0,
        civil_minutes: 31.0,
        civil_seconds: 27.0,
    };
    test_civil_time.test_civil_time_to_decimal_hours();
    test_civil_time.test_decimal_hours_to_civil_time();
    test_civil_time.test_decimal_time_parts();

    let mut test_local_civil_time = DTT::TestLocalCivilTimeScaffold {
        lct_hours: 3.0,
        lct_minutes: 37.0,
        lct_seconds: 0.0,
        is_daylight_savings: true,
        zone_correction: 4,
        local_day: 1.0,
        local_month: 7,
        local_year: 2013,
    };
    test_local_civil_time.test_local_civil_time_to_universal_time();
    test_local_civil_time.test_universal_time_to_local_civil_time();

    let mut test_universal_time_sidereal_time = DTT::TestUniversalTimeSiderealTimeScaffold {
        ut_hours: 14.0,
        ut_minutes: 36.0,
        ut_seconds: 51.67,
        gw_day: 22.0,
        gw_month: 4,
        gw_year: 1980,
    };
    test_universal_time_sidereal_time.test_universal_time_to_greenwich_sidereal_time();
    test_universal_time_sidereal_time.test_greenwich_sidereal_time_to_universal_time();

    let mut test_greenwich_sidereal_local_sidereal =
        DTT::TestGreenwichSiderealLocalSiderealScaffold {
            gst_hours: 4.0,
            gst_minutes: 40.0,
            gst_seconds: 5.23,
            geographical_longitude: -64.0,
        };
    test_greenwich_sidereal_local_sidereal.test_greenwich_sidereal_time_to_local_sidereal_time();
    test_greenwich_sidereal_local_sidereal.test_local_sidereal_time_to_greenwich_sidereal_time();

    DTT::test_julian_date_to_day_of_week();
}

pub fn run_coordinate_tests() {
    let mut test_angle_decimal_degrees = CST::TestAngleDecimalDegreesScaffold {
        degrees: 182.0,
        minutes: 31.0,
        seconds: 27.0,
    };
    test_angle_decimal_degrees.test_angle_to_decimal_degrees();
    test_angle_decimal_degrees.test_decimal_degrees_to_angle();

    let mut test_right_ascension_hour_angle = CST::TestRightAscensionHourAngleScaffold {
        ra_hours: 18.0,
        ra_minutes: 32.0,
        ra_seconds: 21.0,
        lct_hours: 14.0,
        lct_minutes: 36.0,
        lct_seconds: 51.67,
        is_daylight_saving: false,
        zone_correction: -4,
        local_day: 22.0,
        local_month: 4,
        local_year: 1980,
        geographical_longitude: -64.0,
    };
    test_right_ascension_hour_angle.test_right_ascension_to_hour_angle();
    test_right_ascension_hour_angle.test_hour_angle_to_right_ascension();

    let mut test_equatorial_horizon = CST::TestEquatorialHorizonScaffold {
        hour_angle_hours: 5.0,
        hour_angle_minutes: 51.0,
        hour_angle_seconds: 44.0,
        declination_degrees: 23.0,
        declination_minutes: 13.0,
        declination_seconds: 10.0,
        geographical_latitude: 52.0,
    };
    test_equatorial_horizon.test_equatorial_coordinates_to_horizon_coordinates();
    test_equatorial_horizon.test_horizon_coordinates_to_equatorial_coordinates();

    let mut test_ecliptic = CST::TestEclipticScaffold {
        ecliptic_longitude_degrees: 139.0,
        ecliptic_longitude_minutes: 41.0,
        ecliptic_longitude_seconds: 10.0,
        ecliptic_latitude_degrees: 4.0,
        ecliptic_latitude_minutes: 52.0,
        ecliptic_latitude_seconds: 31.0,
        greenwich_day: 6.0,
        greenwich_month: 7,
        greenwich_year: 2009,
    };
    test_ecliptic.test_mean_obliquity_of_the_ecliptic();
    test_ecliptic.test_ecliptic_coordinate_to_equatorial_coordinate();
    test_ecliptic.test_equatorial_coordinate_to_ecliptic_coordinate();

    let mut test_galactic = CST::TestGalacticScaffold {
        ra_hours: 10.0,
        ra_minutes: 21.0,
        ra_seconds: 0.0,
        dec_degrees: 10.0,
        dec_minutes: 3.0,
        dec_seconds: 11.0,
    };
    test_galactic.test_equatorial_coordinate_to_galactic_coordinate();
    test_galactic.test_galactic_coordinate_to_equatorial_coordinate();
}

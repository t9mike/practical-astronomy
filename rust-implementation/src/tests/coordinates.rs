use crate::lib::coordinates as CS;
use crate::lib::util;

pub struct TestAngleDecimalDegreesScaffold {
    pub degrees: f64,
    pub minutes: f64,
    pub seconds: f64,
}

impl TestAngleDecimalDegreesScaffold {
    pub fn test_angle_to_decimal_degrees(&mut self) {
        let decimal_degrees = util::round_f64(
            CS::angle_to_decimal_degrees(self.degrees, self.minutes, self.seconds),
            6,
        );

        println!(
            "Angle to decimal degrees: [DMS] {}d {}m {}s = [Decimal Degrees] {}",
            self.degrees, self.minutes, self.seconds, decimal_degrees
        );

        assert_eq!(decimal_degrees, 182.524167, "Decimal Degrees");
    }

    pub fn test_decimal_degrees_to_angle(&mut self) {
        let decimal_degrees = util::round_f64(
            CS::angle_to_decimal_degrees(self.degrees, self.minutes, self.seconds),
            6,
        );

        let (degrees, minutes, seconds) = CS::decimal_degrees_to_angle(decimal_degrees);

        println!(
            "Decimal degrees to angle: [Decimal Degrees] {} = [DMS] {}d {}m {}s",
            decimal_degrees, degrees, minutes, seconds
        );

        assert_eq!(degrees, 182.0, "Degrees");
        assert_eq!(minutes, 31.0, "Minutes");
        assert_eq!(seconds, 27.0, "Seconds");
    }
}

pub struct TestRightAscensionHourAngleScaffold {
    pub ra_hours: f64,
    pub ra_minutes: f64,
    pub ra_seconds: f64,
    pub lct_hours: f64,
    pub lct_minutes: f64,
    pub lct_seconds: f64,
    pub is_daylight_saving: bool,
    pub zone_correction: i32,
    pub local_day: f64,
    pub local_month: u32,
    pub local_year: u32,
    pub geographical_longitude: f64,
}

impl TestRightAscensionHourAngleScaffold {
    pub fn test_right_ascension_to_hour_angle(&mut self) {
        let (hour_angle_hours, hour_angle_minutes, hour_angle_seconds) =
            CS::right_ascension_to_hour_angle(
                self.ra_hours,
                self.ra_minutes,
                self.ra_seconds,
                self.lct_hours,
                self.lct_minutes,
                self.lct_seconds,
                self.is_daylight_saving,
                self.zone_correction,
                self.local_day,
                self.local_month,
                self.local_year,
                self.geographical_longitude,
            );

        println!(
			"Right ascension to hour angle: [RA] {}:{}:{} [LCT] {}:{}:{} [DST?] {} [ZC] {} [Local Date] {}/{}/{} [Geog Long] {} = [HA] {}:{}:{}",
			self.ra_hours,
			self.ra_minutes,
			self.ra_seconds,
			self.lct_hours,
			self.lct_minutes,
			self.lct_seconds,
			self.is_daylight_saving,
			self.zone_correction,
			self.local_month,
			self.local_day,
			self.local_year,
			self.geographical_longitude,
			hour_angle_hours,
			hour_angle_minutes,
			hour_angle_seconds
		);

        assert_eq!(hour_angle_hours, 9.0, "Hour Angle Hours");
        assert_eq!(hour_angle_minutes, 52.0, "Hour Angle Minutes");
        assert_eq!(hour_angle_seconds, 23.66, "Hour Angle Seconds");
    }
    pub fn test_hour_angle_to_right_ascension(&mut self) {
        let (hour_angle_hours, hour_angle_minutes, hour_angle_seconds) =
            CS::right_ascension_to_hour_angle(
                self.ra_hours,
                self.ra_minutes,
                self.ra_seconds,
                self.lct_hours,
                self.lct_minutes,
                self.lct_seconds,
                self.is_daylight_saving,
                self.zone_correction,
                self.local_day,
                self.local_month,
                self.local_year,
                self.geographical_longitude,
            );

        let (right_ascension_hours, right_ascension_minutes, right_ascension_seconds) =
            CS::hour_angle_to_right_ascension(
                hour_angle_hours,
                hour_angle_minutes,
                hour_angle_seconds,
                self.lct_hours,
                self.lct_minutes,
                self.lct_seconds,
                self.is_daylight_saving,
                self.zone_correction,
                self.local_day,
                self.local_month,
                self.local_year,
                self.geographical_longitude,
            );

        println!(
			"Hour angle to right ascension: [HA] {}:{}:{} [LCT] {}:{}:{} [DST?] {} [ZC] {} [Local Date] {}/{}/{} [Geog Long] {} = [RA] {}:{}:{}",
			hour_angle_hours,
			hour_angle_minutes,
			hour_angle_seconds,
			self.lct_hours,
			self.lct_minutes,
			self.lct_seconds,
			self.is_daylight_saving,
			self.zone_correction,
			self.local_month,
			self.local_day,
			self.local_year,
			self.geographical_longitude,
			right_ascension_hours,
			right_ascension_minutes,
			right_ascension_seconds,
		);

        assert_eq!(right_ascension_hours, 18.0, "Right Ascension Hours");
        assert_eq!(right_ascension_minutes, 32.0, "Right Ascension Minutes");
        assert_eq!(right_ascension_seconds, 21.0, "Right Ascension Seconds");
    }
}

pub struct TestEquatorialHorizonScaffold {
    pub hour_angle_hours: f64,
    pub hour_angle_minutes: f64,
    pub hour_angle_seconds: f64,
    pub declination_degrees: f64,
    pub declination_minutes: f64,
    pub declination_seconds: f64,
    pub geographical_latitude: f64,
}

impl TestEquatorialHorizonScaffold {
    pub fn test_equatorial_coordinates_to_horizon_coordinates(&mut self) {
        let (
            azimuth_degrees,
            azimuth_minutes,
            azimuth_seconds,
            altitude_degrees,
            altitude_minutes,
            altitude_seconds,
        ) = CS::equatorial_coordinates_to_horizon_coordinates(
            self.hour_angle_hours,
            self.hour_angle_minutes,
            self.hour_angle_seconds,
            self.declination_degrees,
            self.declination_minutes,
            self.declination_seconds,
            self.geographical_latitude,
        );

        println!(
			"Equatorial coordinates to horizon coordinates: [HA] {}:{}:{} [Declination] {}d {}m {}s [Geographical Latitude] {} = [Azimuth] {}d {}m {}s [Altitude] {}d {}m {}s",
			self.hour_angle_hours,
			self.hour_angle_minutes,
			self.hour_angle_seconds,
			self.declination_degrees,
			self.declination_minutes,
			self.declination_seconds,
			self.geographical_latitude,
			azimuth_degrees,
			azimuth_minutes,
			azimuth_seconds,
			altitude_degrees,
			altitude_minutes,
			altitude_seconds
		);

        assert_eq!(azimuth_degrees, 283.0, "Azimuth Degrees");
        assert_eq!(azimuth_minutes, 16.0, "Azimuth Minutes");
        assert_eq!(azimuth_seconds, 15.7, "Azimuth Seconds");
        assert_eq!(altitude_degrees, 19.0, "Altitude Degrees");
        assert_eq!(altitude_minutes, 20.0, "Altitude Minutes");
        assert_eq!(altitude_seconds, 3.64, "Altitude Seconds");
    }

    pub fn test_horizon_coordinates_to_equatorial_coordinates(&mut self) {
        let (
            azimuth_degrees,
            azimuth_minutes,
            azimuth_seconds,
            altitude_degrees,
            altitude_minutes,
            altitude_seconds,
        ) = CS::equatorial_coordinates_to_horizon_coordinates(
            self.hour_angle_hours,
            self.hour_angle_minutes,
            self.hour_angle_seconds,
            self.declination_degrees,
            self.declination_minutes,
            self.declination_seconds,
            self.geographical_latitude,
        );

        let (
            hour_angle_hours,
            hour_angle_minutes,
            hour_angle_seconds,
            declination_degrees,
            declination_minutes,
            declination_seconds,
        ) = CS::horizon_coordinates_to_equatorial_coordinates(
            azimuth_degrees,
            azimuth_minutes,
            azimuth_seconds,
            altitude_degrees,
            altitude_minutes,
            altitude_seconds,
            self.geographical_latitude,
        );

        println!(
			"Horizon coordinates to equatorial coordinates: [Azimuth] {}d {}m {}s [Altitude] {}d {}m {}s [Geographical Latitude] {} = [HA] {}:{}:{} [Declination] {}d {}m {}s",
			azimuth_degrees,
			azimuth_minutes,
			azimuth_seconds,
			altitude_degrees,
			altitude_minutes,
			altitude_seconds,
			self.geographical_latitude,
			hour_angle_hours,
			hour_angle_minutes,
			hour_angle_seconds,
			declination_degrees,
			declination_minutes,
			declination_seconds,
		);

        assert_eq!(hour_angle_hours, 5.0, "Hour Angle Hours");
        assert_eq!(hour_angle_minutes, 51.0, "Hour Angle Minutes");
        assert_eq!(hour_angle_seconds, 44.0, "Hour Angle Seconds");
        assert_eq!(declination_degrees, 23.0, "Declination Degrees");
        assert_eq!(declination_minutes, 13.0, "Declination Minutes");
        assert_eq!(declination_seconds, 10.0, "Declination Seconds");
    }
}

pub struct TestEclipticScaffold {
    pub ecliptic_longitude_degrees: f64,
    pub ecliptic_longitude_minutes: f64,
    pub ecliptic_longitude_seconds: f64,
    pub ecliptic_latitude_degrees: f64,
    pub ecliptic_latitude_minutes: f64,
    pub ecliptic_latitude_seconds: f64,
    pub greenwich_day: f64,
    pub greenwich_month: u32,
    pub greenwich_year: u32,
}
impl TestEclipticScaffold {
    pub fn test_mean_obliquity_of_the_ecliptic(&mut self) {
        let obliquity = util::round_f64(
            CS::mean_obliquity_of_the_ecliptic(
                self.greenwich_day,
                self.greenwich_month,
                self.greenwich_year,
            ),
            8,
        );

        println!(
            "Mean obliquity of the ecliptic: [Greenwich Date] {}/{}/{} = [Obliquity] {}",
            self.greenwich_month, self.greenwich_day, self.greenwich_year, obliquity
        );

        assert_eq!(obliquity, 23.43805531, "Obliquity");
    }

    pub fn test_ecliptic_coordinate_to_equatorial_coordinate(&mut self) {
        let (ra_hours, ra_minutes, ra_seconds, dec_degrees, dec_minutes, dec_seconds) =
            CS::ecliptic_coordinate_to_equatorial_coordinate(
                self.ecliptic_longitude_degrees,
                self.ecliptic_longitude_minutes,
                self.ecliptic_longitude_seconds,
                self.ecliptic_latitude_degrees,
                self.ecliptic_latitude_minutes,
                self.ecliptic_latitude_seconds,
                self.greenwich_day,
                self.greenwich_month,
                self.greenwich_year,
            );

        println!(
			"Ecliptic coordinates to equatorial coordinates: [Ecliptic] [Longitude] {}d {}m {}s [Latitude] {}d {}m {}s [Greenwich Date] {}/{}/{} = [Right Ascension] {}:{}:{} [Declination] {}d {}m {}s",
			self.ecliptic_longitude_degrees,
			self.ecliptic_longitude_minutes,
			self.ecliptic_longitude_seconds,
			self.ecliptic_latitude_degrees,
			self.ecliptic_latitude_minutes,
			self.ecliptic_latitude_seconds,
			self.greenwich_month,
			self.greenwich_day,
			self.greenwich_year,
			ra_hours,
			ra_minutes,
			ra_seconds,
			dec_degrees,
			dec_minutes,
			dec_seconds
		);

        assert_eq!(ra_hours, 9.0, "RA Hours");
        assert_eq!(ra_minutes, 34.0, "RA Minutes");
        assert_eq!(ra_seconds, 53.4, "RA Seconds");
        assert_eq!(dec_degrees, 19.0, "Dec Degrees");
        assert_eq!(dec_minutes, 32.0, "Dec Minutes");
        assert_eq!(dec_seconds, 8.52, "Dec Seconds");
    }

    pub fn test_equatorial_coordinate_to_ecliptic_coordinate(&mut self) {
        let (ra_hours, ra_minutes, ra_seconds, dec_degrees, dec_minutes, dec_seconds) =
            CS::ecliptic_coordinate_to_equatorial_coordinate(
                self.ecliptic_longitude_degrees,
                self.ecliptic_longitude_minutes,
                self.ecliptic_longitude_seconds,
                self.ecliptic_latitude_degrees,
                self.ecliptic_latitude_minutes,
                self.ecliptic_latitude_seconds,
                self.greenwich_day,
                self.greenwich_month,
                self.greenwich_year,
            );

        let (ecl_long_deg, ecl_long_min, ecl_long_sec, ecl_lat_deg, ecl_lat_min, ecl_lat_sec) =
            CS::equatorial_coordinate_to_ecliptic_coordinate(
                ra_hours,
                ra_minutes,
                ra_seconds,
                dec_degrees,
                dec_minutes,
                dec_seconds,
                self.greenwich_day,
                self.greenwich_month,
                self.greenwich_year,
            );

        println!(
			"Equatorial coordinates to ecliptic coordinates: [Right Ascension] {}:{}:{} [Declination] {}d {}m {}s [Greenwich Date] {}/{}/{} = [Ecliptic] [Longitude] {}d {}m {}s [Latitude] {}d {}m {}s",
			ra_hours,
			ra_minutes,
			ra_seconds,
			dec_degrees,
			dec_minutes,
			dec_seconds,
			self.greenwich_month,
			self.greenwich_day,
			self.greenwich_year,
			ecl_long_deg,
			ecl_long_min,
			ecl_long_sec,
			ecl_lat_deg,
			ecl_lat_min,
			ecl_lat_sec
		);

        assert_eq!(ecl_long_deg, 139.0, "Ecliptic Longitude Degrees");
        assert_eq!(ecl_long_min, 41.0, "Ecliptic Longitude Minutes");
        assert_eq!(ecl_long_sec, 9.97, "Ecliptic Longitude Seconds");
        assert_eq!(ecl_lat_deg, 4.0, "Ecliptic Latitude Degrees");
        assert_eq!(ecl_lat_min, 52.0, "Ecliptic Latitude Minutes");
        assert_eq!(ecl_lat_sec, 30.99, "Ecliptic Latitude Seconds");
    }
}

pub struct TestGalacticScaffold {
    pub ra_hours: f64,
    pub ra_minutes: f64,
    pub ra_seconds: f64,
    pub dec_degrees: f64,
    pub dec_minutes: f64,
    pub dec_seconds: f64,
}
impl TestGalacticScaffold {
    pub fn test_equatorial_coordinate_to_galactic_coordinate(&mut self) {
        let (gal_long_deg, gal_long_min, gal_long_sec, gal_lat_deg, gal_lat_min, gal_lat_sec) =
            CS::equatorial_coordinate_to_galactic_coordinate(
                self.ra_hours,
                self.ra_minutes,
                self.ra_seconds,
                self.dec_degrees,
                self.dec_minutes,
                self.dec_seconds,
            );

        println!(
			"Equatorial coordinate to galactic coordinate: [RA] {}:{}:{} [Dec] {}d {}m {}s = [Galactic] [Long] {}d {}m {}s [Lat] {}d {}m {}s",
			self.ra_hours,
			self.ra_minutes,
			self.ra_seconds,
			self.dec_degrees,
			self.dec_minutes,
			self.dec_seconds,
			gal_long_deg,
			gal_long_min,
			gal_long_sec,
			gal_lat_deg,
			gal_lat_min,
			gal_lat_sec
		);

        assert_eq!(gal_long_deg, 232.0, "Galactic Longitude Degrees");
        assert_eq!(gal_long_min, 14.0, "Galactic Longitude Minutes");
        assert_eq!(gal_long_sec, 52.38, "Galactic Longitude Seconds");
        assert_eq!(gal_lat_deg, 51.0, "Galactic Latitude Degrees");
        assert_eq!(gal_lat_min, 7.0, "Galactic Latitude Minutes");
        assert_eq!(gal_lat_sec, 20.16, "Galactic Latitude Seconds");
    }

    pub fn test_galactic_coordinate_to_equatorial_coordinate(&mut self) {
        let (gal_long_deg, gal_long_min, gal_long_sec, gal_lat_deg, gal_lat_min, gal_lat_sec) =
            CS::equatorial_coordinate_to_galactic_coordinate(
                self.ra_hours,
                self.ra_minutes,
                self.ra_seconds,
                self.dec_degrees,
                self.dec_minutes,
                self.dec_seconds,
            );

        let (ra_hours, ra_minutes, ra_seconds, dec_degrees, dec_minutes, dec_seconds) =
            CS::galactic_coordinate_to_equatorial_coordinate(
                gal_long_deg,
                gal_long_min,
                gal_long_sec,
                gal_lat_deg,
                gal_lat_min,
                gal_lat_sec,
            );

        println!(
			"Galactic coordinate to equatorial coordinate: [Galactic] [Long] {}d {}m {}s [Lat] {}d {}m {}s = [RA] {}:{}:{} [Dec] {}d {}m {}s", 
			gal_long_deg,
			gal_long_min,
			gal_long_sec,
			gal_lat_deg,
			gal_lat_min,
			gal_lat_sec,
			ra_hours,
			ra_minutes,
			ra_seconds,
			dec_degrees,
			dec_minutes,
			dec_seconds,
		);

        assert_eq!(ra_hours, 10.0, "Right Ascension Hours");
        assert_eq!(ra_minutes, 21.0, "Right Ascension Minutes");
        assert_eq!(ra_seconds, 0.0, "Right Ascension Seconds");
        assert_eq!(dec_degrees, 10.0, "Declination Degrees");
        assert_eq!(dec_minutes, 3.0, "Declination Degrees");
        assert_eq!(dec_seconds, 11.0, "Declination Seconds");
    }
}

pub fn test_angle_between_two_objects(
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
) {
    let (angle_deg, angle_min, angle_sec) = CS::angle_between_two_objects(
        ra_long_1_hour_deg,
        ra_long_1_min,
        ra_long_1_sec,
        dec_lat_1_deg,
        dec_lat_1_min,
        dec_lat_1_sec,
        ra_long_2_hour_deg,
        ra_long_2_min,
        ra_long_2_sec,
        dec_lat_2_deg,
        dec_lat_2_min,
        dec_lat_2_sec,
        hour_or_degree.to_string(),
    );

    println!(
		"Angle between two objects: [Object 1] [RA Long] {}{} {}m {}s [Dec Lat] {}d {}m {}s [Object 2] [RA Long] {}{} {}m {}s [Dec Lat] {}d {}m {}s [Hour or Degree?] {} = [Angle] {}d {}m {}s",
		ra_long_1_hour_deg,
		if hour_or_degree == "H" {"h"} else {"d"},
		ra_long_1_min,
		ra_long_1_sec,
		dec_lat_1_deg,
		dec_lat_1_min,
		dec_lat_1_sec,
		ra_long_2_hour_deg,
		if hour_or_degree == "H" {"h"} else {"d"},
		ra_long_2_min,
		ra_long_2_sec,
		dec_lat_2_deg,
		dec_lat_2_min,
		dec_lat_2_sec,
		hour_or_degree,
		angle_deg,
		angle_min,
		angle_sec
	);

    assert_eq!(angle_deg, 23.0, "Angle Degrees");
    assert_eq!(angle_min, 40.0, "Angle Minutes");
    assert_eq!(angle_sec, 25.86, "Angle Seconds");
}

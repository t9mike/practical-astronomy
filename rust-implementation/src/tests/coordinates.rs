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

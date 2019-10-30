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

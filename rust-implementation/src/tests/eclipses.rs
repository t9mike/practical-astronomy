use crate::lib::eclipses as ECL;

pub struct TestLunarEclipseScaffold {
    pub local_date_day: f64,
    pub local_date_month: u32,
    pub local_date_year: u32,
    pub is_daylight_saving: bool,
    pub zone_correction_hours: i32,
}
impl TestLunarEclipseScaffold {
    pub fn test_lunar_eclipse_occurrence(&mut self) {
        let (status, event_date_day, event_date_month, event_date_year) =
            ECL::lunar_eclipse_occurrence(
                self.local_date_day,
                self.local_date_month,
                self.local_date_year,
                self.is_daylight_saving,
                self.zone_correction_hours,
            );

        println!(
			"Lunar eclipse occurrence: [Local Date] {}/{}/{} [DST?] {} [Zone Correction] {} = [Status] {:} [Event Date] {}/{}/{}",
			self.local_date_month,
			self.local_date_day,
			self.local_date_year,
			self.is_daylight_saving,
			self.zone_correction_hours,
			status,
			event_date_month,
			event_date_day,
			event_date_year
		);

        assert_eq!(status, "Lunar eclipse certain", "Lunar eclipse status");
        assert_eq!(event_date_day, 4.0, "Lunar eclipse event date (day)");
        assert_eq!(event_date_month, 4, "Lunar eclipse event date (month)");
        assert_eq!(event_date_year, 2015, "Lunar eclipse event date (year)");
    }
    pub fn test_lunar_eclipse_circumstances(&mut self) {
        let (
            lunar_eclipse_certain_date_day,
            lunar_eclipse_certain_date_month,
            lunar_eclipse_certain_date_year,
            ut_start_pen_phase_hour,
            ut_start_pen_phase_minutes,
            ut_start_umbral_phase_hour,
            ut_start_umbral_phase_minutes,
            ut_start_total_phase_hour,
            ut_start_total_phase_minutes,
            ut_mid_eclipse_hour,
            ut_mid_eclipse_minutes,
            ut_end_total_phase_hour,
            ut_end_total_phase_minutes,
            ut_end_umbral_phase_hour,
            ut_end_umbral_phase_minutes,
            ut_end_pen_phase_hour,
            ut_end_pen_phase_minutes,
            eclipse_magnitude,
        ) = ECL::lunar_eclipse_circumstances(
            self.local_date_day,
            self.local_date_month,
            self.local_date_year,
            self.is_daylight_saving,
            self.zone_correction_hours,
        );

        println!(
			"Lunar eclipse circumstances: [Local Date] {}/{}/{} [DST?] {} [Zone Correction] {} hours = [Certain Date] {}/{}/{} [Start] [Penumbral] {}:{} [Umbral] {}:{} [Total] {}:{} [Mid] {}:{} [End] [Total] {}:{} [Umbral] {}:{} [Penumbral] {}:{} [Magnitude] {}",
			self.local_date_month,
			self.local_date_day,
			self.local_date_year,
			self.is_daylight_saving,
			self.zone_correction_hours,
			lunar_eclipse_certain_date_month,
			lunar_eclipse_certain_date_day,
			lunar_eclipse_certain_date_year,
			ut_start_pen_phase_hour,
			ut_start_pen_phase_minutes,
			ut_start_umbral_phase_hour,
			ut_start_umbral_phase_minutes,
			ut_start_total_phase_hour,
			ut_start_total_phase_minutes,
			ut_mid_eclipse_hour,
			ut_mid_eclipse_minutes,
			ut_end_total_phase_hour,
			ut_end_total_phase_minutes,
			ut_end_umbral_phase_hour,
			ut_end_umbral_phase_minutes,
			ut_end_pen_phase_hour,
			ut_end_pen_phase_minutes,
			eclipse_magnitude
		);

        assert_eq!(lunar_eclipse_certain_date_day, 4.0, "Eclipse Date (day)");
        assert_eq!(lunar_eclipse_certain_date_month, 4, "Eclipse Date (month)");
        assert_eq!(lunar_eclipse_certain_date_year, 2015, "Eclipse Date (year)");
        assert_eq!(ut_start_pen_phase_hour, 9.0, "Start Penumbral Phase (hour)");
        assert_eq!(
            ut_start_pen_phase_minutes, 0.0,
            "Start Penumbral Phase (minutes)"
        );
        assert_eq!(
            ut_start_umbral_phase_hour, 10.0,
            "Start Umbral Phase (hour)"
        );
        assert_eq!(
            ut_start_umbral_phase_minutes, 16.0,
            "Start Umbral Phase (minutes)"
        );
        assert_eq!(ut_start_total_phase_hour, 11.0, "Start Total Phase (hour)");
        assert_eq!(
            ut_start_total_phase_minutes, 55.0,
            "Start Total Phase (minutes)"
        );
        assert_eq!(ut_mid_eclipse_hour, 12.0, "Mid Eclipse (hour)");
        assert_eq!(ut_mid_eclipse_minutes, 1.0, "Mid Eclipse (minutes)");
        assert_eq!(ut_end_total_phase_hour, 12.0, "End Total Phase (hour)");
        assert_eq!(ut_end_total_phase_minutes, 7.0, "End Total Phase (minutes)");
        assert_eq!(ut_end_umbral_phase_hour, 13.0, "End Umbral Phase (hour)");
        assert_eq!(
            ut_end_umbral_phase_minutes, 46.0,
            "End Umbral Phase (minutes)"
        );
        assert_eq!(ut_end_pen_phase_hour, 15.0, "End Penumbral Phase (hour)");
        assert_eq!(
            ut_end_pen_phase_minutes, 1.0,
            "End Penumbral Phase (minutes)"
        );
        assert_eq!(eclipse_magnitude, 1.01, "Eclipse Magnitude");
    }
}

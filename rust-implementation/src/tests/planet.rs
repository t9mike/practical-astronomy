use crate::lib::planet as CP;

pub struct TestPositionOfPlanetScaffold {
    pub lct_hour: f64,
    pub lct_minute: f64,
    pub lct_second: f64,
    pub is_daylight_saving: bool,
    pub zone_correction_hours: i32,
    pub local_date_day: f64,
    pub local_date_month: u32,
    pub local_date_year: u32,
    pub planet_name: String,
}
impl TestPositionOfPlanetScaffold {
    pub fn test_approximate_position_of_planet(&mut self) {
        let (
            planet_ra_hour,
            planet_ra_min,
            planet_ra_sec,
            planet_dec_deg,
            planet_dec_min,
            planet_dec_sec,
        ) = CP::approximate_position_of_planet(
            self.lct_hour,
            self.lct_minute,
            self.lct_second,
            self.is_daylight_saving,
            self.zone_correction_hours,
            self.local_date_day,
            self.local_date_month,
            self.local_date_year,
            self.planet_name.to_string(),
        );

        println!(
			"Approximate position of planet: [Local Civil Time] {}:{}:{} [DST?] {} [Zone Correction] {} [Local Date] {}/{}/{} [Planet Name] {} = [RA] {}h {}m {}s [Dec] {}d {}m {}s",
			self.lct_hour,
			self.lct_minute,
			self.lct_second,
			self.is_daylight_saving,
			self.zone_correction_hours,
			self.local_date_month,
			self.local_date_day,
			self.local_date_year,
			self.planet_name,
			planet_ra_hour,
			planet_ra_min,
			planet_ra_sec,
			planet_dec_deg,
			planet_dec_min,
			planet_dec_sec
		);

        assert_eq!(planet_ra_hour, 11.0, "Planet Right Ascension (hour)");
        assert_eq!(planet_ra_min, 11.0, "Planet Right Ascension (minutes)");
        assert_eq!(planet_ra_sec, 13.8, "Planet Right Ascension (seconds)");
        assert_eq!(planet_dec_deg, 6.0, "Planet Declination (degrees)");
        assert_eq!(planet_dec_min, 21.0, "Planet Declination (minutes)");
        assert_eq!(planet_dec_sec, 25.1, "Planet Declination (seconds)");
    }

    pub fn test_precise_position_of_planet(&mut self) {
        let (
            planet_ra_hour,
            planet_ra_min,
            planet_ra_sec,
            planet_dec_deg,
            planet_dec_min,
            planet_dec_sec,
        ) = CP::precise_position_of_planet(
            self.lct_hour,
            self.lct_minute,
            self.lct_second,
            self.is_daylight_saving,
            self.zone_correction_hours,
            self.local_date_day,
            self.local_date_month,
            self.local_date_year,
            self.planet_name.to_string(),
        );

        println!(
			"Precise position of planet: [Local Civil Time] {}:{}:{} [DST?] {} [Zone Correction] {} [Local Date] {}/{}/{} [Planet Name] {} = [RA] {}h {}m {}s [Dec] {}d {}m {}s",
			self.lct_hour,
			self.lct_minute,
			self.lct_second,
			self.is_daylight_saving,
			self.zone_correction_hours,
			self.local_date_month,
			self.local_date_day,
			self.local_date_year,
			self.planet_name,
			planet_ra_hour,
			planet_ra_min,
			planet_ra_sec,
			planet_dec_deg,
			planet_dec_min,
			planet_dec_sec
		);

        assert_eq!(planet_ra_hour, 11.0, "Planet Right Ascension (hour)");
        assert_eq!(planet_ra_min, 10.0, "Planet Right Ascension (minutes)");
        assert_eq!(planet_ra_sec, 31.52, "Planet Right Ascension (seconds)");
        assert_eq!(planet_dec_deg, 6.0, "Planet Declination (degrees)");
        assert_eq!(planet_dec_min, 25.0, "Planet Declination (minutes)");
        assert_eq!(planet_dec_sec, 46.25, "Planet Declination (seconds)");
    }
}

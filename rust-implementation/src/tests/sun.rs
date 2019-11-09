use crate::lib::sun as CS;

pub fn test_approximate_position_of_sun(
    lct_hours: f64,
    lct_minutes: f64,
    lct_seconds: f64,
    local_day: f64,
    local_month: u32,
    local_year: u32,
    is_daylight_saving: bool,
    zone_correction: i32,
) {
    let (sun_ra_hour, sun_ra_min, sun_ra_sec, sun_dec_deg, sun_dec_min, sun_dec_sec) =
        CS::approximate_position_of_sun(
            lct_hours,
            lct_minutes,
            lct_seconds,
            local_day,
            local_month,
            local_year,
            is_daylight_saving,
            zone_correction,
        );

    println!(
		"Approximate position of sun: [Local Civil Time] {}:{}:{} [Local Date] {}/{}/{} [DST?] {} [Zone Correction] {} = [Sun] [RA] {}h {}m {}s [Dec] {}d {}m {}s",
		lct_hours,
		lct_minutes,
		lct_seconds,
		local_month,
		local_day,
		local_year,
		is_daylight_saving,
		zone_correction,
		sun_ra_hour,
		sun_ra_min,
		sun_ra_sec,
		sun_dec_deg,
		sun_dec_min,
		sun_dec_sec
	);

    assert_eq!(sun_ra_hour, 8.0, "Sun RA Hour");
    assert_eq!(sun_ra_min, 23.0, "Sun RA Minutes");
    assert_eq!(sun_ra_sec, 33.73, "Sun RA Seconds");
    assert_eq!(sun_dec_deg, 19.0, "Sun Dec Degrees");
    assert_eq!(sun_dec_min, 21.0, "Sun Dec Minutes");
    assert_eq!(sun_dec_sec, 14.32, "Sun Dec Seconds");
}

pub fn test_precise_position_of_sun(
    lct_hours: f64,
    lct_minutes: f64,
    lct_seconds: f64,
    local_day: f64,
    local_month: u32,
    local_year: u32,
    is_daylight_saving: bool,
    zone_correction: i32,
) {
    let (sun_ra_hour, sun_ra_min, sun_ra_sec, sun_dec_deg, sun_dec_min, sun_dec_sec) =
        CS::precise_position_of_sun(
            lct_hours,
            lct_minutes,
            lct_seconds,
            local_day,
            local_month,
            local_year,
            is_daylight_saving,
            zone_correction,
        );

    println!(
		"Precise position of sun: [Local Civil Time] {}:{}:{} [Local Date] {}/{}/{} [DST?] {} [Zone Correction] {} = [Sun] [RA] {}h {}m {}s [Dec] {}d {}m {}s",
		lct_hours,
		lct_minutes,
		lct_seconds,
		local_month,
		local_day,
		local_year,
		is_daylight_saving,
		zone_correction,
		sun_ra_hour,
		sun_ra_min,
		sun_ra_sec,
		sun_dec_deg,
		sun_dec_min,
		sun_dec_sec
	);

    assert_eq!(sun_ra_hour, 8.0, "Sun RA Hour");
    assert_eq!(sun_ra_min, 26.0, "Sun RA Minutes");
    assert_eq!(sun_ra_sec, 3.84, "Sun RA Seconds");
    assert_eq!(sun_dec_deg, 19.0, "Sun Dec Degrees");
    assert_eq!(sun_dec_min, 12.0, "Sun Dec Minutes");
    assert_eq!(sun_dec_sec, 49.68, "Sun Dec Seconds");
}

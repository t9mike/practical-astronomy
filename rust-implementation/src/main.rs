extern crate clap;
extern crate num;

use crate::lib::coordinates as CS;
use crate::lib::macros as MA;

mod lib;
mod testrunner;
mod tests;

use clap::{App, Arg};

fn main() {
    let matches = App::new("Practical Astronomy")
        .arg(
            Arg::with_name("tests")
                .short("t")
                .long("tests")
                .takes_value(false)
                .help("Run unit tests"),
        )
        .arg(
            Arg::with_name("canned")
                .short("c")
                .long("canned")
                .takes_value(false)
                .help("Run canned observer test"),
        )
        .get_matches();

    if matches.is_present("canned") {
        // Local coordinates (Dayton, OH)
        let latitude_input = 39.78;
        let longitude_input = -84.2;

        // Local date/time
        let lct_hours = 20.0;
        let lct_minutes = 0.0;
        let lct_seconds = 0.0;
        let is_daylight_saving = false;
        let zone_correction = -5;
        let local_day = 17.0;
        let local_month = 12;
        let local_year = 2019;

        // Sirius
        let right_ascension_input = 6.75257;
        let declination_input = -16.7131;

        let (hour_angle_hours, hour_angle_minutes, hour_angle_seconds) =
            CS::right_ascension_to_hour_angle(
                MA::dh_hour(right_ascension_input) as f64,
                MA::dh_min(right_ascension_input) as f64,
                MA::dh_sec(right_ascension_input),
                lct_hours,
                lct_minutes,
                lct_seconds,
                is_daylight_saving,
                zone_correction,
                local_day,
                local_month,
                local_year,
                longitude_input,
            );

        let (
            azimuth_degrees,
            azimuth_minutes,
            azimuth_seconds,
            altitude_degrees,
            altitude_minutes,
            altitude_seconds,
        ) = CS::equatorial_coordinates_to_horizon_coordinates(
            hour_angle_hours,
            hour_angle_minutes,
            hour_angle_seconds,
            MA::dd_deg(declination_input),
            MA::dd_min(declination_input),
            MA::dd_sec(declination_input),
            latitude_input,
        );

        println!(
            "Observing results: [Azimuth] {} degrees {} minutes {} seconds [Altitude] {} degrees {} minutes {} seconds",
            azimuth_degrees,
            azimuth_minutes,
            azimuth_seconds,
            altitude_degrees,
            altitude_minutes,
            altitude_seconds
        )
    }
    if matches.is_present("tests") {
        testrunner::run_tests();
    }
}

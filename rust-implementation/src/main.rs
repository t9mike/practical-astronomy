extern crate clap;
extern crate num;
extern crate sqlite;

use crate::lib::coordinates as CS;
use crate::lib::macros as MA;

mod lib;
mod testrunner;
mod tests;

use clap::{App, Arg};

macro_rules! ok(($result:expr) => ($result.unwrap()));

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

        let connection = ok!(sqlite::open("hygdata.db"));

        let statement = "select ProperName,Magnitude,RightAscension,Declination from hygdata where Magnitude <= 6 and ProperName != 'Sol' order by Magnitude";

        ok!(connection.iterate(statement, |pairs| {

        let right_ascension_input = pairs[2].1.unwrap().to_string().parse::<f64>().unwrap();
        let declination_input = pairs[3].1.unwrap().to_string().parse::<f64>().unwrap();

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
            "Observing results for {}: [Right Ascension/Declination] {}/{} = [Azimuth] {} degrees {} minutes {} seconds [Altitude] {} degrees {} minutes {} seconds",
            pairs[0].1.unwrap(),
            right_ascension_input,
            declination_input,
            azimuth_degrees,
            azimuth_minutes,
            azimuth_seconds,
            altitude_degrees,
            altitude_minutes,
            altitude_seconds
        );

        return true;
        }));
    }
    if matches.is_present("tests") {
        testrunner::run_tests();
    }
}

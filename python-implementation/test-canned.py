#!/usr/bin/python3

import sqlite3

import lib.pa_coordinate as PC
import lib.pa_macro as MA

# Local coordinates (Dayton, OH)
latitude_input = 39.78
longitude_input = -84.2

# Local date/time
lct_hours = 20.0
lct_minutes = 0.0
lct_seconds = 0.0
is_daylight_saving = False
zone_correction = -5
local_day = 17.0
local_month = 12
local_year = 2019

conn = sqlite3.connect('hygdata.db')

cursor = conn.execute('select ProperName,RightAscension,Declination from hygdata where Magnitude <= 6 order by Magnitude')
for row in cursor:
	right_ascension_input = row[1]
	declination_input = row[2]

	hour_angle_hours, hour_angle_minutes, hour_angle_seconds = PC.right_ascension_to_hour_angle(MA.dh_hour(right_ascension_input), MA.dh_min(right_ascension_input), MA.dh_sec(right_ascension_input), lct_hours, lct_minutes, lct_seconds, is_daylight_saving, zone_correction, local_day, local_month, local_year, longitude_input)

	azimuth_degrees, azimuth_minutes, azimuth_seconds, altitude_degrees, altitude_minutes, altitude_seconds = PC.equatorial_coordinates_to_horizon_coordinates(hour_angle_hours, hour_angle_minutes, hour_angle_seconds, MA.dd_deg(declination_input), MA.dd_min(declination_input), MA.dd_sec(declination_input), latitude_input)

	print(f"{row[0]}: [Right Ascension/Declination] {right_ascension_input}/{declination_input} = [Altitude] {altitude_degrees} degrees {altitude_minutes} minutes {altitude_seconds} [Azimuth] {azimuth_degrees} degrees {azimuth_minutes} minutes {azimuth_seconds} seconds")
	# print(f"{row[0]}: [Right Ascension/Declination] {row[1]}/{row[2]}")

conn.close()

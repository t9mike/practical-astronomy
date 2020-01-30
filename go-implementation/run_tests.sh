#!/usr/bin/bash

function usage_instructions {
	echo ""
	echo "Usage:  run_tests <test_name>"
	echo ""
	echo "Valid test names:"
	echo "  all         -- Run all tests"
	echo "  coordinate  -- Run coordinate tests"
	echo "  datetime    -- Run datetime tests"
	echo "  sun         -- Run sun tests"
	echo "  macro       -- Run macro tests"
	echo "  util        -- Run util tests"
}

function exec_test {
	case "$1" in
		"all")
			run_coordinate_tests
			run_datetime_tests
			run_sun_tests
			run_macro_tests
			run_util_tests
			;;
		"coordinate")
			run_coordinate_tests
			;;
		"datetime")
			run_datetime_tests
			;;
		"sun")
			run_sun_tests
			;;
		"macro")
			run_macro_tests
			;;
		"util")
			run_util_tests
			;;
		*)
			echo "Not a valid test name"
			usage_instructions				
			exit 1
			;;
	esac

}

function run_test_return {
	go test
	cd ../..
}

function run_coordinate_tests {
	cd lib/coordinates
	run_test_return
}

function run_datetime_tests {
	cd lib/datetime
	run_test_return
}

function run_sun_tests {
	cd lib/sun
	run_test_return
}

function run_macro_tests {
	cd lib/macros
	run_test_return
}

function run_util_tests {
	cd lib/util
	run_test_return
}

if [ -z "$1" ]
then
	echo "Not specified: name of test"
	usage_instructions	
	exit 1
fi

exec_test $1

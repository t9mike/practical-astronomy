default:
	@echo 'Targets:'
	@echo '  format  --  Beautify Go source files.'
	@echo '  tests   --  Run unit tests.'

format:
	@gofmt -s -w .

tests:
	@./run_macro_tests.sh
	@./run_datetime_tests.sh
	@./run_coordinate_tests.sh
	@./run_sun_tests.sh
	@./run_util_tests.sh

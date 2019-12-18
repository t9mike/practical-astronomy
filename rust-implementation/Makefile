default:
	@echo 'Targets:'
	@echo '  canned   -- Run with --canned argument.'
	@echo '  hygdata  -- Retrieve the hygdata database.'
	@echo '  tests    -- Run tests.'
	@echo '  tests-so -- Run tests, but suppress program output. (Useful for reviewing build warnings.)'

canned:
	@cargo run -- -c

hygdata:
	@wget --no-check-certificate https://github.com/jfcarr-astronomy/HYG-Database-SQLite/raw/master/hygdata.db

tests:
	@cargo run -- -t

tests-so:
	@cargo run -- -t > /dev/null

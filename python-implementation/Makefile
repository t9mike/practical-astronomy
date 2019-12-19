default:
	@echo 'Targets:'
	@echo '  all-tests                  -- Run all unit tests.'
	@echo '  test-easter                -- Run Easter unit tests.'
	@echo '  test-day-number            -- Run day number unit tests.'
	@echo '  test-julian                -- Run Julian date unit tests.'
	@echo '  test-time                  -- Run time conversion unit tests.'
	@echo '  test-coordinate            -- Run coordinate unit tests.'
	@echo '  test-sun                   -  Run Sun unit tests.'
	@echo '  test-planet-comet-binary   -- Run planet, comet, and binary star unit tests.'
	@echo '  test-moon                  -- Run Moon unit tests.'
	@echo '  test-eclipses              -- Run eclipse unit tests.'
	@echo '  hygdata                    -- Retrieve the hygdata database.'
	@echo '  canned                     -- Run multi-star horizon coordinate calculations.'

all-tests: test-easter test-day-number test-julian test-time test-coordinate test-sun test-planet-comet-binary test-moon test-eclipses

test-easter:
	@./test-date-of-easter.py -v

test-day-number:
	@./test-day-number.py

test-julian:
	@./test-julian.py

test-time:
	@./test-time.py

test-coordinate:
	@./test-coordinate.py

test-sun:
	@./test-sun.py

test-planet-comet-binary:
	@./test-planet-comet-binary.py

test-moon:
	@./test-moon.py

test-eclipses:
	@./test-eclipses.py

hygdata:
	@wget --no-check-certificate https://github.com/jfcarr-astronomy/HYG-Database-SQLite/raw/master/hygdata.db

canned:
	@./test-canned.py

default:
	@echo 'Targets:'
	@echo ' view            View documentation in Firefox.'
	@echo ' build-python    Build documentation from Python sources.'
	@echo ' build-rust      Build documentation from Rust sources.'

view: build
	@firefox ../docs/python/index.html &

build-python:
	@-rm -rf ../docs/python
	@doxygen

build-rust:
	@./build-rust.sh

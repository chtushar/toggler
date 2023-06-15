build:
	@sh -c './scripts/build.sh'

run: build
	'./bin/toggler'

build-api:
	@sh -c './scripts/build-api.sh'

tests:
	@sh -c './scripts/run-tests.sh'

build:
	@sh -c './scripts/build.sh'

run: build
	'./bin/toggler'
 

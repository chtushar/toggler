# cd internal/ui && npm run build && cd ../.. || exit 1

OUTPUT='bin/toggler'

echo "Building app $VERSION"

npm run build --prefix ./dashboard || exit 1
go build -o "$OUTPUT" ./main.go || exit 1

echo "Build successful. Output -> ${OUTPUT}"
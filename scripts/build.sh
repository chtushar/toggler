OUTPUT='bin/toggler'

echo "Building app $VERSION"

go build -o "$OUTPUT" ./cmd/main.go || exit 1

echo "Build successful. Output -> ${OUTPUT}"
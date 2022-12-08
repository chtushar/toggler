cd internal/ui && npm run build && cd ../.. || exit 1

OUTPUT='bin/toggler'

echo "Building app $VERSION"

go build -o "$OUTPUT" ./cmd/main.go || exit 1

echo "Build successful. Output -> ${OUTPUT}"
build:
    rm -f ./flat/*.go # Remove flatbuffer files

    ./flatbuffers-schema/binaries/flatc --go --gen-object-api --gen-all \
    --go-namespace flat \
    -o ./ ./flatbuffers-schema/schema/rlbot.fbs
    # Generate flatbuffer files

    for f in ./flat/*.go; do mv "$f" "./flat/generated_$(basename "$f")"; done

    go fmt
    go build # Build

#!/bin/bash

SPEC_FILE="auth-spec.yml" 
CONFIG="specgen_config.yaml" 
GENERATOR="go-server"    
OUTPUT_DIR="generated" 
COPY_PATH="src/generated" 

if ! command -v openapi-generator-cli &> /dev/null; then
    echo "Error: openapi-generator-cli is not installed."
    echo "Install it from https://openapi-generator.tech/docs/installation/"
    exit 1
fi

openapi-generator-cli generate \
    -i "$SPEC_FILE" \
    -g "$GENERATOR" \
    -o "$OUTPUT_DIR" \
    --config "$CONFIG"

mkdir -p $COPY_PATH
cp -R $OUTPUT_DIR/go/api.go $COPY_PATH
cp -R $OUTPUT_DIR/go/error.go $COPY_PATH
cp -R $OUTPUT_DIR/go/helpers.go $COPY_PATH
cp -R $OUTPUT_DIR/go/impl.go $COPY_PATH
cp -R $OUTPUT_DIR/go/model*.go $COPY_PATH
cp -R $OUTPUT_DIR/go/routers.go $COPY_PATH
rm -rf $OUTPUT_DIR

if [ $? -eq 0 ]; then
    echo "Code generation completed successfully"
else
    echo "Code generation failed."
    exit 1
fi
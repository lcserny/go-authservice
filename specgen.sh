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
cp -R $OUTPUT_DIR/go/*.go $COPY_PATH
rm -rf $OUTPUT_DIR

# Fix imports for errors package, openapi generator bug
FIX_FILES=$(grep -rl "errors.Is" $COPY_PATH)
for FILE in $FIX_FILES; do
    if grep -q "encoding/json" "$FILE"; then
        sed -i '/encoding\/json/a "errors"\r\n"io"' "$FILE"
        echo "Fixed errors package in: $FILE"
    fi
done

if [ $? -eq 0 ]; then
    echo "Code generation completed successfully"
else
    echo "Code generation failed."
    exit 1
fi
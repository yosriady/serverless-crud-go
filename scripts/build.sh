#!/usr/bin/env bash

echo "Compiling functions to bin/handlers/ ..."

cd src/handlers/
for f in *.go; do
  filename="${f%.go}"
  GOOS=linux go build -o "../../bin/handlers/$filename" ${f}
  echo "✓ Compiled $filename"
done

echo "Done."
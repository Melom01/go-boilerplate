#!/bin/bash
echo "Sorting imports in all Go files, excluding the generated DI files..."
find . -type f -name '*.go' ! -path './di/*_gen.go' -exec goimports -w {} +
echo "Done!"

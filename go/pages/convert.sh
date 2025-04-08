#!/usr/bin/bash

for file in *.md; do
    basename=$(basename "$file" .md)
    pandoc "$file" -o "$basename.html"
    pandoc "$file" -o "$basename.docx"
done
#!/bin/bash

FILES=(
  "screenshot-node-1-full.png"
  "screenshot-node-2-full.png"
  "screenshot-nodes-full.png"
  "screenshot-pod-1-full.png"
  "screenshot-pod-2-full.png"
  "screenshot-pods-full.png"
)

HEIGHT=450

for file in "${FILES[@]}"; do
  if [[ ! -f "$file" ]]; then
    echo "Skipping $file (not found)"
    continue
  fi

  base="${file%-full.png}"
  output="${base}-thumb.png"

  echo "Creating thumbnail: $output"

  # sips is "scriptable image processing system" included in macOS
  sips --resampleHeight "$HEIGHT" "$file" --out "$output"
done

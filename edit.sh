#!/bin/bash

if [ $# -ne 1 ]; then
  echo "Usage: $0 <folder>"
  exit 1
fi

FOLDER="$1"

MODFILE1="$FOLDER/go.mod"
MODFILE2="$FOLDER/README.md"

REPO="rajatsandeepsen/beckn-go"

for MODFILE in "$MODFILE1" "$MODFILE2"; do
  if [ ! -f "$MODFILE" ]; then
    echo "File not found: $MODFILE"
    continue
  fi

  # Detect OS for sed compatibility
  if sed --version >/dev/null 2>&1; then
    # GNU sed (Linux)
    sed -i "s|github.com/GIT_USER_ID/GIT_REPO_ID|github.com/$REPO/$FOLDER|g" "$MODFILE"
  else
    # BSD sed (macOS)
    sed -i '' "s|github.com/GIT_USER_ID/GIT_REPO_ID|github.com/$REPO/$FOLDER|g" "$MODFILE"
  fi

  echo "Replacement completed for file: $MODFILE"
done
#!/bin/bash

# Ensure we have at least two arguments
if [ "$#" -lt 2 ]; then
    echo "Usage: $0 <database-url> <file-pattern> ..."
    echo "Example: $0 'postgres://username:password@localhost:5432/yourdbname?sslmode=disable' '*/sqlc/schema.sql'"
    exit 1
fi

DATABASE_URL=$1
shift 1
PATTERNS=$@

# Loop through each pattern and apply the SQL files
for PATTERN in $PATTERNS; do
    for FILE in $PATTERN; do
        if [ -f "$FILE" ]; then
            echo "Applying $FILE..."
            psql "$DATABASE_URL" -f "$FILE"
            if [ $? -ne 0 ]; then
                echo "Error occurred while applying $FILE"
                exit 1
            fi
        else
            echo "No matching files for pattern: $PATTERN"
        fi
    done
done

echo "All DDL files have been successfully applied."

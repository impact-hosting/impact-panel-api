#!/bin/bash

cd /impact


echo "Building API"
go build -o ./impact_api
chmod +x ./impact_api

bash /scripts/setup_db.sh

echo "Starting API"
./impact_api
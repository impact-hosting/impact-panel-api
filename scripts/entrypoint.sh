#!/bin/bash

cd /impact

go build -o ./impact_api

bash ./impact_api &

/bin/bash
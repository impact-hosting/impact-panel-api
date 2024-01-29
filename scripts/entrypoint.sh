#!/bin/bash

cd /impact

go build -o ./dest/impact_api

bash ./dest/impact_api &

/bin/bash
#!/bin/bash

echo "Inserting test data"
mariadb --host $DB_HOST --port $DB_PORT --user $DB_USER --password $DB_PASSWORD < /impact/database/test_data.sql
#!/bin/bash

METRICS_FILE=data/metrics_from_special_app.txt

# create the metrics file if it doesn't exist
if [ ! -f "${METRICS_FILE}" ]; then
    touch "${METRICS_FILE}"
fi

# generate dummy metrics and write to file
for i in {1..10}; do
    metric_name="metric_${i}"
    metric_value=$((RANDOM % 1000))
    echo "${metric_name}=${metric_value}" >> "${METRICS_FILE}"
done
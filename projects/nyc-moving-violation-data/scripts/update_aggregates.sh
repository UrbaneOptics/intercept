#!/bin/bash

cd ./data/csv/ && \
echo "Aggregating data" && \
for file in ./**/*.csv; do cat "$file" >> merged.csv; done && \
echo "Processing aggregates" && \
ruby ../../scripts/retrieve_violation_names.rb merged.csv && \
ruby ../../scripts/retrieve_violation_count.rb && \

echo "Cleaning up artifacts" && \
rm merged.csv && \
echo "Aggregation complete. Validate and commit changes."

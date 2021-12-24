#!/bin/bash

# Run from CSV directory

# Aggregate all csv's
for file in **/*.csv; do cat "$file" >> merged.csv; done    

# Collect names and output to file
ruby ../../scripts/retrieve_violation_names.rb merged.csv

rm merged.csv

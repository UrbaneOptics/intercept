#!/bin/bash

read -p 'Year(YYYY): ' YEAR
read -p 'Month(MM): ' MONTH

curl https://www1.nyc.gov/assets/nypd/downloads/excel/traffic_data/archive/$YEAR-$MONTH-sum-excel.zip -o $YEAR-$MONTH-sum-excel.zip && \
unzip $YEAR-$MONTH-sum-excel.zip -d ./data/csv/$YEAR\_$MONTH\_sum/ && \
rm $YEAR-$MONTH-sum-excel.zip && \
cd ./data/csv/$YEAR\_$MONTH\_sum && \
sh ../../../scripts/excel_to_csv.sh && \
sh ../../../scripts/format_csv_files.sh

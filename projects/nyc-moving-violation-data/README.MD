# NYC Moving Violation Data
NYC takes monthly tallies of every precinct. This data is publicly available via NYC's [Moving Violations and Summonses](https://www1.nyc.gov/site/nypd/stats/traffic-data/traffic-data-moving.page) page. The data formats they provide are Excel(xlsx) and PDF. Unfortunately, the raw data has some improper formatting making it difficult to parse. 

This repository aims to format and present this data in a standardized way to make it easier to parse.  
(1) CSV files are provided in replacement to the standard Excel files. These CSV files are formatted to remove data that is unnecessary and standardize the format across all months.  
(2) [JSON files are provided as well](./data/json/aggregates#precinct-violation-aggregates).

## The data 
The cleaned up moving violation data can be found in the [data/csv](https://github.com/UrbaneOptics/nyc-moving-violation-data/tree/master/data/csv) folder. This contains all of the monthly precinct data in a consistent format.

Data for all of the moving violations the NYPD takes tallys under can be found in the [data/moving_violation_types](https://github.com/UrbaneOptics/nyc-moving-violation-data/blob/master/data/moving_violation_types.csv).

# Notes on data
On July 2013, the 121st precinct was created. Summons issues prior to that date will be
reflected in the totals for the 120th or 122nd precinct, depending on where they were issuedThe 121st precinct will only have totals from July 2013 onwards.# Progress
Data for years 2011-2019 appears to be fully normalized in CSV format. These CSVs should be able to be processed in a standard fashion.If you notice any issues with the data, please open a ticket or e-mail joe@urbaneoptics.com

The process for importing and normalizing a dataset is currently
```
Steps:
1) Manually rename folder to standard format in the csv folder(ie. 2020_08_sum).
2) cd into folder
3) run sh ../../../scripts/excel_to_csv.sh
4) run sh ../../../scripts/format_csv_files.sh
5) cd back to /csv folder
6) run sh ../../scripts/remove_filename_prefix.sh 
```

After a new dataset has been imported, violation names and precincts need to be updated. The JSON file generated
helps us see how often violation names are used, and how precincts are represented. Violation names used change over time, and occasionally precinct names as well. For example, the Midtown North Precinct is represented as both `MTN` and `Midtown North` at different times.

1) cd into the /csv folder 
2) Run `sh ../../scripts/update_violations_and_precincts.sh` to generate a JSON containing all the violations and precinct names


To update all aggregates

```
folders=(*)              
for i in "${folders[@]}"
do     
cd $i             
csv_files=(*.csv)
cd ..                      
for j in "${csv_files[@]}"                          
  ruby ../../scripts/retrieve_violation_count.rb $j
done
```

[Deprecated] To process all folders, from the `/data/ directory, run the following:
```
folders=(*.csv)
for i in "${folders[@]}"
do
  cd $i
  sh ../../scripts/excel_to_csv.sh 
  sh ../../scripts/format_csv_files.sh
  cd ..
done
```
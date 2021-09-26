# NYC Moving Violation Data
NYC takes monthly tallies of every precinct. This data is publicly available via NYC's [Moving Violations and Summonses](https://www1.nyc.gov/site/nypd/stats/traffic-data/traffic-data-moving.page) page. The data formats they provide are Excel(xlsx) and PDF. Unfortunately, the raw data has some improper formatting. In addition to that, a new spreadsheet file exists for each precinct for every month. This makes it difficult to parse moving violation data in any meaningful way without potentially incurring hours of manual work to aggregate and parse the data.

This repository aims to format and present this data in a standardized way. To do this, I've:  
(1) Provided CSV files to replace the shared Excel files. These CSV files are formatted to remove data that is unnecessary and standardize the format across all months.
(2) [JSON files are provided as well](./data/json/aggregates#precinct-violation-aggregates).

## The data 
The cleaned up moving violation data can be found in the [data/csv](https://github.com/UrbaneOptics/nyc-moving-violation-data/tree/master/data/csv) folder. This contains all of the monthly precinct data in a consistent format.

Data for all of the moving violations that the NYPD tallies can be found in the [data/moving_violation_types](https://github.com/UrbaneOptics/nyc-moving-violation-data/blob/master/data/moving_violation_types.csv).

# Notes on data

| Summary | Start Date | Note 
---|---|---
Precinct 121 Created | July 2013(07/2013) |  On July 2013, the 121st precinct was created. Summons issued prior to that date will be reflected in the totals for the 120th or 122nd precinct, depending on where they were issued. The 121st precinct will only have totals from July 2013 onwards.


If you notice any issues with the data, please open a ticket or e-mail joe@urbaneoptics.com

# Importing new datasets
The NYPD appears to update moving violation datasets on about a 2 month delay. So you should expect to see a July dataset appear sometime in September. To import a new dataset when they're posted on the [Moving Violations and Summonses](https://www1.nyc.gov/site/nypd/stats/traffic-data/traffic-data-moving.page) page, you should:

1) Download and extract the dataset to the `./data/csv/` folder in the following format: `<YYYY>_<MM>_sum`
2) cd into the folder you just extracted(ex. `cd ./data/csv/2021_07_sum`)
3) run `sh ../../../scripts/excel_to_csv.sh` to convert the excel files to CSV(NOTE: gnumeric is required here through `brew install gnumeric` if on OSX)
4) run `sh ../../../scripts/format_csv_files.sh` to normalize the csv files
5) run `cd ..` to go back to the main `csv/` folder 
6) run `sh ../../scripts/remove_filename_prefix.sh` to rename all of the newly formatted files

After a new dataset has been imported, violation names and precincts need to be updated. The JSON file generated
helps us see how often violation names are used, and how precincts are represented. Violation names used change over time, and occasionally precinct names as well. For example, the Midtown North Precinct is represented as both `MTN` and `Midtown North` at different times.

1) cd into the `/csv` folder 
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
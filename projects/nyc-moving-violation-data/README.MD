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


If you notice any issues with the data, please open a new issue in this repository

# Importing new datasets
## **Prerequisites**
- Gnumeric(brew install gnumeric)

## **Commands**

### **Downloading and extracting new datasets to CSV**
```shell
# Run from the root directory
sh scripts/download_extract_format.sh
```

The script will prompt you to enter a year and month and will then begin downloading, extracting and formatting the dataset into the `data/csv` directory.

**Note:** The NYPD appears to update the [Moving Violations and Summonses archive page](https://www1.nyc.gov/site/nypd/stats/traffic-data/traffic-data-moving.page) on about a 2 month delay. So you should expect to see a July dataset appear sometime in September. It is possible to download a dataset before the archive index page is updated, but there is a risk of this data changing.  To import a new dataset when they're posted on the page, you can run the following from the root directory:


### **Updating JSON aggregates**
When a new dataset has been imported, violation names and precincts need to be updated. The JSON file generated
helps us see how often violation names are used, and how precincts are represented. Violation names used change over time, and occasionally precinct names as well. For example, the Midtown North Precinct is represented as both `MTN` and `Midtown North` at different times.

The [data/json](./data/json) directory holds JSON files which contain a single file per precinct which holds aggregates of each of the moving violations a precinct has written since moving violation data has been public. 

To update all aggregates, run the following:
```shell
# Run from the root directory
sh scripts/update_aggregates.sh
```

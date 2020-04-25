# NYPD-Moving-Violations

[WIP] This project is an early work in progress. Expect major changes to occur frequently. 


The goal of this service is to work as an API to allow a user to query for historical moving violation information written by every precinct. This data is typically stored in Excel or PDF format, and the raw data can be found [here](https://www1.nyc.gov/site/nypd/stats/traffic-data/traffic-data-archive-2020.page). 

The raw data is unstandardized and someone who wishes to analyze historical data for just precinct would have to download and combine an excel file for every month they're interested in. This service aims to make the process obsolete.



# Running the service locally

### Required Software
- Postgresql
- Go

The easiest way to get started is to run the initialization scripts to create and seed your database. These steps assume you have GO set up and Postgres running locally.

From the root directory of this service, run:
```
cd pkg/scripts/

# Create the local DB with a new default user
sh create_local_db.sh 

# Restore a DB dump of historical data
sh restore_backups.sh

# Run the local dev server(defaults to localhost:4000)
go run ./cmd/web
```

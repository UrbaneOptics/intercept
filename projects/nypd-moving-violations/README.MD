# NYPD-Moving-Violations

The goal of this service is to work as an API to allow a user to query for historical moving violation information written by every precinct in New York City(NYC). This data is typically stored in Excel or PDF format, and the raw data can be found on the official [NYPD Traffic Data Archive page](https://www1.nyc.gov/site/nypd/stats/traffic-data/traffic-data-archive-2020.page). 

The raw data is unstandardized and someone who wishes to analyze historical data for just precinct would have to download and combine an excel file for every month they're interested in. This service aims to make the process obsolete by aggregating all of the statistics in an easily queryable relational database(Postgresql).

### Required Software
- Postgresql
- Go (only if running the API)
- Python3 (only if updating historical tallies)

# Importing the data
Running the API is not necessary to analyze information. An up to date database dump is accessible at [pkg/scripts/db.sql](./pkg/scripts/db.sql). You can import and analyze this data with your own tools of choice.

# Database Schema Explained
## Table: moving_violations
The moving violations table exists to keep a record of the names of the moving violations that occur. This table is fairly static but can change if the NYPD records new types of moving violations.
| Column name | Type and constraints | Purpose
---|---|---
id | integer | Primary key
name | varchar(unique) | The name of a moving violation(ex. 'Fail to Signal').

## Table: precincts
The precincts table exists to keep a record of the names of the precincts that exist. This table is fairly static but can change if the NYPD creates a new precinct.
| Column name | Type and constraints | Purpose
---|---|---
id | integer | Primary key
name | varchar(unique) | The name of a precinct(ex. 'Precinct 001').
short_name | varchar(unique) | The short name for a precinct(ex. '001' for 'Precinct 001').
is_aggregate | boolean | Whether the precinct represented is an aggregate of precincts(ex. 'Queens South').


## Table: tallies
The precincts table exists to keep a standardized record of the occurences of a moving violation for a particular year and date for a precinct.
| Column name | Type and constraints | Purpose
---|---|---
id | integer | Primary key
count | integer | The amount of times a moving violation occured.
month | integer | The month a moving violation occured.
year integer | The year a moving violation occured.
precinct_id | integer | references a `precincts` record.
moving_violation_id | integer | references a `moving_violations` record.

# Running the API locally
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

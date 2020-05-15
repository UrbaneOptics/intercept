# Precinct Violation Aggregates
Each file in this folder contains all moving violation data for an individual precinct.
There are currently 97 months worth of data being aggregated. Each time a new month is normalized and added,
the aggregates will be updated to include that month

The format of the data is as follows:
```json
{
  "<YEAR_MONTH_sum>": {
    "<VIOLATION_NAME>": "<COUNT>",
    "<VIOLATION_NAME>": "<COUNT>",
    ...
  },
  ...
}
```

### Field Definitions
`YEAR_MONTH_sum` references the year and month we're focusing on.

`VIOLATION_NAME` references the specific violation we're tracking. `Backing Unsafely` is an example of a violation name.

`COUNT` references the amount of violations that were given by the precinct.
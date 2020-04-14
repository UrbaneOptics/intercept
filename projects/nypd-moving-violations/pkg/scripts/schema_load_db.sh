database="intercept_nypd_mv"
user="nypdmv"

#Execute few psql commands: 
#Note: you can also add -h hostname -U username in the below commands. 
 
echo "Dropping tables"
psql -d $database -U nypdmv << EOF 
DROP TABLE tallies;
DROP TABLE moving_violations;
DROP TABLE precincts;
EOF

echo "Loading tables"
psql -d $database -U nypdmv << EOF 
CREATE TABLE precincts(
  id serial PRIMARY KEY,
  name VARCHAR(50) UNIQUE NOT NULL default '',
  is_aggregate BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE moving_violations(
  id serial PRIMARY KEY, 
  name VARCHAR(100) UNIQUE NOT NULL default ''
);

CREATE TABLE tallies(
  id serial PRIMARY KEY,
  count INT NOT NULL,
  month INT NOT NULL ,
  year INT NOT NULL,
  precinct_id INT REFERENCES precincts(id),
  moving_violation_id INT REFERENCES moving_violations(id)
);
EOF

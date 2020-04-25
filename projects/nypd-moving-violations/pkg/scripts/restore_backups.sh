database="intercept_nypd_mv"
user="nypdmv"

echo "Dropping tables"
psql -d $database -U $user << EOF 
DROP TABLE precincts, moving_violations, tallies;
EOF

echo "Restoring backup"
psql $database  < db.sql

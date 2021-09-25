# This will convert all excel files in the current working directory to CSV files
# and it will delete the excel files afterward.
# Requires gnumeric to be installed(brew install gnumeric)
find . -maxdepth 1 -name "*.xlsx" -exec ssconvert {} --export-type=Gnumeric_stf:stf_csv \;

rm *.xlsx

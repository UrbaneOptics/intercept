# This will remove the 'mv-en-us-' prefix from all filenames existing in folders
# from the directory this is called in.
for filename in **/*.csv; do 
    [ -f "$filename" ] || continue
    mv "$filename" "${filename//mv-en-us-/}"
done

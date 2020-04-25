psql << EOF
CREATE DATABASE intercept_nypd_mv;
CREATE USER nypdmv WITH ENCRYPTED PASSWORD 'password';
GRANT ALL PRIVILEGES ON DATABASE intercept_nypd_mv TO nypdmv;
EOF
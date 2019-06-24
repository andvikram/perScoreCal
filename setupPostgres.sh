set -e
# source .env

# sudo printf "CREATE USER $DEV_USERNAME WITH PASSWORD '$DEV_PASSWORD';\nCREATE DATABASE $DEV_DBNAME WITH OWNER $DEV_USERNAME;" > setupPostgres.sql
# sudo -u postgres psql -f setupPostgres.sql

createdb -h $DEV_HOST -p $DEV_PORT -O $DEV_USERNAME -e $DEV_DBNAME

# go get github.com/tools/godep
# godep restore

go run main.go setupdb

go run main.go serve

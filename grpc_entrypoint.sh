#!bin/sh

set -e
COMMAND=$@

echo "Waiting for the database to be ready..."
max_retries=10
while [ $max_retries -gt 0 ] && ! mysql -h "$MYSQL_HOST" -P "$MYSQL_PORT" -u"$MYSQL_USER" -p"$MYSQL_PASSWORD" "$MYSQL_DB" -e 'SHOW TABLES'; do
  if nc -z $DB_HOST $DB_PORT; then
    echo "Database is ready!"
    break
  fi
  max_retries=$(($max_retries-1))
  sleep 5
done
echo
if [ $max_retries -eq 0 ]; then
  echo >&2 'Could not connect to the database. Aborting...'
  exit 1
fi

exec $COMMAND
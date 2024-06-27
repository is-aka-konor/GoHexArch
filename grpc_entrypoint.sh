#!/bin/sh

set -e
COMMAND=$@

echo 'Waiting for database to be available...'
maxTries=10
while [ "$maxTries" -gt 0 ] && ! mysql --default-auth=mysql_native_password -u $MYSQL_USER --password=$MYSQL_PASSWORD -e 'SHOW DATABASES'; do
    maxTries=$(($maxTries - 1))
    sleep 3
done
echo
if [ "$maxTries" -le 0 ]; then
    echo >&2 'error: unable to contact mysql after 10 tries'
    exit 1
fi
echo 'Connection was successful!'
exec $COMMAND
#!/bin/sh

while true ; do
    VAR0=$(curl -s -i http://localhost:3000/api/v1/users/search | awk 'NR==1{print $2}')

    echo $(curl -i http://localhost:3000/api/v1/users/search)

    echo ${VAR0}
    echo 200
    if [ "$VAR0" -eq 200 ]; then
        break
    fi
    echo "sleeping"
    sleep 2
done

echo "yes"
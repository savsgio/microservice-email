#!/usr/bin/env bash

# SCRIPT TO DELETE DOCKER CONTAINER WHEN EXIST AND NOT ACTIVE

CONTAINER_NAME=$1
CONTAINER_ID=$(docker ps -a -q --filter name=${CONTAINER_NAME})

if [[ ${CONTAINER_ID} != '' ]]; then

    for ID in ${CONTAINER_ID}; do

        ACTIVE=$(docker inspect -f {{.State.Running}} ${ID})

        if [[ ${ACTIVE} != 'true' ]]; then
            docker rm ${ID}
        fi

    done

fi

#!/usr/bin/env bash

. ./environment/docker_scripts/config.sh

echo "Starting server..."

./environment/docker_scripts/check_if_container_exist.sh ${PROJECT_NAME}

if [[ $? > 0 ]]; then
    echo ""
    echo "There was an error verifying the container, check that there are no containers with the name \"${PROJECT_NAME}\" created"
    echo "Command ==> docker ps -a --filter name=${PROJECT_NAME}"
    echo ""
    exit $?
fi

docker-compose -p ${PROJECT_NAME} -f ${COMPOSER_FILE} run --rm --service-ports --name ${PROJECT_NAME} ${PROJECT_NAME} /bin/bash \
    -ci "make get && make run"
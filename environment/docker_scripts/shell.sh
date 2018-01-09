#!/usr/bin/env bash

. ./environment/docker_scripts/config.sh

docker-compose -p ${PROJECT_NAME} -f ${COMPOSER_FILE} run --rm ${PROJECT_NAME} /bin/bash

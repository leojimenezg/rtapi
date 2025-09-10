#!/bin/bash

docker build . -t leojimenezg/rtapi-db:v1.0.0
docker run --name rtapi-db -d -p 3306:3306 --env-file ./db.env leojimenezg/rtapi-db:v1.0.0

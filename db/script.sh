#!/bin/bash

# Build the image (only once)
docker build . -t leojimenezg/rtapi-db:v1.0.0
# Create and run the container
docker run --name rtapi-db -d -p 3306:3306 --env-file ./db.env leojimenezg/rtapi-db:v1.0.0

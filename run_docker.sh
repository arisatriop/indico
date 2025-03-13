#!/bin/bash

# Define variables
IMAGE_NAME="indico-technical-test:latest"
CONTAINER_NAME="indico-technical-test"
DOCKERFILE="Dockerfile"

# Step 1: Build the Docker image
docker rmi -f $IMAGE_NAME
echo "Building the Docker image: $IMAGE_NAME"
docker build -t $IMAGE_NAME -f $DOCKERFILE . 

if [ $? -ne 0 ]; then
    echo "Failed to build the Docker image."
    exit 1
fi

# Step 3: Check if the container already exists
if [ $(docker ps -a --filter "name=^${CONTAINER_NAME}$" --format '{{.Names}}' | wc -l) -eq 1 ]; then
    echo "Container $CONTAINER_NAME already exists. Removing it..."
    docker rm -f $CONTAINER_NAME
fi

# Step 3: Run the Docker container
echo "Running the Docker container:"
docker run --name $CONTAINER_NAME -v .:/app -p 8080:8080 \
    -e APP_ENV=local \
    -e APP_NAME="Warehouse management system" \
    -e DBMS=psql \
    -e DB_HOST=host.docker.internal \
    -e DB_PORT=5432 \
    -e DB_USER=postgres \
    -e DB_PASSWORD=postgres \
    -e DB_NAME=postgres \
    $IMAGE_NAME 



#!/bin/bash

docker compose down
docker rmi backend-backend:latest
docker compose up -d

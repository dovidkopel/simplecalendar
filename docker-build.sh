#!/usr/bin/env bash

docker build -t dovidkopel/aws-sam-local:latest .
docker tag dovidkopel/aws-sam-local:latest dovidkopel/aws-sam-local:0.3.0
docker login -u dovidkopel -p AE86EsqV\!5c@3xhp
docker push dovidkopel/aws-sam-local:latest
docker push dovidkopel/aws-sam-local:0.3.0
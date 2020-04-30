#!/bin/bash
docker-compose exec rabbitmq /usr/local/bin/rabbitmqadmin --username=yjc --password=yjc666 declare exchange name=apiServers type=fanout
docker-compose exec rabbitmq /usr/local/bin/rabbitmqadmin --username=yjc --password=yjc666 declare exchange name=dataServers type=fanout
docker-compose exec rabbitmq rabbitmqctl add_user test test
docker-compose exec rabbitmq rabbitmqctl set_permissions -p / test ".*" ".*" ".*"

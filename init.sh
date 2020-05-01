#!/bin/bash
docker-compose exec rabbitmq /usr/local/bin/rabbitmqadmin --username=yjc --password=yjc666 declare exchange name=apiServers type=fanout
docker-compose exec rabbitmq /usr/local/bin/rabbitmqadmin --username=yjc --password=yjc666 declare exchange name=dataServers type=fanout
docker-compose exec rabbitmq rabbitmqctl add_user test test
docker-compose exec rabbitmq rabbitmqctl set_permissions -p / test ".*" ".*" ".*"
docker-compose exec elastic curl 127.0.0.1:9200/metadata -XDELETE
docker-compose exec elastic curl 127.0.0.1:9200/metadata -H 'Content-type':'application/json' -XPUT -d'{"mappings":{"properties":{"name":{"type":"keyword"},"version":{"type":"integer"},"size":{"type":"integer"},"hash":{"type":"text"}}}}'
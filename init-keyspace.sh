#!/bin/bash

sleep 5  # Wait for Cassandra to start

cqlsh -e "CREATE KEYSPACE restaurant WITH replication = {'class':'SimpleStrategy', 'replication_factor':1};"

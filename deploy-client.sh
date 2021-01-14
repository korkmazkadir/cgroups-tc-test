#!/bin/bash

ip_address=$1
port_number=$2
number_of_nodes=$3

rm client.log
touch client.log

for (( i=1; i<=$number_of_nodes; i++ ))
do  

   ./client-app "${ip_address}" "${port_number}" 2>> "client.log" &

done
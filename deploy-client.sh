#!/bin/bash


read -p "Server IP address: " ip_address
read -p "Base port number: " base_port_number
read -p "Number of servers: " number_of_servers
read -p "Number of client apps: " number_of_nodes

sudo killall -r client-app
rm client.log
touch client.log

for (( i=1; i<=$number_of_nodes; i++ ))
do  

   ./client-app "${ip_address}" "${base_port_number}" "${number_of_servers}" 2>> "client.log" &

done

tail -100f client.log


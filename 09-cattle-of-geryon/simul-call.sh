#!/bin/bash

python3 siege.py -u http://23.23.42.188:3000/api/parts/3 -c 20 -r 50 &
python3 siege.py -u http://23.23.42.188:3000/api/parts/30 -c 20 -r 50 &
python3 siege.py -u http://23.23.42.188:3000/api/parts/300 -c 20 -r 50 &
python3 siege.py -u http://23.23.42.188:3000/api/parts/3000 -c 20 -r 50 &
python3 siege.py -u http://23.23.42.188:3000/api/parts/30000 -c 20 -r 50 &
python3 siege.py -u http://23.23.42.188:3000/api/parts/300000 -c 20 -r 50

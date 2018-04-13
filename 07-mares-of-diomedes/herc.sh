#!/bin/bash
./manger & sleep 5
#kill -9 $(pgrep manger lampon yes crap ruins xanthos)
#pgrep lampon yes crap ruins xanthos
pkill -9 lampon yes crap ruins xanthos
#top -l 1 | grep manger | cut -d" " -f1
ps | grep manger
kill -s KILL $(ps | grep manger | cut -d" " -f1)
rm -r *.poo

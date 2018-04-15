#!/bin/bash

# confirm that the user really wants to run the manger binary
clear; echo $'Releasing the Mares of Diomedes will bring disaster upon us all. Are you sure you want to do this?\n'
read -rsp $'Press enter to continue...\n'
# countdown before running the mares binary
for ((i=3; i>0; i--)) 
do
    clear; echo "Releasing the mares in $i..."; sleep 1 
done
clear; echo "The Mares of Diomedes are upon us!"; sleep 1


# run the manger binary in the background, then wait 5 seconds
./manger &
sleep 5

# tame lampon - using SIGTRAP (-5)
pkill -5 lampon
# tame xanthos - using SIGUSR1 (-10)
pkill -10 xanthos -bash yes
# tame deinos - using SIGSEGV (-11)
pkill -11 deinos crap; rm -f *.poo
# tame podargos - using SIGILL (-4)
pkill -4 podargos ruins


# confirm that the user really wants to release true terror
clear; echo $'Thank the Gods, the Mares have been tamed!\n'
read -rsp $'Press enter to release true terror!\n'
# countdown to armegeddon
for ((i=3; i>0; i--)) 
do
    clear; echo "Releasing true terror in $i..."; sleep 1 
done
clear;


# compile and start armegeddon, then let run for 3 seconds
gcc armegeddon.c -Wall -Werror -Wextra -o armegeddon && ./armegeddon &
sleep 10

# stop armegeddon
kill -s KILL $(pgrep armegeddon)
pkill sh Preview > /dev/null 2>&1
rm -f *.crap




# ---- ALTERNATIVE WAYS OF TAMING THE MARES ----

# # tame the mares and clean the poo in one line with kill and pgrep
# kill -s KILL $(pgrep manger lampon yes crap ruins xanthos -bash); rm -rf *.poo

# # tame the mares and clean the poo in one line with pkill
# pkill -9 manger lampon yes crap ruins xanthos -bash; rm -rf *.poo

# # A BUNCH OF DIFFERENT WAYS:
# kill -s KILL $(ps -Af | grep lampon | awk '{print $2}')
# kill -s KILL $(pgrep xanthos -bash yes)
# pkill -9 deinos crap
# ps | grep ruins | awk '{print $1}' | xargs kill -9
# kill -s KILL $(top -l 1 | grep manger | cut -d" " -f1)
# rm -rf *.poo

# ----------------------------------------------
# Labour 03: Erymanthian Boar

Goal: This project aims to give you a small approach on timed commands and their execution.

> Write a small shell script that when you run it will execute it’s code on the upcoming 
> 21st of December at 8:42am. The content of the script can be anything you want.
> Make sure your script confirms the creation of the timed task by showing the date and time of the execution.  
> *Note: cron is recurring, look for a command that is executed only once at a specified time.*

## How To Run
To run this script on a macOS system: 

1. By default, the **atrun** utility is disabled on macOS. To enable it, run the following command :   
```
sudo launchctl load -w /System/Library/LaunchDaemons/com.apple.atrun.plist
```

2. Run the script
```
sh boaring-schedule.sh
```

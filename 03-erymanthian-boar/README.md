# Labour 03: Erymanthian Boar
**Goal**  
This project aims to give you a small approach on timed commands and their execution.

**Mandatory**  
Write a small shell script that when you run it will execute it’s code on the upcoming 21st of December at 8:42am. The content of the script can be anything you want. Make sure your script confirms the creation of the timed task by showing the date and time of the execution.  

**_*Note:_** *cron is recurring, look for a command that is executed only once at a specified time.*

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

## The 'at' command
Although `crontab` is generally used to set recurring tasks, the `at` command is very useful for scheduling one time tasks.  

**Commands used with `at`**
- **at**: *execute commands at specified time.*
- **atq**: *lists the pending jobs of users.*
- **atrm**: *delete jobs by their job number.*

**Examples**
1. Schedule first job using `at` command
    ```bash
    # will run “sh backup.sh” at 9:00 in the morning.
    $ echo "sh boaring-schedule.sh" | at 9:00 AM
    ```
2. List the scheduled jobs using `atq`
    ```bash
    # executing as root shows all users jobs
    $ atq

    # job id | execution date | execution time | user
    3       2018-03-23 09:00 a root
    5       2018-03-23 10:00 a nmolina
    1       2018-03-23 12:00 a root
    ```
3. Remove scheduled job using `atrm`
    ```bash
    # You can remove any at job with their job id.
    $ atrm 3
    $ atq

    5       2018-03-23 10:00 a nmolina
    1       2018-03-23 12:00 a root
    ```
4. Check the content of scheduled `at` job
    ```bash
    # will show what commands are scheduled for job id 5
    $ at -c 5

    ...
    ...
    sh boaring-schedule.sh
    ```
## Syntax examples for `at`

Schedule task at coming 10:00 AM.
```
$ at 10:00 AM
```
Schedule task at 10:00 AM on coming Sunday.
```
$ at 10:00 AM Sun
```
Schedule task at 10:00 AM on coming 25’th July.
```
$ at 10:00 AM July 25
```
Schedule task at 10:00 AM on 22’nd June 2018.
```
$ at 10:00 AM 6/22/2018
$ at 10:00 AM 6.22.2018
```
Schedule task at 10:00 AM on same date next month.
```
$ at 10:00 AM next month
```
Schedule task at 10:00 AM tomorrow.
```
$ at 10:00 AM tomorrow
```
Schedule task in 24 hours.
```
$ at now + 24 hours
```
Schedule task to execute just after 1 hour.
```
$ at now + 1 hour
```
Schedule task to execute just after 30 minutes.
```
$ at now + 30 minutes
```
Schedule task to execute just after 1 and 2 weeks.
```
$ at now + 1 week
$ at now + 2 weeks
```
Schedule task to execute just after 1 and 2 years.
```
$ at now + 1 year
$ at now + 2 years
```
Schedule task to execute at midnight.
```
$ at midnight
```

Reference:  
https://tecadmin.net/one-time-task-scheduling-using-at-commad-in-linux/#


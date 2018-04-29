# Labour 09: Cattle of Geryon
**Goal**  
For this labour you will write a program like siege. Learn about how siege works and can be used to benchmark a server. The point of this program is to simulate placing a server "under siege." You can use any language.

**Mandatory**  
Release a great flood of simulated clients! As mentioned before, write a program like siege, Test HTTP load and benchmark how the server runs under different loads / attacks.

## **Commands to Run**   
**Run Benchmark using Wizard**
```
$ python3 siege.py
```
**Run Benchmark using Flags**
```
$ python3 siege.py -u http://google.com -c 100 -r 10
```
**Note:**
- **-u:** *url to benchmark*   
    *(default: http://23.23.42.188)*
- **-c:** *number of clients/threads to use*  
    *(default: 100)*
- **-r:** *number of requests per client*  
    *(default: 5)*  

**Usage**
```
usage: siege.py [-h] [-u url] [-c clients] [-r requests]

Simple http load tester and benchmarking utility

optional arguments:
  -h, --help   show this help message and exit
  -u url       url/website benchmark will be used on
  -c clients   number of clients/threads
  -r requests  number of requests per clients
```
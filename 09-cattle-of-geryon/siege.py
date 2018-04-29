import requests
import threading
import sys
import time
import argparse
import os

# ---------------- FUNCTIONS ----------------
def attack():
    hits = 0
    while hits < reqs:
        try:
            r = requests.get(url, timeout=3)
            if r.status_code == 200:
                results['success'] += 1
            else:
                results['failure'] += 1
            results['time'] += float(str(r.elapsed)[5:11])
            size = len(r.content)
            results['size'] += size
        except:
            results['failure'] += 1
        hits += 1
# ---------------- DEFINE FLAGS ----------------
parser = argparse.ArgumentParser(description='Simple http load tester and benchmarking utility')
parser.add_argument('-u', metavar='url', help='url/website benchmark will be used on')
parser.add_argument('-c', metavar='clients', type=int, help='number of clients/threads')
parser.add_argument('-r', metavar='requests', type=int, help='number of requests per clients')
args = parser.parse_args()
# ---------------- SET VARIABLES ----------------
# set parameters
url = args.u if args.u != None else "http://23.23.42.188"
clients = args.c if args.c not in (None, 0) else 100
reqs = args.r if args.r not in (None, 0) else 5
# if no flags provided, run wizard.
if args.u == args.c == args.r == None:
    os.system("clear")    
    url = input(f'What kingdom should we attack?\n(default: {url})\n\nwebsite URL: ') or url
    os.system("clear")
    clients = int(input(f'How many legions should we amass?\n(default: {clients})\n\nclients/threads: ') or clients)
    os.system("clear")
    reqs = int(input(f'How many soldiers should each legion have?\n(default: {reqs})\n\nrequests per client: ') or reqs)
    os.system("clear")
# results parameters
results = {
    'transactions': clients * reqs,
    'success': 0,
    'failure': 0,
    'time': 0,
    'size': 0
}
# ---------------- CODE START ----------------
# Check for valid url
try:
    r = requests.get(url)
    print (f'''
STARTING LOAD TEST
URL:        {url}
Clients:    {clients}
Requests:   {reqs}
''')
except:
    print ('error: bad url <{}>'.format(url))
    sys.exit(1)
# Start benchmark
print ('----------- Running Benchmark -----------')
start_time = time.time()
threads = []
for clients in range (1, clients+1):
    t = threading.Thread(target=attack)
    threads.append(t)
    t.start()
for t in threads:
    t.join() # wait until finished
sys.stdout.flush()
elapsed_time = time.time() - start_time
# Displaying Results
print (f'''
Transactions:               {results['transactions']} hits
Availability:               {(results['success'] * 100) / results['transactions']} %
Successful transactions:    {results['success']}
Failed transactions:        {results['failure']}
Elapsed time:               {round(elapsed_time, 3)} secs
Average response time:      {round(results['time'] / results['transactions'], 3)} secs
Totaled response times:     {round(results['time'], 3)} secs
Data transfered:            {round(results['size'] / 1024, 3)} KB
''')

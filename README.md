# clpsrvlist
- This module gets the status of cluster servers with RESTful API.

## Sample Configuration
```
+--------------+     +------------------+
| client       +--+--+ server1          |
| - clpsrvlist |  |  | - EXPRESSCLUSTER |
+--------------+  |  +------------------+
                  |
                  |  +------------------+
                  +--+ server2          |
                  |  | - EXPRESSCLUSTER |
                  |  +------------------+
                  :
```

## Prerequisite
- Enable RESTful API of EXPRESSCLUSTER.

## How to Use
1. Donwload the module from [the release page](https://github.com/EXPRESSCLUSTER/clpsrvlist/releases).
   - The current version supports Linux OS only.
1. Create servers.json file as below.
   ```json
   [  
       [
           {
               "ipaddress": "192.168.122.1",
               "port": "29009",
               "user": "root",
               "password": "password"    
           }, {
               "ipaddress": "192.168.122.2",
               "port": "29009",
               "user": "root",
               "password": "password"    
           }
       ], [
           {
               "ipaddress": "192.168.122.3",
               "port": "29009",
               "user": "root",
               "password": "password"    
           }, {
               "ipaddress": "192.168.122.4",
               "port": "29009",
               "user": "root",
               "password": "password"    
           }
       ]
   ]
   ```
   - This sample has two clusters. The first cluster has server1 (192.168.122.1) and server2 (192.168.122.2). The second one has server3 (192.168.122.3) and server4 (192.168.122.4).
1. Run clpsrvlist.
   ```
   ./clpsrvlist
   ```
   - If you want to enable error messages, use -d option.
     ```sh
     ./clpsrvlist -d 1
     ```
1. You can get the result as below.
   ```
   server1 , Online
   server2 , Online
   failover1 , Online , server1
   failover2 , Online , server1
   server3 , Online
   server4 , Online
   failover3 , Online , server3
   failover4 , Online , server4
   ```

## How to Build
1. Install Go.
1. Clone the repository.
   ```sh
   git clone https://github.com/EXPRESSCLUSTER/clpsrvlist.git
   ```
1. Build.
   ```sh
   go build
   ```

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
1. Donwload the module.
1. Create servers.json file as below.
   ```json
   {
       "servers": [{
           "user": "root",
           "password": "password",
           "ipaddress": "192.168.122.101",
           "port": "29009",
           "hostname": "server1"
       }, {
           "user": "root",
           "password": "password",
           "ipaddress": "192.168.122.102",
           "port": "29009",
           "hostname": "server2"
       }]
   }   
   ```
1. Run clpsrvlist.
   ```
   ./clpsrvlist
   ```
   - If you want to enable error messages, use -m option.
     ```sh
     ./clpsrvlist -m 1
     ```
1. You can get the result as below.
   ```
   server1 Online
   server2 Online
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

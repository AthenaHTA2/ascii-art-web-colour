#!/bin/bash
# sudo docker stop $(docker ps -a -q)
# sudo docker rm $(docker ps -a -q)
sudo docker system prune -a
y

sudo docker build -t binarycolour .
sudo docker container run -p 8080:8080 -d --name containerbinarycolour binarycolour
sudo docker image ls
sudo docker ps -a
sudo docker exec -it containerbinarycolour /bin/bash
sudo ls -l
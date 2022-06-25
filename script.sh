(!/bin/sh)

sudo docker stop $(docker ps -a -q)
sudo docker rm $(docker ps -a -q)
sudo docker system prune -a
sudo docker y

sudo docker build -t dockerforum .
# sudo docker container run -p 8080:8080 -d --name containerdockerforum dockerforum
sudo docker image ls
sudo docker ps -a
sudo docker exec -it containerdockerforum /bin/bash
sudo ls -l
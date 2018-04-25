#!/bin/bash
docker stop client && docker rm client
docker rmi client && docker rmi $(docker images -a | grep "none")
docker build -t client .
docker run  -itd --name client \
  -v /home/robert/react-techrouteplanner/src/shared/react:/client \
  client
#docker exec -ti client /bin/bash
docker logs -f client

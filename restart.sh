docker login xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx

docker pull xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx

docker rm -f gin-api

docker run -d --restart=always -p 4000:4000 --name gin-api xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx ./gin-api

docker image prune -f
docker container prune -f

docker ps

docker logs -f gin-api

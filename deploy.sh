docker build . -t xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
docker login xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
docker push xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
ssh server "./restart.sh"

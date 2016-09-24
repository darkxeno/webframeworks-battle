

docker rm some-mongo
docker run -v /tmp/mongodb:/data/db --publish 27017:27017 --name some-mongo -d mongo:3.3

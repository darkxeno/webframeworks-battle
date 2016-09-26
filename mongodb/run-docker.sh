
docker stop mongodb-with-testdata
docker rm mongodb-with-testdata
docker run --publish 27017:27017 --name mongodb-with-testdata -d mongodb-with-testdata
#docker run -v /tmp/mongodb:/data/db --publish 27017:27017 --name mongodb-with-testdata -d mongodb-with-testdata

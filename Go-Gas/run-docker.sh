docker stop my-running-app
docker rm my-running-app
docker run -d -p=8181:8181 --env MONGODB_URL=192.168.0.5:27017 --name my-running-app my-golang-app

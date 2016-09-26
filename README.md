# Web Frameworks - battle

Objectives:
* Compare the performance of diferent programming languages / frameworks / libs
* Solve a I/O bound scenario
* Calculate RPS and latency under diferent loads (amount of data 1K / 10K / 100K / 1000k )
* Register memory comsuption ( peak / average )
* Compare code quality of the diferent solutions


Requirements for each test:
* Create a folder following the naming pattern: {language}-{framework}-{scenario}
* Create a GET /test webservices, allowing to limit the amount of mongodb rows like this /test?limit=10
	This webservice should access the mongodb local database and the data from the testData collection
* Be able to receive and environment var like this MONGODB_URL=192.168.0.5:27017
* Create a Dockerfile for the deployment
* Create a run-docker.sh bash script on the test folder root to start the docker container


# Task
Dockerize a simple web application. You can reference provided examples if you are stuck, but try to do as much as possible using only cheatsheet and your own knowledge.

## About the app
- The app is a simple one file go web server that serves static HTML page on port 8080.
- The page is a to-do list app that allows you to add and remove items from the list.
- The app stores to-do items in the postgres database.
- The app connects to the database using environment variables for configuration.

> Note: to compile go code you need to use `go build` command. The resulting binary will have the same name as the folder containing the code.

## Requirements
- Create a Dockerfile that will build the go application and run it.
- Create a docker-compose.yml file that will define two services: the web app and the postgres database.
- The web app service should build the Dockerfile you created.
- The web app service should expose port 8080.
- The web app service should set the necessary environment variables for connecting to the database.
- The postgres service should use the official postgres image.
- The postgres service should set the necessary environment variables for setting up the database (e.g. POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_DB).
- The postgres service should expose port 5432.
- The web app service should depend on the postgres service. (i.e. the database should be started before the web app (this thing is deprecated but still useful for simple cases)) (if it does not work try to figure it out)
- Create `run.sh` script that will build and start the application using docker-compose.

## Env variables for the web app
- POSTGRES_HOST: hostname of the postgres service (use the service name defined in docker-compose.yml)
- POSTGRES_PORT: port of the postgres service (usually 5432)
- POSTGRES_USER: username for the database
- POSTGRES_PASSWORD: password for the database
- POSTGRES_DB: name of the database


If you do everything correctly, you should see the web app running on http://localhost:8080 and be able to add and remove tasks from the to-do list.

Good luck! :)
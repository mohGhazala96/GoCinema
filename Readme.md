# GoCinema #
GoCinema is an online platform that lets the user reserve seats online for different screening times and dates

## 2. Dependancies: ##
- Docker
- GoLang as a RestFUL API
- PostgreSQL as our data storage system
- TheMovieDatabase - TMDb https://developers.themoviedb.org
- Angular 6 as a frontend

## 3. Getting started: ##
1. Register on TMDb as a developer and get your API Key
Sign up here: https://www.themoviedb.org/account/signup 

2. The first thing you should do is to git clone the project: 
```
git clone git@github.com:mohGhazala96/GoCinema.git
```

3. Navigate to the ```env/``` directory

4. Create 2 files and use the example files for reference
```
goenv.env 
pgenv.env
```
5. For more information about how to setup the environmental variables check section 4

6. To run the application check section 5

## 4. Environmental Variables ##

### pgenv.env ###

```
POSTGRES_USER=REPLACE_WITH_DB_USERNAME
POSTGRES_PASSWORD=REPLACE_WITH_DB_PASS
POSTGRES_DB=REPLACE_WITH_DB_NAME
```
##### POSTGRES_USER #####

This is your root /admin account that you will use through out your application to access the postgres database

##### POSTGRES_PASSWORD #####

This is your root/admin password. Make sure that it is something secure

##### POSTGRES_DB #####

This is the database name that you will use to store the tables and data of your application. 


### goenv.env ###
```
DATABASE_HOST=db
DATABASE_USER=REPLACE_WITH_DB_USERNAME
DATABASE_PASSWORD=REPLACE_WITH_PASSWORD 
DATABASE_NAME=REPLACE_WITH_DB_NAME
DATABASE_PORT=5432
WEB_HOST=0.0.0.0
WEB_PORT=3000
API_Key=REPLACE_WITH_API_KEY
```
##### DATABASE_HOST #####

The host of the database that we will connect to

##### DATABASE_USER (Replace the value) #####

The PostgreSQL username that you will connect with. This username should match the POSTGRES_USER

##### DATABASE_PASSWORD (Replace the value) #####

The PostgreSQL password that you will connect with. This password should match the POSTGRES_PASSWORD

##### DATABASE_NAME (Replace the value) #####

This name should match the POSTGRES_DB 

##### DATABASE_PORT #####

The port that we access the database through. 

##### WEB_HOST #####

This sets the backend to listen to requests sent to this ip (0.0.0.0 is global)

##### WEB_PORT #####

The port that the go-app runs through

##### API_KEY (Replace the value) #####

This is your unique API KEY. Get it by following step 1 in Getting Started section

## 5. Running the Application ## 

There are 2 ways to run the platform. Either use docker run or docker-compose. 

Make sure that you followed the steps properly and have `env/goenv.env` and `env/pgenv.env` setup properly

### A. Using docker run ###
This section explains how to run each container individually.

#### First Run: ####
Go to the directory of the project

1. Run the PostgreSQL database first in a terminal instance
```
docker run -p 5432:5432 --name db --env-file ./env/pgenv.env -v ${PWD}/data.sql:/docker-entrypoint-initdb.d/init.sql -w /docker-entrypoint-initdb.d/init.sql -v ${PWD}/postgres-data:/var/lib/postgresql/data -w /var/lib/postgresql/data healthcheck/postgres:alpine
```

2. Make sure that this message appears before moving onto the next step
```
LOG: database system is ready to accept connections
```

3. Run Go-App in a different terminal instance (tab)
```
docker build -t go-app ./go-app
docker run -it --name goapp -p 3000:3000 --env-file ${PWD}/env/goenv.env --link db:postgres go-app
```

4. Run the frontend in a different terminal instance (tab)
To run the frontend for the first time
```
docker build -t angular-app ./frontend
docker run -it --name angularapp -p 4200:4200 -v ${PWD}/frontend:/usr/src/app -w /usr/src/app angular-app
```

**Note: If you want to run them in daemon mode add the -d flag**

#### Stopping the platform ####
1. Stop PostgreSQL
```
docker stop db
```

2. Stop Go-App
```
docker stop goapp
```

3. Stop Frontend
```
docker stop angularapp
```

#### To start the platform again #### 
1. Start PostgreSQL
```
docker start db
```

2. Start Go-App
```
docker start goapp
```

3. Start Frontend
```
docker start angularapp
```

#### If you want to reconfigure everything ####
1. Delete the postgres-data folder
```
rm -r postgres-data/
```
2. Remove the containers that you made
```
docker rm db
docker rm goapp
docker rm angularapp
```

3. Run the platform as if you are running it for the first time (check the section above)

### B. Using docker-compose ###
To run the application for the first time
```
docker-compose up --build
```

To run the application afterwards (without building)
```
docker-compose up
```

To close down the app
```
docker-compose down
```

**If you would like to run the application in daemon mode**
```
docker-compose up --build -d
```
and to inspect the logs

```
docker logs db
docker logs app
docker logs angular-app
```

## Developers: ##
1. Karim ElGhandour 
2. Mohamed Aboughazala
3. Farid Khaled






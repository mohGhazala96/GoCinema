## Setup: ##
The docker-compose file will initialize everything that you need

## Getting Started: ##
To run the app for the first time or when changing the contents of .go
```
docker-compose up --build -d
```

To run the applcation after that
```
docker-compose up -d
```

To close down docker-compose
```
docker-compose down
```

To read the logs from your application
```
docker-compose logs app
```

Docker run:

##FIRST RUN 

First make sure that you edit the environmental variables in the env folder. Check the example env files.
The final env files that the system is looking for:
```
env/goenv.env
env/pgenv.env
```

Second run db and make sure it's built. Only do this command once
```
docker run -p 5432:5432 --name db --env-file ./env/pgenv.env -v ${PWD}/data.sql:/docker-entrypoint-initdb.d/init.sql -w /docker-entrypoint-initdb.d/init.sql -v ${PWD}/postgres-data:/var/lib/postgresql/data -w /var/lib/postgresql/data healthcheck/postgres:alpine
```

To stop the db run:

```
docker stop db
```

To start the db again:

```
docker start db
```

To run the go app for the first time
```
docker build -t go-app ./go-app
docker run -it --name goapp -p 3000:3000 --env-file ${PWD}/env/goenv.env --link db:postgres go-app
```

To run the app afterwards:
```
docker run goapp
```

To run the frontend for the first time
```
docker build -t angular-app ./frontend
docker run -it --name angularapp -p 4200:4200 -v ${PWD}/frontend:/usr/src/app -w /usr/src/app angular-app
```

To run the frontend afterwards:
```
docker run angularapp
```









# add config.json in the project
follow config-example.json
# to run postgres on docker, if you dont have POSTGRES docker will install it for you
docker run -e POSTGRES_PASSWORD=secret -e  POSTGRES_USER=root -e POSTGRES_DB=mycinema -p 5432:5432 -d postgres:11.1
# To send db data from local file to the database
cat data.sql | docker exec -i (Insert Contanier Id Here) psql -h localhost -U root mycinema
# To add go vendor
*change GoPath to be your current path*
export PATH=$YOURPATH:$(go env GOPATH)/bin
*to get your path*
PWD
*to install govendor*
go get -u github.com/kardianos/govendor
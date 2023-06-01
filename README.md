# sample-go-server


## run the docker container
 
```docker pull cassandra:latest```


```docker run --name cassandra -p 9042:9042 -d cassandra:latest```

## create the keyspace and table

``` docker exec -it my-cassandra-container cqlsh ```

```CREATE KEYSPACE IF NOT EXISTS test WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };```

```USE test;```

create tables for user and projects as defined as structs in the controllers file in the cqlsh like this:

```CREATE TABLE IF NOT EXISTS test.users (id UUID, name TEXT, email TEXT, PRIMARY KEY(id));```

```CREATE TABLE IF NOT EXISTS test.projects (id UUID, name TEXT, description TEXT, PRIMARY KEY(id));```

## install the dependencies

```go mod download```

## run the server

```go run main.go```

## now the server is running on port 8080






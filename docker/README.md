# Docker commands

## Steps to run docer database

### Create & run docker image

To build DB image (you can list image by running: "$ docker images -a")

```bash
docker build -t postgres-db ./
```

You can now run the image by:

```bash
docker run -d --name postgresdb-container -p 5432:5432 postgres-db
```

### Connect to the container

```bash
docker exec -it postgresdb-container bash
```

And to login to the postgres instance

```bash
psql -U postgres
```

Create the database

```bash
CREATE DATABASE board;
```

## TODO:

- Add a sql dump to seed the database

# connect-apms

## Compiling

```
$ make clean && make
```

## Building Docker image

Note that only the owner of the repository is allowed to build the image. 

```
$ docker build -t ghcr.io/mrexmelle/connect-apms:${VERSION} .
$ docker push ghcr.io/mrexmelle/connect-apms:${VERSION}
```

## Running

### For local environment

```
$ docker pull postgres:15-alpine
$ docker run \
	-v $PWD/data:/var/lib/postgresql/data \
	-v $PWD/init-db:/docker-entrypoint-initdb.d \
	-p 8081:8081 \
	-e MONGO_INITDB_ROOT_PASSWORD=123
	mongo:15-alpine
$ ./connect-apms serve
```

### For docker environment

```
$ docker compose up
```
Note that you cannot alter the docker image in the container registry. Only the owner of the repository is allowed to do so.

If error happens in `core` service due to failure to connect to database, restart it:
```
$ docker compose restart core
```
The failure happens due to `db` service isn't ready when `core` attempts to connect to it.

# Stub Authentication

## Acceptance Test

| Name | Username | Password | Status Code | Response Body |
|---|---|---|---|---|
| login passed | imkk-000 | 1mkkn@ja* | 200 | `{"status":"ok"}` |
| login failed with wrong password | imkk-000 | wrongP@ssw0rd | 401 | `{"status":"wrong password"}` |
| login failed with not existing username | fakeusername | noempty | 401 | `{"status":"not existing username"}` |

## API Provider

### Post: /login

#### Request

```text
Request Type: Post Form
Request Body: username=imkk-000&password=1mkkn@ja*
```

#### Response

```text
Response Type: JSON
Response Body: {"status": "ok"}
```

## How to test

```sh
go clean -testcache
go test ./service ./api
CGO_ENABLED=0 go build -o bin/app cmd/login/main.go
./bin/app&
go test ./atdd && kill -9 $$
```

## How to build to docker image

```sh
# build image
docker pull golang:1.12
docker pull alpine:3.10.2
docker build -t stub-login:0.0.1 .

# spawn new container
docker run --name login -d -p 8888:8888 stub-login:0.0.1

# stop container
docker stop login

# remove container
docker rm login
```

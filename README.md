![Tests](https://github.com/TikTokTechImmersion/assignment_demo_2023/actions/workflows/test.yml/badge.svg)
# 2023 TikTok Tech Immersion Backend Assignment

## 0. Introduction
This is a submission of the backend assignment of 2023
TikTok Tech Immersion, based upon the [demo template](https://github.com/TikTokTechImmersion/assignment_demo_2023).
The main changes from the template are:
1. Modifications to `http-server` to support parsing of
   URL params.
2. Initialising MySQL database in`DockerDB` container with
   port `3306` exposed.
3. Addition of `db` field to `IMServiceImpl` struct type
   for communication with aforementioned MySQL database.
4. Modifications to `IMServiceImpl` to read and write
   from `IMServiceImpl.db`.

## 1. Setting up
Run the following command the same directory as
`docker-compose.yml` to initialise the necessary Docker
containers:
```bash
docker compose up -d
```

## 2. Syntax
### 2.1 Sending messages
To send a message, we send a `POST` request to
`localhost:8080/api/send`, specifying the following URL
parameters:
* `chat`: Name of chat in the form `[USER A]:[USER B]`
* `sender`: Name of sender
* `text`: Text in message

Example
To send "hi" from a to b, we use the following command:
```bash
curl -X POST 'localhost:8080/api/send?chat=a:b&sender=a&text=hi'
```

### 2.2 Pulling messages
To send a message, we send a `GET` request to
`localhost:8080/api/pull`, specifying the following URL
parameters:
* `chat`: Name of chat in the form `[USER A]:[USER B]`
* `cursor`: Earliest epoch time of messages to retrieve, 0 by default
* `limit`: Maximum number of messages to retrieve, 10 by default
* `reverse`: Boolean value of whether to sort messages in ascending
   order by time

Example
To retrieve all messages between a and b (i.e., chat `a:b`),
we use the following command:
```bash
curl 'localhost:8080/api/pull?chat=a:b&cursor=0&limit=10&reverse=false'
```
* Although we specified `cursor` as 0 and `limit` as 10,
  here, we could omit them. The request will default to
  `cursor=0` and `limit=10` respectively.

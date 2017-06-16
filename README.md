<p align="center">
<img src="https://media2.giphy.com/media/p3R62d6L0WYw0/200w.gif">
</p>

# go-kickable
**go-kickable** is a package created to answer the age old question... "_Can I Kick It?_"

This package also showcases various microservice technologies and patterns.

**What is a kickable?**
Currently only the word "it" is kickable. 


# API

### Install

```shell
go get -u github.com/defstream/go-kicakble
```

### Usage


```go
package main

import (
    "github.com/defstream/go-kickable"
)

func main() {
    response := kickable.CanIKick("it") 
    fmt.Println(response)
}
```

## `github.com/defstream/go-kickable`

### `Yes string`
Yes is the message returned when an item is kickable.

### `No string`
No is the message returned when an item is not kickable.

### `func Kickable(it string) bool`
Returns true if kickable otherwise false.

### `func CanIKick(it string) string`
Returns a string message indicating if it is kickable or not.

**parameters:**
- **it**            {String}    The value to determine as kickable

**returns**: {String} "Yes, yes you can." if `it` is determined to be kickable, otherwise "No.".

# CLI

## Get Source

```shell
git clone http://github.com/defstream/go-kickable
cd go-kickable
```

## Install

```shell
make install
```

## CLI Usage

```shell
cani -kick "it"
```

## Build Showcase
All showcase examples will be placed in the `bin`  folder. Use the `-help` flag or view the README.md below for more information.

```shell
make all
```

# Showcase

### Service
- 😀 [GraphQL](./service/graphql/README.md)
- 😀 [GRPC](service/grpc/README.md)
    - 😩 HTTP Gateway
- 😀 [Go-Micro](service/micro/README.md)
- 😀 [HTTP](service/http/README.md)
    - 😩 CORS
    - 😩 Keep Alives
    - 😩 Timeouts
    - 😩 Caching
- 😩 HTTP2-Push
- 😀 [KITE](service/kite/README.md)
- 😀 [RPC/GOB](service/rpc/gob/README.md)
- 😀 [RPC/JSON](service/rpc/json/README.md)
- 😩 TCP

### Static Websites
- 😩 S3
- 😩 Static File Server

### Cache
- Strategies
    - 😩 Cache Aside
    - 😩 Read-Through
    - 😩 Write-Through
    - 😩 Write-Behind
- 😩 Redis
- 😩 Memcache
- 😩 Varnish

### Relational Data
- 😩 Amazon RDS
- 😩 Amazon Aurora
- 😩 MYSQL 
- 😩 MSSQL 
- 😩 Postgress
- 😩 Dynamo DB

### Search
- 😩 Solr/Lucene 
- 😩 Elastisearch
- 😩 IBM Simple Search

### Email
- 😩 SMTP
- 😩 AWS SES
- 😩 SendGrid
- 😩 Mail Gun

### Authentication
- 😩 JWT
- 😩 Amazon Cognito
- 😩 Auth0
- 😩 Firebase

### Publish/Subscribe
- 😩 ActiveMQ
- 😩 Event Emitter
- 😩 Firebase Cloud Messaging
- 😩 IBM MQ
- 😩 Kafka
- 😩 Kinesis
- 😩 Nats
- 😩 Rabbit MQ
- 😩 Redis
- 😩 SQS
- 😩 SNS
- 😩 Websockets

### Containers
- 😩 AWS ECS
- 😩 Docker
- 😩 ECS
- 😩 Kubernetes
- 😩 Mesos 

### Cloud 
- 😩 AWS EC2
- 😩 AWS Elastic Beanstalk
- 😩 Google App Engine
- 😩 IBM Bluemix

### Serverless
- 😩 AWS Lambda
- 😩 AWS API Gateway
- 😩 Google Cloud Functions
- 😩 IBM OpenWhisk

### Machine Learning


# Maintainers
Hector Gray (Twitter: <a href="https://twitter.com/defstream">@defstream</a>)

# Contribute
Pull Requests welcome. Please make sure all tests pass 😀

# License
MIT
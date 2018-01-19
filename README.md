# Serverless Todo CRUD in Go

This is an example Todo CRUD application written in Go.

## Prerequisites

- [Node.js & NPM](https://github.com/creationix/nvm)
- [Serverless framework](https://serverless.com/framework/docs/providers/aws/guide/installation/): `npm install -g serverless`
- [Go](https://golang.org/dl/)

## Quick Start

0. Clone the repo

```
git clone git@github.com:yosriady/serverless-crud-go.git
cd serverless-crud-go
```

1. Install Go dependencies

```
go get github.com/aws/aws-lambda-go/lambda
```

2. Compile functions as individual binaries for deployment package:

```
./scripts/build.sh
```

3. Deploy!

```
./scripts/deploy.sh
```

4. Test:

```
> curl -X POST https://<hash>.execute-api.<region>.amazonaws.com/dev/todos
[]
```

## TODOs

- [ ] Use https://github.com/golang/dep to manage packages
# Serverless Todo CRUD in Go

This is an example Todo CRUD application written in Go.

## Prerequisites

- [Node.js & NPM](https://github.com/creationix/nvm)
- [Serverless framework](https://serverless.com/framework/docs/providers/aws/guide/installation/): `npm install -g serverless`
- [Go](https://golang.org/dl/)
- [dep](https://github.com/golang/dep): `brew install dep && brew upgrade dep`

## Quick Start

0. Clone the repo

```
git clone git@github.com:yosriady/serverless-crud-go.git
cd serverless-crud-go
```

1. Install Go dependencies

```
dep ensure
```

2. Compile functions as individual binaries for deployment package:

```
./scripts/build.sh
```

> You need to perform this compilation step before deploying.

3. Deploy!

```
serverless deploy
```

> You can perform steps 2 and 3 simultaneously by running `./scripts/deploy.sh`.

4. Test:

> Use [this Postman collection](https://www.getpostman.com/collections/4cff6cc5d40753d75348) or use cURL:

```
> curl https://<hash>.execute-api.<region>.amazonaws.com/dev/todos
{
    "todos": [
        {
            "id": "d3e38e20-5e73-4e24-9390-2747cf5d19b5",
            "description": "buy fruits",
            "done": false,
            "created_at": "2018-01-23 08:48:21.211887436 +0000 UTC m=+0.045616262"
        },
        {
            "id": "1b580cc9-a5fa-4d29-b122-d20274537707",
            "description": "go for a run",
            "done": false,
            "created_at": "2018-01-23 10:30:25.230758674 +0000 UTC m=+0.050585237"
        }
    ]
}
```

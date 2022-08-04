# MINIPULSA

created and coded by elangreza14

### minipulsa arch

![alt arch](https://github.com/elangreza14/minipulsa/blob/master/minipulsaarch.jpeg)

to repopulate this repo just run

before running this repo make sure port 8080, 9000, 9001, 9002, 9003 is free


| Service        | PORT | METHOD     |
| -------------- | ---- | ---------- |
| Authentication | 9000 | UNARY GRPC |
| Product        | 9001 | UNARY GRPC |
| Wallet         | 9002 | UNARY GRPC |
| Order          | 9003 | UNARY GRPC |
| Api Gateway    | 8080 | HTTP REST  |

### API AND GRPC DOCS

1. for arch in `minipulsaarch.jpg` file
1. for rest api swagger for api gateway in `./docs` folder
1. for GRPC in `./pb` folder

### SCENARIO USING MINIPULSA

1. setup environtment and running in docker compose with

   ```bash
   docker-compose up -d
   ```

2. hit authentication api to login into system and get the jwt token `localhost:8080/authentication/login-register`

   ```curl -X 'POST' \
   'http://localhost:8080/authentication/login-register' \
   -H 'accept: application/json' \
   -H 'Content-Type: application/json' \
   -d '{
   "email": "string@string.com",
   "password": "string"
   }'
   ```

3. view product `localhost:8080/product/list`

   ```curl -X 'GET' \
   'http://localhost:8080/product/list' \
   -H 'accept: application/json'
   ```

4. top up the wallet and make sure insert token into authorization header `localhost:8080/wallet/use`

   ```curl -X 'POST' \
   'http://localhost:8080/wallet/use' \
   -H 'accept: application/json' \
   -H 'Authorization: Bearer {{TOKEN}}' \
   -H 'Content-Type: application/json' \
   -d '{
   "amount": 10000
   }'
   ```

5. you can check the wallet with `http://localhost:8080/wallet/detail`

   ```curl -X 'GET' \
   'http://localhost:8080/wallet/detail' \
   -H 'accept: application/json' \
   -H 'Authorization: Bearer {{TOKEN}}'
   ```

6. finally buy some pulsa, and make sure transaction is success. if wallet is not enough money transaction will be cancel `localhost:8080/order/process`

   ```curl -X 'POST' \
   'http://localhost:8080/order/process' \
   -H 'accept: application/json' \
   -H 'Authorization: Bearer {{TOKEN}}' \
   -H 'Content-Type: application/json' \
   -d '{
   "product_id": 1
   }'
   ```

7. you can also check transaction with this `localhost:8080/order/detail`

   ```curl -X 'GET' \
    'http://localhost:8080/order/detail' \
    -H 'accept: application/json' \
    -H 'Authorization: Bearer {{TOKEN}}'
   ```

### this is tree for the project

```

.
├── docker-compose.yaml 
├── docker_postgres_init.sql
├── docs  // public api 
│   └── docs-minipulsa.yaml
├── Makefile
├── minipulsaarch.jpeg
├── pb // internal protocol buffer
│   └── minipulsa.proto
├── readme.md
└── src // microservices
    ├── api-gateway
    ├── authentication
    ├── order
    ├── product
    └── wallet


8 directories


71 directories, 182 files

```

# MINIPULSA REPO

created and coded by MUHAMMAD REZA ELANG ERLANGGA

to repopulate this repo just run

before running this repo make sure port 8080, 9000, 9001, 9002, 9003 is free

### RUN this microservice

    ```
    sudo docker-compose up -d

    ```

| Service        | PORT | METHOD     |
| -------------- | ---- | ---------- |
| Authentication | 9000 | UNARY GRPC |
| Product        | 9001 | UNARY GRPC |
| Wallet         | 9002 | UNARY GRPC |
| Order          | 9003 | UNARY GRPC |
| Api Gateway    | 8080 | HTTP REST  |

### API AND GRPC DOCS

1. for arch in `minipulsaarch.jpg` file
1. for rest api for api gateway in `./docs` folder
1. for GRPC in `./pb` folder

### SCENARIO USING MINIPULSA

1. setup environtment and running in docker compose with

   ```
   sudo docker-compose up -d

   ```

2. hit authentication api to login into system and get the jwt token `localhost:8080/authentication/login-register`

   ```
   curl -X 'POST' \
   'http://localhost:8080/authentication/login-register' \
   -H 'accept: application/json' \
   -H 'Content-Type: application/json' \
   -d '{
   "email": "string@string.com",
   "password": "string"
   }'
   ```

3. view product `localhost:8080/product/list`

   ```
   curl -X 'GET' \
   'http://localhost:8080/product/list' \
   -H 'accept: application/json'
   ```

4. top up the wallet and make sure insert token into authorization header `localhost:8080/wallet/use`

   ```
   curl -X 'POST' \
   'http://localhost:8080/wallet/use' \
   -H 'accept: application/json' \
   -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjViMzA3YWQ1LTI2OWItNGY3Zi1iNDIwLTMxMTBkNzY1Y2RiOSIsImVtYWlsIjoic3RyaW5nQHN0cmluZy5jb20iLCJpc3N1ZWRfYXQiOiIyMDIyLTA0LTI1VDIwOjM4OjU2LjI2NDg5MTU0NVoiLCJleHBpcmVkX2F0IjoiMjAyMi0wNC0yNlQwMjozODo1Ni4yNjQ4OTE2NTVaIn0.ACWuwCXo1nZSYs2Ib4INQNiHFHTd90HlkhGAxQz16ag' \
   -H 'Content-Type: application/json' \
   -d '{
   "amount": 10000
   }'
   ```

5. you can check the wallet with `http://localhost:8080/wallet/detail`

   ```
   curl -X 'GET' \
   'http://localhost:8080/wallet/detail' \
   -H 'accept: application/json' \
   -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjViMzA3YWQ1LTI2OWItNGY3Zi1iNDIwLTMxMTBkNzY1Y2RiOSIsImVtYWlsIjoic3RyaW5nQHN0cmluZy5jb20iLCJpc3N1ZWRfYXQiOiIyMDIyLTA0LTI1VDIwOjM4OjU2LjI2NDg5MTU0NVoiLCJleHBpcmVkX2F0IjoiMjAyMi0wNC0yNlQwMjozODo1Ni4yNjQ4OTE2NTVaIn0.ACWuwCXo1nZSYs2Ib4INQNiHFHTd90HlkhGAxQz16ag'
   ```

6. finally buy some pulsa, and make sure transaction is success. if wallet is not enough money transaction will be cancel `localhost:8080/order/process`

   ```curl -X 'POST' \
   'http://localhost:8080/order/process' \
   -H 'accept: application/json' \
   -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjViMzA3YWQ1LTI2OWItNGY3Zi1iNDIwLTMxMTBkNzY1Y2RiOSIsImVtYWlsIjoic3RyaW5nQHN0cmluZy5jb20iLCJpc3N1ZWRfYXQiOiIyMDIyLTA0LTI1VDIwOjM4OjU2LjI2NDg5MTU0NVoiLCJleHBpcmVkX2F0IjoiMjAyMi0wNC0yNlQwMjozODo1Ni4yNjQ4OTE2NTVaIn0.ACWuwCXo1nZSYs2Ib4INQNiHFHTd90HlkhGAxQz16ag' \
   -H 'Content-Type: application/json' \
   -d '{
   "product_id": 1
   }'
   ```

7. you can also check transaction with this `localhost:8080/order/detail`

   ```curl -X 'GET' \
    'http://localhost:8080/order/detail' \
    -H 'accept: application/json' \
    -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjViMzA3YWQ1LTI2OWItNGY3Zi1iNDIwLTMxMTBkNzY1Y2RiOSIsImVtYWlsIjoic3RyaW5nQHN0cmluZy5jb20iLCJpc3N1ZWRfYXQiOiIyMDIyLTA0LTI1VDIwOjM4OjU2LjI2NDg5MTU0NVoiLCJleHBpcmVkX2F0IjoiMjAyMi0wNC0yNlQwMjozODo1Ni4yNjQ4OTE2NTVaIn0.ACWuwCXo1nZSYs2Ib4INQNiHFHTd90HlkhGAxQz16ag'
   ```

### this is tree for the project

```

.
├── docker-compose.yaml
├── docker_postgres_init.sql
├── docs
│ └── docs-minipulsa.yaml
├── Makefile
├── minipulsaarch.jpeg
├── pb
│ └── minipulsa.proto
├── readme.md
└── src
├── api-gateway
├── authentication
├── order
├── product
└── wallet

71 directories, 182 files

```

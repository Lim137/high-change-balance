LAUNCHING
---
1. Add .env file to the root folder of the project (check ".env example" topic)
2. Run the command below in the root folder of the project
```
go run cmd/balance-tracker/main.go
```
.env example
---
```
API_KEY=11111111111111111111111111111111
PORT=8080
ENV=prod
```
The ENV parameter is optional, it can be local, dev or prod (default value - prod).<br>
The PORT parameter is optional (default value - 8080)

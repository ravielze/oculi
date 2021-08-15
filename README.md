# Oculi
![](https://shields.io/badge/go-v1.16-blue?logo=go)
![](https://img.shields.io/github/issues/ravielze/oculi)
![](https://img.shields.io/github/forks/ravielze/oculi)
![](https://img.shields.io/github/stars/ravielze/oculi)

Backend utilities

# Tech Stacks

1. Echo
2. Gorm
3. Postgresql / Mariadb / Mysql
4. Zap
5. Swagger
6. Uber dig
7. Bcrypt
8. Mockgen
9. JWT
10. Minio for S3 File Storage
11. Excelize
12. QRCode
13. etc.


# Useful Command & Tips
## Mocking

1. Install mockgen
2. Run command
`mockgen --source=input.go -package=packageName --destination=mocks/output.go`


## Regenerate Swagger JSON from API Blueprint

1. Install api2bswagger using `npm install -g apib2swagger`
2. Run `apib2swagger -i ./api/blueprint.apib -o ./resources/external/docs.json`

# Example Project

A simple backend using this tools: [LINK](https://github.com/ravielze/oculi/tree/master/example)

# Todo

1. Test redis cache, implement redis pubsub
2. Add websocket tools
3. Add more custom validator
4. Research rabbitmq, kafka, elastic search, mongodb
5. Add encoding json library
6. Add webclient library

# Contributors

<a href="https://github.com/ravielze/oculi/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=ravielze/oculi" />
</a>
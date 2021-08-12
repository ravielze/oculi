# Oculi
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
11. etc.


# Mocks Command

## Prerequisite

- Install mockgen

- Run command
`mockgen --source=input.go -package=packageName --destination=mocks/output.go`

# Example Project

A simple backend using this tools: [LINK](https://github.com/ravielze/oculi/tree/master/example)

# Todo

1. Test redis cache, implement redis pubsub
2. Add websocket tools
3. Add more custom validator
4. Research rabbitmq, kafka, elastic search, mongodb
5. Add encoding json library
6. Add webclient library
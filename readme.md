# Mini-wallet

How to run for dev?

- Make sure you have working postgresql and setup the table using ```migrate.sql``` included 

- Then setting up the config in ```/config/config.dev.json```

- run this command 
```go
go run app/main.go --env=dev
```
then the application will be started at given ```servicePort```
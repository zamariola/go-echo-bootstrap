# go-echo-bootstrap
The simplest Go project to start coding with echo framework and connect to databases and messaging systems

## Build: 
```shell
#make all [OUTPUT=folder_name]
``` 
will compile to default **build** folder 

## Run:
```shell
#./build/server [--port=9011]
```

Then your server will be visible on **http://localhost:9011**

## Try it:

```shell
#curl http://localhost:9011/status
```
or

[http://localhost:9011/status](http://localhost:9011/status)

## Simple Create / Read operation

To be able to persist information you need to set up database configuration:

```shell
#./build/server --port=9011 \
--db_host=localhost \
--db_port=5431 \
--db_user=mydbuser \
--db_pass=mydbpass \
--db_name=mydb
```

Then you can create:

```shell
#curl localhost:9011/users -H'Content-Type: application/json' -d'{"name":"Name Surname","email":"my@email.com"}'
```

And read:

```shell
#curl localhost:9011/users/Name%20Surname
```

# IMPORTANT:

For the sake of simplicity I've organized all classes inside one package (internal), but this isn't a good practice. Also, I haven't included some common features as database migration, package structure, clean architecture and etc.

If you like to add it, I recommmend checking these links:

- **Db Migrations**: https://github.com/golang-migrate/migrate
- **Package structure**: https://blog.golang.org/package-names
- **Amazon SNQ/SQS**: https://docs.aws.amazon.com/sdk-for-go/api/service/sns/

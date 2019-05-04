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

# IMPORTANT:

For the sake of simplicity I've organized all classes inside one package (internal), but this isn't a good practice. Also, I haven't included some common features as database migration, package structure, clean architecture and etc.

If you like to add it, I recommmend checking these links:

- **Db Migrations**: https://github.com/golang-migrate/migrate
- **Package structure**: https://blog.golang.org/package-names
- **Amazon SNQ/SQS**: https://docs.aws.amazon.com/sdk-for-go/api/service/sns/

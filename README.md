## cssd-admin

### How to run

______

#### Required

* Go
* Mysql

#### Conf

You should modify conf/app.ini

```ini
[database]
Type = mysql
User = root
Password = 
Host = 127.0.0.1:3306
Name = cssd
TablePrefix = cssd_
```

#### Run

```go
cd $GOPATH/src/cssd-admin
go run main.go

// Or if you want reload it hotly just make sure that air you have installed and in you PATH
// At you root directory:
air
```

#### Features

______

* RESTful API
* Gorm
* logging
* Gin
* App configurable
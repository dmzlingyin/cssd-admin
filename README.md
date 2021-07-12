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

```
cd $GOPATH/src/cssd-admin
go run main.go
```

#### Features

______

* RESTful API
* Gorm
* logging
* Gin
* App configurable
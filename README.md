# TheItem

This Application developed using gogen framework code generator
[Gogen Framework](https://github.com/mirzaakhena/gogen)

This application has one entity (`domain_item/model/entity/item.go`)
```go
type Item struct {
    ID              vo.ItemID     `json:"id" bson:"_id"`
    Created         time.Time     `json:"created"`
    Updated         time.Time     `json:"updated"`
    Name            string        `json:"name"`
    Rating          vo.Rating     `json:"rating" `
    Category        vo.Category   `json:"category"`
    ImageURL        vo.StringURL  `json:"image"`
    Reputation      vo.Reputation `json:"reputation"`
    ReputationBadge string        `json:"reputation_badge"`
    Price           int           `json:"price"`
    Availability    int           `json:"availability"`
}
```

Has a six use cases (`domain_item/usecase/`)
```text
1. getallitem      --> Get All Item with filter
2. getoneitem      --> Get Only One Item by ID
3. runitemcreate   --> Create an Item
4. runitemdelete   --> Delete an Item by ID
5. runitempurchase --> Reduce the Availability of Item
6. runitemupdate   --> Update an Item
```

Each use case, published via REST API using [gin-gonic](https://github.com/gin-gonic/gin) (`domain_item/controller/restapi/router.go`)
```text
runitemcreate   --> POST   /api/v1/items             
getallitem      --> GET    /api/v1/items             
getoneitem      --> GET    /api/v1/items/:item_id    
runitemupdate   --> PUT    /api/v1/items/:item_id    
runitemdelete   --> DELETE /api/v1/items/:item_id    
runitempurchase --> POST   /api/v1/items/:item_id/purchase 
```

You can decide to run this application with 3 alternative database 

1. **SQLite** using Gorm (`domain_item/gateway/withmysqldb`)
2. **MySQL** using Gorm (`domain_item/gateway/withsqlitedb`)
3. Native **MongoDB** (`domain_item/gateway/withmongodb`)

By default, it is running with **SQLite** db

## Run backend directly from code using SQLite

After git clone the code, open a terminal (we call it a first terminal), 
then download the dependency by call this command

```shell
$ go mod tidy
```
Run it by
```shell
$ go run main.go appitem
```

then you will see the application is running
```shell
➜  theitem go run main.go appitem
Version 0.0.1
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /ping                           --> theitem/domain_item/controller/restapi.NewController.func1 (3 handlers)
[GIN-debug] GET    /web/*filepath                  --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (3 handlers)
[GIN-debug] HEAD   /web/*filepath                  --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (3 handlers)
[GIN-debug] POST   /api/v1/items                   --> theitem/domain_item/controller/restapi.(*controller).runItemCreateHandler.func1 (6 handlers)
[GIN-debug] GET    /api/v1/items                   --> theitem/domain_item/controller/restapi.(*controller).getAllItemHandler.func1 (6 handlers)
[GIN-debug] GET    /api/v1/items/:item_id          --> theitem/domain_item/controller/restapi.(*controller).getOneItemHandler.func1 (6 handlers)
[GIN-debug] PUT    /api/v1/items/:item_id          --> theitem/domain_item/controller/restapi.(*controller).runItemUpdateHandler.func1 (6 handlers)
[GIN-debug] DELETE /api/v1/items/:item_id          --> theitem/domain_item/controller/restapi.(*controller).runItemDeleteHandler.func1 (6 handlers)
[GIN-debug] POST   /api/v1/items/:item_id/purchase --> theitem/domain_item/controller/restapi.(*controller).runItemPurchaseHandler.func1 (6 handlers)
INFO  0000000000000000 server is running at :8080      restapi.(*gracefullyShutdown).Start:40
```

The API is running on port 8080. The port setting is in `config.json`. 
You can use postman or curl for accessing the API. But it is better to use the UI. Keep reading, because we also provide the UI.

When running with **SQLite** db, we only use the `db_name` field in `config.json` and ignore the other database fields.
```json
{
  "database": {
    "username": "root",
    "password": "12345",
    "host": "localhost",
    "port": "27017",
    "db_name": "itemdb"
  },
  "servers": {
    "appItem": {
      "address": ":8080"
    }
  }
}
```

## Run frontend from code independently (development mode)
This application also has simple user interface (UI) for better experience, interaction or just testing purpose. 
The UI use all the capability of REST API.
The frontend application using vue js under `web/` directory. 
To follow the further instruction of using UI, make sure you already install nodejs in your system.

In order to run the UI, open new second terminal, change directory

```shell 
$ cd web
```

then install the dependency by running this command 
```shell
$ npm install
```

While the backend apps is still running, run this command

```shell
$ npm run dev
```

Then you will see this output
```shell
➜  web npm run dev

> simulator@0.0.0 dev
> vite


  VITE v3.2.5  ready in 222 ms

  ➜  Local:   http://localhost:5173/web/
  ➜  Network: use --host to expose
```

Open your browser then access `http://localhost:5173/web/`

## Run frontend via backend
You can run the frontend without running it separately. 
In this case the backend will support the frontend as a webserver. 
All you need to do is build the web package distribution first.
Stop the frontend application in second terminal (if frontend is still running), 
then run this command

```shell
$ npm run build
```
The command will create a bundled web application in folder `web/dist/`.

Back to first terminal, stop the backend by `ctrl+c`, then re-run it again
```shell
$ go run main.go appitem
```

Now, open your browser then access `http://localhost:8080/web/`

## Run backend directly from code using MySQL 

Open file `application/app_appitem.go` then change the code, from this

```go
datasource := withsqlitedb.NewGateway(log, appData, cfg)
//datasource := withmysqldb.NewGateway(log, appData, cfg)
//datasource := withmongodb.NewGateway(log, appData, cfg)

```

Into this

```go
//datasource := withsqlitedb.NewGateway(log, appData, cfg)
datasource := withmysqldb.NewGateway(log, appData, cfg)
//datasource := withmongodb.NewGateway(log, appData, cfg)
```

Open `config.json`
```json
{
  "database": {
    "username": "root",
    "password": "12345",
    "host": "localhost",
    "port": "3306",
    "db_name": "itemdb"
  },
  "servers": {
    "appItem": {
      "address": ":8080"
    }
  }
}
```
adjust the config as necessary (for example : username and password)

## Run backend directly from code using MongoDB

Actually the process is same as using MySQL, we only switch the code in `application/app_appitem.go` into this
```go
//datasource := withsqlitedb.NewGateway(log, appData, cfg)
//datasource := withmysqldb.NewGateway(log, appData, cfg)
datasource := withmongodb.NewGateway(log, appData, cfg)
```
And then adjust the `config.json` as necessary (please keep in mind by default, mongodb use port 27017)

## Run with Docker

Before running with docker, first decide whether you want to run it by **SQLite**, **MySQL** or **MongoDB** 
by switching the implementation in `application/app_appitem.go`.

By default (NOT using docker), this application use `config.json`. 
In docker version, it uses different config which is in `config.prod.json`.
You can change the setting via `docker-compose.yml`.

Currently, file `docker-compose.yml` specify 2 database image (mongodb and mysql) and 1 application image (myapp).

### using Docker and SQLite
Since **SQLite** is a simple embedded database, we don't need to use any image. 
By default, it will just run. 

### using Docker and MySQL
You need to enable this part on `docker-compose.yml`
```text
mysqldb:
    image: mysql
    restart: always
    environment:
        - MYSQL_ROOT_PASSWORD=12345
        - MYSQL_DATABASE=itemdb
    ports:
        - "3306:3306"
```
Adjust the `config.prod.json` as necessary

### using Docker and MongoDB
You need to enable this part on `docker-compose.yml`
```text
mongodb:
    image: mongo
    ports:
        - "27017:27017"
    environment:
        - MONGO_INITDB_ROOT_USERNAME=root
        - MONGO_INITDB_ROOT_PASSWORD=12345
```
Adjust the `config.prod.json` as necessary

Run the docker compose (you may add `-d` for running it in background) 

```shell
$ docker compose up
```

Open browser then access
```shell
http://localhost:8080/web/
```


## Sample Payload

#### Create new item
```text
POST   /api/v1/items

REQUEST
{
    "name": "the item name",
    "rating": 2,
    "category":  "cartoon",
    "image": "http://image.aa",
    "reputation":  34,
    "price": 5000,
    "availability":  10
}

RESPONSE OK
{
  "success": true,
  "errorCode": "",
  "errorMessage": "",
  "data": {},
  "traceId": "KR9HW32UT28N0VKW"
}

RESPONSE FAIL
{
  "success": false,
  "errorCode": "ER0006",
  "errorMessage": "name length must greater than 10",
  "data": null,
  "traceId": "HO6ONN4SICA1UHTY"
}

RESPONSE FAIL
{
  "success": false,
  "errorCode": "ER0004",
  "errorMessage": "invalid rating value. must be integer between 0..5",
  "data": null,
  "traceId": "NT325V3DN8MLDC0A"
}

RESPONSE FAIL
{
  "success": false,
  "errorCode": "ER0003",
  "errorMessage": "invalid category. must be one of [photo sketch cartoon animation]",
  "data": null,
  "traceId": "NJXGNC72X60GLUBW"
}

RESPONSE FAIL
{
  "success": false,
  "errorCode": "ER0002",
  "errorMessage": "invalid url for 'image'",
  "data": null,
  "traceId": "91T18FT3SFFFLNVA"
}

RESPONSE FAIL
{
  "success": false,
  "errorCode": "ER0001",
  "errorMessage": "out of range reputation. must between 0 to 1000",
  "data": null,
  "traceId": "DS4YVNSPGCGHPRMF"
}

RESPONSE FAIL
{
  "success": false,
  "errorCode": "ER0011",
  "errorMessage": "price must greater or equal zero",
  "data": null,
  "traceId": "FXA7LLI93HOX9HFF"
}


```

#### Display all the item
```text
GET    /api/v1/items
    page=1&
    size=2&
    rating=3&
    reputation_badge=yellow&
    availability_more=0&
    availability_less=100&
    category=photo

RESPONSE OK
{
  "success": true,
  "errorCode": "",
  "errorMessage": "",
  "data": {
    "count": 2,
    "items": [
      {
        "id": "0caf9621-aab4-4fc8-a133-47fc98ec36cf",
        "created": "2023-02-12T09:21:46.947388+07:00",
        "updated": "2023-02-12T09:21:46.947388+07:00",
        "name": "the first item",
        "rating": 2,
        "category": "animation",
        "image": "http://image.aa",
        "reputation": 34,
        "reputation_badge": "red",
        "price": 5000,
        "availability": 10
      },
      {
        "id": "de5aafdc-3361-4e69-83ea-5529b21f255e",
        "created": "2023-02-12T09:22:25.311257+07:00",
        "updated": "2023-02-12T09:22:25.311257+07:00",
        "name": "the second item",
        "rating": 2,
        "category": "cartoon",
        "image": "http://image.aa",
        "reputation": 34,
        "reputation_badge": "red",
        "price": 5000,
        "availability": 10
      }
    ]
  },
  "traceId": "K2VRLG7OC6AP5IGN"
}
```

#### Get one item by id
```text
GET    /api/v1/items/0caf9621-aab4-4fc8-a133-47fc98ec36cf

RESPONSE
{
  "success": true,
  "errorCode": "",
  "errorMessage": "",
  "data": {
    "item": {
      "id": "0caf9621-aab4-4fc8-a133-47fc98ec36cf",
      "created": "2023-02-12T09:21:46.947388+07:00",
      "updated": "2023-02-12T09:21:46.947388+07:00",
      "name": "the first item",
      "rating": 2,
      "category": "animation",
      "image": "http://image.aa",
      "reputation": 34,
      "reputation_badge": "red",
      "price": 5000,
      "availability": 10
    }
  },
  "traceId": "1W9DAWJPNSFCWY38"
}

RESPONSE FAIL
{
  "success": false,
  "errorCode": "ER0007",
  "errorMessage": "unavailable item with id 'abcd9621-aab4-4fc8-a133-47fc98ec61de'",
  "data": null,
  "traceId": "EH863OJWJUGB2ERC"
}

```

#### Update the item
```text
PUT    /api/v1/items/0caf9621-aab4-4fc8-a133-47fc98ec36cf

REQUEST
{
  "name": "the changes name",
  "category":"sketch",
  "image": "http://whatever.com",
  "price": 15000
}

RESPONSE
{
  "success": true,
  "errorCode": "",
  "errorMessage": "",
  "data": {},
  "traceId": "TB9HW79UT28I2V6P"
}

RESPONSE FAIL
{
  "success": false,
  "errorCode": "ER0006",
  "errorMessage": "name length must greater than 10",
  "data": null,
  "traceId": "HO6ONN4SICA1UHTY"
}

RESPONSE FAIL
{
  "success": false,
  "errorCode": "ER0003",
  "errorMessage": "invalid category. must be one of [photo sketch cartoon animation]",
  "data": null,
  "traceId": "NJXGNC72X60GLUBW"
}

RESPONSE FAIL
{
  "success": false,
  "errorCode": "ER0002",
  "errorMessage": "invalid url for 'image'",
  "data": null,
  "traceId": "91T18FT3SFFFLNVA"
}

RESPONSE FAIL
{
  "success": false,
  "errorCode": "ER0011",
  "errorMessage": "price must greater or equal zero",
  "data": null,
  "traceId": "FXA7LLI93HOX9HFF"
}

```

#### Delete the item
```text
DELETE /api/v1/items/0caf9621-aab4-4fc8-a133-47fc98ec36cf

RESPONSE OK
{
  "success": true,
  "errorCode": "",
  "errorMessage": "",
  "data": {},
  "traceId": "AHPHIPWUNX147SPN"
}

RESPONSE FAIL
{
  "success": false,
  "errorCode": "ER0007",
  "errorMessage": "unavailable item with id 'abcd9621-aab4-4fc8-a133-47fc98ec61de'",
  "data": null,
  "traceId": "EH863OJWJUGB2ERC"
}
```

#### Purchase the Item
```text
POST   /api/v1/items/de5aafdc-3361-4e69-83ea-5529b21f255e/purchase

REQUEST
{
  "quantity": 2
} 

RESPONSE OK
{
  "success": true,
  "errorCode": "",
  "errorMessage": "",
  "data": {},
  "traceId": "V1Z2A4IW93LH79CZ"
}

RESPONSE FAIL
{
  "success": false,
  "errorCode": "ER0008",
  "errorMessage": "unavailable item stock. requested 20 but availability is 10",
  "data": null,
  "traceId": "Q50VJPFV91UCL9Z1"
}
```

## Error Codes

All error codes listed in `domain_item/model/errorenum/error_codes.go`

```text
UnknownError                ER0000 unknown error
OutOfRangeReputation        ER0001 out of range reputation. must between 0 to 1000
InvalidURL                  ER0002 invalid url for '%s'
InvalidCategory             ER0003 invalid category. must be one of %v
InvalidRatingValue          ER0004 invalid rating value. must be integer between 0..5
ForbiddenWord               ER0005 word '%s' is not allowed
NameLengthMustGreaterThan   ER0006 name length must greater than %d
UnavailableItem             ER0007 unavailable item with id '%s'
UnavailableItemStock        ER0008 unavailable item stock. requested %d but availability is %d
ItemNameAlreadyExist        ER0009 item with name '%s' already exist
InvalidReputationBadge      ER0010 invalid reputation badge
PriceMustGreaterOrEqualZero ER0011 price must greater or equal zero
```
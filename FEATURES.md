## Clean Architecture applied by Gogen Framework 
![Gogen Framework Architecture](https://media.licdn.com/dms/image/C5622AQH21VrTSV3EKw/feedshare-shrink_2048_1536/0/1671333904862?e=1678924800&v=beta&t=EEnUpIwS7wWn3ShGk1TvCWAKX-5BskEcszUNe8lX81s)

## Dependency Inversion Principle

The Dependency Inversion Principle is a principle of object-oriented design which states that high-level modules should not depend on low-level modules, 
but both should depend on abstractions. This helps to decouple the implementation details of a system from its higher-level requirements, 
making it easier to maintain and extend the system over time.

A design pattern that provides a way to implement the Dependency Inversion Principle called as Dependency Injection. 
It involves injecting the required dependencies into a class, rather than having the class create or manage its own dependencies. 
The basic idea behind Dependency Injection is to separate the concerns of an object into two parts: the object's behavior and the dependencies it requires to perform that behavior.

In Dependency Injection, the dependencies are provided to the object from the outside, rather than being created or managed internally. 
This allows for greater flexibility and maintainability, because individual components can be substituted or modified without affecting the rest of the system.

There are several ways to implement Dependency Injection, including constructor injection, setter injection, and interface injection. 
The specific method used depends on the requirements of the system, but the overall goal is to provide a way to manage dependencies in a flexible and maintainable way.

In gogen framework, you can find the dependency injection in [`application/app_appitem.go`](https://github.com/mirzaakhena/theitem/blob/main/application/app_appitem.go) as a constructor injection
where we inject the `datasource` object which instantiated from one one various `NewGateway` method into each `NewUsecase` method (as a constructor).
```go
datasource := withsqlitedb.NewGateway(log, appData, cfg)
...
getallitem.NewUsecase(datasource)
```
In the use case side, `NewUsecase` method assign the reference value to the `Outport` interface which own by `getAllItemInteractor` struct

```go
type getAllItemInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &getAllItemInteractor{
		outport: outputPort,
	}
}
```

Use case `getAllItemInteractor` indirectly depend on `gateway`.
Instead of creating an instance of the `gateway` struct inside the `getAllItemInteractor` struct, 
the required `gateway` instance is passed into the constructor of the `getAllItemInteractor` struct as a parameter. 
This is an example of constructor injection, where the dependencies are provided to the object via its constructor. 
It allows us to use any other implementation of `gateway`.

## Non Anemic Domain Model

Non Anemic Domain Model is a design pattern in object-oriented programming where the domain objects contain both data and behavior. 
It is considered to be a more robust and maintainable design, because it follows the Single Responsibility Principle, 
which states that a class should have only one reason to change. 
When the behavior of a domain object is implemented inside the object itself, 
it becomes easier to understand how the object behaves, and how it can be used to solve specific problems.

You can find it in many place, mostly in `model/entity` or `model/vo` like

- [`domain_item/model/entity/item.go`](https://github.com/mirzaakhena/theitem/blob/main/domain_item/model/entity/item.go)
- [`domain_item/model/vo/item.go`](https://github.com/mirzaakhena/theitem/blob/main/domain_item/model/entity/item.go)
- [`domain_item/model/vo/category.go`](https://github.com/mirzaakhena/theitem/blob/main/domain_item/model/vo/category.go)
- [`domain_item/model/vo/item_id.go`](https://github.com/mirzaakhena/theitem/blob/main/domain_item/model/vo/item_id.go)
- [`domain_item/model/vo/rating.go`](https://github.com/mirzaakhena/theitem/blob/main/domain_item/model/vo/rating.go)
- [`domain_item/model/vo/reputation.go`](https://github.com/mirzaakhena/theitem/blob/main/domain_item/model/vo/reputation.go)
- [`domain_item/model/vo/string_url.go`](https://github.com/mirzaakhena/theitem/blob/main/domain_item/model/vo/string_url.go)

Let see one sample in `domain_item/model/entity/item.go`

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

func (r *Item) Purchase(quantity int) error {
	if r.Availability < quantity {
		return errorenum.UnavailableItemStock.Var(quantity, r.Availability)
	}
	r.Availability = r.Availability - quantity
	return nil
}
```

In above code, `Item` not only describe the fields required but also can have a method. In this case it has a method `Purchase`

Another interesting example is in value object `domain_item/model/vo/reputation.go`

```go

type Reputation int

func (r Reputation) Validate() error {
    if r < 0 || r > 1000 {
        return errorenum.OutOfRangeReputation
    }
    return nil
}

func (r Reputation) Int() int {
    return int(r)
}

func (r Reputation) Badge() string {
    if r <= 500 {
        return "red"
    }
    if r <= 799 {
        return "yellow"
    }
    return "green"
}

```

In that example, `Reputation` not only as a new int type but it also has a several method like Validate(), Int() and Badge()


## Interface Segregation Principle

The Interface Segregation Principle is one of the SOLID principles of object-oriented programming, which state a set of design guidelines to create maintainable and scalable software systems. The Interface Segregation Principle states that:

"Clients should not be forced to depend on interfaces they do not use."

In other words, this principle suggests that interfaces should be fine-grained and client-specific, rather than being broad and all-encompassing. When an interface defines too many methods, clients that only use a small subset of the methods are still required to depend on the entire interface, resulting in unnecessary coupling. By creating multiple, smaller interfaces that are tailored to specific sets of clients, the design can be made more flexible and maintainable.

The Interface Segregation Principle helps to promote the separation of concerns and loose coupling within a software system, making it easier to modify and extend individual components without affecting others.

In gogen framework, you find this principle in every `Outport` interface which compose other interface from `domain_item/model/repository` or `domain_item/model/service`. 
It works like extending an interface from another interface.
`Outport` only compose all method that will be used by Interactor. 

So instead of creating an interface like this :

```go
type ItemRepo interface {
	SaveItem(ctx context.Context, obj *entity.Item) error 
	FindAllItem(ctx context.Context, page, size int, query ItemQueryFilter) ([]*entity.Item, int64, error) 
	FindOneItem(ctx context.Context, itemID vo.ItemID) (*entity.Item, error) 
	FindOneItemByName(ctx context.Context, name string) (*entity.Item, error) 
	HardDeleteOneItem(ctx context.Context, item *entity.Item) error
}
```

We split into "one interface one method" way, like this:
```go
type SaveItemRepo interface {
	SaveItem(ctx context.Context, obj *entity.Item) error
}

type ExistInForbiddenNameListRepo interface {
	ExistInForbiddenNameList(ctx context.Context, name string) bool
}

type FindAllItemRepo interface {
	FindAllItem(ctx context.Context, page, size int, query ItemQueryFilter) ([]*entity.Item, int64, error)
}

type FindOneItemRepo interface {
	FindOneItem(ctx context.Context, itemID vo.ItemID) (*entity.Item, error)
}

type FindOneItemByNameRepo interface {
	FindOneItemByName(ctx context.Context, name string) (*entity.Item, error)
}

type DeleteOneItemRepo interface {
	FindOneItemRepo
	HardDeleteOneItem(ctx context.Context, item *entity.Item) error
}
```

Then `Outport` only compose some method. For example in 

```go
// domain_item/usecase/getoneitem

type Outport interface {
	repository.FindOneItemRepo
}
```

```go
// domain_item/usecase/runitemcreate

type Outport interface {
	repository.SaveItemRepo
	repository.ExistInForbiddenNameListRepo
	repository.FindOneItemByNameRepo
}
```

```go
// domain_item/usecase/runitempurchase

type Outport interface {
	repository.SaveItemRepo
	repository.FindOneItemRepo
}
```

```go
// domain_item/usecase/runitemupdate

type Outport interface {
	repository.SaveItemRepo
	repository.ExistInForbiddenNameListRepo
	repository.FindOneItemRepo
	repository.FindOneItemByNameRepo
}
```

by this approach, it make the testing easier

## Clean Architecture

Gogen applied clean architecture. From layer dependency perspective we can see that 
- `use case` layer depend on `model` 
- `gateway` layer depend on `use case` 
- `controller` layer depend on `use case`
- `application` depend on `usecase`, `gateway` and `controller`

Dependency, can be seen by an `import` statement on each file. Or we can prove it by trying to delete some layer. For the example :
1. if you delete the layer `controller` or `gateway`, the only part who will get an error is `application`, the other layer will not get affected
2. if you delete the layer `usecase` then layer `application`, `controller` and `gateway` will get an error.
3. if you delete the `model`, all layer will get an error

We can see the integration from `application/app_appitem.go`

```go
func (appItem) Run() error {

	const appName = "appItem"

	cfg := config.ReadConfig()

	appData := gogen.NewApplicationData(appName)

	log := logger.NewSimpleJSONLogger(appData)

	datasource := withsqlitedb.NewGateway(log, appData, cfg)
	//datasource := withmysqldb.NewGateway(log, appData, cfg)
	//datasource := withmongodb.NewGateway(log, appData, cfg)

	primaryDriver := restapi.NewController(appData, log, cfg)
    //primaryDriver := restapi2.NewController(appData, log, cfg)

	primaryDriver.AddUsecase(
		//
		getallitem.NewUsecase(datasource),
		getoneitem.NewUsecase(datasource),
		runitemcreate.NewUsecase(datasource),
		runitemdelete.NewUsecase(datasource),
		runitempurchase.NewUsecase(datasource),
		runitemupdate.NewUsecase(datasource),
	)

	primaryDriver.RegisterRouter()

	primaryDriver.Start()

	return nil
}
```
In this project demonstration, we have 
- 3 alternative gateway (sqlitem, mysql, mongodb)
- 2 alternative controller (gin, echo)









package tenantapigowrapper

const (
	entityOrder = "order"
)

type APISingleEntity interface {
	getEntitySingleName() string
}

type APIListEntity interface {
	getEntitySingleName() string
	getEntityListName() string
}

type Order struct {
	Ref      int      `json:"ref"`
	Customer Customer `json:"customer"`
}

func (order Order) getEntitySingleName() string {
	return entityOrder
}

type Orders []Order

func (orders Orders) getEntitySingleName() string {
	return entityOrder
}

func (orders Orders) getEntityListName() string {
	return entityOrder
}

type Customer struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

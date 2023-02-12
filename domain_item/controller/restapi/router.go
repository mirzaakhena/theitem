package restapi

func (r *controller) RegisterRouter() {
	resource := r.Router.Group("/api/v1", r.authentication())
	resource.POST("/items", r.authorization(), r.runItemCreateHandler())
	resource.GET("/items", r.authorization(), r.getAllItemHandler())
	resource.GET("/items/:item_id", r.authorization(), r.getOneItemHandler())
	resource.PUT("/items/:item_id", r.authorization(), r.runItemUpdateHandler())
	resource.DELETE("/items/:item_id", r.authorization(), r.runItemDeleteHandler())
	resource.POST("/items/:item_id/purchase", r.authorization(), r.runItemPurchaseHandler())
}

package restapi2

func (r *controller) RegisterRouter() {
	resource := r.Router.Group("/api/v1")
	resource.POST("/items", r.runItemCreateHandler())
	resource.GET("/items", r.getAllItemHandler())
	resource.GET("/items/:item_id", r.getOneItemHandler())
	resource.POST("/items/:item_id", r.runItemUpdateHandler())
	resource.POST("/items/:item_id", r.runItemDeleteHandler())
	resource.POST("/items/:item_id/purchase", r.runItemPurchaseHandler())
	//!
}

package restapi2

func (r *controller) RegisterRouter() {
	resource := r.Router.Group("/api/v1")
	resource.GET("/getallitem", r.getAllItemHandler())
	resource.GET("/getoneitem", r.getOneItemHandler())
	resource.POST("/runitemcreate", r.runItemCreateHandler())
	resource.POST("/runitemdelete", r.runItemDeleteHandler())
	resource.POST("/runitempurchase", r.runItemPurchaseHandler())
	resource.POST("/runitemupdate", r.runItemUpdateHandler())
	//!

}

package restapi2

func (r *controller) RegisterRouter() {
	resource := r.Router.Group("/api/v1", r.authentication())
	resource.GET("/getallitem", r.authorization(), r.getAllItemHandler())
	resource.GET("/getoneitem", r.authorization(), r.getOneItemHandler())
	resource.POST("/runitemcreate", r.authorization(), r.runItemCreateHandler())
	resource.POST("/runitemdelete", r.authorization(), r.runItemDeleteHandler())
	resource.POST("/runitempurchase", r.authorization(), r.runItemPurchaseHandler())
	resource.POST("/runitemupdate", r.authorization(), r.runItemUpdateHandler())
	//!

}

package web

type CollectionUpdateRequest struct {
	Id   		int    `validate:"required"`
	Name 		string `validate:"required,max=200,min=1" json:"name"`
	Point 	int		 `validate:"required,max=200,min=1" number:"point"`
	Status 	bool 	 `validate:"required" boolean:"status"`
}

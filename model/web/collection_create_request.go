package web

type CollectionCreateRequest struct {
	Name 		string 		`validate:"required,min=1,max=100" json:"name"`
	Point 	int     	`validate:"required" json:"point"`
	Status 	bool      `validate:"required" json:"status"`
}

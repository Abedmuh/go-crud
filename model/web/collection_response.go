package web

type CollectionResponse struct {
	Id   		int    `json:"id"`
	Name 		string `json:"name"`
	Point 	int		 `json:"point"`
	Status 	bool 	 `json:"status"`
}

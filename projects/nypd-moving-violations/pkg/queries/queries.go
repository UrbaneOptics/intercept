package queries

type TalliesRequest struct {
	PrecinctIDs []int `json:"precinct_ids"`
	Page        int   `json:"page"`
	PerPage     int   `json:"per_page"`
}

// NewTalliesRequest create new default instance of a TalliesRequest
func NewTalliesRequest() TalliesRequest {
	req := TalliesRequest{}
	req.Page = 1
	req.PerPage = 1000
	return req
}

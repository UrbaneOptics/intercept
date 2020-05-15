package queries

// TalliesRequest is a struct used to handle querying for tallies
type TalliesRequest struct {
	PrecinctIDs        []int `json:"precinct_ids"`
	MovingViolationIDs []int `json:"moving_violation_ids"`
	Page               int   `json:"page"`
	PerPage            int   `json:"per_page"`
	StartYear          int   `json:"start_year"`
	EndYear            int   `json:"end_year"`
}

// NewTalliesRequest create new default instance of a TalliesRequest
func NewTalliesRequest() TalliesRequest {
	req := TalliesRequest{}
	req.Page = 1
	req.PerPage = 1000
	req.StartYear = 2011
	req.EndYear = 2020

	return req
}

package queries

type TalliesRequest struct {
	PrecinctIDs []int `json:"precinct_ids"`
	Page        int   `json:"page"`
	PerPage     int   `json:"per_page"`
}

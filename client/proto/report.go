package proto

// ReportResponse model.
type ReportResponse struct {
	Total       int64 `json:"total"`
	TotalIn     int64 `json:"total_in"`
	TotalOut    int64 `json:"total_out"`
	Success     int64 `json:"success"`
	SuccessIn   int64 `json:"success_in"`
	SuccessOut  int64 `json:"success_out"`
	Canceled    int64 `json:"canceled"`
	CanceledIn  int64 `json:"canceled_in"`
	CanceledOut int64 `json:"canceled_out"`
	Expired     int64 `json:"expired"`
	Failed      int64 `json:"failed"`
	FailedIn    int64 `json:"failed_in"`
	FailedOut   int64 `json:"failed_out"`
}

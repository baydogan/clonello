package models

type CreateListRequest struct {
	BoardID string `json:"board_id"`
	Title   string `json:"title"`
}

type List struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	BoardID string `json:"board_id"`
}

type GetListsResponse struct {
	Lists []List `json:"lists"`
}

type CreateListResponse struct {
	ListID string `json:"list_id"`
}

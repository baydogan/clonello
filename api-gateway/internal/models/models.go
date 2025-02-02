package models

type CreateBoardRequest struct {
	UserID string `json:"owner_id"`
	Name   string `json:"name"`
}

type CreateBoardResponse struct {
	BoardID string `json:"board_id"`
}

type GetBoardsResponse struct {
	Boards []Board `json:"boards"`
}

type Board struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	OwnerID string `json:"owner_id"`
}

package common


//
type RequestId struct {
	FileId   string `json:"file_id"`
	Uploader string `json:"uploader"`
}

//
type Piece struct {
	Id    string  `json:"piece_id"`
	Size  int64   `json:"piece_size"`
	Price float64 `json:"price"`
}

//
type RequestInfo struct {
	Id          RequestId `json:"id"`
	Pieces      []Piece   `json:"pieces"`
	Copies      int64     `json:"num_of_copies"`
	NodeId      string    `json:"node_id"`
	FileName    string    `json:"filename"`
	Description string    `json:"description"`
}


//
type TVFileInfo struct {
	FileHash    string   `json:"file_hash"`
	FileName    string   `json:"file_name"`
	FileSize    int64    `json:"file_size"`
	Pieces      []Piece  `json:"pieces"`
	Copies      int      `json:"copies"`
	Price       float64  `json:"price"`
	Description string   `json:"description"`
	NodeId      string   `json:"node_id"`
}


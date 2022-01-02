package payloads

type AlbumAttributes struct {
	Title     string  `json:"title"`
	Artist    string  `json:"artist"`
	Price     float64 `json:"price"`
	CreatedAt string  `json:"createdAt"`
}

type ResourceObject struct {
	Id         string      `json:"id"`
	Attributes interface{} `json:"attributes"`
}

type Response struct {
	Data ResourceObject `json:"data"`
}

type ListResponse struct {
	Data []ResourceObject `json:"data"`
	Meta interface{}      `json:"meta"`
}

type ListMeta struct {
	Count int `json:"count"`
}

func MakeResponse(obj ResourceObject) Response {
	return Response{
		Data: obj,
	}
}

func MakeListResponse(objs []ResourceObject) ListResponse {
	return ListResponse{
		Data: objs,
		Meta: ListMeta{
			Count: len(objs),
		},
	}
}

package payloads

type ResourceObject struct {
	Id            string                  `json:"id"`
	Type          string                  `json:"type"`
	Attributes    interface{}             `json:"attributes"`
	Relationships ResponseRelationshipMap `json:"relationships"`
}

type ResponseRelationshipMap map[string]ResponseRelationshipModel

type ResponseRelationshipModel struct {
	Data ResourceLinkage `json:"data"`
}

type ResourceLinkage []ResourceIdentifier

type ResourceIdentifier struct {
	Id   string `json:"id"`
	Type string `json:"type"`
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

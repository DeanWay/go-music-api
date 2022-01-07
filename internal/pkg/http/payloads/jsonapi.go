package payloads

type ResponseResourceObject struct {
	Id            string                  `json:"id"`
	Type          string                  `json:"type"`
	Attributes    interface{}             `json:"attributes"`
	Relationships ResponseRelationshipMap `json:"relationships"`
}

type RequestRelationshipMap map[string]RequestRelationshipModel

type RequestRelationshipModel struct {
	Data ResourceLinkage `json:"data"`
}

type ResponseRelationshipMap map[string]ResponseRelationshipModel

type ResponseRelationshipModel struct {
	Data  ResourceLinkage   `json:"data"`
	Meta  interface{}       `json:"meta,omitempty"`
	Links map[string]string `json:"links,omitempty"`
}

type ResourceLinkage []ResourceIdentifier

type ResourceIdentifier struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type Response struct {
	Data ResponseResourceObject `json:"data"`
}

type ListResponse struct {
	Data []ResponseResourceObject `json:"data"`
	Meta interface{}              `json:"meta"`
}

type ListMeta struct {
	Count int `json:"count"`
}

func MakeResponse(obj ResponseResourceObject) Response {
	return Response{
		Data: obj,
	}
}

func MakeListResponse(objs []ResponseResourceObject) ListResponse {
	return ListResponse{
		Data: objs,
		Meta: ListMeta{
			Count: len(objs),
		},
	}
}

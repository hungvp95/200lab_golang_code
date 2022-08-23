package common

type successRes struct {
	Data   any `json:"data"`
	Paging any `json:"paging,omitempty"`
	Filter any `json:"filter,omitempty"`
}

func NewSuccessRes(data, paging, filter any) successRes  {
	return successRes{
		Data: data, 
		Paging: paging, 
		Filter: filter,
	}
}

func NewDataSuccessRes(data any) successRes {
	return successRes{
		Data: data,
		Paging: nil,
		Filter: nil,
	}
}

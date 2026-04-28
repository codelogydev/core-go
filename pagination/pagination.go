package pagination

type Params struct {
	Page  int `form:"page"  json:"page"`
	Limit int `form:"limit" json:"limit"`
}

type Meta struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}

type Response struct {
	Data interface{} `json:"data"`
	Meta Meta        `json:"meta"`
}

func (p *Params) Normalize() {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.Limit <= 0 || p.Limit > 100 {
		p.Limit = 10
	}
}

func (p *Params) Offset() int {
	return (p.Page - 1) * p.Limit
}

func NewResponse(data interface{}, total int, params Params) Response {
	pages := total / params.Limit
	if total%params.Limit > 0 {
		pages++
	}
	return Response{
		Data: data,
		Meta: Meta{
			Page:       params.Page,
			Limit:      params.Limit,
			Total:      total,
			TotalPages: pages,
		},
	}
}

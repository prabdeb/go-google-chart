package googlechart

type (
	//ColsType is Google Chart standard Cols type
	ColsType struct {
		ID      string `json:"id"`
		Label   string `json:"label"`
		Pattern string `json:"pattern"`
		Type    string `json:"type"`
	}
	//RowType is Google Chart standard Rows type
	RowType struct {
		C []RowCType `json:"c"`
	}
	//RowCType is Google Chart standard Rows.C type
	RowCType struct {
		V interface{} `json:"v"`
		F interface{} `json:"f"`
	}
	//Chart is Google Chart standard type
	Chart struct {
		Cols []ColsType `json:"cols"`
		Rows []RowType  `json:"rows"`
	}
)

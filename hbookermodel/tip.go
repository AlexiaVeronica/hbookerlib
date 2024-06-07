package hbookermodel

type Tip struct {
	Code string `json:"code"`
	Tip  string `json:"tip"`
}

func (u *Tip) GetCode() string {
	return u.Code
}

func (u *Tip) GetTip() string {
	return u.Tip
}

func (u *Tip) IsSuccess() bool {
	return u.Code == "100000"
}

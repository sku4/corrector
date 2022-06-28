package speller

type Request struct {
	Lang   string   `json:"lang"`
	Format string   `json:"format"`
	Texts  []string `json:"text"`
}

func NewRequest() (pd *Request) {
	pd = &Request{}
	pd.Lang = "ru"
	pd.Format = "plain"
	return
}

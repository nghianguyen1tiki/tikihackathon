package crawl

type Data struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Rating      float32 `json:"avgRating"`
	Servings    int     `json:"servings"`
	TotalTime   int     `json:"totalTime"`
	Ingredients []struct {
		Unit struct {
			Unit  string `json:"unit"`
			Value string `json:"value"`
		} `json:"unit"`
		Name     *string `json:"name"`
		Quantity string `json:"quantity"`
	} `json:"ingredients"`
	Steps []struct {
		Content string `json:"content"`
		Photos  [][]struct {
			URL    string `json:"url"`
			Height int    `json:"height"`
			Width  int    `json:"width"`
		} `json:"photos"`
	} `json:"steps"`
	Photos [][]struct {
		URL    string `json:"url"`
		Height int    `json:"height"`
		Width  int    `json:"width"`
	} `json:"photos"`
}

type Response struct {
	Data Data `json:"data"`
	Code int  `json:"code"`
}

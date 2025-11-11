package analog

// AnalogsRequest represents request for analogs search
type AnalogsRequest struct {
	User     string     `json:"user"`
	Password string     `json:"password"`
	Service  string     `json:"service"`
	Action   string     `json:"action"`
	Rating   int        `json:"rating"`
	Lang     string     `json:"lng"`
	Data     [][]string `json:"data"`
}

// AnalogsResponse represents response for analogs search
type AnalogsResponse struct {
	Error string      `json:"error"`
	Time  float64     `json:"time"`
	Data  interface{} `json:"data"` // This will be a map with article as key and brand info as value
}

// AnalogsAdvRequest represents request for advanced analogs search
type AnalogsAdvRequest struct {
	User     string   `json:"user"`
	Password string   `json:"password"`
	Service  string   `json:"service"`
	Action   string   `json:"action"`
	Rating   int      `json:"rating"`
	Lang     string   `json:"lng"`
	Data     []string `json:"data"`
}

// AnalogsAdvResponse represents response for advanced analogs search
type AnalogsAdvResponse struct {
	Error string      `json:"error"`
	Time  float64     `json:"time"`
	Data  interface{} `json:"data"` // This will be a map with articleBrand as key and analog info as value
}

// ProducersAdvRequest represents request for advanced producers search
type ProducersAdvRequest struct {
	User     string   `json:"user"`
	Password string   `json:"password"`
	Service  string   `json:"service"`
	Action   string   `json:"action"`
	Rating   int      `json:"rating"`
	Data     []string `json:"data"`
}

// ProducersAdvResponse represents response for advanced producers search
type ProducersAdvResponse struct {
	Error string      `json:"error"`
	Time  float64     `json:"time"`
	Data  interface{} `json:"data"` // This will be a map with articleBrand as key and producer info as value
}

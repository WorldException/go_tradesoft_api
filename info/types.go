package info

// BrandByArticle represents brand information for article search
type BrandByArticle struct {
	Brand string `json:"brand"`
	Name  string `json:"name"`
}

// BrandByBarcode represents brand information for barcode search
type BrandByBarcode struct {
	Code  string `json:"code"`
	Brand string `json:"brand"`
	Name  string `json:"name"`
}

// PartInfoResult represents part information result
type PartInfoResult struct {
	Code             string            `json:"code"`
	Brand            string            `json:"brand"`
	Name             string            `json:"name"`
	Barcode          string            `json:"barCode"`
	PrePacking       int               `json:"prePacking"`
	ReplaceCode      string            `json:"replaceCode"`
	Weight           float64           `json:"weight"`
	Parameters       []Parameter       `json:"parameters"`
	BrandInfo        BrandInfo         `json:"brandInfo"`
	Apply            []ApplyItem       `json:"apply"`
	OriginalReplaces []OriginalReplace `json:"originalReplaces"`
	PartsInSet       []PartInSet       `json:"partsInSet"`
	Images           []string          `json:"images"`
}

// Parameter represents part parameter
type Parameter struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// BrandInfo represents brand information
type BrandInfo struct {
	Key        string           `json:"key"`
	Name       string           `json:"name"`
	Site       string           `json:"site"`
	Additional []AdditionalInfo `json:"additional"`
}

// AdditionalInfo represents additional brand info
type AdditionalInfo struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// ApplyItem represents applicability information
type ApplyItem struct {
	Brand        string `json:"brand"`
	Model        string `json:"model"`
	Modification string `json:"modif"`
	DateFrom     int64  `json:"dateFrom"`
	DateTo       *int64 `json:"dateTo,omitempty"`
	KW           int    `json:"kw"`
	HP           int    `json:"hp"`
	CC           int    `json:"cc"`
	Body         string `json:"body"`
	EngineCode   string `json:"engCode"`
}

// OriginalReplace represents original replacement information
type OriginalReplace struct {
	Brand string   `json:"brand"`
	Codes []string `json:"codes"`
}

// PartInSet represents part in set information
type PartInSet struct {
	Code  string `json:"code"`
	Brand string `json:"brand"`
	Name  string `json:"name"`
}

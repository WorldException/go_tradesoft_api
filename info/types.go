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

// PartInfoResponse represents the response for part info request
type PartInfoResponse struct {
	Result *PartInfoResult `json:"result,omitempty"`
	Error  *string         `json:"error,omitempty"`
}

// PartInfoRequest represents the request for part info
type PartInfoRequest struct {
	User     string        `json:"user"`
	Password string        `json:"password"`
	Service  string        `json:"service"`
	Action   string        `json:"action"`
	Param    PartInfoParam `json:"param"`
}

// PartInfoParam represents parameters for part info request
type PartInfoParam struct {
	Code  string `json:"code"`
	Brand string `json:"brand"`
	Lang  string `json:"lang"`
}

// BrandsByArticleResponse represents the response for brands by article request
type BrandsByArticleResponse struct {
	Result []BrandByArticle `json:"result,omitempty"`
	Error  *string          `json:"error,omitempty"`
}

// BrandsByArticleRequest represents the request for brands by article
type BrandsByArticleRequest struct {
	User     string               `json:"user"`
	Password string               `json:"password"`
	Service  string               `json:"service"`
	Action   string               `json:"action"`
	Param    BrandsByArticleParam `json:"param"`
}

// BrandsByArticleParam represents parameters for brands by article request
type BrandsByArticleParam struct {
	Code string `json:"code"`
	Lang string `json:"lang"`
}

// BrandsByBarcodeResponse represents the response for brands by barcode request
type BrandsByBarcodeResponse struct {
	Result []BrandByBarcode `json:"result,omitempty"`
	Error  *string          `json:"error,omitempty"`
}

// BrandsByBarcodeRequest represents the request for brands by barcode
type BrandsByBarcodeRequest struct {
	User     string                `json:"user"`
	Password string                `json:"password"`
	Service  string                `json:"service"`
	Action   string                `json:"action"`
	Params   BrandsByBarcodeParams `json:"params"`
}

// BrandsByBarcodeParams represents parameters for brands by barcode request
type BrandsByBarcodeParams struct {
	Barcode int64 `json:"barcode"`
}

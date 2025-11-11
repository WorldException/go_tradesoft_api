package provider

import (
	"tradesoft/common"
)

// ProducerListRequest represents request for producer list
type ProducerListRequest struct {
	User      string             `json:"user"`
	Password  string             `json:"password"`
	Service   string             `json:"service"`
	Action    string             `json:"action"`
	Timelimit int                `json:"timelimit"`
	Container []ProducerListItem `json:"container"`
}

// ProducerListItem represents item in producer list request
type ProducerListItem struct {
	Provider common.ProviderID `json:"provider"`
	Login    string            `json:"login"`
	Password string            `json:"password"`
	Code     string            `json:"code"`
	Options  common.Options    `json:"options,omitempty"`
}

// ProducerListResponse represents response for producer list
type ProducerListResponse struct {
	Error     string                  `json:"error"`
	Time      float64                 `json:"time"`
	Container []ProducerListContainer `json:"container"`
}

// ProducerListContainer represents container in producer list response
type ProducerListContainer struct {
	Cache    common.Cache       `json:"cache"`
	Provider common.ProviderID  `json:"provider"`
	Time     float64            `json:"time"`
	Data     []ProducerListData `json:"data"`
}

// ProducerListData represents data in producer list response
type ProducerListData struct {
	Code     string `json:"code"`
	Producer string `json:"producer"`
	Caption  string `json:"caption"`
}

// PriceListRequest represents request for price list
type PriceListRequest struct {
	User      string          `json:"user"`
	Password  string          `json:"password"`
	Service   string          `json:"service"`
	Action    string          `json:"action"`
	Timelimit int             `json:"timelimit"`
	Container []PriceListItem `json:"container"`
}

// PriceListItem represents item in price list request
type PriceListItem struct {
	Provider common.ProviderID `json:"provider"`
	Login    string            `json:"login"`
	Password string            `json:"password"`
	Code     string            `json:"code"`
	Producer string            `json:"producer"`
	Options  common.Options    `json:"options,omitempty"`
}

// PriceListResponse represents response for price list
type PriceListResponse struct {
	Error     string               `json:"error"`
	Time      float64              `json:"time"`
	Container []PriceListContainer `json:"container"`
}

// PriceListContainer represents container in price list response
type PriceListContainer struct {
	Cache    common.Cache      `json:"cache"`
	Provider common.ProviderID `json:"provider"`
	Time     float64           `json:"time"`
	Data     []PriceListData   `json:"data"`
}

// PriceListData represents data in price list response
type PriceListData struct {
	Code              string              `json:"code"`
	Producer          string              `json:"producer"`
	Price             float64             `json:"price"`
	CurrencyCode      string              `json:"currencycode"`
	Caption           string              `json:"caption"`
	Rest              string              `json:"rest"`
	DeliveryDays      string              `json:"deliverydays"`
	DeliveryDaysMin   int                 `json:"deliverydays_min"`
	DeliveryDaysMax   int                 `json:"deliverydays_max"`
	Direction         string              `json:"direction"`
	Updated           int64               `json:"updated"`
	MinQuantity       int                 `json:"minquantity"`
	Amount            int                 `json:"amount"`
	ProviderRating    int                 `json:"provider_rating"`
	WarProviderRating int                 `json:"war_provider_rating"`
	ProviderCode      common.ProviderCode `json:"providercode"`
	Weight            float64             `json:"weight"`
	ItemType          string              `json:"itemtype"`
	ItemHash          common.ItemHash     `json:"itemHash"`
	DeliveryInfo      string              `json:"delivery"`
	DeliveryCost      float64             `json:"delivery_cost"`
	DeliveryCurrency  string              `json:"delivery_currency"`
	Return            string              `json:"return"`
	ReturnInfo        string              `json:"return_info"`
	IsDealer          bool                `json:"isDealer"`
	IsUsed            bool                `json:"isUsed"`
	Image             string              `json:"image"`
}

// AdditionalPartInfoRequest represents request for additional part info
type AdditionalPartInfoRequest struct {
	User      string                   `json:"user"`
	Password  string                   `json:"password"`
	Service   string                   `json:"service"`
	Action    string                   `json:"action"`
	Container []AdditionalPartInfoItem `json:"container"`
}

// AdditionalPartInfoItem represents item in additional part info request
type AdditionalPartInfoItem struct {
	Provider     common.ProviderID   `json:"provider"`
	Login        string              `json:"login"`
	Password     string              `json:"password"`
	ProviderCode common.ProviderCode `json:"providerCode"`
}

// AdditionalPartInfoResponse represents response for additional part info
type AdditionalPartInfoResponse struct {
	Error     string                        `json:"error"`
	Time      float64                       `json:"time"`
	Container []AdditionalPartInfoContainer `json:"container"`
}

// AdditionalPartInfoContainer represents container in additional part info response
type AdditionalPartInfoContainer struct {
	Provider     common.ProviderID   `json:"provider"`
	Time         float64             `json:"time"`
	ProviderCode common.ProviderCode `json:"providerCode"`
	Name         string              `json:"name"`
	Images       []string            `json:"images"`
	Description  string              `json:"description"`
}

// OptionsListRequest represents request for options list
type OptionsListRequest struct {
	User      string            `json:"user"`
	Password  string            `json:"password"`
	Service   string            `json:"service"`
	Action    string            `json:"action"`
	Container []OptionsListItem `json:"container"`
}

// OptionsListItem represents item in options list request
type OptionsListItem struct {
	Provider common.ProviderID `json:"provider"`
	Login    string            `json:"login"`
	Password string            `json:"password"`
}

// OptionsListResponse represents response for options list
type OptionsListResponse struct {
	Error     string                 `json:"error"`
	Time      float64                `json:"time"`
	Container []OptionsListContainer `json:"container"`
}

// OptionsListContainer represents container in options list response
type OptionsListContainer struct {
	Cache     common.Cache      `json:"cache"`
	Provider  common.ProviderID `json:"provider"`
	Time      float64           `json:"time"`
	MakeOrder string            `json:"makeOrder"`
	Data      []OptionsListData `json:"data"`
}

// OptionsListData represents data in options list response
type OptionsListData struct {
	Name        string   `json:"name"`
	Title       string   `json:"title"`
	Type        string   `json:"type"`
	Multiple    bool     `json:"multiple"`
	Options     []string `json:"options"`
	Default     string   `json:"default"`
	SortIndex   int      `json:"sortIndex"`
	Description string   `json:"description"`
}

// ProviderListRequest represents request for provider list
type ProviderListRequest struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Service  string `json:"service"`
	Action   string `json:"action"`
}

// ProviderListResponse represents response for provider list
type ProviderListResponse struct {
	Error         string              `json:"error"`
	Time          float64             `json:"time"`
	Data          []SupplierData      `json:"data"`
	GeoDictionary []GeoDictionaryItem `json:"geoDictionary"`
}

// PreOrderSearchRequest represents request for pre-order search
type PreOrderSearchRequest struct {
	User      string               `json:"user"`
	Password  string               `json:"password"`
	Service   string               `json:"service"`
	Action    string               `json:"action"`
	Timelimit int                  `json:"timelimit"`
	Container []PreOrderSearchItem `json:"container"`
}

// PreOrderSearchItem represents item in pre-order search request
type PreOrderSearchItem struct {
	Provider common.ProviderID `json:"provider"`
	Login    string            `json:"login"`
	Password string            `json:"password"`
	Code     string            `json:"code"`
	Producer string            `json:"producer"`
	ItemHash common.ItemHash   `json:"itemHash"`
	Options  common.Options    `json:"options,omitempty"`
}

// PreOrderSearchResponse represents response for pre-order search
type PreOrderSearchResponse struct {
	Error     string                    `json:"error"`
	Time      float64                   `json:"time"`
	Container []PreOrderSearchContainer `json:"container"`
}

// PreOrderSearchContainer represents container in pre-order search response
type PreOrderSearchContainer struct {
	Cache    common.Cache             `json:"cache"`
	Provider common.ProviderID        `json:"provider"`
	Time     float64                  `json:"time"`
	Items    []PreOrderSearchItemData `json:"items"`
}

// PreOrderSearchItemData represents item data in pre-order search response
type PreOrderSearchItemData struct {
	Code              string              `json:"code"`
	Producer          string              `json:"producer"`
	Price             float64             `json:"price"`
	CurrencyCode      string              `json:"currencycode"`
	Caption           string              `json:"caption"`
	Rest              string              `json:"rest"`
	DeliveryDays      string              `json:"deliverydays"`
	DeliveryDaysMin   int                 `json:"deliverydays_min"`
	DeliveryDaysMax   int                 `json:"deliverydays_max"`
	Direction         string              `json:"direction"`
	Updated           int64               `json:"updated"`
	MinQuantity       int                 `json:"minquantity"`
	Amount            int                 `json:"amount"`
	ProviderRating    int                 `json:"provider_rating"`
	WarProviderRating int                 `json:"war_provider_rating"`
	ProviderCode      common.ProviderCode `json:"providercode"`
	Weight            float64             `json:"weight"`
	ItemType          string              `json:"itemtype"`
	ItemHash          common.ItemHash     `json:"itemHash"`
	DeliveryInfo      string              `json:"delivery"`
	DeliveryCost      float64             `json:"delivery_cost"`
	DeliveryCurrency  string              `json:"delivery_currency"`
	ItemID            common.ItemID       `json:"itemId"`
	Return            string              `json:"return"`
	ReturnInfo        string              `json:"return_info"`
	IsDealer          bool                `json:"isDealer"`
	IsUsed            bool                `json:"isUsed"`
	Image             string              `json:"image"`
}

// MakeOrderOfflineRequest represents request for make order offline
type MakeOrderOfflineRequest struct {
	User      string           `json:"user"`
	Password  string           `json:"password"`
	Service   string           `json:"service"`
	Action    string           `json:"action"`
	Timelimit int              `json:"timelimit"`
	Param     []MakeOrderParam `json:"param"`
}

// MakeOrderParam represents parameter in make order request
type MakeOrderParam struct {
	Provider          common.ProviderID `json:"provider"`
	Login             string            `json:"login"`
	Password          string            `json:"password"`
	Comment           string            `json:"comment"`
	ClientOrderNumber string            `json:"clientOrderNumber"`
	Items             []MakeOrderItem   `json:"items"`
	Options           common.Options    `json:"options,omitempty"`
}

// MakeOrderItem represents item in make order request
type MakeOrderItem struct {
	ItemID    common.ItemID `json:"itemId"`
	Quantity  int           `json:"quantity"`
	Reference string        `json:"reference"`
	Comment   string        `json:"comment"`
}

// MakeOrderOfflineResponse represents response for make order offline
type MakeOrderOfflineResponse struct {
	Error  string          `json:"error"`
	Time   float64         `json:"time"`
	Result MakeOrderResult `json:"result"`
}

// MakeOrderResult represents result in make order response
type MakeOrderResult struct {
	Provider    common.ProviderID     `json:"provider"`
	OrderStatus string                `json:"orderStatus"`
	Time        float64               `json:"time"`
	Items       []MakeOrderItemResult `json:"items"`
}

// MakeOrderItemResult represents item result in make order response
type MakeOrderItemResult struct {
	ItemID      common.ItemID      `json:"itemId"`
	OrderItemID common.OrderItemID `json:"orderItemId"`
	Error       string             `json:"error"`
}

// StatusListRequest represents request for status list
type StatusListRequest struct {
	User      string           `json:"user"`
	Password  string           `json:"password"`
	Service   string           `json:"service"`
	Action    string           `json:"action"`
	Container []StatusListItem `json:"container"`
}

// StatusListItem represents item in status list request
type StatusListItem struct {
	Provider common.ProviderID `json:"provider"`
	Login    string            `json:"login"`
	Password string            `json:"password"`
}

// StatusListResponse represents response for status list
type StatusListResponse struct {
	Error     string                `json:"error"`
	Time      float64               `json:"time"`
	Container []StatusListContainer `json:"container"`
}

// StatusListContainer represents container in status list response
type StatusListContainer struct {
	Provider   common.ProviderID `json:"provider"`
	Error      string            `json:"error"`
	StatusList []Status          `json:"statusList"`
	Time       float64           `json:"time"`
}

// Status represents status of item
type Status struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// ItemsStatusRequest represents request for items status
type ItemsStatusRequest struct {
	User      string            `json:"user"`
	Password  string            `json:"password"`
	Service   string            `json:"service"`
	Action    string            `json:"action"`
	Container []ItemsStatusItem `json:"container"`
}

// ItemsStatusItem represents item in items status request
type ItemsStatusItem struct {
	Provider common.ProviderID    `json:"provider"`
	Login    string               `json:"login"`
	Password string               `json:"password"`
	Items    []common.OrderItemID `json:"items"`
	Options  common.Options       `json:"options,omitempty"`
}

// ItemsStatusResponse represents response for items status
type ItemsStatusResponse struct {
	Error     string                 `json:"error"`
	Time      float64                `json:"time"`
	Container []ItemsStatusContainer `json:"container"`
}

// ItemsStatusContainer represents container in items status response
type ItemsStatusContainer struct {
	Provider common.ProviderID              `json:"provider"`
	Error    string                         `json:"error"`
	Items    map[string]ItemsStatusItemData `json:"items"`
	Time     float64                        `json:"time"`
}

// ItemsStatusItemData represents item data in items status response
type ItemsStatusItemData struct {
	ProviderItemId      common.OrderItemID `json:"providerItemId"`
	ProviderOrderNumber string             `json:"providerOrderNumber"`
	ProviderID          string             `json:"providerId"`
	Quantity            int                `json:"quantity"`
	StateID             string             `json:"stateId"`
	StateName           string             `json:"stateName"`
	Price               float64            `json:"price"`
	CurrencyCode        string             `json:"currencycode"`
	Code                string             `json:"code"`
	Brand               string             `json:"brand"`
	Substates           []Substate         `json:"substates,omitempty"`
	Error               string             `json:"error"`
}

// Substate represents substate in items status response
type Substate struct {
	InnerID   string `json:"innerId,omitempty"`
	Quantity  int    `json:"quantity"`
	StateID   string `json:"stateId"`
	StateName string `json:"stateName"`
}

// SupplierData represents supplier data
type SupplierData struct {
	Name              common.ProviderID  `json:"name"`
	Active            bool               `json:"active"`
	OrderActive       bool               `json:"orderActive"`
	OrderFeature      bool               `json:"orderFeature"`
	SyncFeature       bool               `json:"syncFeature"`
	ActiveTo          *common.CustomTime `json:"activeTo,omitempty"` // format "04.03.2014"
	Title             string             `json:"title"`
	Description       string             `json:"description"`
	ContractRequired  bool               `json:"contractRequired"`
	AgreementText     string             `json:"agreementText"`
	OnlyAgreementText string             `json:"onlyAgreementText"`
	LoginInfoText     string             `json:"loginInfoText"`
	SiteUrl           string             `json:"siteUrl"`
	IconUrl           string             `json:"iconUrl"`
	GeoIds            []string           `json:"geoIds"`
}

// GeoDictionaryItem represents geo dictionary item
type GeoDictionaryItem struct {
	Key  string `json:"key"`
	Name string `json:"name"`
	Type string `json:"type"`
}

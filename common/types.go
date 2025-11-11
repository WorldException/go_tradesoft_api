package common

// ProviderID represents supplier ID
type ProviderID string

// ProviderCode represents part id on the supplier's website
type ProviderCode string

// ItemID represents unique part ID from preliminary price search
type ItemID string

// OrderItemID represents internal ID of the item for upcoming item synchronization
type OrderItemID string

// RequestError represents request error
type RequestError string

// ItemHash represents position key obtained by GetPriceList method
type ItemHash string

// Time represents request processing time
type Time float64

// Cache represents cache flag (0 - data from provider, 1 - data from service cache)
type Cache int

// Options represents provider options
type Options map[string]string

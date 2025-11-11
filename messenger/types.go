package messenger

// SendSMSRequest represents request for sending SMS
type SendSMSRequest struct {
	User     string   `json:"user"`
	Password string   `json:"password"`
	Service  string   `json:"service"`
	Action   string   `json:"action"`
	Param    SMSParam `json:"param"`
}

// SMSParam represents SMS parameters
type SMSParam struct {
	PhoneNumber string `json:"phoneNumber"`
	Text        string `json:"text"`
}

// SendSMSResponse represents response for sending SMS
type SendSMSResponse struct {
	Error  string    `json:"error"`
	Time   float64   `json:"time"`
	Result SMSResult `json:"result"`
}

// SMSResult represents SMS result
type SMSResult struct {
	SMSID       string `json:"smsId"`
	Sender      string `json:"sender"`
	PhoneNumber int64  `json:"phoneNumber"`
	Segments    int    `json:"segments"`
	Error       string `json:"error"`
}

// SMSStatusRequest represents request for getting SMS status
type SMSStatusRequest struct {
	User     string   `json:"user"`
	Password string   `json:"password"`
	Service  string   `json:"service"`
	Action   string   `json:"action"`
	Param    []string `json:"param"`
}

// SMSStatusResponse represents response for getting SMS status
type SMSStatusResponse struct {
	Error  string         `json:"error"`
	Time   float64        `json:"time"`
	Result []StatusResult `json:"result"`
}

// StatusResult represents SMS status result
type StatusResult struct {
	SMSID    string `json:"smsId"`
	Status   string `json:"status"`
	Sender   string `json:"sender"`
	Segments int    `json:"segments"`
	Error    string `json:"error"`
}

// SMSBalanceRequest represents request for getting SMS balance
type SMSBalanceRequest struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Service  string `json:"service"`
	Action   string `json:"action"`
}

// SMSBalanceResponse represents response for getting SMS balance
type SMSBalanceResponse struct {
	Error  string           `json:"error"`
	Time   float64          `json:"time"`
	Result SMSBalanceResult `json:"result"`
}

// SMSBalanceResult represents SMS balance result
type SMSBalanceResult struct {
	Balance int `json:"balance"`
}

// SMSBalance2Request represents request for getting SMS balance 2
type SMSBalance2Request struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Service  string `json:"service"`
	Action   string `json:"action"`
}

// SMSBalance2Response represents response for getting SMS balance 2
type SMSBalance2Response struct {
	Error  string              `json:"error"`
	Time   float64             `json:"time"`
	Result []SMSBalance2Result `json:"result"`
}

// SMSBalance2Result represents SMS balance 2 result
type SMSBalance2Result struct {
	Key     string `json:"key"`
	Name    string `json:"name"`
	Balance int    `json:"balance"`
}

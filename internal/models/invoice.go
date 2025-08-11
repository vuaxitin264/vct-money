package models

type Item struct {
    Name   string  `json:"name"`
    Qty    int     `json:"qty"`
    Price  float64 `json:"price"`
    Notes  string  `json:"notes"`
}

type Surcharge struct {
    Label  string  `json:"label"`
    Amount float64 `json:"amount"`
}

type Invoice struct {
    ID         int64       `json:"id"`
    OrderCode  string      `json:"order_code"`
    CustomerID int64       `json:"customer_id"`
    Items      []Item      `json:"items"`
    WeightLbs  float64     `json:"weight_lbs"`
    Subtotal   float64     `json:"subtotal"`
    Surcharges []Surcharge `json:"surcharges"`
    Total      float64     `json:"total"`
    Currency   string      `json:"currency"`

    // Money transfer (NEW)
    TransferAmount float64 `json:"transfer_amount"`
    TransferFee    float64 `json:"transfer_fee"`
    TotalRecipient float64 `json:"total_recipient"`
    InvoiceDate    string  `json:"invoice_date"` // YYYY-MM-DD
    BankInfoVN     string  `json:"bank_info_vn"`

    Status     string `json:"status"`
    Notes      string `json:"notes"`
    CreatedAt  string `json:"created_at"`
    UpdatedAt  string `json:"updated_at"`
}

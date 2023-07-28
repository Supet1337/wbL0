package internal

import (
	"fmt"
	"time"
)

type Error struct {
	Message string
	Code    int
}

func (err *Error) Error() string {
	return fmt.Sprintf("Code: %d\nMessage: %s\n", err.Code, err.Message)
}

type Order struct {
	OrderUid    string `json:"order_uid" validate:"nonzero"`
	TrackNumber string `json:"track_number" validate:"nonzero"`
	Entry       string `json:"entry" validate:"nonzero"`
	Delivery    struct {
		Name    string `json:"name" validate:"nonzero"`
		Phone   string `json:"phone" validate:"nonzero"`
		Zip     string `json:"zip" validate:"nonzero"`
		City    string `json:"city" validate:"nonzero"`
		Address string `json:"address" validate:"nonzero"`
		Region  string `json:"region" validate:"nonzero"`
		Email   string `json:"email" validate:"regexp=^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$"`
	} `json:"delivery" validate:"nonzero"`
	Payment struct {
		Transaction  string `json:"transaction" validate:"nonzero"`
		RequestId    string `json:"request_id" validate:"nonzero"`
		Currency     string `json:"currency" validate:"regexp=^[A-Z]{3}$"`
		Provider     string `json:"provider" validate:"nonzero"`
		Amount       int    `json:"amount" validate:"min=0,max=1000000"`
		PaymentDt    int64  `json:"payment_dt" validate:"nonzero"`
		Bank         string `json:"bank" validate:"nonzero"`
		DeliveryCost int    `json:"delivery_cost" validate:"min=0,max=1000000"`
		GoodsTotal   int    `json:"goods_total" validate:"min=0,max=1000000"`
		CustomFee    int    `json:"custom_fee" validate:"min=0,max=1000000"`
	} `json:"payment" validate:"nonzero"`
	Items []struct {
		ChrtId      int    `json:"chrt_id" validate:"nonzero"`
		TrackNumber string `json:"track_number" validate:"nonzero"`
		Price       int    `json:"price" validate:"min=0,max=1000000"`
		Rid         string `json:"rid" validate:"nonzero"`
		Name        string `json:"name" validate:"nonzero"`
		Sale        int    `json:"sale" validate:"min=0,max=100"`
		Size        string `json:"size" validate:"nonzero"`
		TotalPrice  int    `json:"total_price" validate:"min=0,max=1000000"`
		NmId        int    `json:"nm_id" validate:"nonzero"`
		Brand       string `json:"brand" validate:"nonzero"`
		Status      int    `json:"status" validate:"min=0,max=10"`
	} `json:"items"`
	Locale            string    `json:"locale" validate:"nonzero"`
	InternalSignature string    `json:"internal_signature" validate:"nonzero"`
	CustomerId        string    `json:"customer_id" validate:"nonzero"`
	DeliveryService   string    `json:"delivery_service" validate:"nonzero"`
	Shardkey          string    `json:"shardkey" validate:"nonzero"`
	SmId              int       `json:"sm_id" validate:"min=0,max=100000"`
	DateCreated       time.Time `json:"date_created" validate:"nonzero"`
	OofShard          string    `json:"oof_shard" validate:"nonzero"`
}

package polymodels

import (
	"encoding/binary"
	"github.com/zeebo/xxh3"
	"time"
)

type Ticker struct {
	Ticker             string    `json:"ticker"`
	Name               string    `json:"name"`
	Market             string    `json:"market"`
	Locale             string    `json:"locale"`
	Active             bool      `json:"active"`
	CurrencySymbol     string    `json:"currency_symbol"`
	CurrencyName       string    `json:"currency_name"`
	BaseCurrencySymbol string    `json:"base_currency_symbol"`
	BaseCurrencyName   string    `json:"base_currency_name"`
	PrimaryExchange    string    `json:"primary_exchange"`
	Type               string    `json:"type"`
	Cik                string    `json:"cik"`
	CompositeFigi      string    `json:"composite_figi"`
	ShareClassFigi     string    `json:"share_class_figi"`
	LastUpdatedUtc     time.Time `json:"last_updated_utc"`
	DelistedUtc        time.Time `json:"delisted_utc"`
}

func (t *Ticker) Hash() uint64 {

	var buff [41]byte

	copy(buff[0:0+8], t.Ticker)
	copy(buff[8:8+8], t.Market)
	copy(buff[16:16+8], t.Locale)
	binary.BigEndian.PutUint64(buff[24:24+8], uint64(t.LastUpdatedUtc.UnixNano()))
	binary.BigEndian.PutUint64(buff[32:32+8], uint64(t.DelistedUtc.UnixNano()))
	if t.Active == true {
		buff[40] = 1
	}

	return xxh3.Hash(buff[:])
}

package polymodels

import (
	"encoding/binary"
	"github.com/zeebo/xxh3"
)

type HistoricEquityTrade struct {
	Ticker            string  `json:"-" csv:"ticker"`
	SIPTimestamp      int64   `json:"t" csv:"sip_timestamp,omitempty"`
	ExchangeTimestamp int64   `json:"y" csv:"exchange_timestamp,omitempty"`
	TRFTimestamp      int64   `json:"f" csv:"trf_timestamp,omitempty"`
	TRFID             int32   `json:"r" csv:"trf_id,omitempty"`
	Correction        int     `json:"e" csv:"correction,omitempty"`
	SequenceNumber    int     `json:"q" csv:"sequence_number,omitempty"`
	Id                string  `json:"i" csv:"trade_id,omitempty"`
	OriginalID        string  `json:"I" csv:"original_id,omitempty"`
	ExchangeID        int32   `json:"x" csv:"exchange_id,omitempty"`
	Volume            int64   `json:"s" csv:"volume,omitempty"`
	Conditions        []int32 `json:"c" csv:"conditions,omitempty"`
	Price             float64 `json:"p" csv:"price,omitempty"`
	Tape              int8    `json:"z" csv:"tape,omitempty"`
}

func (et *HistoricEquityTrade) Reset() {
	*et = HistoricEquityTrade{}
}

// Hash generates a unique hash for the trade object.
func (et *HistoricEquityTrade) Hash() uint64 {

	var buff [56]byte

	copy(buff[0:0+8], et.Ticker)
	binary.BigEndian.PutUint64(buff[8:8+8], uint64(et.SIPTimestamp))
	binary.BigEndian.PutUint64(buff[16:16+8], uint64(et.SequenceNumber))
	copy(buff[24:24+8], et.Id)
	binary.BigEndian.PutUint64(buff[32:32+8], hashInt32(et.Conditions))
	binary.BigEndian.PutUint64(buff[40:40+8], uint64(et.Correction))
	binary.BigEndian.PutUint64(buff[48:48+8], uint64(et.ExchangeID))

	return xxh3.Hash(buff[:])
}

func (et *HistoricEquityTrade) Equal(other *HistoricEquityTrade) bool {

	if et.SIPTimestamp != other.SIPTimestamp {
		return false
	}

	if et.ExchangeTimestamp != other.ExchangeTimestamp {
		return false
	}

	if et.TRFTimestamp != other.TRFTimestamp {
		return false
	}

	if et.TRFID != other.TRFID {
		return false
	}

	if et.Correction != other.Correction {
		return false
	}

	if et.SequenceNumber != other.SequenceNumber {
		return false
	}

	if et.Id != other.Id {
		return false
	}

	if et.OriginalID != other.OriginalID {
		return false
	}

	if et.ExchangeID != other.ExchangeID {
		return false
	}

	if et.Volume != other.Volume {
		return false
	}

	if et.Price != other.Price {
		return false
	}

	if et.Tape != other.Tape {
		return false
	}

	if et.Ticker != other.Ticker {
		return false
	}

	if equalIntArr(et.Conditions, other.Conditions) == false {
		return false
	}

	return true
}

package polymodels

import (
	"encoding/binary"
	"github.com/zeebo/xxh3"
	"github.com/zeroallox/go-lib-polygon-io-uo/polypsv"
)

type HistoricEquityTrade struct {
	Ticker            string  `json:"-"`
	SIPTimestamp      int64   `json:"t"`
	ExchangeTimestamp int64   `json:"y"`
	TRFTimestamp      int64   `json:"f"`
	TRFID             int32   `json:"r"`
	Correction        int     `json:"e"`
	SequenceNumber    int     `json:"q"`
	Id                string  `json:"i"`
	OriginalID        string  `json:"I"`
	ExchangeID        int32   `json:"x"`
	Volume            int64   `json:"s"`
	Conditions        []int32 `json:"c"`
	Price             float64 `json:"p"`
	Tape              int8    `json:"z"`
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

func (et *HistoricEquityTrade) ToPSV(psvw *polypsv.Writer) error {

	psvw.WriteNonEmptyString(et.Ticker).Sep()
	psvw.WriteNonZeroInt(et.SIPTimestamp).Sep()
	psvw.WriteNonZeroInt(et.ExchangeTimestamp).Sep()
	psvw.WriteNonZeroInt(et.TRFTimestamp).Sep()
	psvw.WriteNonZeroUint(uint64(et.TRFID)).Sep()
	psvw.WriteNonZeroUint(uint64(et.Correction)).Sep()
	psvw.WriteNonZeroInt(int64(et.SequenceNumber)).Sep()
	psvw.WriteNonEmptyString(et.Id).Sep()
	psvw.WriteNonEmptyString(et.OriginalID).Sep()
	psvw.WriteNonZeroUint(uint64(et.ExchangeID)).Sep()
	psvw.WriteNonZeroInt(et.Volume).Sep()
	psvw.WriteInt32Array(et.Conditions).Sep()
	psvw.WriteNonZeroFloat(et.Price).Sep()
	psvw.WriteNonZeroUint(uint64(et.Tape))

	return nil
}

func (et *HistoricEquityTrade) FromPSV(src [][]byte) error {

	var err error = nil

	et.Ticker = polypsv.ReadString(src[0])

	et.SIPTimestamp, err = polypsv.ReadInt(src[1])
	if err != nil {
		return err
	}

	et.ExchangeTimestamp, err = polypsv.ReadInt(src[2])
	if err != nil {
		return err
	}

	et.TRFTimestamp, err = polypsv.ReadInt(src[3])
	if err != nil {
		return err
	}

	trfID, err := polypsv.ReadInt(src[4])
	if err != nil {
		return err
	}
	et.TRFID = int32(trfID)

	corr, err := polypsv.ReadInt(src[5])
	if err != nil {
		return err
	}
	et.Correction = int(corr)

	sq, err := polypsv.ReadInt(src[6])
	if err != nil {
		return err
	}
	et.SequenceNumber = int(sq)

	et.Id = polypsv.ReadString(src[7])
	et.OriginalID = polypsv.ReadString(src[8])

	exchangeID, err := polypsv.ReadInt(src[9])
	if err != nil {
		return err
	}
	et.ExchangeID = int32(exchangeID)

	vol, err := polypsv.ReadInt(src[10])
	if err != nil {
		return err
	}
	et.Volume = vol

	cond, err := polypsv.ReadInt32Array(src[11])
	if err != nil {
		return err
	}
	et.Conditions = cond

	price, err := polypsv.ReadFloat(src[12])
	if err != nil {
		return err
	}
	et.Price = price

	tape, err := polypsv.ReadInt(src[13])
	if err != nil {
		return err
	}

	et.Tape = int8(tape)

	return nil
}

func (et *HistoricEquityTrade) IsEqual(other *HistoricEquityTrade) bool {

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

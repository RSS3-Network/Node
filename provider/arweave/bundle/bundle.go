package bundle

import (
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"io"

	"github.com/hamba/avro"
	"github.com/samber/lo"
)

type Decoder struct {
	reader io.Reader
	buffer io.Reader
	schema avro.Schema
	header Header
	cursor uint32
}

func (d *Decoder) DecodeHeader() (*Header, error) {
	if err := d.decodeHeaderDataItemNumbers(); err != nil {
		return nil, fmt.Errorf("decode data item numbers: %w", err)
	}

	if err := d.decodeHeaderDataInfos(); err != nil {
		return nil, fmt.Errorf("decode data infos: %w", err)
	}

	header := d.header

	return &header, nil
}

func (d *Decoder) DecodeDataItem() (*DataItem, error) {
	dataItemInfo := d.header.DataItemInfos[d.cursor]

	d.buffer = io.LimitReader(d.reader, int64(dataItemInfo.Size))

	var dataItem DataItem

	if err := d.decodeDataItemSignature(&dataItem); err != nil {
		return nil, fmt.Errorf("decode data item info signature: %w", err)
	}

	if err := d.decodeDataItemOwner(&dataItem); err != nil {
		return nil, fmt.Errorf("decode data item info owner: %w", err)
	}

	if err := d.decodeDataItemTarget(&dataItem); err != nil {
		return nil, fmt.Errorf("decode data item info target: %w", err)
	}

	if err := d.decodeDataItemAnchor(&dataItem); err != nil {
		return nil, fmt.Errorf("decode data item info anchor: %w", err)
	}

	if err := d.decodeDataItemTags(&dataItem); err != nil {
		return nil, fmt.Errorf("decode data item info tags: %w", err)
	}

	dataItem.dataReader = d.buffer

	d.cursor++

	return &dataItem, nil
}

func (d *Decoder) Next() bool {
	return d.cursor < d.header.Numbers
}

func (d *Decoder) decodeHeaderDataItemNumbers() error {
	buffer := make([]byte, 32)

	_, err := io.ReadFull(d.reader, buffer)
	if err != nil {
		return err
	}

	d.header.Numbers = binary.LittleEndian.Uint32(buffer)

	return nil
}

func (d *Decoder) decodeHeaderDataInfos() error {
	d.header.DataItemInfos = make([]DataItemInfo, 0, d.header.Numbers)

	buffer := make([]byte, 64)

	for index := uint32(0); index < d.header.Numbers; index++ {
		if _, err := io.ReadFull(d.reader, buffer); err != nil {
			return err
		}

		dataItemInfo := DataItemInfo{
			Size: binary.LittleEndian.Uint32(buffer[:32]),
			ID:   base64.RawURLEncoding.EncodeToString(buffer[32:]),
		}

		d.header.DataItemInfos = append(d.header.DataItemInfos, dataItemInfo)
	}

	return nil
}

func (d *Decoder) decodeDataItemSignature(dataItem *DataItem) error {
	buffer := make([]byte, 2)

	if _, err := io.ReadFull(d.buffer, buffer); err != nil {
		return err
	}

	dataItem.SignatureType = binary.LittleEndian.Uint16(buffer)

	signatureType, supported := d.lookupDataItemSignatureType(dataItem)
	if !supported {
		return fmt.Errorf("unsupported signature type: %d", dataItem.SignatureType)
	}

	buffer = make([]byte, signatureType.SignatureLength)

	if _, err := io.ReadFull(d.buffer, buffer); err != nil {
		return err
	}

	dataItem.Signature = base64.RawURLEncoding.EncodeToString(buffer)

	return nil
}

func (d *Decoder) decodeDataItemOwner(dataItem *DataItem) error {
	var buffer []byte

	signatureType, supported := d.lookupDataItemSignatureType(dataItem)
	if !supported {
		return fmt.Errorf("unsupported signature type: %d", dataItem.SignatureType)
	}

	buffer = make([]byte, signatureType.PublicKeyLength)

	if _, err := io.ReadFull(d.buffer, buffer); err != nil {
		return err
	}

	dataItem.Owner = base64.RawURLEncoding.EncodeToString(buffer)

	return nil
}

func (d *Decoder) decodeDataItemPresenceField(buffer []byte) (bool, error) {
	if _, err := io.ReadFull(d.buffer, buffer[:1]); err != nil {
		return false, err
	}

	switch flag := buffer[0]; flag {
	case 0x0:
		return false, nil
	case 0x1:
		_, err := io.ReadFull(d.buffer, buffer[1:])

		return true, err
	default:
		return false, fmt.Errorf("invalid presence byte: %d", flag)
	}
}

func (d *Decoder) decodeDataItemTarget(dataItem *DataItem) error {
	buffer := make([]byte, 32+1)

	presence, err := d.decodeDataItemPresenceField(buffer)
	if err != nil {
		return err
	}

	if presence {
		dataItem.Target = base64.RawURLEncoding.EncodeToString(buffer[1:])
	}

	return nil
}

func (d *Decoder) decodeDataItemAnchor(dataItem *DataItem) error {
	buffer := make([]byte, 32+1)

	presence, err := d.decodeDataItemPresenceField(buffer)
	if err != nil {
		return err
	}

	if presence {
		dataItem.Anchor = base64.RawURLEncoding.EncodeToString(buffer[1:])
	}

	return nil
}

func (d *Decoder) decodeDataItemTags(dataItem *DataItem) error {
	buffer := make([]byte, 8)

	if _, err := io.ReadFull(d.buffer, buffer); err != nil {
		return err
	}

	dataItem.TagsNumber = binary.LittleEndian.Uint16(buffer)

	if _, err := io.ReadFull(d.buffer, buffer); err != nil {
		return err
	}

	dataItem.TagsBytes = binary.LittleEndian.Uint16(buffer)

	buffer = make([]byte, dataItem.TagsBytes)

	if _, err := io.ReadFull(d.buffer, buffer); err != nil {
		return err
	}

	return avro.Unmarshal(d.schema, buffer, &dataItem.Tags)
}

func (d *Decoder) lookupDataItemSignatureType(dataItem *DataItem) (*SignatureType, bool) {
	signatureType, exist := signatureTypeMap[dataItem.SignatureType]
	if !exist {
		return nil, false
	}

	return &signatureType, true
}

func NewDecoder(reader io.Reader) *Decoder {
	decoder := Decoder{
		reader: reader,
		schema: lo.Must(avro.Parse(`{"type":"array","items":{"type":"record","name":"Tag","fields":[{"name":"name","type":"bytes"},{"name":"value","type":"bytes"}]}}`)),
	}

	return &decoder
}

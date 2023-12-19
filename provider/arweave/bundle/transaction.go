package bundle

import "io"

type Transaction struct {
	Header    Header
	DataItems []DataItem
}

type Header struct {
	Numbers       uint32
	DataItemInfos []DataItemInfo
}

type DataItemInfo struct {
	Size uint32
	ID   string
}

var _ io.Reader = (*DataItem)(nil)

type DataItem struct {
	SignatureType uint16
	Signature     string
	Owner         string
	Target        string
	Anchor        string
	TagsNumber    uint16
	TagsBytes     uint16
	Tags          []Tag
	Data          []byte

	dataReader io.Reader
}

func (d *DataItem) Read(buffer []byte) (int, error) {
	return d.dataReader.Read(buffer)
}

type Tag struct {
	Name  []byte `avro:"name"`
	Value []byte `avro:"value"`
}

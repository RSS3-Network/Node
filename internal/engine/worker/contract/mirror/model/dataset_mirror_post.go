package model

type DatasetMirrorPostTransformer interface {
	Import(post *DatasetMirrorPost) error
	Export() (*DatasetMirrorPost, error)
}

type DatasetMirrorPost struct {
	TransactionID        string
	OriginContentDigital string
}

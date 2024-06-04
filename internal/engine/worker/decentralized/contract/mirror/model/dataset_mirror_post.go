package model

// DatasetMirrorPostTransformer interface that wraps the basic Import and Export methods.
type DatasetMirrorPostTransformer interface {
	Import(post *DatasetMirrorPost) error
	Export() (*DatasetMirrorPost, error)
}

// DatasetMirrorPost represents a mirror post for revise logic check.
type DatasetMirrorPost struct {
	TransactionID       string
	OriginContentDigest string
}

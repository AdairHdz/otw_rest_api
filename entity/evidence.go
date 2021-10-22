package entity

type Evidence struct {
	EntityUUID
	FileName string
	FileExtension string
	ReviewID string
	Review Review
}
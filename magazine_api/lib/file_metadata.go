package lib

import "github.com/aws/aws-sdk-go-v2/feature/s3/manager"

// UploadMetadata metadata received after uploading file
type UploadMetadata struct {
	FieldName      string
	URL            string
	FileName       string
	FileUID        string
	Size           int64
	UploadResponse *manager.UploadOutput
}

type UploadedFiles []UploadMetadata

func (f UploadedFiles) GetFile(fieldName string) UploadMetadata {
	for _, file := range f {
		if file.FieldName == fieldName {
			return file
		}
	}
	return UploadMetadata{}
}

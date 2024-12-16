package dto

import "mime/multipart"

// FileFormRequest handles file uploads.
// `swaggerignore:"true"` excludes it from Swagger docs.
type FileFormRequest struct {
	File *multipart.FileHeader `form:"file" binding:"required" swaggerignore:"true"`
}

// UploadFileRequest extends FileFormRequest with a required description.
type UploadFileRequest struct {
	FileFormRequest
	Description string `form:"desciption" binding:"required"`
}

// CreateFileRequest is used to create a new file record.
type CreateFileRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Directory   string `json:"directory"`
	MediaType   string `json:"mediaType"`
}

// UpdateFileRequest updates a file's description.
type UpdateFileRequest struct {
	Description string `json:"description"`
}

// FileResponse represents the response for a file record.
type FileResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Directory   string `json:"directory"`
	MediaType   string `json:"mediaType"`
}

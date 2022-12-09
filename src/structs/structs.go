package structs

// ScopeInformation contains the information about the scope for this service
type ScopeInformation struct {
	JSONSchema       string `json:"$schema"`
	ScopeName        string `json:"name"`
	ScopeDescription string `json:"description"`
	ScopeValue       string `json:"scopeStringValue"`
}

// RequestError contains all information about an error which shall be sent back to the client
type RequestError struct {
	HttpStatus       int    `json:"httpCode"`
	HttpError        string `json:"httpError"`
	ErrorCode        string `json:"error"`
	ErrorTitle       string `json:"errorName"`
	ErrorDescription string `json:"errorDescription"`
}

type FileInformation struct {
	// UploadTime specifies at which time the upload to the server was started
	UploadTime float64 `json:"uploadTime"`
	// Name contains the file name which has been set by the user uploading the file
	Name string `json:"fileName"`
	// MIMEType contains the MIME type of the file
	MIMEType string `json:"mimeType"`
	// Size contains the file size in bytes
	Size int `json:"size"`
	// Path contains the path under which the file is stored
	Path string `json:"path"`
	// DownloadPath contains the path under which the file may be accessed by other users
	DownloadPath string `json:"downloadPath"`
}

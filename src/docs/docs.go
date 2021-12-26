package main

// swagger:parameters key
type recordIdWrapper struct {
	// The key of the record for which the operation relates
	// in: path
	// required: true
	KEY string `json:"string"`
}

// swagger:parameters value
type recordValueWrapper struct {
	// The key of the record for which the operation relates
	// in: body
	// required: true
	VALUE string `"string"`
}

// swagger:response noContentResponse
type noContentResponseWrapper struct {
}

// swagger:response commonResponse
type commonResponseWrapper struct {
	// The message
	// in: body
	Body struct {
		// The response message
		//
		// Required: true
		// Example: Given key not found
		Message string
	}
}

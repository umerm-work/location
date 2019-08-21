package io


// Response helper method
func SuccessMessage(data interface{}, msg ... string) Response {
	newMessage := "Success"
	if len(msg) > 0 {
		newMessage = msg[0]
	}
	return Response{
		Data:    data,
		Success: true,
		Message: newMessage,
	}
}
func FailureMessage(err ErrorResponse, msg ... string) Response {
	newMessage := "Failure"
	if len(msg) > 0 {
		newMessage = msg[0]
	}
	return Response{
		Success: false,
		Message: newMessage,
		Error:   err,
	}
}

func ErrorMessage(usserMesg string, err error) ErrorResponse {
	return ErrorResponse{
		UserMessage:     usserMesg,
		InternalMessage: err,
	}
}

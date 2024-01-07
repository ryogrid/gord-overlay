// Code generated by ogen, DO NOT EDIT.

package api_internal

import (
	"fmt"
)

func (s *ErrorStatusCode) Error() string {
	return fmt.Sprintf("code %d: %+v", s.StatusCode, s.Response)
}

// Represents error object.
// Ref: #/definitions/Error
type Error struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

// GetCode returns the value of Code.
func (s *Error) GetCode() int64 {
	return s.Code
}

// GetMessage returns the value of Message.
func (s *Error) GetMessage() string {
	return s.Message
}

// SetCode sets the value of Code.
func (s *Error) SetCode(val int64) {
	s.Code = val
}

// SetMessage sets the value of Message.
func (s *Error) SetMessage(val string) {
	s.Message = val
}

// ErrorStatusCode wraps Error with StatusCode.
type ErrorStatusCode struct {
	StatusCode int
	Response   Error
}

// GetStatusCode returns the value of StatusCode.
func (s *ErrorStatusCode) GetStatusCode() int {
	return s.StatusCode
}

// GetResponse returns the value of Response.
func (s *ErrorStatusCode) GetResponse() Error {
	return s.Response
}

// SetStatusCode sets the value of StatusCode.
func (s *ErrorStatusCode) SetStatusCode(val int) {
	s.StatusCode = val
}

// SetResponse sets the value of Response.
func (s *ErrorStatusCode) SetResponse(val Error) {
	s.Response = val
}

// InternalServiceDeleteValueInnerOK is response for InternalServiceDeleteValueInner operation.
type InternalServiceDeleteValueInnerOK struct{}

// NewOptBool returns new OptBool with value set to v.
func NewOptBool(v bool) OptBool {
	return OptBool{
		Value: v,
		Set:   true,
	}
}

// OptBool is optional bool.
type OptBool struct {
	Value bool
	Set   bool
}

// IsSet returns true if OptBool was set.
func (o OptBool) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptBool) Reset() {
	var v bool
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptBool) SetTo(v bool) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptBool) Get() (v bool, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptBool) Or(d bool) bool {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// Ref: #/definitions/serverGetValueInnerResponse
type ServerGetValueInnerResponse struct {
	Value   OptString `json:"value"`
	Success OptBool   `json:"success"`
}

// GetValue returns the value of Value.
func (s *ServerGetValueInnerResponse) GetValue() OptString {
	return s.Value
}

// GetSuccess returns the value of Success.
func (s *ServerGetValueInnerResponse) GetSuccess() OptBool {
	return s.Success
}

// SetValue sets the value of Value.
func (s *ServerGetValueInnerResponse) SetValue(val OptString) {
	s.Value = val
}

// SetSuccess sets the value of Success.
func (s *ServerGetValueInnerResponse) SetSuccess(val OptBool) {
	s.Success = val
}

// Ref: #/definitions/serverNode
type ServerNode struct {
	Host OptString `json:"host"`
}

// GetHost returns the value of Host.
func (s *ServerNode) GetHost() OptString {
	return s.Host
}

// SetHost sets the value of Host.
func (s *ServerNode) SetHost(val OptString) {
	s.Host = val
}

// Ref: #/definitions/serverNodes
type ServerNodes struct {
	Nodes []ServerNode `json:"nodes"`
}

// GetNodes returns the value of Nodes.
func (s *ServerNodes) GetNodes() []ServerNode {
	return s.Nodes
}

// SetNodes sets the value of Nodes.
func (s *ServerNodes) SetNodes(val []ServerNode) {
	s.Nodes = val
}

// Ref: #/definitions/serverPutValueInnerResponse
type ServerPutValueInnerResponse struct {
	Success OptBool `json:"success"`
}

// GetSuccess returns the value of Success.
func (s *ServerPutValueInnerResponse) GetSuccess() OptBool {
	return s.Success
}

// SetSuccess sets the value of Success.
func (s *ServerPutValueInnerResponse) SetSuccess(val OptBool) {
	s.Success = val
}

// Ref: #/definitions/serverSuccessResponse
type ServerSuccessResponse struct {
	Success OptBool `json:"success"`
}

// GetSuccess returns the value of Success.
func (s *ServerSuccessResponse) GetSuccess() OptBool {
	return s.Success
}

// SetSuccess sets the value of Success.
func (s *ServerSuccessResponse) SetSuccess(val OptBool) {
	s.Success = val
}

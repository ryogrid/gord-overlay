// Code generated by ogen, DO NOT EDIT.

package api_internal

// InternalServiceDeleteValueInnerOK is response for InternalServiceDeleteValueInner operation.
type InternalServiceDeleteValueInnerOK struct{}

// InternalServiceFindClosestPrecedingNodeOK is response for InternalServiceFindClosestPrecedingNode operation.
type InternalServiceFindClosestPrecedingNodeOK struct{}

// InternalServiceFindSuccessorByListOK is response for InternalServiceFindSuccessorByList operation.
type InternalServiceFindSuccessorByListOK struct{}

// InternalServiceFindSuccessorByTableOK is response for InternalServiceFindSuccessorByTable operation.
type InternalServiceFindSuccessorByTableOK struct{}

// InternalServiceGetValueInnerOK is response for InternalServiceGetValueInner operation.
type InternalServiceGetValueInnerOK struct{}

// InternalServiceNotifyOK is response for InternalServiceNotify operation.
type InternalServiceNotifyOK struct{}

// InternalServicePingOK is response for InternalServicePing operation.
type InternalServicePingOK struct{}

// InternalServicePredecessorOK is response for InternalServicePredecessor operation.
type InternalServicePredecessorOK struct{}

// InternalServicePutValueInnerOK is response for InternalServicePutValueInner operation.
type InternalServicePutValueInnerOK struct{}

// InternalServiceSuccessorsOK is response for InternalServiceSuccessors operation.
type InternalServiceSuccessorsOK struct{}

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

package perl

import "encoding/json"

type PerlFunction[Result any] struct {
	params      P
	returnValue Result
}

// custom bool (for unmarshalling)
type Bool bool

func (b *Bool) UnmarshalJSON(data []byte) error {
	var intValue int
	if err := json.Unmarshal(data, &intValue); err != nil {
		return err
	}

	*b = Bool(intValue == 1)
	return nil
}

type Void int

// Params
type P map[string]interface{}

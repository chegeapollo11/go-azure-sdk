package entityqueries

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomEntityQuery interface {
}

func unmarshalCustomEntityQueryImplementation(input []byte) (CustomEntityQuery, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling CustomEntityQuery into map[string]interface: %+v", err)
	}

	value, ok := temp["kind"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "Activity") {
		var out ActivityCustomEntityQuery
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ActivityCustomEntityQuery: %+v", err)
		}
		return out, nil
	}

	type RawCustomEntityQueryImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawCustomEntityQueryImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}

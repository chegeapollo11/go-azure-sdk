package checkdataconnectorrequirements

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataConnectorsCheckRequirements interface {
}

func unmarshalDataConnectorsCheckRequirementsImplementation(input []byte) (DataConnectorsCheckRequirements, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DataConnectorsCheckRequirements into map[string]interface: %+v", err)
	}

	value, ok := temp["kind"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "AzureActiveDirectory") {
		var out AADCheckRequirements
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AADCheckRequirements: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "AzureAdvancedThreatProtection") {
		var out AATPCheckRequirements
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AATPCheckRequirements: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "AzureSecurityCenter") {
		var out ASCCheckRequirements
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ASCCheckRequirements: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "AmazonWebServicesCloudTrail") {
		var out AwsCloudTrailCheckRequirements
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AwsCloudTrailCheckRequirements: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "AmazonWebServicesS3") {
		var out AwsS3CheckRequirements
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AwsS3CheckRequirements: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Dynamics365") {
		var out Dynamics365CheckRequirements
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Dynamics365CheckRequirements: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "IOT") {
		var out IoTCheckRequirements
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IoTCheckRequirements: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "MicrosoftCloudAppSecurity") {
		var out MCASCheckRequirements
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MCASCheckRequirements: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "MicrosoftDefenderAdvancedThreatProtection") {
		var out MDATPCheckRequirements
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MDATPCheckRequirements: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "MicrosoftThreatIntelligence") {
		var out MSTICheckRequirements
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MSTICheckRequirements: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "MicrosoftThreatProtection") {
		var out MtpCheckRequirements
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MtpCheckRequirements: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Office365Project") {
		var out Office365ProjectCheckRequirements
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Office365ProjectCheckRequirements: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "OfficeATP") {
		var out OfficeATPCheckRequirements
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OfficeATPCheckRequirements: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "OfficeIRM") {
		var out OfficeIRMCheckRequirements
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OfficeIRMCheckRequirements: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "OfficePowerBI") {
		var out OfficePowerBICheckRequirements
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OfficePowerBICheckRequirements: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "ThreatIntelligence") {
		var out TICheckRequirements
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TICheckRequirements: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "ThreatIntelligenceTaxii") {
		var out TiTaxiiCheckRequirements
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TiTaxiiCheckRequirements: %+v", err)
		}
		return out, nil
	}

	type RawDataConnectorsCheckRequirementsImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawDataConnectorsCheckRequirementsImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}

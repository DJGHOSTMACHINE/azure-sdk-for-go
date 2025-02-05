//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatabox

import "encoding/json"

func unmarshalCopyLogDetailsClassification(rawMsg json.RawMessage) (CopyLogDetailsClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b CopyLogDetailsClassification
	switch m["copyLogDetailsType"] {
	case string(ClassDiscriminatorDataBox):
		b = &DataBoxAccountCopyLogDetails{}
	case string(ClassDiscriminatorDataBoxCustomerDisk):
		b = &DataBoxCustomerDiskCopyLogDetails{}
	case string(ClassDiscriminatorDataBoxDisk):
		b = &DataBoxDiskCopyLogDetails{}
	case string(ClassDiscriminatorDataBoxHeavy):
		b = &DataBoxHeavyAccountCopyLogDetails{}
	default:
		b = &CopyLogDetails{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalCopyLogDetailsClassificationArray(rawMsg json.RawMessage) ([]CopyLogDetailsClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]CopyLogDetailsClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalCopyLogDetailsClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalCopyLogDetailsClassificationMap(rawMsg json.RawMessage) (map[string]CopyLogDetailsClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages map[string]json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fMap := make(map[string]CopyLogDetailsClassification, len(rawMessages))
	for key, rawMessage := range rawMessages {
		f, err := unmarshalCopyLogDetailsClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fMap[key] = f
	}
	return fMap, nil
}

func unmarshalDataAccountDetailsClassification(rawMsg json.RawMessage) (DataAccountDetailsClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b DataAccountDetailsClassification
	switch m["dataAccountType"] {
	case string(DataAccountTypeManagedDisk):
		b = &ManagedDiskDetails{}
	case string(DataAccountTypeStorageAccount):
		b = &StorageAccountDetails{}
	default:
		b = &DataAccountDetails{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalDataAccountDetailsClassificationArray(rawMsg json.RawMessage) ([]DataAccountDetailsClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]DataAccountDetailsClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalDataAccountDetailsClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalDataAccountDetailsClassificationMap(rawMsg json.RawMessage) (map[string]DataAccountDetailsClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages map[string]json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fMap := make(map[string]DataAccountDetailsClassification, len(rawMessages))
	for key, rawMessage := range rawMessages {
		f, err := unmarshalDataAccountDetailsClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fMap[key] = f
	}
	return fMap, nil
}

func unmarshalDatacenterAddressResponseClassification(rawMsg json.RawMessage) (DatacenterAddressResponseClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b DatacenterAddressResponseClassification
	switch m["datacenterAddressType"] {
	case string(DatacenterAddressTypeDatacenterAddressInstruction):
		b = &DatacenterAddressInstructionResponse{}
	case string(DatacenterAddressTypeDatacenterAddressLocation):
		b = &DatacenterAddressLocationResponse{}
	default:
		b = &DatacenterAddressResponse{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalDatacenterAddressResponseClassificationArray(rawMsg json.RawMessage) ([]DatacenterAddressResponseClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]DatacenterAddressResponseClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalDatacenterAddressResponseClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalDatacenterAddressResponseClassificationMap(rawMsg json.RawMessage) (map[string]DatacenterAddressResponseClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages map[string]json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fMap := make(map[string]DatacenterAddressResponseClassification, len(rawMessages))
	for key, rawMessage := range rawMessages {
		f, err := unmarshalDatacenterAddressResponseClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fMap[key] = f
	}
	return fMap, nil
}

func unmarshalJobDetailsClassification(rawMsg json.RawMessage) (JobDetailsClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b JobDetailsClassification
	switch m["jobDetailsType"] {
	case string(ClassDiscriminatorDataBox):
		b = &DataBoxJobDetails{}
	case string(ClassDiscriminatorDataBoxCustomerDisk):
		b = &DataBoxCustomerDiskJobDetails{}
	case string(ClassDiscriminatorDataBoxDisk):
		b = &DataBoxDiskJobDetails{}
	case string(ClassDiscriminatorDataBoxHeavy):
		b = &DataBoxHeavyJobDetails{}
	default:
		b = &JobDetails{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalJobDetailsClassificationArray(rawMsg json.RawMessage) ([]JobDetailsClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]JobDetailsClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalJobDetailsClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalJobDetailsClassificationMap(rawMsg json.RawMessage) (map[string]JobDetailsClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages map[string]json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fMap := make(map[string]JobDetailsClassification, len(rawMessages))
	for key, rawMessage := range rawMessages {
		f, err := unmarshalJobDetailsClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fMap[key] = f
	}
	return fMap, nil
}

func unmarshalJobSecretsClassification(rawMsg json.RawMessage) (JobSecretsClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b JobSecretsClassification
	switch m["jobSecretsType"] {
	case string(ClassDiscriminatorDataBox):
		b = &DataboxJobSecrets{}
	case string(ClassDiscriminatorDataBoxCustomerDisk):
		b = &CustomerDiskJobSecrets{}
	case string(ClassDiscriminatorDataBoxDisk):
		b = &DataBoxDiskJobSecrets{}
	case string(ClassDiscriminatorDataBoxHeavy):
		b = &DataBoxHeavyJobSecrets{}
	default:
		b = &JobSecrets{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalJobSecretsClassificationArray(rawMsg json.RawMessage) ([]JobSecretsClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]JobSecretsClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalJobSecretsClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalJobSecretsClassificationMap(rawMsg json.RawMessage) (map[string]JobSecretsClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages map[string]json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fMap := make(map[string]JobSecretsClassification, len(rawMessages))
	for key, rawMessage := range rawMessages {
		f, err := unmarshalJobSecretsClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fMap[key] = f
	}
	return fMap, nil
}

func unmarshalScheduleAvailabilityRequestClassification(rawMsg json.RawMessage) (ScheduleAvailabilityRequestClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ScheduleAvailabilityRequestClassification
	switch m["skuName"] {
	case string(SKUNameDataBox):
		b = &DataBoxScheduleAvailabilityRequest{}
	case string(SKUNameDataBoxDisk):
		b = &DiskScheduleAvailabilityRequest{}
	case string(SKUNameDataBoxHeavy):
		b = &HeavyScheduleAvailabilityRequest{}
	default:
		b = &ScheduleAvailabilityRequest{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalScheduleAvailabilityRequestClassificationArray(rawMsg json.RawMessage) ([]ScheduleAvailabilityRequestClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]ScheduleAvailabilityRequestClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalScheduleAvailabilityRequestClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalScheduleAvailabilityRequestClassificationMap(rawMsg json.RawMessage) (map[string]ScheduleAvailabilityRequestClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages map[string]json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fMap := make(map[string]ScheduleAvailabilityRequestClassification, len(rawMessages))
	for key, rawMessage := range rawMessages {
		f, err := unmarshalScheduleAvailabilityRequestClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fMap[key] = f
	}
	return fMap, nil
}

func unmarshalValidationInputRequestClassification(rawMsg json.RawMessage) (ValidationInputRequestClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ValidationInputRequestClassification
	switch m["validationType"] {
	case string(ValidationInputDiscriminatorValidateAddress):
		b = &ValidateAddress{}
	case string(ValidationInputDiscriminatorValidateCreateOrderLimit):
		b = &CreateOrderLimitForSubscriptionValidationRequest{}
	case string(ValidationInputDiscriminatorValidateDataTransferDetails):
		b = &DataTransferDetailsValidationRequest{}
	case string(ValidationInputDiscriminatorValidatePreferences):
		b = &PreferencesValidationRequest{}
	case string(ValidationInputDiscriminatorValidateSKUAvailability):
		b = &SKUAvailabilityValidationRequest{}
	case string(ValidationInputDiscriminatorValidateSubscriptionIsAllowedToCreateJob):
		b = &SubscriptionIsAllowedToCreateJobValidationRequest{}
	default:
		b = &ValidationInputRequest{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalValidationInputRequestClassificationArray(rawMsg json.RawMessage) ([]ValidationInputRequestClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]ValidationInputRequestClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalValidationInputRequestClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalValidationInputRequestClassificationMap(rawMsg json.RawMessage) (map[string]ValidationInputRequestClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages map[string]json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fMap := make(map[string]ValidationInputRequestClassification, len(rawMessages))
	for key, rawMessage := range rawMessages {
		f, err := unmarshalValidationInputRequestClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fMap[key] = f
	}
	return fMap, nil
}

func unmarshalValidationInputResponseClassification(rawMsg json.RawMessage) (ValidationInputResponseClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ValidationInputResponseClassification
	switch m["validationType"] {
	case string(ValidationInputDiscriminatorValidateAddress):
		b = &AddressValidationProperties{}
	case string(ValidationInputDiscriminatorValidateCreateOrderLimit):
		b = &CreateOrderLimitForSubscriptionValidationResponseProperties{}
	case string(ValidationInputDiscriminatorValidateDataTransferDetails):
		b = &DataTransferDetailsValidationResponseProperties{}
	case string(ValidationInputDiscriminatorValidatePreferences):
		b = &PreferencesValidationResponseProperties{}
	case string(ValidationInputDiscriminatorValidateSKUAvailability):
		b = &SKUAvailabilityValidationResponseProperties{}
	case string(ValidationInputDiscriminatorValidateSubscriptionIsAllowedToCreateJob):
		b = &SubscriptionIsAllowedToCreateJobValidationResponseProperties{}
	default:
		b = &ValidationInputResponse{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalValidationInputResponseClassificationArray(rawMsg json.RawMessage) ([]ValidationInputResponseClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]ValidationInputResponseClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalValidationInputResponseClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalValidationInputResponseClassificationMap(rawMsg json.RawMessage) (map[string]ValidationInputResponseClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages map[string]json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fMap := make(map[string]ValidationInputResponseClassification, len(rawMessages))
	for key, rawMessage := range rawMessages {
		f, err := unmarshalValidationInputResponseClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fMap[key] = f
	}
	return fMap, nil
}

func unmarshalValidationRequestClassification(rawMsg json.RawMessage) (ValidationRequestClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ValidationRequestClassification
	switch m["validationCategory"] {
	case "JobCreationValidation":
		b = &CreateJobValidations{}
	default:
		b = &ValidationRequest{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalValidationRequestClassificationArray(rawMsg json.RawMessage) ([]ValidationRequestClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]ValidationRequestClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalValidationRequestClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalValidationRequestClassificationMap(rawMsg json.RawMessage) (map[string]ValidationRequestClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages map[string]json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fMap := make(map[string]ValidationRequestClassification, len(rawMessages))
	for key, rawMessage := range rawMessages {
		f, err := unmarshalValidationRequestClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fMap[key] = f
	}
	return fMap, nil
}

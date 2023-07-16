package schema

import (
	"encoding/json"
	"fmt"
)

func (dm *DefinitionMap) UnmarshalJSON(data []byte) (err error) {
	// Definitions must be LexiconPrimaryTypes.
	temp := make(map[string]json.RawMessage, 2)
	err = json.Unmarshal(data, &temp)
	if err != nil {
		return
	}
	*dm = make(DefinitionMap)

	for k, v := range temp {
		// peek at the type.
		tempValue := map[string]interface{}{}
		if err := json.Unmarshal(v, &tempValue); err != nil {
			return err
		}
		if t, ok := tempValue["type"]; ok {
			ts, ok := t.(string)
			if !ok {
				return fmt.Errorf("type of %s was not a string: %T", k, t)
			}
			concreteType := NewDefinition(SchemaType(ts))
			if concreteType == nil {
				return fmt.Errorf("type of %s isn't known as a definition: %s", k, ts)
			}

			if err = json.Unmarshal(v, concreteType); err != nil {
				return fmt.Errorf("%s: %w", k, err)
			}
			(*dm)[k] = &concreteType
		} else {
			fmt.Printf("key %s is not typed. has value %v\n", k, tempValue)
		}
	}

	return
}

func (fm *FieldMap) UnmarshalJSON(data []byte) (err error) {
	// Properties must be LexiconFieldTypes.
	temp := make(map[string]json.RawMessage, 2)
	err = json.Unmarshal(data, &temp)
	if err != nil {
		return
	}
	*fm = make(FieldMap)

	for k, v := range temp {
		// peek at the type.
		tempValue := map[string]interface{}{}
		if err := json.Unmarshal(v, &tempValue); err != nil {
			return err
		}
		if t, ok := tempValue["type"]; ok {
			ts, ok := t.(string)
			if !ok {
				return fmt.Errorf("type of %s was not a string: %T", k, t)
			}
			concreteType := NewField(SchemaType(ts))
			if concreteType == nil {
				return fmt.Errorf("type of %s isn't known as a field: %s", k, ts)
			}

			if err = json.Unmarshal(v, concreteType); err != nil {
				return fmt.Errorf("%s: %w", k, err)
			}
			(*fm)[k] = &concreteType
		} else {
			fmt.Printf("key %s is not typed. has value %v\n", k, tempValue)
		}
	}

	return
}

func (fd *FieldDefinition) UnmarshalJSON(data []byte) (err error) {
	temp := make(map[string]interface{}, 2)
	err = json.Unmarshal(data, &temp)
	if err != nil {
		return
	}

	if t, ok := temp["type"]; ok {
		ts, ok := t.(string)
		if !ok {
			return fmt.Errorf("type of field was not a string: %T", t)
		}
		concreteType := NewField(SchemaType(ts))
		if concreteType == nil {
			return fmt.Errorf("type of field isn't known as a field: %s", ts)
		}

		if err = json.Unmarshal(data, concreteType); err != nil {
			return fmt.Errorf("items: %w", err)
		}
		fd.Field = concreteType
	} else {
		fmt.Printf("field is not typed. has value %v\n", temp)
	}

	return
}

func (fd *FieldDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(fd.Field)
}

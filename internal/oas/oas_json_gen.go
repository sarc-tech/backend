// Code generated by ogen, DO NOT EDIT.

package oas

import (
	"math/bits"
	"strconv"
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"

	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/validate"
)

// Encode encodes AddIncidentsApplicationJSON as json.
func (s *AddIncidentsApplicationJSON) Encode(e *jx.Encoder) {
	unwrapped := (*Incident)(s)

	unwrapped.Encode(e)
}

// Decode decodes AddIncidentsApplicationJSON from json.
func (s *AddIncidentsApplicationJSON) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode AddIncidentsApplicationJSON to nil")
	}
	var unwrapped Incident
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = AddIncidentsApplicationJSON(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *AddIncidentsApplicationJSON) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *AddIncidentsApplicationJSON) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes AddIncidentsApplicationXWwwFormUrlencoded as json.
func (s *AddIncidentsApplicationXWwwFormUrlencoded) Encode(e *jx.Encoder) {
	unwrapped := (*Incident)(s)

	unwrapped.Encode(e)
}

// Decode decodes AddIncidentsApplicationXWwwFormUrlencoded from json.
func (s *AddIncidentsApplicationXWwwFormUrlencoded) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode AddIncidentsApplicationXWwwFormUrlencoded to nil")
	}
	var unwrapped Incident
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = AddIncidentsApplicationXWwwFormUrlencoded(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *AddIncidentsApplicationXWwwFormUrlencoded) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *AddIncidentsApplicationXWwwFormUrlencoded) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes AddStatusApplicationJSON as json.
func (s *AddStatusApplicationJSON) Encode(e *jx.Encoder) {
	unwrapped := (*Status)(s)

	unwrapped.Encode(e)
}

// Decode decodes AddStatusApplicationJSON from json.
func (s *AddStatusApplicationJSON) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode AddStatusApplicationJSON to nil")
	}
	var unwrapped Status
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = AddStatusApplicationJSON(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *AddStatusApplicationJSON) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *AddStatusApplicationJSON) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes AddStatusApplicationXWwwFormUrlencoded as json.
func (s *AddStatusApplicationXWwwFormUrlencoded) Encode(e *jx.Encoder) {
	unwrapped := (*Status)(s)

	unwrapped.Encode(e)
}

// Decode decodes AddStatusApplicationXWwwFormUrlencoded from json.
func (s *AddStatusApplicationXWwwFormUrlencoded) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode AddStatusApplicationXWwwFormUrlencoded to nil")
	}
	var unwrapped Status
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = AddStatusApplicationXWwwFormUrlencoded(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *AddStatusApplicationXWwwFormUrlencoded) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *AddStatusApplicationXWwwFormUrlencoded) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Incident) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Incident) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("id")
		e.Str(s.ID)
	}
	{
		e.FieldStart("region")
		e.Str(s.Region)
	}
	{
		e.FieldStart("fio")
		e.Str(s.Fio)
	}
	{
		e.FieldStart("statusId")
		e.Str(s.StatusId)
	}
	{
		e.FieldStart("date")
		e.Str(s.Date)
	}
}

var jsonFieldsNameOfIncident = [5]string{
	0: "id",
	1: "region",
	2: "fio",
	3: "statusId",
	4: "date",
}

// Decode decodes Incident from json.
func (s *Incident) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Incident to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.ID = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "region":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.Region = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"region\"")
			}
		case "fio":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				v, err := d.Str()
				s.Fio = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"fio\"")
			}
		case "statusId":
			requiredBitSet[0] |= 1 << 3
			if err := func() error {
				v, err := d.Str()
				s.StatusId = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"statusId\"")
			}
		case "date":
			requiredBitSet[0] |= 1 << 4
			if err := func() error {
				v, err := d.Str()
				s.Date = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"date\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Incident")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00011111,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfIncident) {
					name = jsonFieldsNameOfIncident[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Incident) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Incident) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *IncidentsResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *IncidentsResponse) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("trackingId")
		e.Str(s.TrackingId)
	}
	{
		e.FieldStart("status")
		e.Str(s.Status)
	}
	{
		e.FieldStart("data")
		e.ArrStart()
		for _, elem := range s.Data {
			elem.Encode(e)
		}
		e.ArrEnd()
	}
	{
		if s.Statuses != nil {
			e.FieldStart("statuses")
			e.ArrStart()
			for _, elem := range s.Statuses {
				elem.Encode(e)
			}
			e.ArrEnd()
		}
	}
}

var jsonFieldsNameOfIncidentsResponse = [4]string{
	0: "trackingId",
	1: "status",
	2: "data",
	3: "statuses",
}

// Decode decodes IncidentsResponse from json.
func (s *IncidentsResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode IncidentsResponse to nil")
	}
	var requiredBitSet [1]uint8
	s.setDefaults()

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "trackingId":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.TrackingId = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"trackingId\"")
			}
		case "status":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.Status = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"status\"")
			}
		case "data":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				s.Data = make([]Incident, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem Incident
					if err := elem.Decode(d); err != nil {
						return err
					}
					s.Data = append(s.Data, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"data\"")
			}
		case "statuses":
			if err := func() error {
				s.Statuses = make([]Status, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem Status
					if err := elem.Decode(d); err != nil {
						return err
					}
					s.Statuses = append(s.Statuses, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"statuses\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode IncidentsResponse")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000111,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfIncidentsResponse) {
					name = jsonFieldsNameOfIncidentsResponse[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *IncidentsResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *IncidentsResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes time.Time as json.
func (o OptDate) Encode(e *jx.Encoder, format func(*jx.Encoder, time.Time)) {
	if !o.Set {
		return
	}
	format(e, o.Value)
}

// Decode decodes time.Time from json.
func (o *OptDate) Decode(d *jx.Decoder, format func(*jx.Decoder) (time.Time, error)) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptDate to nil")
	}
	o.Set = true
	v, err := format(d)
	if err != nil {
		return err
	}
	o.Value = v
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptDate) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e, json.EncodeDate)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptDate) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d, json.DecodeDate)
}

// Encode encodes string as json.
func (o OptString) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Str(string(o.Value))
}

// Decode decodes string from json.
func (o *OptString) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptString to nil")
	}
	o.Set = true
	v, err := d.Str()
	if err != nil {
		return err
	}
	o.Value = string(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptString) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptString) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Status) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Status) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("id")
		e.Str(s.ID)
	}
	{
		e.FieldStart("name")
		e.Str(s.Name)
	}
}

var jsonFieldsNameOfStatus = [2]string{
	0: "id",
	1: "name",
}

// Decode decodes Status from json.
func (s *Status) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Status to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.ID = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "name":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.Name = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"name\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Status")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000011,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfStatus) {
					name = jsonFieldsNameOfStatus[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Status) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Status) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *StatususResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *StatususResponse) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("trackingId")
		e.Str(s.TrackingId)
	}
	{
		e.FieldStart("status")
		e.Str(s.Status)
	}
	{
		e.FieldStart("data")
		e.ArrStart()
		for _, elem := range s.Data {
			elem.Encode(e)
		}
		e.ArrEnd()
	}
}

var jsonFieldsNameOfStatususResponse = [3]string{
	0: "trackingId",
	1: "status",
	2: "data",
}

// Decode decodes StatususResponse from json.
func (s *StatususResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode StatususResponse to nil")
	}
	var requiredBitSet [1]uint8
	s.setDefaults()

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "trackingId":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.TrackingId = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"trackingId\"")
			}
		case "status":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.Status = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"status\"")
			}
		case "data":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				s.Data = make([]Status, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem Status
					if err := elem.Decode(d); err != nil {
						return err
					}
					s.Data = append(s.Data, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"data\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode StatususResponse")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000111,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfStatususResponse) {
					name = jsonFieldsNameOfStatususResponse[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *StatususResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *StatususResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes UpdateIncidentsApplicationJSON as json.
func (s *UpdateIncidentsApplicationJSON) Encode(e *jx.Encoder) {
	unwrapped := (*Incident)(s)

	unwrapped.Encode(e)
}

// Decode decodes UpdateIncidentsApplicationJSON from json.
func (s *UpdateIncidentsApplicationJSON) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode UpdateIncidentsApplicationJSON to nil")
	}
	var unwrapped Incident
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = UpdateIncidentsApplicationJSON(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *UpdateIncidentsApplicationJSON) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *UpdateIncidentsApplicationJSON) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes UpdateIncidentsApplicationXWwwFormUrlencoded as json.
func (s *UpdateIncidentsApplicationXWwwFormUrlencoded) Encode(e *jx.Encoder) {
	unwrapped := (*Incident)(s)

	unwrapped.Encode(e)
}

// Decode decodes UpdateIncidentsApplicationXWwwFormUrlencoded from json.
func (s *UpdateIncidentsApplicationXWwwFormUrlencoded) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode UpdateIncidentsApplicationXWwwFormUrlencoded to nil")
	}
	var unwrapped Incident
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = UpdateIncidentsApplicationXWwwFormUrlencoded(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *UpdateIncidentsApplicationXWwwFormUrlencoded) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *UpdateIncidentsApplicationXWwwFormUrlencoded) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes UpdateStatusApplicationJSON as json.
func (s *UpdateStatusApplicationJSON) Encode(e *jx.Encoder) {
	unwrapped := (*Status)(s)

	unwrapped.Encode(e)
}

// Decode decodes UpdateStatusApplicationJSON from json.
func (s *UpdateStatusApplicationJSON) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode UpdateStatusApplicationJSON to nil")
	}
	var unwrapped Status
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = UpdateStatusApplicationJSON(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *UpdateStatusApplicationJSON) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *UpdateStatusApplicationJSON) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes UpdateStatusApplicationXWwwFormUrlencoded as json.
func (s *UpdateStatusApplicationXWwwFormUrlencoded) Encode(e *jx.Encoder) {
	unwrapped := (*Status)(s)

	unwrapped.Encode(e)
}

// Decode decodes UpdateStatusApplicationXWwwFormUrlencoded from json.
func (s *UpdateStatusApplicationXWwwFormUrlencoded) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode UpdateStatusApplicationXWwwFormUrlencoded to nil")
	}
	var unwrapped Status
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = UpdateStatusApplicationXWwwFormUrlencoded(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *UpdateStatusApplicationXWwwFormUrlencoded) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *UpdateStatusApplicationXWwwFormUrlencoded) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *User) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *User) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("id")
		e.Str(s.ID)
	}
	{
		e.FieldStart("yandexId")
		e.Str(s.YandexId)
	}
	{
		if s.Surname.Set {
			e.FieldStart("surname")
			s.Surname.Encode(e)
		}
	}
	{
		e.FieldStart("name")
		e.Str(s.Name)
	}
	{
		if s.Patronymic.Set {
			e.FieldStart("patronymic")
			s.Patronymic.Encode(e)
		}
	}
	{
		if s.CallSign.Set {
			e.FieldStart("callSign")
			s.CallSign.Encode(e)
		}
	}
	{
		if s.Gender.Set {
			e.FieldStart("gender")
			s.Gender.Encode(e)
		}
	}
	{
		if s.Birthdate.Set {
			e.FieldStart("birthdate")
			s.Birthdate.Encode(e, json.EncodeDate)
		}
	}
	{
		if s.Vk.Set {
			e.FieldStart("vk")
			s.Vk.Encode(e)
		}
	}
	{
		if s.Telegram.Set {
			e.FieldStart("telegram")
			s.Telegram.Encode(e)
		}
	}
	{
		if s.Email.Set {
			e.FieldStart("email")
			s.Email.Encode(e)
		}
	}
	{
		if s.Phone.Set {
			e.FieldStart("phone")
			s.Phone.Encode(e)
		}
	}
	{
		e.FieldStart("approval")
		e.Bool(s.Approval)
	}
	{
		e.FieldStart("role")
		e.Str(s.Role)
	}
}

var jsonFieldsNameOfUser = [14]string{
	0:  "id",
	1:  "yandexId",
	2:  "surname",
	3:  "name",
	4:  "patronymic",
	5:  "callSign",
	6:  "gender",
	7:  "birthdate",
	8:  "vk",
	9:  "telegram",
	10: "email",
	11: "phone",
	12: "approval",
	13: "role",
}

// Decode decodes User from json.
func (s *User) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode User to nil")
	}
	var requiredBitSet [2]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.ID = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "yandexId":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.YandexId = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"yandexId\"")
			}
		case "surname":
			if err := func() error {
				s.Surname.Reset()
				if err := s.Surname.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"surname\"")
			}
		case "name":
			requiredBitSet[0] |= 1 << 3
			if err := func() error {
				v, err := d.Str()
				s.Name = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"name\"")
			}
		case "patronymic":
			if err := func() error {
				s.Patronymic.Reset()
				if err := s.Patronymic.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"patronymic\"")
			}
		case "callSign":
			if err := func() error {
				s.CallSign.Reset()
				if err := s.CallSign.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"callSign\"")
			}
		case "gender":
			if err := func() error {
				s.Gender.Reset()
				if err := s.Gender.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"gender\"")
			}
		case "birthdate":
			if err := func() error {
				s.Birthdate.Reset()
				if err := s.Birthdate.Decode(d, json.DecodeDate); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"birthdate\"")
			}
		case "vk":
			if err := func() error {
				s.Vk.Reset()
				if err := s.Vk.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"vk\"")
			}
		case "telegram":
			if err := func() error {
				s.Telegram.Reset()
				if err := s.Telegram.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"telegram\"")
			}
		case "email":
			if err := func() error {
				s.Email.Reset()
				if err := s.Email.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"email\"")
			}
		case "phone":
			if err := func() error {
				s.Phone.Reset()
				if err := s.Phone.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"phone\"")
			}
		case "approval":
			requiredBitSet[1] |= 1 << 4
			if err := func() error {
				v, err := d.Bool()
				s.Approval = bool(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"approval\"")
			}
		case "role":
			requiredBitSet[1] |= 1 << 5
			if err := func() error {
				v, err := d.Str()
				s.Role = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"role\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode User")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [2]uint8{
		0b00001011,
		0b00110000,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfUser) {
					name = jsonFieldsNameOfUser[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *User) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *User) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *UsersResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *UsersResponse) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("trackingId")
		e.Str(s.TrackingId)
	}
	{
		e.FieldStart("status")
		e.Str(s.Status)
	}
	{
		e.FieldStart("data")
		e.ArrStart()
		for _, elem := range s.Data {
			elem.Encode(e)
		}
		e.ArrEnd()
	}
}

var jsonFieldsNameOfUsersResponse = [3]string{
	0: "trackingId",
	1: "status",
	2: "data",
}

// Decode decodes UsersResponse from json.
func (s *UsersResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode UsersResponse to nil")
	}
	var requiredBitSet [1]uint8
	s.setDefaults()

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "trackingId":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.TrackingId = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"trackingId\"")
			}
		case "status":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.Status = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"status\"")
			}
		case "data":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				s.Data = make([]User, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem User
					if err := elem.Decode(d); err != nil {
						return err
					}
					s.Data = append(s.Data, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"data\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode UsersResponse")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000111,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfUsersResponse) {
					name = jsonFieldsNameOfUsersResponse[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *UsersResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *UsersResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

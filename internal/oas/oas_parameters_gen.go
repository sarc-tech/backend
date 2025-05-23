// Code generated by ogen, DO NOT EDIT.

package oas

import (
	"net/http"
	"net/url"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// DeleteIncidentParams is parameters of deleteIncident operation.
type DeleteIncidentParams struct {
	// Request id to delete.
	IncidentId int
}

func unpackDeleteIncidentParams(packed middleware.Parameters) (params DeleteIncidentParams) {
	{
		key := middleware.ParameterKey{
			Name: "incidentId",
			In:   "path",
		}
		params.IncidentId = packed[key].(int)
	}
	return params
}

func decodeDeleteIncidentParams(args [1]string, argsEscaped bool, r *http.Request) (params DeleteIncidentParams, _ error) {
	// Decode path: incidentId.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "incidentId",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(val)
				if err != nil {
					return err
				}

				params.IncidentId = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "incidentId",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// GetIncidentByIdParams is parameters of getIncidentById operation.
type GetIncidentByIdParams struct {
	// ID of Incidents to return.
	IncidentId int
}

func unpackGetIncidentByIdParams(packed middleware.Parameters) (params GetIncidentByIdParams) {
	{
		key := middleware.ParameterKey{
			Name: "incidentId",
			In:   "path",
		}
		params.IncidentId = packed[key].(int)
	}
	return params
}

func decodeGetIncidentByIdParams(args [1]string, argsEscaped bool, r *http.Request) (params GetIncidentByIdParams, _ error) {
	// Decode path: incidentId.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "incidentId",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(val)
				if err != nil {
					return err
				}

				params.IncidentId = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "incidentId",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

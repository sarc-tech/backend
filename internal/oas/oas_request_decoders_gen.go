// Code generated by ogen, DO NOT EDIT.

package oas

import (
	"io"
	"mime"
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"go.uber.org/multierr"

	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

func (s *Server) decodeAddIncidentsRequest(r *http.Request) (
	req AddIncidentsReq,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = multierr.Append(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = multierr.Append(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "application/json":
		if r.ContentLength == 0 {
			return req, close, validate.ErrBodyRequired
		}
		buf, err := io.ReadAll(r.Body)
		if err != nil {
			return req, close, err
		}

		if len(buf) == 0 {
			return req, close, validate.ErrBodyRequired
		}

		d := jx.DecodeBytes(buf)

		var request AddIncidentsApplicationJSON
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			if err := d.Skip(); err != io.EOF {
				return errors.New("unexpected trailing data")
			}
			return nil
		}(); err != nil {
			err = &ogenerrors.DecodeBodyError{
				ContentType: ct,
				Body:        buf,
				Err:         err,
			}
			return req, close, err
		}
		return &request, close, nil
	case ct == "application/x-www-form-urlencoded":
		if r.ContentLength == 0 {
			return req, close, validate.ErrBodyRequired
		}
		form, err := ht.ParseForm(r)
		if err != nil {
			return req, close, errors.Wrap(err, "parse form")
		}

		var request AddIncidentsApplicationXWwwFormUrlencoded
		{
			var unwrapped Incident
			q := uri.NewQueryDecoder(form)
			{
				cfg := uri.QueryParameterDecodingConfig{
					Name:    "id",
					Style:   uri.QueryStyleForm,
					Explode: true,
				}
				if err := q.HasParam(cfg); err == nil {
					if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						unwrapped.ID = c
						return nil
					}); err != nil {
						return req, close, errors.Wrap(err, "decode \"id\"")
					}
				} else {
					return req, close, errors.Wrap(err, "query")
				}
			}
			{
				cfg := uri.QueryParameterDecodingConfig{
					Name:    "region",
					Style:   uri.QueryStyleForm,
					Explode: true,
				}
				if err := q.HasParam(cfg); err == nil {
					if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						unwrapped.Region = c
						return nil
					}); err != nil {
						return req, close, errors.Wrap(err, "decode \"region\"")
					}
				} else {
					return req, close, errors.Wrap(err, "query")
				}
			}
			{
				cfg := uri.QueryParameterDecodingConfig{
					Name:    "fio",
					Style:   uri.QueryStyleForm,
					Explode: true,
				}
				if err := q.HasParam(cfg); err == nil {
					if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						unwrapped.Fio = c
						return nil
					}); err != nil {
						return req, close, errors.Wrap(err, "decode \"fio\"")
					}
				} else {
					return req, close, errors.Wrap(err, "query")
				}
			}
			{
				cfg := uri.QueryParameterDecodingConfig{
					Name:    "statusId",
					Style:   uri.QueryStyleForm,
					Explode: true,
				}
				if err := q.HasParam(cfg); err == nil {
					if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						unwrapped.StatusId = c
						return nil
					}); err != nil {
						return req, close, errors.Wrap(err, "decode \"statusId\"")
					}
				} else {
					return req, close, errors.Wrap(err, "query")
				}
			}
			{
				cfg := uri.QueryParameterDecodingConfig{
					Name:    "date",
					Style:   uri.QueryStyleForm,
					Explode: true,
				}
				if err := q.HasParam(cfg); err == nil {
					if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						unwrapped.Date = c
						return nil
					}); err != nil {
						return req, close, errors.Wrap(err, "decode \"date\"")
					}
				} else {
					return req, close, errors.Wrap(err, "query")
				}
			}
			request = AddIncidentsApplicationXWwwFormUrlencoded(unwrapped)
		}
		return &request, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeAddStatusRequest(r *http.Request) (
	req AddStatusReq,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = multierr.Append(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = multierr.Append(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "application/json":
		if r.ContentLength == 0 {
			return req, close, validate.ErrBodyRequired
		}
		buf, err := io.ReadAll(r.Body)
		if err != nil {
			return req, close, err
		}

		if len(buf) == 0 {
			return req, close, validate.ErrBodyRequired
		}

		d := jx.DecodeBytes(buf)

		var request AddStatusApplicationJSON
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			if err := d.Skip(); err != io.EOF {
				return errors.New("unexpected trailing data")
			}
			return nil
		}(); err != nil {
			err = &ogenerrors.DecodeBodyError{
				ContentType: ct,
				Body:        buf,
				Err:         err,
			}
			return req, close, err
		}
		return &request, close, nil
	case ct == "application/x-www-form-urlencoded":
		if r.ContentLength == 0 {
			return req, close, validate.ErrBodyRequired
		}
		form, err := ht.ParseForm(r)
		if err != nil {
			return req, close, errors.Wrap(err, "parse form")
		}

		var request AddStatusApplicationXWwwFormUrlencoded
		{
			var unwrapped Status
			q := uri.NewQueryDecoder(form)
			{
				cfg := uri.QueryParameterDecodingConfig{
					Name:    "id",
					Style:   uri.QueryStyleForm,
					Explode: true,
				}
				if err := q.HasParam(cfg); err == nil {
					if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						unwrapped.ID = c
						return nil
					}); err != nil {
						return req, close, errors.Wrap(err, "decode \"id\"")
					}
				} else {
					return req, close, errors.Wrap(err, "query")
				}
			}
			{
				cfg := uri.QueryParameterDecodingConfig{
					Name:    "name",
					Style:   uri.QueryStyleForm,
					Explode: true,
				}
				if err := q.HasParam(cfg); err == nil {
					if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						unwrapped.Name = c
						return nil
					}); err != nil {
						return req, close, errors.Wrap(err, "decode \"name\"")
					}
				} else {
					return req, close, errors.Wrap(err, "query")
				}
			}
			request = AddStatusApplicationXWwwFormUrlencoded(unwrapped)
		}
		return &request, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeUpdateIncidentsRequest(r *http.Request) (
	req UpdateIncidentsReq,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = multierr.Append(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = multierr.Append(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "application/json":
		if r.ContentLength == 0 {
			return req, close, validate.ErrBodyRequired
		}
		buf, err := io.ReadAll(r.Body)
		if err != nil {
			return req, close, err
		}

		if len(buf) == 0 {
			return req, close, validate.ErrBodyRequired
		}

		d := jx.DecodeBytes(buf)

		var request UpdateIncidentsApplicationJSON
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			if err := d.Skip(); err != io.EOF {
				return errors.New("unexpected trailing data")
			}
			return nil
		}(); err != nil {
			err = &ogenerrors.DecodeBodyError{
				ContentType: ct,
				Body:        buf,
				Err:         err,
			}
			return req, close, err
		}
		return &request, close, nil
	case ct == "application/x-www-form-urlencoded":
		if r.ContentLength == 0 {
			return req, close, validate.ErrBodyRequired
		}
		form, err := ht.ParseForm(r)
		if err != nil {
			return req, close, errors.Wrap(err, "parse form")
		}

		var request UpdateIncidentsApplicationXWwwFormUrlencoded
		{
			var unwrapped Incident
			q := uri.NewQueryDecoder(form)
			{
				cfg := uri.QueryParameterDecodingConfig{
					Name:    "id",
					Style:   uri.QueryStyleForm,
					Explode: true,
				}
				if err := q.HasParam(cfg); err == nil {
					if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						unwrapped.ID = c
						return nil
					}); err != nil {
						return req, close, errors.Wrap(err, "decode \"id\"")
					}
				} else {
					return req, close, errors.Wrap(err, "query")
				}
			}
			{
				cfg := uri.QueryParameterDecodingConfig{
					Name:    "region",
					Style:   uri.QueryStyleForm,
					Explode: true,
				}
				if err := q.HasParam(cfg); err == nil {
					if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						unwrapped.Region = c
						return nil
					}); err != nil {
						return req, close, errors.Wrap(err, "decode \"region\"")
					}
				} else {
					return req, close, errors.Wrap(err, "query")
				}
			}
			{
				cfg := uri.QueryParameterDecodingConfig{
					Name:    "fio",
					Style:   uri.QueryStyleForm,
					Explode: true,
				}
				if err := q.HasParam(cfg); err == nil {
					if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						unwrapped.Fio = c
						return nil
					}); err != nil {
						return req, close, errors.Wrap(err, "decode \"fio\"")
					}
				} else {
					return req, close, errors.Wrap(err, "query")
				}
			}
			{
				cfg := uri.QueryParameterDecodingConfig{
					Name:    "statusId",
					Style:   uri.QueryStyleForm,
					Explode: true,
				}
				if err := q.HasParam(cfg); err == nil {
					if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						unwrapped.StatusId = c
						return nil
					}); err != nil {
						return req, close, errors.Wrap(err, "decode \"statusId\"")
					}
				} else {
					return req, close, errors.Wrap(err, "query")
				}
			}
			{
				cfg := uri.QueryParameterDecodingConfig{
					Name:    "date",
					Style:   uri.QueryStyleForm,
					Explode: true,
				}
				if err := q.HasParam(cfg); err == nil {
					if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						unwrapped.Date = c
						return nil
					}); err != nil {
						return req, close, errors.Wrap(err, "decode \"date\"")
					}
				} else {
					return req, close, errors.Wrap(err, "query")
				}
			}
			request = UpdateIncidentsApplicationXWwwFormUrlencoded(unwrapped)
		}
		return &request, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeUpdateStatusRequest(r *http.Request) (
	req UpdateStatusReq,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = multierr.Append(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = multierr.Append(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "application/json":
		if r.ContentLength == 0 {
			return req, close, validate.ErrBodyRequired
		}
		buf, err := io.ReadAll(r.Body)
		if err != nil {
			return req, close, err
		}

		if len(buf) == 0 {
			return req, close, validate.ErrBodyRequired
		}

		d := jx.DecodeBytes(buf)

		var request UpdateStatusApplicationJSON
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			if err := d.Skip(); err != io.EOF {
				return errors.New("unexpected trailing data")
			}
			return nil
		}(); err != nil {
			err = &ogenerrors.DecodeBodyError{
				ContentType: ct,
				Body:        buf,
				Err:         err,
			}
			return req, close, err
		}
		return &request, close, nil
	case ct == "application/x-www-form-urlencoded":
		if r.ContentLength == 0 {
			return req, close, validate.ErrBodyRequired
		}
		form, err := ht.ParseForm(r)
		if err != nil {
			return req, close, errors.Wrap(err, "parse form")
		}

		var request UpdateStatusApplicationXWwwFormUrlencoded
		{
			var unwrapped Status
			q := uri.NewQueryDecoder(form)
			{
				cfg := uri.QueryParameterDecodingConfig{
					Name:    "id",
					Style:   uri.QueryStyleForm,
					Explode: true,
				}
				if err := q.HasParam(cfg); err == nil {
					if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						unwrapped.ID = c
						return nil
					}); err != nil {
						return req, close, errors.Wrap(err, "decode \"id\"")
					}
				} else {
					return req, close, errors.Wrap(err, "query")
				}
			}
			{
				cfg := uri.QueryParameterDecodingConfig{
					Name:    "name",
					Style:   uri.QueryStyleForm,
					Explode: true,
				}
				if err := q.HasParam(cfg); err == nil {
					if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						unwrapped.Name = c
						return nil
					}); err != nil {
						return req, close, errors.Wrap(err, "decode \"name\"")
					}
				} else {
					return req, close, errors.Wrap(err, "query")
				}
			}
			request = UpdateStatusApplicationXWwwFormUrlencoded(unwrapped)
		}
		return &request, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeUpdateUserRequest(r *http.Request) (
	req *User,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = multierr.Append(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = multierr.Append(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "application/json":
		if r.ContentLength == 0 {
			return req, close, validate.ErrBodyRequired
		}
		buf, err := io.ReadAll(r.Body)
		if err != nil {
			return req, close, err
		}

		if len(buf) == 0 {
			return req, close, validate.ErrBodyRequired
		}

		d := jx.DecodeBytes(buf)

		var request User
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			if err := d.Skip(); err != io.EOF {
				return errors.New("unexpected trailing data")
			}
			return nil
		}(); err != nil {
			err = &ogenerrors.DecodeBodyError{
				ContentType: ct,
				Body:        buf,
				Err:         err,
			}
			return req, close, err
		}
		return &request, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}

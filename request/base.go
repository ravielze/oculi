package request

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gofrs/uuid"
	"github.com/ravielze/oculi/common/encoding/radix36"
	"github.com/ravielze/oculi/persistent/sql"
)

type (
	Base struct {
		ctx               context.Context
		db                sql.API
		tx                sql.API
		errors            []error
		responseCode      int
		data              map[string]string
		requestIdentifier uint64
	}
)

func (r *Base) Identifier() uint64 {
	return r.requestIdentifier
}

func (r *Base) GetContext() context.Context {
	return r.ctx
}

func (r *Base) HasError() bool {
	return len(r.errors) > 0
}

func (r *Base) AddError(responseCode int, err ...error) {
	if r.responseCode < 400 {
		r.responseCode = responseCode
	}
	r.errors = append(r.errors, err...)
}

func (r *Base) SetResponseCode(code int) {
	r.responseCode = code
}

func (r *Base) ResponseCode() int {
	return r.responseCode
}
func (r *Base) Error() error {
	if len(r.errors) > 0 {
		return r.errors[0]
	}
	return nil
}

func (r *Base) HasTransaction() bool {
	return r.tx != nil
}

func (r *Base) Transaction() sql.API {
	if r.tx == nil {
		return r.db
	}
	return r.tx
}

func (r *Base) NewTransaction() sql.API {
	r.tx = r.db.Begin()
	return r.tx
}

func (r *Base) CommitTransaction() sql.API {
	if r.tx == nil {
		return r.db
	}

	r.tx.Commit()
	return r.tx
}

func (r *Base) RollbackTransaction() sql.API {
	if r.tx == nil {
		return r.db
	}

	r.tx.Rollback()
	return r.tx
}

func NewBaseWithIdentifier(identifier uint64, db sql.API) Context {
	return &Base{
		ctx:               context.Background(),
		db:                db,
		tx:                nil,
		errors:            make([]error, 0),
		responseCode:      200,
		data:              make(map[string]string, 5),
		requestIdentifier: identifier,
	}
}

func NewBase(db sql.API) Context {
	return &Base{
		ctx:               context.Background(),
		db:                db,
		tx:                nil,
		errors:            make([]error, 0),
		responseCode:      200,
		data:              make(map[string]string, 5),
		requestIdentifier: 0,
	}
}

func (r *Base) SetContext(ctx context.Context) Context {
	r.ctx = ctx
	return r
}

func (r *Base) ParseString(key, value string) Context {
	if !r.HasError() {
		r.data[key] = value
	}
	return r
}
func (r *Base) ParseStringOrDefault(key, value, def string) Context {
	if !r.HasError() {
		if len(value) == 0 || len(strings.TrimSpace(value)) == 0 {
			value = def
		}
		r.data[key] = value
	}
	return r
}

func (r *Base) ParseUUID(key, value string) Context {
	if !r.HasError() {
		if len(value) == 0 || len(strings.TrimSpace(value)) == 0 {
			r.AddError(http.StatusBadRequest, errors.New(ErrMissingValue+key))
		} else {
			uuidParsed := uuid.FromStringOrNil(value)
			if strings.EqualFold(value, "default") {
				r.data[key] = "default"
			} else if uuidParsed == uuid.Nil {
				r.AddError(http.StatusBadRequest, errors.New(ErrValueNotUUID+key))
			} else {
				r.data[key] = uuidParsed.String()
			}
		}
	}
	return r
}

func (r *Base) Parse36(key, value string) Context {
	if !r.HasError() {
		p := strings.ToUpper(value)
		if strings.EqualFold(p, "default") {
			r.data[key] = "default"
		} else if len(p) == 0 || len(strings.TrimSpace(p)) == 0 {
			r.AddError(http.StatusBadRequest, errors.New(ErrMissingValue+key))
		} else {
			if data, err := radix36.NewRadix36(p); err != nil {
				r.AddError(http.StatusBadRequest, errors.New(ErrValueNotBase36+key))
			} else {
				r.data[key] = data.String()
			}
		}
	}
	return r
}

func (r *Base) ParseUUID36(key, value string) Context {
	if !r.HasError() {
		p := value
		if strings.EqualFold(p, "default") {
			r.data[key] = "default"
		} else if len(p) == 0 || len(strings.TrimSpace(p)) == 0 {
			r.AddError(http.StatusBadRequest, errors.New(ErrMissingValue+key))
		} else {
			data, err := radix36.NewFromUUIDString(p)
			if err != nil {
				r.AddError(http.StatusBadRequest, errors.New(ErrValueNotUUID+key))
			} else {
				r.data[key] = data.String()
			}
		}
	}
	return r
}

func (r *Base) Parse36UUID(key, value string) Context {
	if !r.HasError() {

		p := value
		if strings.EqualFold(p, "default") {
			r.data[key] = "default"
		} else if len(p) == 0 || len(strings.TrimSpace(p)) == 0 {
			r.AddError(http.StatusBadRequest, errors.New(ErrMissingValue+key))
		} else {
			if data, err := radix36.NewRadix36(p); err != nil {
				r.AddError(http.StatusBadRequest, errors.New(ErrValueNotBase36+key))
			} else {
				r.data[key] = data.ToUUID().String()
			}
		}
	}
	return r
}

func (r *Base) ParseBoolean(key, value string, def bool) Context {
	if !r.HasError() {

		q := value
		if (len(q) == 0 || len(strings.TrimSpace(q)) == 0) ||
			(q != strconv.FormatBool(false) && q != strconv.FormatBool(true)) {
			q = strconv.FormatBool(def)
		}

		r.data[key] = q
	}
	return r
}

func (r *Base) Data() *map[string]string {
	if r.HasError() {
		return nil
	}
	return &r.data
}

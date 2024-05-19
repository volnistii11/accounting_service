// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: transfer.proto

package transfer

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on CreateMoneyTransferRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateMoneyTransferRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateMoneyTransferRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateMoneyTransferRequestMultiError, or nil if none found.
func (m *CreateMoneyTransferRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateMoneyTransferRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetInfo() == nil {
		err := CreateMoneyTransferRequestValidationError{
			field:  "Info",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetInfo()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateMoneyTransferRequestValidationError{
					field:  "Info",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateMoneyTransferRequestValidationError{
					field:  "Info",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetInfo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateMoneyTransferRequestValidationError{
				field:  "Info",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateMoneyTransferRequestMultiError(errors)
	}

	return nil
}

// CreateMoneyTransferRequestMultiError is an error wrapping multiple
// validation errors returned by CreateMoneyTransferRequest.ValidateAll() if
// the designated constraints aren't met.
type CreateMoneyTransferRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateMoneyTransferRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateMoneyTransferRequestMultiError) AllErrors() []error { return m }

// CreateMoneyTransferRequestValidationError is the validation error returned
// by CreateMoneyTransferRequest.Validate if the designated constraints aren't met.
type CreateMoneyTransferRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateMoneyTransferRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateMoneyTransferRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateMoneyTransferRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateMoneyTransferRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateMoneyTransferRequestValidationError) ErrorName() string {
	return "CreateMoneyTransferRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateMoneyTransferRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateMoneyTransferRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateMoneyTransferRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateMoneyTransferRequestValidationError{}

// Validate checks the field values on MoneyTransferInfo with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *MoneyTransferInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MoneyTransferInfo with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// MoneyTransferInfoMultiError, or nil if none found.
func (m *MoneyTransferInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *MoneyTransferInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetRequestId() <= 0 {
		err := MoneyTransferInfoValidationError{
			field:  "RequestId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetFromUserId() <= 0 {
		err := MoneyTransferInfoValidationError{
			field:  "FromUserId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetToUserId() <= 0 {
		err := MoneyTransferInfoValidationError{
			field:  "ToUserId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetSum() <= 0 {
		err := MoneyTransferInfoValidationError{
			field:  "Sum",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return MoneyTransferInfoMultiError(errors)
	}

	return nil
}

// MoneyTransferInfoMultiError is an error wrapping multiple validation errors
// returned by MoneyTransferInfo.ValidateAll() if the designated constraints
// aren't met.
type MoneyTransferInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MoneyTransferInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MoneyTransferInfoMultiError) AllErrors() []error { return m }

// MoneyTransferInfoValidationError is the validation error returned by
// MoneyTransferInfo.Validate if the designated constraints aren't met.
type MoneyTransferInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MoneyTransferInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MoneyTransferInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MoneyTransferInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MoneyTransferInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MoneyTransferInfoValidationError) ErrorName() string {
	return "MoneyTransferInfoValidationError"
}

// Error satisfies the builtin error interface
func (e MoneyTransferInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMoneyTransferInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MoneyTransferInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MoneyTransferInfoValidationError{}
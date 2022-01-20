// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/account/service/v1/account.proto

package v1

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

// Validate checks the field values on CreateAccountRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateAccountRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateAccountRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateAccountRequestMultiError, or nil if none found.
func (m *CreateAccountRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateAccountRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetName()) < 1 {
		err := CreateAccountRequestValidationError{
			field:  "Name",
			reason: "value length must be at least 1 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Sex

	if err := m._validateEmail(m.GetEmail()); err != nil {
		err = CreateAccountRequestValidationError{
			field:  "Email",
			reason: "value must be a valid email address",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Age

	if len(errors) > 0 {
		return CreateAccountRequestMultiError(errors)
	}

	return nil
}

func (m *CreateAccountRequest) _validateHostname(host string) error {
	s := strings.ToLower(strings.TrimSuffix(host, "."))

	if len(host) > 253 {
		return errors.New("hostname cannot exceed 253 characters")
	}

	for _, part := range strings.Split(s, ".") {
		if l := len(part); l == 0 || l > 63 {
			return errors.New("hostname part must be non-empty and cannot exceed 63 characters")
		}

		if part[0] == '-' {
			return errors.New("hostname parts cannot begin with hyphens")
		}

		if part[len(part)-1] == '-' {
			return errors.New("hostname parts cannot end with hyphens")
		}

		for _, r := range part {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
				return fmt.Errorf("hostname parts can only contain alphanumeric characters or hyphens, got %q", string(r))
			}
		}
	}

	return nil
}

func (m *CreateAccountRequest) _validateEmail(addr string) error {
	a, err := mail.ParseAddress(addr)
	if err != nil {
		return err
	}
	addr = a.Address

	if len(addr) > 254 {
		return errors.New("email addresses cannot exceed 254 characters")
	}

	parts := strings.SplitN(addr, "@", 2)

	if len(parts[0]) > 64 {
		return errors.New("email address local phrase cannot exceed 64 characters")
	}

	return m._validateHostname(parts[1])
}

// CreateAccountRequestMultiError is an error wrapping multiple validation
// errors returned by CreateAccountRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateAccountRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateAccountRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateAccountRequestMultiError) AllErrors() []error { return m }

// CreateAccountRequestValidationError is the validation error returned by
// CreateAccountRequest.Validate if the designated constraints aren't met.
type CreateAccountRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateAccountRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateAccountRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateAccountRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateAccountRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateAccountRequestValidationError) ErrorName() string {
	return "CreateAccountRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateAccountRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateAccountRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateAccountRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateAccountRequestValidationError{}

// Validate checks the field values on CreateAccountReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateAccountReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateAccountReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateAccountReplyMultiError, or nil if none found.
func (m *CreateAccountReply) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateAccountReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return CreateAccountReplyMultiError(errors)
	}

	return nil
}

// CreateAccountReplyMultiError is an error wrapping multiple validation errors
// returned by CreateAccountReply.ValidateAll() if the designated constraints
// aren't met.
type CreateAccountReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateAccountReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateAccountReplyMultiError) AllErrors() []error { return m }

// CreateAccountReplyValidationError is the validation error returned by
// CreateAccountReply.Validate if the designated constraints aren't met.
type CreateAccountReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateAccountReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateAccountReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateAccountReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateAccountReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateAccountReplyValidationError) ErrorName() string {
	return "CreateAccountReplyValidationError"
}

// Error satisfies the builtin error interface
func (e CreateAccountReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateAccountReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateAccountReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateAccountReplyValidationError{}

// Validate checks the field values on UpdateAccountRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateAccountRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateAccountRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateAccountRequestMultiError, or nil if none found.
func (m *UpdateAccountRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateAccountRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UpdateAccountRequestMultiError(errors)
	}

	return nil
}

// UpdateAccountRequestMultiError is an error wrapping multiple validation
// errors returned by UpdateAccountRequest.ValidateAll() if the designated
// constraints aren't met.
type UpdateAccountRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateAccountRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateAccountRequestMultiError) AllErrors() []error { return m }

// UpdateAccountRequestValidationError is the validation error returned by
// UpdateAccountRequest.Validate if the designated constraints aren't met.
type UpdateAccountRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateAccountRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateAccountRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateAccountRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateAccountRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateAccountRequestValidationError) ErrorName() string {
	return "UpdateAccountRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateAccountRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateAccountRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateAccountRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateAccountRequestValidationError{}

// Validate checks the field values on UpdateAccountReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateAccountReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateAccountReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateAccountReplyMultiError, or nil if none found.
func (m *UpdateAccountReply) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateAccountReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UpdateAccountReplyMultiError(errors)
	}

	return nil
}

// UpdateAccountReplyMultiError is an error wrapping multiple validation errors
// returned by UpdateAccountReply.ValidateAll() if the designated constraints
// aren't met.
type UpdateAccountReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateAccountReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateAccountReplyMultiError) AllErrors() []error { return m }

// UpdateAccountReplyValidationError is the validation error returned by
// UpdateAccountReply.Validate if the designated constraints aren't met.
type UpdateAccountReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateAccountReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateAccountReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateAccountReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateAccountReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateAccountReplyValidationError) ErrorName() string {
	return "UpdateAccountReplyValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateAccountReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateAccountReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateAccountReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateAccountReplyValidationError{}

// Validate checks the field values on DeleteAccountRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteAccountRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteAccountRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteAccountRequestMultiError, or nil if none found.
func (m *DeleteAccountRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteAccountRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeleteAccountRequestMultiError(errors)
	}

	return nil
}

// DeleteAccountRequestMultiError is an error wrapping multiple validation
// errors returned by DeleteAccountRequest.ValidateAll() if the designated
// constraints aren't met.
type DeleteAccountRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteAccountRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteAccountRequestMultiError) AllErrors() []error { return m }

// DeleteAccountRequestValidationError is the validation error returned by
// DeleteAccountRequest.Validate if the designated constraints aren't met.
type DeleteAccountRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteAccountRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteAccountRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteAccountRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteAccountRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteAccountRequestValidationError) ErrorName() string {
	return "DeleteAccountRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteAccountRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteAccountRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteAccountRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteAccountRequestValidationError{}

// Validate checks the field values on DeleteAccountReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteAccountReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteAccountReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteAccountReplyMultiError, or nil if none found.
func (m *DeleteAccountReply) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteAccountReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeleteAccountReplyMultiError(errors)
	}

	return nil
}

// DeleteAccountReplyMultiError is an error wrapping multiple validation errors
// returned by DeleteAccountReply.ValidateAll() if the designated constraints
// aren't met.
type DeleteAccountReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteAccountReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteAccountReplyMultiError) AllErrors() []error { return m }

// DeleteAccountReplyValidationError is the validation error returned by
// DeleteAccountReply.Validate if the designated constraints aren't met.
type DeleteAccountReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteAccountReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteAccountReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteAccountReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteAccountReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteAccountReplyValidationError) ErrorName() string {
	return "DeleteAccountReplyValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteAccountReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteAccountReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteAccountReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteAccountReplyValidationError{}

// Validate checks the field values on GetAccountRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetAccountRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetAccountRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetAccountRequestMultiError, or nil if none found.
func (m *GetAccountRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetAccountRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return GetAccountRequestMultiError(errors)
	}

	return nil
}

// GetAccountRequestMultiError is an error wrapping multiple validation errors
// returned by GetAccountRequest.ValidateAll() if the designated constraints
// aren't met.
type GetAccountRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetAccountRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetAccountRequestMultiError) AllErrors() []error { return m }

// GetAccountRequestValidationError is the validation error returned by
// GetAccountRequest.Validate if the designated constraints aren't met.
type GetAccountRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetAccountRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetAccountRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetAccountRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetAccountRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetAccountRequestValidationError) ErrorName() string {
	return "GetAccountRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetAccountRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetAccountRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetAccountRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetAccountRequestValidationError{}

// Validate checks the field values on GetAccountReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetAccountReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetAccountReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetAccountReplyMultiError, or nil if none found.
func (m *GetAccountReply) ValidateAll() error {
	return m.validate(true)
}

func (m *GetAccountReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetAccountReplyMultiError(errors)
	}

	return nil
}

// GetAccountReplyMultiError is an error wrapping multiple validation errors
// returned by GetAccountReply.ValidateAll() if the designated constraints
// aren't met.
type GetAccountReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetAccountReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetAccountReplyMultiError) AllErrors() []error { return m }

// GetAccountReplyValidationError is the validation error returned by
// GetAccountReply.Validate if the designated constraints aren't met.
type GetAccountReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetAccountReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetAccountReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetAccountReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetAccountReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetAccountReplyValidationError) ErrorName() string { return "GetAccountReplyValidationError" }

// Error satisfies the builtin error interface
func (e GetAccountReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetAccountReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetAccountReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetAccountReplyValidationError{}

// Validate checks the field values on ListAccountRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListAccountRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListAccountRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListAccountRequestMultiError, or nil if none found.
func (m *ListAccountRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListAccountRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ListAccountRequestMultiError(errors)
	}

	return nil
}

// ListAccountRequestMultiError is an error wrapping multiple validation errors
// returned by ListAccountRequest.ValidateAll() if the designated constraints
// aren't met.
type ListAccountRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListAccountRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListAccountRequestMultiError) AllErrors() []error { return m }

// ListAccountRequestValidationError is the validation error returned by
// ListAccountRequest.Validate if the designated constraints aren't met.
type ListAccountRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListAccountRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListAccountRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListAccountRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListAccountRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListAccountRequestValidationError) ErrorName() string {
	return "ListAccountRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListAccountRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListAccountRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListAccountRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListAccountRequestValidationError{}

// Validate checks the field values on ListAccountReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListAccountReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListAccountReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListAccountReplyMultiError, or nil if none found.
func (m *ListAccountReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ListAccountReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetResults() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListAccountReplyValidationError{
						field:  fmt.Sprintf("Results[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListAccountReplyValidationError{
						field:  fmt.Sprintf("Results[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListAccountReplyValidationError{
					field:  fmt.Sprintf("Results[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListAccountReplyMultiError(errors)
	}

	return nil
}

// ListAccountReplyMultiError is an error wrapping multiple validation errors
// returned by ListAccountReply.ValidateAll() if the designated constraints
// aren't met.
type ListAccountReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListAccountReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListAccountReplyMultiError) AllErrors() []error { return m }

// ListAccountReplyValidationError is the validation error returned by
// ListAccountReply.Validate if the designated constraints aren't met.
type ListAccountReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListAccountReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListAccountReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListAccountReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListAccountReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListAccountReplyValidationError) ErrorName() string { return "ListAccountReplyValidationError" }

// Error satisfies the builtin error interface
func (e ListAccountReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListAccountReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListAccountReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListAccountReplyValidationError{}

// Validate checks the field values on ListAccountReply_Account with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListAccountReply_Account) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListAccountReply_Account with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListAccountReply_AccountMultiError, or nil if none found.
func (m *ListAccountReply_Account) ValidateAll() error {
	return m.validate(true)
}

func (m *ListAccountReply_Account) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Sex

	// no validation rules for Email

	// no validation rules for Age

	if len(errors) > 0 {
		return ListAccountReply_AccountMultiError(errors)
	}

	return nil
}

// ListAccountReply_AccountMultiError is an error wrapping multiple validation
// errors returned by ListAccountReply_Account.ValidateAll() if the designated
// constraints aren't met.
type ListAccountReply_AccountMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListAccountReply_AccountMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListAccountReply_AccountMultiError) AllErrors() []error { return m }

// ListAccountReply_AccountValidationError is the validation error returned by
// ListAccountReply_Account.Validate if the designated constraints aren't met.
type ListAccountReply_AccountValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListAccountReply_AccountValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListAccountReply_AccountValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListAccountReply_AccountValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListAccountReply_AccountValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListAccountReply_AccountValidationError) ErrorName() string {
	return "ListAccountReply_AccountValidationError"
}

// Error satisfies the builtin error interface
func (e ListAccountReply_AccountValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListAccountReply_Account.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListAccountReply_AccountValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListAccountReply_AccountValidationError{}

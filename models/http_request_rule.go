// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2019 HAProxy Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HTTPRequestRule HTTP Request Rule
//
// HAProxy HTTP request rule configuration (corresponds to http-request directives)
// swagger:model http_request_rule
type HTTPRequestRule struct {

	// return headers
	ReturnHeaders []*HTTPRequestRuleReturnHdrsItems0 `json:"return_hdrs"`

	// acl file
	// Pattern: ^[^\s]+$
	ACLFile string `json:"acl_file,omitempty"`

	// acl keyfmt
	// Pattern: ^[^\s]+$
	ACLKeyfmt string `json:"acl_keyfmt,omitempty"`

	// auth realm
	// Pattern: ^[^\s]+$
	AuthRealm string `json:"auth_realm,omitempty"`

	// cache name
	// Pattern: ^[^\s]+$
	CacheName string `json:"cache_name,omitempty"`

	// capture id
	CaptureID *int64 `json:"capture_id,omitempty"`

	// capture len
	CaptureLen int64 `json:"capture_len,omitempty"`

	// capture sample
	// Pattern: ^[^\s]+$
	CaptureSample string `json:"capture_sample,omitempty"`

	// cond
	// Enum: [if unless]
	Cond string `json:"cond,omitempty"`

	// cond test
	CondTest string `json:"cond_test,omitempty"`

	// deny status
	// Enum: [200 400 403 405 408 425 429 500 502 503 504]
	DenyStatus *int64 `json:"deny_status,omitempty"`

	// expr
	Expr string `json:"expr,omitempty"`

	// hdr format
	HdrFormat string `json:"hdr_format,omitempty"`

	// hdr match
	HdrMatch string `json:"hdr_match,omitempty"`

	// hdr name
	HdrName string `json:"hdr_name,omitempty"`

	// hint format
	// Pattern: ^[^\s]+$
	HintFormat string `json:"hint_format,omitempty"`

	// hint name
	// Pattern: ^[^\s]+$
	HintName string `json:"hint_name,omitempty"`

	// index
	// Required: true
	Index *int64 `json:"index"`

	// log level
	// Enum: [emerg alert crit err warning notice info debug silent]
	LogLevel string `json:"log_level,omitempty"`

	// lua action
	// Pattern: ^[^\s]+$
	LuaAction string `json:"lua_action,omitempty"`

	// lua params
	LuaParams string `json:"lua_params,omitempty"`

	// map file
	// Pattern: ^[^\s]+$
	MapFile string `json:"map_file,omitempty"`

	// map keyfmt
	// Pattern: ^[^\s]+$
	MapKeyfmt string `json:"map_keyfmt,omitempty"`

	// map valuefmt
	// Pattern: ^[^\s]+$
	MapValuefmt string `json:"map_valuefmt,omitempty"`

	// mark value
	// Pattern: ^(0x[0-9A-Fa-f]+|[0-9]+)$
	MarkValue string `json:"mark_value,omitempty"`

	// method fmt
	// Pattern: ^[^\s]+$
	MethodFmt string `json:"method_fmt,omitempty"`

	// nice value
	// Maximum: 1024
	// Minimum: -1024
	NiceValue int64 `json:"nice_value,omitempty"`

	// path fmt
	// Pattern: ^[^\s]+$
	PathFmt string `json:"path_fmt,omitempty"`

	// path match
	// Pattern: ^[^\s]+$
	PathMatch string `json:"path_match,omitempty"`

	// protocol
	// Enum: [ipv4 ipv6]
	Protocol string `json:"protocol,omitempty"`

	// query fmt
	QueryFmt string `json:"query-fmt,omitempty"`

	// redir code
	// Enum: [301 302 303 307 308]
	RedirCode *int64 `json:"redir_code,omitempty"`

	// redir option
	RedirOption string `json:"redir_option,omitempty"`

	// redir type
	// Enum: [location prefix scheme]
	RedirType string `json:"redir_type,omitempty"`

	// redir value
	// Pattern: ^[^\s]+$
	RedirValue string `json:"redir_value,omitempty"`

	// resolvers
	Resolvers string `json:"resolvers,omitempty"`

	// return content
	ReturnContent string `json:"return_content,omitempty"`

	// return content format
	// Enum: [default-errorfile errorfile errorfiles file lf-file string lf-string]
	ReturnContentFormat string `json:"return_content_format,omitempty"`

	// return content type
	ReturnContentType *string `json:"return_content_type,omitempty"`

	// return status code
	// Maximum: 599
	// Minimum: 200
	ReturnStatusCode *int64 `json:"return_status_code,omitempty"`

	// sc expr
	ScExpr string `json:"sc_expr,omitempty"`

	// sc id
	ScID int64 `json:"sc_id,omitempty"`

	// sc int
	ScInt *int64 `json:"sc_int,omitempty"`

	// service name
	ServiceName string `json:"service_name,omitempty"`

	// spoe engine
	// Pattern: ^[^\s]+$
	SpoeEngine string `json:"spoe_engine,omitempty"`

	// spoe group
	// Pattern: ^[^\s]+$
	SpoeGroup string `json:"spoe_group,omitempty"`

	// strict mode
	// Enum: [on off]
	StrictMode string `json:"strict_mode,omitempty"`

	// tos value
	// Pattern: ^(0x[0-9A-Fa-f]+|[0-9]+)$
	TosValue string `json:"tos_value,omitempty"`

	// track sc0 key
	// Pattern: ^[^\s]+$
	TrackSc0Key string `json:"track-sc0-key,omitempty"`

	// track sc0 table
	// Pattern: ^[^\s]+$
	TrackSc0Table string `json:"track-sc0-table,omitempty"`

	// track sc1 key
	// Pattern: ^[^\s]+$
	TrackSc1Key string `json:"track-sc1-key,omitempty"`

	// track sc1 table
	// Pattern: ^[^\s]+$
	TrackSc1Table string `json:"track-sc1-table,omitempty"`

	// track sc2 key
	// Pattern: ^[^\s]+$
	TrackSc2Key string `json:"track-sc2-key,omitempty"`

	// track sc2 table
	// Pattern: ^[^\s]+$
	TrackSc2Table string `json:"track-sc2-table,omitempty"`

	// type
	// Required: true
	// Enum: [allow deny auth redirect tarpit add-header replace-header replace-value del-header set-header set-log-level set-path replace-path set-query set-uri set-var send-spoe-group add-acl del-acl capture track-sc0 track-sc1 track-sc2 set-map del-map cache-use disable-l7-retry early-hint replace-uri sc-inc-gpc0 sc-inc-gpc1 do-resolve set-dst set-dst-port sc-set-gpt0 set-mark set-nice set-method set-priority-class set-priority-offset set-src set-src-por wait-for-handshake set-tos silent-drop unset-var strict-mode lua use-service return]
	Type string `json:"type"`

	// uri fmt
	URIFmt string `json:"uri-fmt,omitempty"`

	// uri match
	URIMatch string `json:"uri-match,omitempty"`

	// var expr
	VarExpr string `json:"var_expr,omitempty"`

	// var name
	// Pattern: ^[^\s]+$
	VarName string `json:"var_name,omitempty"`

	// var scope
	// Pattern: ^[^\s]+$
	VarScope string `json:"var_scope,omitempty"`
}

// Validate validates this http request rule
func (m *HTTPRequestRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReturnHeaders(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateACLFile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateACLKeyfmt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthRealm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCacheName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCaptureSample(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCond(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDenyStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHintFormat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHintName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIndex(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogLevel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLuaAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMapFile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMapKeyfmt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMapValuefmt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMarkValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMethodFmt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNiceValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePathFmt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePathMatch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtocol(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRedirCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRedirType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRedirValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReturnContentFormat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReturnStatusCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpoeEngine(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpoeGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStrictMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTosValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrackSc0Key(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrackSc0Table(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrackSc1Key(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrackSc1Table(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrackSc2Key(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrackSc2Table(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVarName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVarScope(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HTTPRequestRule) validateReturnHeaders(formats strfmt.Registry) error {

	if swag.IsZero(m.ReturnHeaders) { // not required
		return nil
	}

	for i := 0; i < len(m.ReturnHeaders); i++ {
		if swag.IsZero(m.ReturnHeaders[i]) { // not required
			continue
		}

		if m.ReturnHeaders[i] != nil {
			if err := m.ReturnHeaders[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("return_hdrs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HTTPRequestRule) validateACLFile(formats strfmt.Registry) error {

	if swag.IsZero(m.ACLFile) { // not required
		return nil
	}

	if err := validate.Pattern("acl_file", "body", string(m.ACLFile), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPRequestRule) validateACLKeyfmt(formats strfmt.Registry) error {

	if swag.IsZero(m.ACLKeyfmt) { // not required
		return nil
	}

	if err := validate.Pattern("acl_keyfmt", "body", string(m.ACLKeyfmt), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPRequestRule) validateAuthRealm(formats strfmt.Registry) error {

	if swag.IsZero(m.AuthRealm) { // not required
		return nil
	}

	if err := validate.Pattern("auth_realm", "body", string(m.AuthRealm), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPRequestRule) validateCacheName(formats strfmt.Registry) error {

	if swag.IsZero(m.CacheName) { // not required
		return nil
	}

	if err := validate.Pattern("cache_name", "body", string(m.CacheName), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPRequestRule) validateCaptureSample(formats strfmt.Registry) error {

	if swag.IsZero(m.CaptureSample) { // not required
		return nil
	}

	if err := validate.Pattern("capture_sample", "body", string(m.CaptureSample), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

var httpRequestRuleTypeCondPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["if","unless"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		httpRequestRuleTypeCondPropEnum = append(httpRequestRuleTypeCondPropEnum, v)
	}
}

const (

	// HTTPRequestRuleCondIf captures enum value "if"
	HTTPRequestRuleCondIf string = "if"

	// HTTPRequestRuleCondUnless captures enum value "unless"
	HTTPRequestRuleCondUnless string = "unless"
)

// prop value enum
func (m *HTTPRequestRule) validateCondEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, httpRequestRuleTypeCondPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *HTTPRequestRule) validateCond(formats strfmt.Registry) error {

	if swag.IsZero(m.Cond) { // not required
		return nil
	}

	// value enum
	if err := m.validateCondEnum("cond", "body", m.Cond); err != nil {
		return err
	}

	return nil
}

var httpRequestRuleTypeDenyStatusPropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[200,400,403,405,408,425,429,500,502,503,504]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		httpRequestRuleTypeDenyStatusPropEnum = append(httpRequestRuleTypeDenyStatusPropEnum, v)
	}
}

// prop value enum
func (m *HTTPRequestRule) validateDenyStatusEnum(path, location string, value int64) error {
	if err := validate.Enum(path, location, value, httpRequestRuleTypeDenyStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *HTTPRequestRule) validateDenyStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.DenyStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateDenyStatusEnum("deny_status", "body", *m.DenyStatus); err != nil {
		return err
	}

	return nil
}

func (m *HTTPRequestRule) validateHintFormat(formats strfmt.Registry) error {

	if swag.IsZero(m.HintFormat) { // not required
		return nil
	}

	if err := validate.Pattern("hint_format", "body", string(m.HintFormat), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPRequestRule) validateHintName(formats strfmt.Registry) error {

	if swag.IsZero(m.HintName) { // not required
		return nil
	}

	if err := validate.Pattern("hint_name", "body", string(m.HintName), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPRequestRule) validateIndex(formats strfmt.Registry) error {

	if err := validate.Required("index", "body", m.Index); err != nil {
		return err
	}

	return nil
}

var httpRequestRuleTypeLogLevelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["emerg","alert","crit","err","warning","notice","info","debug","silent"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		httpRequestRuleTypeLogLevelPropEnum = append(httpRequestRuleTypeLogLevelPropEnum, v)
	}
}

const (

	// HTTPRequestRuleLogLevelEmerg captures enum value "emerg"
	HTTPRequestRuleLogLevelEmerg string = "emerg"

	// HTTPRequestRuleLogLevelAlert captures enum value "alert"
	HTTPRequestRuleLogLevelAlert string = "alert"

	// HTTPRequestRuleLogLevelCrit captures enum value "crit"
	HTTPRequestRuleLogLevelCrit string = "crit"

	// HTTPRequestRuleLogLevelErr captures enum value "err"
	HTTPRequestRuleLogLevelErr string = "err"

	// HTTPRequestRuleLogLevelWarning captures enum value "warning"
	HTTPRequestRuleLogLevelWarning string = "warning"

	// HTTPRequestRuleLogLevelNotice captures enum value "notice"
	HTTPRequestRuleLogLevelNotice string = "notice"

	// HTTPRequestRuleLogLevelInfo captures enum value "info"
	HTTPRequestRuleLogLevelInfo string = "info"

	// HTTPRequestRuleLogLevelDebug captures enum value "debug"
	HTTPRequestRuleLogLevelDebug string = "debug"

	// HTTPRequestRuleLogLevelSilent captures enum value "silent"
	HTTPRequestRuleLogLevelSilent string = "silent"
)

// prop value enum
func (m *HTTPRequestRule) validateLogLevelEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, httpRequestRuleTypeLogLevelPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *HTTPRequestRule) validateLogLevel(formats strfmt.Registry) error {

	if swag.IsZero(m.LogLevel) { // not required
		return nil
	}

	// value enum
	if err := m.validateLogLevelEnum("log_level", "body", m.LogLevel); err != nil {
		return err
	}

	return nil
}

func (m *HTTPRequestRule) validateLuaAction(formats strfmt.Registry) error {

	if swag.IsZero(m.LuaAction) { // not required
		return nil
	}

	if err := validate.Pattern("lua_action", "body", string(m.LuaAction), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPRequestRule) validateMapFile(formats strfmt.Registry) error {

	if swag.IsZero(m.MapFile) { // not required
		return nil
	}

	if err := validate.Pattern("map_file", "body", string(m.MapFile), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPRequestRule) validateMapKeyfmt(formats strfmt.Registry) error {

	if swag.IsZero(m.MapKeyfmt) { // not required
		return nil
	}

	if err := validate.Pattern("map_keyfmt", "body", string(m.MapKeyfmt), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPRequestRule) validateMapValuefmt(formats strfmt.Registry) error {

	if swag.IsZero(m.MapValuefmt) { // not required
		return nil
	}

	if err := validate.Pattern("map_valuefmt", "body", string(m.MapValuefmt), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPRequestRule) validateMarkValue(formats strfmt.Registry) error {

	if swag.IsZero(m.MarkValue) { // not required
		return nil
	}

	if err := validate.Pattern("mark_value", "body", string(m.MarkValue), `^(0x[0-9A-Fa-f]+|[0-9]+)$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPRequestRule) validateMethodFmt(formats strfmt.Registry) error {

	if swag.IsZero(m.MethodFmt) { // not required
		return nil
	}

	if err := validate.Pattern("method_fmt", "body", string(m.MethodFmt), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPRequestRule) validateNiceValue(formats strfmt.Registry) error {

	if swag.IsZero(m.NiceValue) { // not required
		return nil
	}

	if err := validate.MinimumInt("nice_value", "body", int64(m.NiceValue), -1024, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("nice_value", "body", int64(m.NiceValue), 1024, false); err != nil {
		return err
	}

	return nil
}

func (m *HTTPRequestRule) validatePathFmt(formats strfmt.Registry) error {

	if swag.IsZero(m.PathFmt) { // not required
		return nil
	}

	if err := validate.Pattern("path_fmt", "body", string(m.PathFmt), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPRequestRule) validatePathMatch(formats strfmt.Registry) error {

	if swag.IsZero(m.PathMatch) { // not required
		return nil
	}

	if err := validate.Pattern("path_match", "body", string(m.PathMatch), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

var httpRequestRuleTypeProtocolPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ipv4","ipv6"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		httpRequestRuleTypeProtocolPropEnum = append(httpRequestRuleTypeProtocolPropEnum, v)
	}
}

const (

	// HTTPRequestRuleProtocolIPV4 captures enum value "ipv4"
	HTTPRequestRuleProtocolIPV4 string = "ipv4"

	// HTTPRequestRuleProtocolIPV6 captures enum value "ipv6"
	HTTPRequestRuleProtocolIPV6 string = "ipv6"
)

// prop value enum
func (m *HTTPRequestRule) validateProtocolEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, httpRequestRuleTypeProtocolPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *HTTPRequestRule) validateProtocol(formats strfmt.Registry) error {

	if swag.IsZero(m.Protocol) { // not required
		return nil
	}

	// value enum
	if err := m.validateProtocolEnum("protocol", "body", m.Protocol); err != nil {
		return err
	}

	return nil
}

var httpRequestRuleTypeRedirCodePropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[301,302,303,307,308]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		httpRequestRuleTypeRedirCodePropEnum = append(httpRequestRuleTypeRedirCodePropEnum, v)
	}
}

// prop value enum
func (m *HTTPRequestRule) validateRedirCodeEnum(path, location string, value int64) error {
	if err := validate.Enum(path, location, value, httpRequestRuleTypeRedirCodePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *HTTPRequestRule) validateRedirCode(formats strfmt.Registry) error {

	if swag.IsZero(m.RedirCode) { // not required
		return nil
	}

	// value enum
	if err := m.validateRedirCodeEnum("redir_code", "body", *m.RedirCode); err != nil {
		return err
	}

	return nil
}

var httpRequestRuleTypeRedirTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["location","prefix","scheme"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		httpRequestRuleTypeRedirTypePropEnum = append(httpRequestRuleTypeRedirTypePropEnum, v)
	}
}

const (

	// HTTPRequestRuleRedirTypeLocation captures enum value "location"
	HTTPRequestRuleRedirTypeLocation string = "location"

	// HTTPRequestRuleRedirTypePrefix captures enum value "prefix"
	HTTPRequestRuleRedirTypePrefix string = "prefix"

	// HTTPRequestRuleRedirTypeScheme captures enum value "scheme"
	HTTPRequestRuleRedirTypeScheme string = "scheme"
)

// prop value enum
func (m *HTTPRequestRule) validateRedirTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, httpRequestRuleTypeRedirTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *HTTPRequestRule) validateRedirType(formats strfmt.Registry) error {

	if swag.IsZero(m.RedirType) { // not required
		return nil
	}

	// value enum
	if err := m.validateRedirTypeEnum("redir_type", "body", m.RedirType); err != nil {
		return err
	}

	return nil
}

func (m *HTTPRequestRule) validateRedirValue(formats strfmt.Registry) error {

	if swag.IsZero(m.RedirValue) { // not required
		return nil
	}

	if err := validate.Pattern("redir_value", "body", string(m.RedirValue), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

var httpRequestRuleTypeReturnContentFormatPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["default-errorfile","errorfile","errorfiles","file","lf-file","string","lf-string"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		httpRequestRuleTypeReturnContentFormatPropEnum = append(httpRequestRuleTypeReturnContentFormatPropEnum, v)
	}
}

const (

	// HTTPRequestRuleReturnContentFormatDefaultErrorfile captures enum value "default-errorfile"
	HTTPRequestRuleReturnContentFormatDefaultErrorfile string = "default-errorfile"

	// HTTPRequestRuleReturnContentFormatErrorfile captures enum value "errorfile"
	HTTPRequestRuleReturnContentFormatErrorfile string = "errorfile"

	// HTTPRequestRuleReturnContentFormatErrorfiles captures enum value "errorfiles"
	HTTPRequestRuleReturnContentFormatErrorfiles string = "errorfiles"

	// HTTPRequestRuleReturnContentFormatFile captures enum value "file"
	HTTPRequestRuleReturnContentFormatFile string = "file"

	// HTTPRequestRuleReturnContentFormatLfFile captures enum value "lf-file"
	HTTPRequestRuleReturnContentFormatLfFile string = "lf-file"

	// HTTPRequestRuleReturnContentFormatString captures enum value "string"
	HTTPRequestRuleReturnContentFormatString string = "string"

	// HTTPRequestRuleReturnContentFormatLfString captures enum value "lf-string"
	HTTPRequestRuleReturnContentFormatLfString string = "lf-string"
)

// prop value enum
func (m *HTTPRequestRule) validateReturnContentFormatEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, httpRequestRuleTypeReturnContentFormatPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *HTTPRequestRule) validateReturnContentFormat(formats strfmt.Registry) error {

	if swag.IsZero(m.ReturnContentFormat) { // not required
		return nil
	}

	// value enum
	if err := m.validateReturnContentFormatEnum("return_content_format", "body", m.ReturnContentFormat); err != nil {
		return err
	}

	return nil
}

func (m *HTTPRequestRule) validateReturnStatusCode(formats strfmt.Registry) error {

	if swag.IsZero(m.ReturnStatusCode) { // not required
		return nil
	}

	if err := validate.MinimumInt("return_status_code", "body", int64(*m.ReturnStatusCode), 200, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("return_status_code", "body", int64(*m.ReturnStatusCode), 599, false); err != nil {
		return err
	}

	return nil
}

func (m *HTTPRequestRule) validateSpoeEngine(formats strfmt.Registry) error {

	if swag.IsZero(m.SpoeEngine) { // not required
		return nil
	}

	if err := validate.Pattern("spoe_engine", "body", string(m.SpoeEngine), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPRequestRule) validateSpoeGroup(formats strfmt.Registry) error {

	if swag.IsZero(m.SpoeGroup) { // not required
		return nil
	}

	if err := validate.Pattern("spoe_group", "body", string(m.SpoeGroup), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

var httpRequestRuleTypeStrictModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["on","off"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		httpRequestRuleTypeStrictModePropEnum = append(httpRequestRuleTypeStrictModePropEnum, v)
	}
}

const (

	// HTTPRequestRuleStrictModeOn captures enum value "on"
	HTTPRequestRuleStrictModeOn string = "on"

	// HTTPRequestRuleStrictModeOff captures enum value "off"
	HTTPRequestRuleStrictModeOff string = "off"
)

// prop value enum
func (m *HTTPRequestRule) validateStrictModeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, httpRequestRuleTypeStrictModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *HTTPRequestRule) validateStrictMode(formats strfmt.Registry) error {

	if swag.IsZero(m.StrictMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateStrictModeEnum("strict_mode", "body", m.StrictMode); err != nil {
		return err
	}

	return nil
}

func (m *HTTPRequestRule) validateTosValue(formats strfmt.Registry) error {

	if swag.IsZero(m.TosValue) { // not required
		return nil
	}

	if err := validate.Pattern("tos_value", "body", string(m.TosValue), `^(0x[0-9A-Fa-f]+|[0-9]+)$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPRequestRule) validateTrackSc0Key(formats strfmt.Registry) error {

	if swag.IsZero(m.TrackSc0Key) { // not required
		return nil
	}

	if err := validate.Pattern("track-sc0-key", "body", string(m.TrackSc0Key), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPRequestRule) validateTrackSc0Table(formats strfmt.Registry) error {

	if swag.IsZero(m.TrackSc0Table) { // not required
		return nil
	}

	if err := validate.Pattern("track-sc0-table", "body", string(m.TrackSc0Table), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPRequestRule) validateTrackSc1Key(formats strfmt.Registry) error {

	if swag.IsZero(m.TrackSc1Key) { // not required
		return nil
	}

	if err := validate.Pattern("track-sc1-key", "body", string(m.TrackSc1Key), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPRequestRule) validateTrackSc1Table(formats strfmt.Registry) error {

	if swag.IsZero(m.TrackSc1Table) { // not required
		return nil
	}

	if err := validate.Pattern("track-sc1-table", "body", string(m.TrackSc1Table), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPRequestRule) validateTrackSc2Key(formats strfmt.Registry) error {

	if swag.IsZero(m.TrackSc2Key) { // not required
		return nil
	}

	if err := validate.Pattern("track-sc2-key", "body", string(m.TrackSc2Key), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPRequestRule) validateTrackSc2Table(formats strfmt.Registry) error {

	if swag.IsZero(m.TrackSc2Table) { // not required
		return nil
	}

	if err := validate.Pattern("track-sc2-table", "body", string(m.TrackSc2Table), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

var httpRequestRuleTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["allow","deny","auth","redirect","tarpit","add-header","replace-header","replace-value","del-header","set-header","set-log-level","set-path","replace-path","set-query","set-uri","set-var","send-spoe-group","add-acl","del-acl","capture","track-sc0","track-sc1","track-sc2","set-map","del-map","cache-use","disable-l7-retry","early-hint","replace-uri","sc-inc-gpc0","sc-inc-gpc1","do-resolve","set-dst","set-dst-port","sc-set-gpt0","set-mark","set-nice","set-method","set-priority-class","set-priority-offset","set-src","set-src-por","wait-for-handshake","set-tos","silent-drop","unset-var","strict-mode","lua","use-service","return"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		httpRequestRuleTypeTypePropEnum = append(httpRequestRuleTypeTypePropEnum, v)
	}
}

const (

	// HTTPRequestRuleTypeAllow captures enum value "allow"
	HTTPRequestRuleTypeAllow string = "allow"

	// HTTPRequestRuleTypeDeny captures enum value "deny"
	HTTPRequestRuleTypeDeny string = "deny"

	// HTTPRequestRuleTypeAuth captures enum value "auth"
	HTTPRequestRuleTypeAuth string = "auth"

	// HTTPRequestRuleTypeRedirect captures enum value "redirect"
	HTTPRequestRuleTypeRedirect string = "redirect"

	// HTTPRequestRuleTypeTarpit captures enum value "tarpit"
	HTTPRequestRuleTypeTarpit string = "tarpit"

	// HTTPRequestRuleTypeAddHeader captures enum value "add-header"
	HTTPRequestRuleTypeAddHeader string = "add-header"

	// HTTPRequestRuleTypeReplaceHeader captures enum value "replace-header"
	HTTPRequestRuleTypeReplaceHeader string = "replace-header"

	// HTTPRequestRuleTypeReplaceValue captures enum value "replace-value"
	HTTPRequestRuleTypeReplaceValue string = "replace-value"

	// HTTPRequestRuleTypeDelHeader captures enum value "del-header"
	HTTPRequestRuleTypeDelHeader string = "del-header"

	// HTTPRequestRuleTypeSetHeader captures enum value "set-header"
	HTTPRequestRuleTypeSetHeader string = "set-header"

	// HTTPRequestRuleTypeSetLogLevel captures enum value "set-log-level"
	HTTPRequestRuleTypeSetLogLevel string = "set-log-level"

	// HTTPRequestRuleTypeSetPath captures enum value "set-path"
	HTTPRequestRuleTypeSetPath string = "set-path"

	// HTTPRequestRuleTypeReplacePath captures enum value "replace-path"
	HTTPRequestRuleTypeReplacePath string = "replace-path"

	// HTTPRequestRuleTypeSetQuery captures enum value "set-query"
	HTTPRequestRuleTypeSetQuery string = "set-query"

	// HTTPRequestRuleTypeSetURI captures enum value "set-uri"
	HTTPRequestRuleTypeSetURI string = "set-uri"

	// HTTPRequestRuleTypeSetVar captures enum value "set-var"
	HTTPRequestRuleTypeSetVar string = "set-var"

	// HTTPRequestRuleTypeSendSpoeGroup captures enum value "send-spoe-group"
	HTTPRequestRuleTypeSendSpoeGroup string = "send-spoe-group"

	// HTTPRequestRuleTypeAddACL captures enum value "add-acl"
	HTTPRequestRuleTypeAddACL string = "add-acl"

	// HTTPRequestRuleTypeDelACL captures enum value "del-acl"
	HTTPRequestRuleTypeDelACL string = "del-acl"

	// HTTPRequestRuleTypeCapture captures enum value "capture"
	HTTPRequestRuleTypeCapture string = "capture"

	// HTTPRequestRuleTypeTrackSc0 captures enum value "track-sc0"
	HTTPRequestRuleTypeTrackSc0 string = "track-sc0"

	// HTTPRequestRuleTypeTrackSc1 captures enum value "track-sc1"
	HTTPRequestRuleTypeTrackSc1 string = "track-sc1"

	// HTTPRequestRuleTypeTrackSc2 captures enum value "track-sc2"
	HTTPRequestRuleTypeTrackSc2 string = "track-sc2"

	// HTTPRequestRuleTypeSetMap captures enum value "set-map"
	HTTPRequestRuleTypeSetMap string = "set-map"

	// HTTPRequestRuleTypeDelMap captures enum value "del-map"
	HTTPRequestRuleTypeDelMap string = "del-map"

	// HTTPRequestRuleTypeCacheUse captures enum value "cache-use"
	HTTPRequestRuleTypeCacheUse string = "cache-use"

	// HTTPRequestRuleTypeDisableL7Retry captures enum value "disable-l7-retry"
	HTTPRequestRuleTypeDisableL7Retry string = "disable-l7-retry"

	// HTTPRequestRuleTypeEarlyHint captures enum value "early-hint"
	HTTPRequestRuleTypeEarlyHint string = "early-hint"

	// HTTPRequestRuleTypeReplaceURI captures enum value "replace-uri"
	HTTPRequestRuleTypeReplaceURI string = "replace-uri"

	// HTTPRequestRuleTypeScIncGpc0 captures enum value "sc-inc-gpc0"
	HTTPRequestRuleTypeScIncGpc0 string = "sc-inc-gpc0"

	// HTTPRequestRuleTypeScIncGpc1 captures enum value "sc-inc-gpc1"
	HTTPRequestRuleTypeScIncGpc1 string = "sc-inc-gpc1"

	// HTTPRequestRuleTypeDoResolve captures enum value "do-resolve"
	HTTPRequestRuleTypeDoResolve string = "do-resolve"

	// HTTPRequestRuleTypeSetDst captures enum value "set-dst"
	HTTPRequestRuleTypeSetDst string = "set-dst"

	// HTTPRequestRuleTypeSetDstPort captures enum value "set-dst-port"
	HTTPRequestRuleTypeSetDstPort string = "set-dst-port"

	// HTTPRequestRuleTypeScSetGpt0 captures enum value "sc-set-gpt0"
	HTTPRequestRuleTypeScSetGpt0 string = "sc-set-gpt0"

	// HTTPRequestRuleTypeSetMark captures enum value "set-mark"
	HTTPRequestRuleTypeSetMark string = "set-mark"

	// HTTPRequestRuleTypeSetNice captures enum value "set-nice"
	HTTPRequestRuleTypeSetNice string = "set-nice"

	// HTTPRequestRuleTypeSetMethod captures enum value "set-method"
	HTTPRequestRuleTypeSetMethod string = "set-method"

	// HTTPRequestRuleTypeSetPriorityClass captures enum value "set-priority-class"
	HTTPRequestRuleTypeSetPriorityClass string = "set-priority-class"

	// HTTPRequestRuleTypeSetPriorityOffset captures enum value "set-priority-offset"
	HTTPRequestRuleTypeSetPriorityOffset string = "set-priority-offset"

	// HTTPRequestRuleTypeSetSrc captures enum value "set-src"
	HTTPRequestRuleTypeSetSrc string = "set-src"

	// HTTPRequestRuleTypeSetSrcPor captures enum value "set-src-por"
	HTTPRequestRuleTypeSetSrcPor string = "set-src-por"

	// HTTPRequestRuleTypeWaitForHandshake captures enum value "wait-for-handshake"
	HTTPRequestRuleTypeWaitForHandshake string = "wait-for-handshake"

	// HTTPRequestRuleTypeSetTos captures enum value "set-tos"
	HTTPRequestRuleTypeSetTos string = "set-tos"

	// HTTPRequestRuleTypeSilentDrop captures enum value "silent-drop"
	HTTPRequestRuleTypeSilentDrop string = "silent-drop"

	// HTTPRequestRuleTypeUnsetVar captures enum value "unset-var"
	HTTPRequestRuleTypeUnsetVar string = "unset-var"

	// HTTPRequestRuleTypeStrictMode captures enum value "strict-mode"
	HTTPRequestRuleTypeStrictMode string = "strict-mode"

	// HTTPRequestRuleTypeLua captures enum value "lua"
	HTTPRequestRuleTypeLua string = "lua"

	// HTTPRequestRuleTypeUseService captures enum value "use-service"
	HTTPRequestRuleTypeUseService string = "use-service"

	// HTTPRequestRuleTypeReturn captures enum value "return"
	HTTPRequestRuleTypeReturn string = "return"
)

// prop value enum
func (m *HTTPRequestRule) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, httpRequestRuleTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *HTTPRequestRule) validateType(formats strfmt.Registry) error {

	if err := validate.RequiredString("type", "body", string(m.Type)); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *HTTPRequestRule) validateVarName(formats strfmt.Registry) error {

	if swag.IsZero(m.VarName) { // not required
		return nil
	}

	if err := validate.Pattern("var_name", "body", string(m.VarName), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPRequestRule) validateVarScope(formats strfmt.Registry) error {

	if swag.IsZero(m.VarScope) { // not required
		return nil
	}

	if err := validate.Pattern("var_scope", "body", string(m.VarScope), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HTTPRequestRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HTTPRequestRule) UnmarshalBinary(b []byte) error {
	var res HTTPRequestRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HTTPRequestRuleReturnHdrsItems0 HTTP request rule return hdrs items0
// swagger:model HTTPRequestRuleReturnHdrsItems0
type HTTPRequestRuleReturnHdrsItems0 struct {

	// fmt
	// Required: true
	Fmt *string `json:"fmt"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this HTTP request rule return hdrs items0
func (m *HTTPRequestRuleReturnHdrsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFmt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HTTPRequestRuleReturnHdrsItems0) validateFmt(formats strfmt.Registry) error {

	if err := validate.Required("fmt", "body", m.Fmt); err != nil {
		return err
	}

	return nil
}

func (m *HTTPRequestRuleReturnHdrsItems0) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HTTPRequestRuleReturnHdrsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HTTPRequestRuleReturnHdrsItems0) UnmarshalBinary(b []byte) error {
	var res HTTPRequestRuleReturnHdrsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

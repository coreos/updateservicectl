// Package update provides access to the .
//
// Usage example:
//
//   import "github.com/coreos/updateservicectl/client/update/v1"
//   ...
//   updateService, err := update.New(oauthHttpClient)
package update // import "github.com/coreos/updateservicectl/client/update/v1"

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	context "golang.org/x/net/context"
	ctxhttp "golang.org/x/net/context/ctxhttp"
	gensupport "google.golang.org/api/gensupport"
	googleapi "google.golang.org/api/googleapi"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled
var _ = ctxhttp.Do

const apiId = "update:v1"
const apiName = "update"
const apiVersion = "v1"
const basePath = "http://internal/_ah/api/update/v1/"

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Admin = NewAdminService(s)
	s.App = NewAppService(s)
	s.Appversion = NewAppversionService(s)
	s.Channel = NewChannelService(s)
	s.Client = NewClientService(s)
	s.Clientupdate = NewClientupdateService(s)
	s.Group = NewGroupService(s)
	s.Upstream = NewUpstreamService(s)
	s.Util = NewUtilService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Admin *AdminService

	App *AppService

	Appversion *AppversionService

	Channel *ChannelService

	Client *ClientService

	Clientupdate *ClientupdateService

	Group *GroupService

	Upstream *UpstreamService

	Util *UtilService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewAdminService(s *Service) *AdminService {
	rs := &AdminService{s: s}
	return rs
}

type AdminService struct {
	s *Service
}

func NewAppService(s *Service) *AppService {
	rs := &AppService{s: s}
	rs.Package = NewAppPackageService(s)
	return rs
}

type AppService struct {
	s *Service

	Package *AppPackageService
}

func NewAppPackageService(s *Service) *AppPackageService {
	rs := &AppPackageService{s: s}
	return rs
}

type AppPackageService struct {
	s *Service
}

func NewAppversionService(s *Service) *AppversionService {
	rs := &AppversionService{s: s}
	return rs
}

type AppversionService struct {
	s *Service
}

func NewChannelService(s *Service) *ChannelService {
	rs := &ChannelService{s: s}
	return rs
}

type ChannelService struct {
	s *Service
}

func NewClientService(s *Service) *ClientService {
	rs := &ClientService{s: s}
	return rs
}

type ClientService struct {
	s *Service
}

func NewClientupdateService(s *Service) *ClientupdateService {
	rs := &ClientupdateService{s: s}
	return rs
}

type ClientupdateService struct {
	s *Service
}

func NewGroupService(s *Service) *GroupService {
	rs := &GroupService{s: s}
	rs.Percent = NewGroupPercentService(s)
	rs.Requests = NewGroupRequestsService(s)
	rs.Rollout = NewGroupRolloutService(s)
	return rs
}

type GroupService struct {
	s *Service

	Percent *GroupPercentService

	Requests *GroupRequestsService

	Rollout *GroupRolloutService
}

func NewGroupPercentService(s *Service) *GroupPercentService {
	rs := &GroupPercentService{s: s}
	return rs
}

type GroupPercentService struct {
	s *Service
}

func NewGroupRequestsService(s *Service) *GroupRequestsService {
	rs := &GroupRequestsService{s: s}
	rs.Events = NewGroupRequestsEventsService(s)
	rs.Versions = NewGroupRequestsVersionsService(s)
	return rs
}

type GroupRequestsService struct {
	s *Service

	Events *GroupRequestsEventsService

	Versions *GroupRequestsVersionsService
}

func NewGroupRequestsEventsService(s *Service) *GroupRequestsEventsService {
	rs := &GroupRequestsEventsService{s: s}
	return rs
}

type GroupRequestsEventsService struct {
	s *Service
}

func NewGroupRequestsVersionsService(s *Service) *GroupRequestsVersionsService {
	rs := &GroupRequestsVersionsService{s: s}
	return rs
}

type GroupRequestsVersionsService struct {
	s *Service
}

func NewGroupRolloutService(s *Service) *GroupRolloutService {
	rs := &GroupRolloutService{s: s}
	rs.Active = NewGroupRolloutActiveService(s)
	return rs
}

type GroupRolloutService struct {
	s *Service

	Active *GroupRolloutActiveService
}

func NewGroupRolloutActiveService(s *Service) *GroupRolloutActiveService {
	rs := &GroupRolloutActiveService{s: s}
	return rs
}

type GroupRolloutActiveService struct {
	s *Service
}

func NewUpstreamService(s *Service) *UpstreamService {
	rs := &UpstreamService{s: s}
	return rs
}

type UpstreamService struct {
	s *Service
}

func NewUtilService(s *Service) *UtilService {
	rs := &UtilService{s: s}
	return rs
}

type UtilService struct {
	s *Service
}

type AdminListUsersResp struct {
	Users []*AdminUser `json:"users,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Users") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Users") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *AdminListUsersResp) MarshalJSON() ([]byte, error) {
	type NoMethod AdminListUsersResp
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type AdminUser struct {
	Token string `json:"token,omitempty"`

	User string `json:"user,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Token") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Token") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *AdminUser) MarshalJSON() ([]byte, error) {
	type NoMethod AdminUser
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type AdminUserReq struct {
	UserName string `json:"userName,omitempty"`

	// ForceSendFields is a list of field names (e.g. "UserName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "UserName") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *AdminUserReq) MarshalJSON() ([]byte, error) {
	type NoMethod AdminUserReq
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type App struct {
	Description string `json:"description,omitempty"`

	Id string `json:"id,omitempty"`

	Label string `json:"label,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Description") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Description") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *App) MarshalJSON() ([]byte, error) {
	type NoMethod App
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type AppChannel struct {
	AppId string `json:"appId,omitempty"`

	DateCreated string `json:"dateCreated,omitempty"`

	Label string `json:"label,omitempty"`

	Publish bool `json:"publish,omitempty"`

	Upstream string `json:"upstream,omitempty"`

	Version string `json:"version,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "AppId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AppId") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *AppChannel) MarshalJSON() ([]byte, error) {
	type NoMethod AppChannel
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type AppInsertReq struct {
	Description string `json:"description,omitempty"`

	Id string `json:"id,omitempty"`

	Label string `json:"label,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Description") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Description") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *AppInsertReq) MarshalJSON() ([]byte, error) {
	type NoMethod AppInsertReq
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type AppListResp struct {
	Items []*App `json:"items,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Items") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Items") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *AppListResp) MarshalJSON() ([]byte, error) {
	type NoMethod AppListResp
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type AppUpdateReq struct {
	Description string `json:"description,omitempty"`

	Id string `json:"id,omitempty"`

	Label string `json:"label,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Description") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Description") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *AppUpdateReq) MarshalJSON() ([]byte, error) {
	type NoMethod AppUpdateReq
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type AppVersionItem struct {
	AppId string `json:"appId,omitempty"`

	Count int64 `json:"count,omitempty"`

	GroupId string `json:"groupId,omitempty"`

	Version string `json:"version,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AppId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AppId") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *AppVersionItem) MarshalJSON() ([]byte, error) {
	type NoMethod AppVersionItem
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type AppVersionList struct {
	Items []*AppVersionItem `json:"items,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Items") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Items") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *AppVersionList) MarshalJSON() ([]byte, error) {
	type NoMethod AppVersionList
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ChannelListResp struct {
	Items []*AppChannel `json:"items,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Items") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Items") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ChannelListResp) MarshalJSON() ([]byte, error) {
	type NoMethod ChannelListResp
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ChannelRequest struct {
	AppId string `json:"appId,omitempty"`

	Label string `json:"label,omitempty"`

	Publish bool `json:"publish,omitempty"`

	Version string `json:"version,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "AppId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AppId") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ChannelRequest) MarshalJSON() ([]byte, error) {
	type NoMethod ChannelRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ClientCountResp struct {
	Count int64 `json:"count,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Count") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Count") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ClientCountResp) MarshalJSON() ([]byte, error) {
	type NoMethod ClientCountResp
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ClientHistoryItem struct {
	DateTime int64 `json:"dateTime,omitempty,string"`

	ErrorCode string `json:"errorCode,omitempty"`

	EventResult string `json:"eventResult,omitempty"`

	EventType string `json:"eventType,omitempty"`

	GroupId string `json:"groupId,omitempty"`

	InstallSource string `json:"installSource,omitempty"`

	Version string `json:"version,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DateTime") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DateTime") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ClientHistoryItem) MarshalJSON() ([]byte, error) {
	type NoMethod ClientHistoryItem
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ClientHistoryResp struct {
	Items []*ClientHistoryItem `json:"items,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Items") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Items") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ClientHistoryResp) MarshalJSON() ([]byte, error) {
	type NoMethod ClientHistoryResp
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ClientUpdate struct {
	AppId string `json:"appId,omitempty"`

	ClientId string `json:"clientId,omitempty"`

	ErrorCode string `json:"errorCode,omitempty"`

	EventResult string `json:"eventResult,omitempty"`

	EventType string `json:"eventType,omitempty"`

	GroupId string `json:"groupId,omitempty"`

	LastSeen string `json:"lastSeen,omitempty"`

	Oem string `json:"oem,omitempty"`

	Version string `json:"version,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AppId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AppId") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ClientUpdate) MarshalJSON() ([]byte, error) {
	type NoMethod ClientUpdate
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ClientUpdateList struct {
	Items []*ClientUpdate `json:"items,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Items") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Items") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ClientUpdateList) MarshalJSON() ([]byte, error) {
	type NoMethod ClientUpdateList
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type Frame struct {
	Duration int64 `json:"duration,omitempty,string"`

	Percent float64 `json:"percent,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Duration") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Duration") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Frame) MarshalJSON() ([]byte, error) {
	type NoMethod Frame
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *Frame) UnmarshalJSON(data []byte) error {
	type NoMethod Frame
	var s1 struct {
		Percent gensupport.JSONFloat64 `json:"percent"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.Percent = float64(s1.Percent)
	return nil
}

type GenerateUuidResp struct {
	Uuid string `json:"uuid,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Uuid") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Uuid") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GenerateUuidResp) MarshalJSON() ([]byte, error) {
	type NoMethod GenerateUuidResp
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type Group struct {
	AppId string `json:"appId,omitempty"`

	ChannelId string `json:"channelId,omitempty"`

	Id string `json:"id,omitempty"`

	Label string `json:"label,omitempty"`

	OemBlacklist string `json:"oemBlacklist,omitempty"`

	RolloutActive bool `json:"rolloutActive,omitempty"`

	UpdatePercent float64 `json:"updatePercent,omitempty"`

	UpdatesPaused bool `json:"updatesPaused,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "AppId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AppId") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Group) MarshalJSON() ([]byte, error) {
	type NoMethod Group
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *Group) UnmarshalJSON(data []byte) error {
	type NoMethod Group
	var s1 struct {
		UpdatePercent gensupport.JSONFloat64 `json:"updatePercent"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.UpdatePercent = float64(s1.UpdatePercent)
	return nil
}

type GroupList struct {
	Items []*Group `json:"items,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Items") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Items") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GroupList) MarshalJSON() ([]byte, error) {
	type NoMethod GroupList
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type GroupPercent struct {
	AppId string `json:"appId,omitempty"`

	Id string `json:"id,omitempty"`

	UpdatePercent float64 `json:"updatePercent,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "AppId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AppId") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GroupPercent) MarshalJSON() ([]byte, error) {
	type NoMethod GroupPercent
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *GroupPercent) UnmarshalJSON(data []byte) error {
	type NoMethod GroupPercent
	var s1 struct {
		UpdatePercent gensupport.JSONFloat64 `json:"updatePercent"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.UpdatePercent = float64(s1.UpdatePercent)
	return nil
}

type GroupRequestsItem struct {
	Result string `json:"result,omitempty"`

	Type string `json:"type,omitempty"`

	Values []*GroupRequestsValues `json:"values,omitempty"`

	Version string `json:"version,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Result") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Result") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GroupRequestsItem) MarshalJSON() ([]byte, error) {
	type NoMethod GroupRequestsItem
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type GroupRequestsRollup struct {
	Items []*GroupRequestsItem `json:"items,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Items") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Items") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GroupRequestsRollup) MarshalJSON() ([]byte, error) {
	type NoMethod GroupRequestsRollup
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type GroupRequestsValues struct {
	Count int64 `json:"count,omitempty,string"`

	Timestamp int64 `json:"timestamp,omitempty,string"`

	// ForceSendFields is a list of field names (e.g. "Count") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Count") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GroupRequestsValues) MarshalJSON() ([]byte, error) {
	type NoMethod GroupRequestsValues
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type Package struct {
	AppId string `json:"appId,omitempty"`

	DateCreated string `json:"dateCreated,omitempty"`

	MetadataSignatureRsa string `json:"metadataSignatureRsa,omitempty"`

	MetadataSize string `json:"metadataSize,omitempty"`

	ReleaseNotes string `json:"releaseNotes,omitempty"`

	Required bool `json:"required,omitempty"`

	Sha1Sum string `json:"sha1Sum,omitempty"`

	Sha256Sum string `json:"sha256Sum,omitempty"`

	Size string `json:"size,omitempty"`

	Url string `json:"url,omitempty"`

	Version string `json:"version,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "AppId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AppId") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Package) MarshalJSON() ([]byte, error) {
	type NoMethod Package
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type PackageList struct {
	Items []*Package `json:"items,omitempty"`

	Total int64 `json:"total,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Items") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Items") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *PackageList) MarshalJSON() ([]byte, error) {
	type NoMethod PackageList
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type PublicPackageItem struct {
	AppId string `json:"AppId,omitempty"`

	Packages []*Package `json:"packages,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AppId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AppId") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *PublicPackageItem) MarshalJSON() ([]byte, error) {
	type NoMethod PublicPackageItem
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type PublicPackageList struct {
	Items []*PublicPackageItem `json:"items,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Items") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Items") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *PublicPackageList) MarshalJSON() ([]byte, error) {
	type NoMethod PublicPackageList
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type Rollout struct {
	AppId string `json:"appId,omitempty"`

	GroupId string `json:"groupId,omitempty"`

	Rollout []*Frame `json:"rollout,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "AppId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AppId") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Rollout) MarshalJSON() ([]byte, error) {
	type NoMethod Rollout
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type RolloutActive struct {
	Active bool `json:"active,omitempty"`

	AppId string `json:"appId,omitempty"`

	GroupId string `json:"groupId,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Active") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Active") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *RolloutActive) MarshalJSON() ([]byte, error) {
	type NoMethod RolloutActive
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type Upstream struct {
	Id string `json:"id,omitempty"`

	Label string `json:"label,omitempty"`

	Url string `json:"url,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Id") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Id") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Upstream) MarshalJSON() ([]byte, error) {
	type NoMethod Upstream
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type UpstreamListResp struct {
	Items []*Upstream `json:"items,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Items") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Items") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *UpstreamListResp) MarshalJSON() ([]byte, error) {
	type NoMethod UpstreamListResp
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type UpstreamSyncResp struct {
	Detail string `json:"detail,omitempty"`

	Status string `json:"status,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Detail") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Detail") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *UpstreamSyncResp) MarshalJSON() ([]byte, error) {
	type NoMethod UpstreamSyncResp
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// method id "update.admin.createUser":

type AdminCreateUserCall struct {
	s            *Service
	adminuserreq *AdminUserReq
	urlParams_   gensupport.URLParams
	ctx_         context.Context
	header_      http.Header
}

// CreateUser: Create a new user.
func (r *AdminService) CreateUser(adminuserreq *AdminUserReq) *AdminCreateUserCall {
	c := &AdminCreateUserCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.adminuserreq = adminuserreq
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AdminCreateUserCall) Fields(s ...googleapi.Field) *AdminCreateUserCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AdminCreateUserCall) Context(ctx context.Context) *AdminCreateUserCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AdminCreateUserCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AdminCreateUserCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.adminuserreq)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "admin/user")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "update.admin.createUser" call.
// Exactly one of *AdminUser or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *AdminUser.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *AdminCreateUserCall) Do(opts ...googleapi.CallOption) (*AdminUser, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &AdminUser{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Create a new user.",
	//   "httpMethod": "POST",
	//   "id": "update.admin.createUser",
	//   "path": "admin/user",
	//   "request": {
	//     "$ref": "AdminUserReq",
	//     "parameterName": "resource"
	//   },
	//   "response": {
	//     "$ref": "AdminUser"
	//   }
	// }

}

// method id "update.admin.deleteUser":

type AdminDeleteUserCall struct {
	s          *Service
	userName   string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// DeleteUser: Delete a user.
func (r *AdminService) DeleteUser(userName string) *AdminDeleteUserCall {
	c := &AdminDeleteUserCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.userName = userName
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AdminDeleteUserCall) Fields(s ...googleapi.Field) *AdminDeleteUserCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AdminDeleteUserCall) Context(ctx context.Context) *AdminDeleteUserCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AdminDeleteUserCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AdminDeleteUserCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "admin/user/{userName}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"userName": c.userName,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "update.admin.deleteUser" call.
// Exactly one of *AdminUser or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *AdminUser.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *AdminDeleteUserCall) Do(opts ...googleapi.CallOption) (*AdminUser, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &AdminUser{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Delete a user.",
	//   "httpMethod": "DELETE",
	//   "id": "update.admin.deleteUser",
	//   "parameterOrder": [
	//     "userName"
	//   ],
	//   "parameters": {
	//     "userName": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "admin/user/{userName}",
	//   "response": {
	//     "$ref": "AdminUser"
	//   }
	// }

}

// method id "update.admin.genToken":

type AdminGenTokenCall struct {
	s            *Service
	userName     string
	adminuserreq *AdminUserReq
	urlParams_   gensupport.URLParams
	ctx_         context.Context
	header_      http.Header
}

// GenToken: Generate a new token.
func (r *AdminService) GenToken(userName string, adminuserreq *AdminUserReq) *AdminGenTokenCall {
	c := &AdminGenTokenCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.userName = userName
	c.adminuserreq = adminuserreq
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AdminGenTokenCall) Fields(s ...googleapi.Field) *AdminGenTokenCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AdminGenTokenCall) Context(ctx context.Context) *AdminGenTokenCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AdminGenTokenCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AdminGenTokenCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.adminuserreq)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "admin/user/{userName}/token/new")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"userName": c.userName,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "update.admin.genToken" call.
// Exactly one of *AdminUser or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *AdminUser.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *AdminGenTokenCall) Do(opts ...googleapi.CallOption) (*AdminUser, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &AdminUser{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Generate a new token.",
	//   "httpMethod": "PUT",
	//   "id": "update.admin.genToken",
	//   "parameterOrder": [
	//     "userName"
	//   ],
	//   "parameters": {
	//     "userName": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "admin/user/{userName}/token/new",
	//   "request": {
	//     "$ref": "AdminUserReq",
	//     "parameterName": "resource"
	//   },
	//   "response": {
	//     "$ref": "AdminUser"
	//   }
	// }

}

// method id "update.admin.getUser":

type AdminGetUserCall struct {
	s            *Service
	userName     string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// GetUser: Get a user.
func (r *AdminService) GetUser(userName string) *AdminGetUserCall {
	c := &AdminGetUserCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.userName = userName
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AdminGetUserCall) Fields(s ...googleapi.Field) *AdminGetUserCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *AdminGetUserCall) IfNoneMatch(entityTag string) *AdminGetUserCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AdminGetUserCall) Context(ctx context.Context) *AdminGetUserCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AdminGetUserCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AdminGetUserCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "admin/user/{userName}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"userName": c.userName,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "update.admin.getUser" call.
// Exactly one of *AdminUser or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *AdminUser.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *AdminGetUserCall) Do(opts ...googleapi.CallOption) (*AdminUser, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &AdminUser{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Get a user.",
	//   "httpMethod": "GET",
	//   "id": "update.admin.getUser",
	//   "parameterOrder": [
	//     "userName"
	//   ],
	//   "parameters": {
	//     "userName": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "admin/user/{userName}",
	//   "response": {
	//     "$ref": "AdminUser"
	//   }
	// }

}

// method id "update.admin.listUsers":

type AdminListUsersCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// ListUsers: List Users.
func (r *AdminService) ListUsers() *AdminListUsersCall {
	c := &AdminListUsersCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AdminListUsersCall) Fields(s ...googleapi.Field) *AdminListUsersCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *AdminListUsersCall) IfNoneMatch(entityTag string) *AdminListUsersCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AdminListUsersCall) Context(ctx context.Context) *AdminListUsersCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AdminListUsersCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AdminListUsersCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "admin/user")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "update.admin.listUsers" call.
// Exactly one of *AdminListUsersResp or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *AdminListUsersResp.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *AdminListUsersCall) Do(opts ...googleapi.CallOption) (*AdminListUsersResp, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &AdminListUsersResp{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List Users.",
	//   "httpMethod": "GET",
	//   "id": "update.admin.listUsers",
	//   "path": "admin/user",
	//   "response": {
	//     "$ref": "AdminListUsersResp"
	//   }
	// }

}

// method id "update.app.delete":

type AppDeleteCall struct {
	s          *Service
	id         string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Delete an application.
func (r *AppService) Delete(id string) *AppDeleteCall {
	c := &AppDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.id = id
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AppDeleteCall) Fields(s ...googleapi.Field) *AppDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AppDeleteCall) Context(ctx context.Context) *AppDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AppDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AppDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "apps/{id}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "update.app.delete" call.
// Exactly one of *App or error will be non-nil. Any non-2xx status code
// is an error. Response headers are in either
// *App.ServerResponse.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *AppDeleteCall) Do(opts ...googleapi.CallOption) (*App, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &App{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Delete an application.",
	//   "httpMethod": "DELETE",
	//   "id": "update.app.delete",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "apps/{id}",
	//   "response": {
	//     "$ref": "App"
	//   }
	// }

}

// method id "update.app.get":

type AppGetCall struct {
	s            *Service
	id           string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Get an application.
func (r *AppService) Get(id string) *AppGetCall {
	c := &AppGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.id = id
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AppGetCall) Fields(s ...googleapi.Field) *AppGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *AppGetCall) IfNoneMatch(entityTag string) *AppGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AppGetCall) Context(ctx context.Context) *AppGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AppGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AppGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "apps/{id}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "update.app.get" call.
// Exactly one of *App or error will be non-nil. Any non-2xx status code
// is an error. Response headers are in either
// *App.ServerResponse.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *AppGetCall) Do(opts ...googleapi.CallOption) (*App, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &App{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Get an application.",
	//   "httpMethod": "GET",
	//   "id": "update.app.get",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "apps/{id}",
	//   "response": {
	//     "$ref": "App"
	//   }
	// }

}

// method id "update.app.insert":

type AppInsertCall struct {
	s            *Service
	appinsertreq *AppInsertReq
	urlParams_   gensupport.URLParams
	ctx_         context.Context
	header_      http.Header
}

// Insert: Insert an application.
func (r *AppService) Insert(appinsertreq *AppInsertReq) *AppInsertCall {
	c := &AppInsertCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.appinsertreq = appinsertreq
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AppInsertCall) Fields(s ...googleapi.Field) *AppInsertCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AppInsertCall) Context(ctx context.Context) *AppInsertCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AppInsertCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AppInsertCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.appinsertreq)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "apps")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "update.app.insert" call.
// Exactly one of *App or error will be non-nil. Any non-2xx status code
// is an error. Response headers are in either
// *App.ServerResponse.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *AppInsertCall) Do(opts ...googleapi.CallOption) (*App, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &App{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Insert an application.",
	//   "httpMethod": "POST",
	//   "id": "update.app.insert",
	//   "path": "apps",
	//   "request": {
	//     "$ref": "AppInsertReq",
	//     "parameterName": "resource"
	//   },
	//   "response": {
	//     "$ref": "App"
	//   }
	// }

}

// method id "update.app.list":

type AppListCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: List all application.
func (r *AppService) List() *AppListCall {
	c := &AppListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AppListCall) Fields(s ...googleapi.Field) *AppListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *AppListCall) IfNoneMatch(entityTag string) *AppListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AppListCall) Context(ctx context.Context) *AppListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AppListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AppListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "apps")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "update.app.list" call.
// Exactly one of *AppListResp or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *AppListResp.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *AppListCall) Do(opts ...googleapi.CallOption) (*AppListResp, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &AppListResp{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List all application.",
	//   "httpMethod": "GET",
	//   "id": "update.app.list",
	//   "path": "apps",
	//   "response": {
	//     "$ref": "AppListResp"
	//   }
	// }

}

// method id "update.app.patch":

type AppPatchCall struct {
	s            *Service
	id           string
	appupdatereq *AppUpdateReq
	urlParams_   gensupport.URLParams
	ctx_         context.Context
	header_      http.Header
}

// Patch: Update an application. This method supports patch semantics.
func (r *AppService) Patch(id string, appupdatereq *AppUpdateReq) *AppPatchCall {
	c := &AppPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.id = id
	c.appupdatereq = appupdatereq
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AppPatchCall) Fields(s ...googleapi.Field) *AppPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AppPatchCall) Context(ctx context.Context) *AppPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AppPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AppPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.appupdatereq)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "apps/{id}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "update.app.patch" call.
// Exactly one of *App or error will be non-nil. Any non-2xx status code
// is an error. Response headers are in either
// *App.ServerResponse.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *AppPatchCall) Do(opts ...googleapi.CallOption) (*App, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &App{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Update an application. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "update.app.patch",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "apps/{id}",
	//   "request": {
	//     "$ref": "AppUpdateReq"
	//   },
	//   "response": {
	//     "$ref": "App"
	//   }
	// }

}

// method id "update.app.update":

type AppUpdateCall struct {
	s            *Service
	id           string
	appupdatereq *AppUpdateReq
	urlParams_   gensupport.URLParams
	ctx_         context.Context
	header_      http.Header
}

// Update: Update an application.
func (r *AppService) Update(id string, appupdatereq *AppUpdateReq) *AppUpdateCall {
	c := &AppUpdateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.id = id
	c.appupdatereq = appupdatereq
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AppUpdateCall) Fields(s ...googleapi.Field) *AppUpdateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AppUpdateCall) Context(ctx context.Context) *AppUpdateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AppUpdateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AppUpdateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.appupdatereq)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "apps/{id}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "update.app.update" call.
// Exactly one of *App or error will be non-nil. Any non-2xx status code
// is an error. Response headers are in either
// *App.ServerResponse.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *AppUpdateCall) Do(opts ...googleapi.CallOption) (*App, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &App{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Update an application.",
	//   "httpMethod": "PATCH",
	//   "id": "update.app.update",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "apps/{id}",
	//   "request": {
	//     "$ref": "AppUpdateReq",
	//     "parameterName": "resource"
	//   },
	//   "response": {
	//     "$ref": "App"
	//   }
	// }

}

// method id "update.app.package.delete":

type AppPackageDeleteCall struct {
	s          *Service
	appId      string
	version    string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Delete an package.
func (r *AppPackageService) Delete(appId string, version string) *AppPackageDeleteCall {
	c := &AppPackageDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.appId = appId
	c.version = version
	return c
}

// MetadataSignatureRsa sets the optional parameter
// "metadataSignatureRsa":
func (c *AppPackageDeleteCall) MetadataSignatureRsa(metadataSignatureRsa string) *AppPackageDeleteCall {
	c.urlParams_.Set("metadataSignatureRsa", metadataSignatureRsa)
	return c
}

// MetadataSize sets the optional parameter "metadataSize":
func (c *AppPackageDeleteCall) MetadataSize(metadataSize string) *AppPackageDeleteCall {
	c.urlParams_.Set("metadataSize", metadataSize)
	return c
}

// ReleaseNotes sets the optional parameter "releaseNotes":
func (c *AppPackageDeleteCall) ReleaseNotes(releaseNotes string) *AppPackageDeleteCall {
	c.urlParams_.Set("releaseNotes", releaseNotes)
	return c
}

// Required sets the optional parameter "required":
func (c *AppPackageDeleteCall) Required(required bool) *AppPackageDeleteCall {
	c.urlParams_.Set("required", fmt.Sprint(required))
	return c
}

// Sha1Sum sets the optional parameter "sha1Sum":
func (c *AppPackageDeleteCall) Sha1Sum(sha1Sum string) *AppPackageDeleteCall {
	c.urlParams_.Set("sha1Sum", sha1Sum)
	return c
}

// Sha256Sum sets the optional parameter "sha256Sum":
func (c *AppPackageDeleteCall) Sha256Sum(sha256Sum string) *AppPackageDeleteCall {
	c.urlParams_.Set("sha256Sum", sha256Sum)
	return c
}

// Size sets the optional parameter "size":
func (c *AppPackageDeleteCall) Size(size string) *AppPackageDeleteCall {
	c.urlParams_.Set("size", size)
	return c
}

// Url sets the optional parameter "url":
func (c *AppPackageDeleteCall) Url(url string) *AppPackageDeleteCall {
	c.urlParams_.Set("url", url)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AppPackageDeleteCall) Fields(s ...googleapi.Field) *AppPackageDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AppPackageDeleteCall) Context(ctx context.Context) *AppPackageDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AppPackageDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AppPackageDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "apps/{appId}/packages/{version}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"appId":   c.appId,
		"version": c.version,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "update.app.package.delete" call.
// Exactly one of *Package or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Package.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *AppPackageDeleteCall) Do(opts ...googleapi.CallOption) (*Package, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Package{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Delete an package.",
	//   "httpMethod": "DELETE",
	//   "id": "update.app.package.delete",
	//   "parameterOrder": [
	//     "appId",
	//     "version"
	//   ],
	//   "parameters": {
	//     "appId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "metadataSignatureRsa": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "metadataSize": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "releaseNotes": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "required": {
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "sha1Sum": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sha256Sum": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "size": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "url": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "version": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "apps/{appId}/packages/{version}",
	//   "response": {
	//     "$ref": "Package"
	//   }
	// }

}

// method id "update.app.package.insert":

type AppPackageInsertCall struct {
	s          *Service
	appId      string
	version    string
	package_   *Package
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Insert: Insert a new package version.
func (r *AppPackageService) Insert(appId string, version string, package_ *Package) *AppPackageInsertCall {
	c := &AppPackageInsertCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.appId = appId
	c.version = version
	c.package_ = package_
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AppPackageInsertCall) Fields(s ...googleapi.Field) *AppPackageInsertCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AppPackageInsertCall) Context(ctx context.Context) *AppPackageInsertCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AppPackageInsertCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AppPackageInsertCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.package_)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "apps/{appId}/packages/{version}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"appId":   c.appId,
		"version": c.version,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "update.app.package.insert" call.
// Exactly one of *Package or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Package.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *AppPackageInsertCall) Do(opts ...googleapi.CallOption) (*Package, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Package{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Insert a new package version.",
	//   "httpMethod": "POST",
	//   "id": "update.app.package.insert",
	//   "parameterOrder": [
	//     "appId",
	//     "version"
	//   ],
	//   "parameters": {
	//     "appId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "version": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "apps/{appId}/packages/{version}",
	//   "request": {
	//     "$ref": "Package",
	//     "parameterName": "resource"
	//   },
	//   "response": {
	//     "$ref": "Package"
	//   }
	// }

}

// method id "update.app.package.list":

type AppPackageListCall struct {
	s            *Service
	appId        string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: List all of the package versions.
func (r *AppPackageService) List(appId string) *AppPackageListCall {
	c := &AppPackageListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.appId = appId
	return c
}

// Limit sets the optional parameter "limit":
func (c *AppPackageListCall) Limit(limit int64) *AppPackageListCall {
	c.urlParams_.Set("limit", fmt.Sprint(limit))
	return c
}

// Skip sets the optional parameter "skip":
func (c *AppPackageListCall) Skip(skip int64) *AppPackageListCall {
	c.urlParams_.Set("skip", fmt.Sprint(skip))
	return c
}

// Version sets the optional parameter "version":
func (c *AppPackageListCall) Version(version string) *AppPackageListCall {
	c.urlParams_.Set("version", version)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AppPackageListCall) Fields(s ...googleapi.Field) *AppPackageListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *AppPackageListCall) IfNoneMatch(entityTag string) *AppPackageListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AppPackageListCall) Context(ctx context.Context) *AppPackageListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AppPackageListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AppPackageListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "apps/{appId}/packages")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"appId": c.appId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "update.app.package.list" call.
// Exactly one of *PackageList or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *PackageList.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *AppPackageListCall) Do(opts ...googleapi.CallOption) (*PackageList, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &PackageList{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List all of the package versions.",
	//   "httpMethod": "GET",
	//   "id": "update.app.package.list",
	//   "parameterOrder": [
	//     "appId"
	//   ],
	//   "parameters": {
	//     "appId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "limit": {
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "skip": {
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "version": {
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "apps/{appId}/packages",
	//   "response": {
	//     "$ref": "PackageList"
	//   }
	// }

}

// method id "update.app.package.publicList":

type AppPackagePublicListCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// PublicList: List all of the publicly available published packages.
func (r *AppPackageService) PublicList() *AppPackagePublicListCall {
	c := &AppPackagePublicListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AppPackagePublicListCall) Fields(s ...googleapi.Field) *AppPackagePublicListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *AppPackagePublicListCall) IfNoneMatch(entityTag string) *AppPackagePublicListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AppPackagePublicListCall) Context(ctx context.Context) *AppPackagePublicListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AppPackagePublicListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AppPackagePublicListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "public/packages")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "update.app.package.publicList" call.
// Exactly one of *PublicPackageList or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *PublicPackageList.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *AppPackagePublicListCall) Do(opts ...googleapi.CallOption) (*PublicPackageList, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &PublicPackageList{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List all of the publicly available published packages.",
	//   "httpMethod": "GET",
	//   "id": "update.app.package.publicList",
	//   "path": "public/packages",
	//   "response": {
	//     "$ref": "PublicPackageList"
	//   }
	// }

}

// method id "update.appversion.list":

type AppversionListCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: List Client updates grouped by app/version.
func (r *AppversionService) List() *AppversionListCall {
	c := &AppversionListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// AppId sets the optional parameter "appId":
func (c *AppversionListCall) AppId(appId string) *AppversionListCall {
	c.urlParams_.Set("appId", appId)
	return c
}

// DateEnd sets the optional parameter "dateEnd":
func (c *AppversionListCall) DateEnd(dateEnd int64) *AppversionListCall {
	c.urlParams_.Set("dateEnd", fmt.Sprint(dateEnd))
	return c
}

// DateStart sets the optional parameter "dateStart":
func (c *AppversionListCall) DateStart(dateStart int64) *AppversionListCall {
	c.urlParams_.Set("dateStart", fmt.Sprint(dateStart))
	return c
}

// EventResult sets the optional parameter "eventResult":
func (c *AppversionListCall) EventResult(eventResult string) *AppversionListCall {
	c.urlParams_.Set("eventResult", eventResult)
	return c
}

// EventType sets the optional parameter "eventType":
func (c *AppversionListCall) EventType(eventType string) *AppversionListCall {
	c.urlParams_.Set("eventType", eventType)
	return c
}

// GroupId sets the optional parameter "groupId":
func (c *AppversionListCall) GroupId(groupId string) *AppversionListCall {
	c.urlParams_.Set("groupId", groupId)
	return c
}

// Oem sets the optional parameter "oem":
func (c *AppversionListCall) Oem(oem string) *AppversionListCall {
	c.urlParams_.Set("oem", oem)
	return c
}

// Version sets the optional parameter "version":
func (c *AppversionListCall) Version(version string) *AppversionListCall {
	c.urlParams_.Set("version", version)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AppversionListCall) Fields(s ...googleapi.Field) *AppversionListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *AppversionListCall) IfNoneMatch(entityTag string) *AppversionListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AppversionListCall) Context(ctx context.Context) *AppversionListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AppversionListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AppversionListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "appversions")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "update.appversion.list" call.
// Exactly one of *AppVersionList or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *AppVersionList.ServerResponse.Header or (if a response was returned
// at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *AppversionListCall) Do(opts ...googleapi.CallOption) (*AppVersionList, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &AppVersionList{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List Client updates grouped by app/version.",
	//   "httpMethod": "GET",
	//   "id": "update.appversion.list",
	//   "parameters": {
	//     "appId": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "dateEnd": {
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "dateStart": {
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "eventResult": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "eventType": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "groupId": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "oem": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "version": {
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "appversions",
	//   "response": {
	//     "$ref": "AppVersionList"
	//   }
	// }

}

// method id "update.channel.delete":

type ChannelDeleteCall struct {
	s          *Service
	appId      string
	label      string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Delete a channel.
func (r *ChannelService) Delete(appId string, label string) *ChannelDeleteCall {
	c := &ChannelDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.appId = appId
	c.label = label
	return c
}

// Publish sets the optional parameter "publish":
func (c *ChannelDeleteCall) Publish(publish bool) *ChannelDeleteCall {
	c.urlParams_.Set("publish", fmt.Sprint(publish))
	return c
}

// Version sets the optional parameter "version":
func (c *ChannelDeleteCall) Version(version string) *ChannelDeleteCall {
	c.urlParams_.Set("version", version)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ChannelDeleteCall) Fields(s ...googleapi.Field) *ChannelDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ChannelDeleteCall) Context(ctx context.Context) *ChannelDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ChannelDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ChannelDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "apps/{appId}/channels/{label}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"appId": c.appId,
		"label": c.label,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "update.channel.delete" call.
// Exactly one of *ChannelRequest or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *ChannelRequest.ServerResponse.Header or (if a response was returned
// at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ChannelDeleteCall) Do(opts ...googleapi.CallOption) (*ChannelRequest, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ChannelRequest{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Delete a channel.",
	//   "httpMethod": "DELETE",
	//   "id": "update.channel.delete",
	//   "parameterOrder": [
	//     "appId",
	//     "label"
	//   ],
	//   "parameters": {
	//     "appId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "label": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "publish": {
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "version": {
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "apps/{appId}/channels/{label}",
	//   "response": {
	//     "$ref": "ChannelRequest"
	//   }
	// }

}

// method id "update.channel.insert":

type ChannelInsertCall struct {
	s              *Service
	appId          string
	channelrequest *ChannelRequest
	urlParams_     gensupport.URLParams
	ctx_           context.Context
	header_        http.Header
}

// Insert: Insert a channel.
func (r *ChannelService) Insert(appId string, channelrequest *ChannelRequest) *ChannelInsertCall {
	c := &ChannelInsertCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.appId = appId
	c.channelrequest = channelrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ChannelInsertCall) Fields(s ...googleapi.Field) *ChannelInsertCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ChannelInsertCall) Context(ctx context.Context) *ChannelInsertCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ChannelInsertCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ChannelInsertCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.channelrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "apps/{appId}/channels")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"appId": c.appId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "update.channel.insert" call.
// Exactly one of *AppChannel or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *AppChannel.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ChannelInsertCall) Do(opts ...googleapi.CallOption) (*AppChannel, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &AppChannel{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Insert a channel.",
	//   "httpMethod": "POST",
	//   "id": "update.channel.insert",
	//   "parameterOrder": [
	//     "appId"
	//   ],
	//   "parameters": {
	//     "appId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "apps/{appId}/channels",
	//   "request": {
	//     "$ref": "ChannelRequest",
	//     "parameterName": "resource"
	//   },
	//   "response": {
	//     "$ref": "AppChannel"
	//   }
	// }

}

// method id "update.channel.list":

type ChannelListCall struct {
	s            *Service
	appId        string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: List channels.
func (r *ChannelService) List(appId string) *ChannelListCall {
	c := &ChannelListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.appId = appId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ChannelListCall) Fields(s ...googleapi.Field) *ChannelListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ChannelListCall) IfNoneMatch(entityTag string) *ChannelListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ChannelListCall) Context(ctx context.Context) *ChannelListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ChannelListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ChannelListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "apps/{appId}/channels")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"appId": c.appId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "update.channel.list" call.
// Exactly one of *ChannelListResp or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *ChannelListResp.ServerResponse.Header or (if a response was returned
// at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ChannelListCall) Do(opts ...googleapi.CallOption) (*ChannelListResp, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ChannelListResp{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List channels.",
	//   "httpMethod": "GET",
	//   "id": "update.channel.list",
	//   "parameterOrder": [
	//     "appId"
	//   ],
	//   "parameters": {
	//     "appId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "apps/{appId}/channels",
	//   "response": {
	//     "$ref": "ChannelListResp"
	//   }
	// }

}

// method id "update.channel.publicList":

type ChannelPublicListCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// PublicList: List all publicly available published channels.
func (r *ChannelService) PublicList() *ChannelPublicListCall {
	c := &ChannelPublicListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ChannelPublicListCall) Fields(s ...googleapi.Field) *ChannelPublicListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ChannelPublicListCall) IfNoneMatch(entityTag string) *ChannelPublicListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ChannelPublicListCall) Context(ctx context.Context) *ChannelPublicListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ChannelPublicListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ChannelPublicListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "public/channels")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "update.channel.publicList" call.
// Exactly one of *ChannelListResp or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *ChannelListResp.ServerResponse.Header or (if a response was returned
// at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ChannelPublicListCall) Do(opts ...googleapi.CallOption) (*ChannelListResp, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ChannelListResp{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List all publicly available published channels.",
	//   "httpMethod": "GET",
	//   "id": "update.channel.publicList",
	//   "path": "public/channels",
	//   "response": {
	//     "$ref": "ChannelListResp"
	//   }
	// }

}

// method id "update.channel.update":

type ChannelUpdateCall struct {
	s              *Service
	appId          string
	label          string
	channelrequest *ChannelRequest
	urlParams_     gensupport.URLParams
	ctx_           context.Context
	header_        http.Header
}

// Update: Update a channel.
func (r *ChannelService) Update(appId string, label string, channelrequest *ChannelRequest) *ChannelUpdateCall {
	c := &ChannelUpdateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.appId = appId
	c.label = label
	c.channelrequest = channelrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ChannelUpdateCall) Fields(s ...googleapi.Field) *ChannelUpdateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ChannelUpdateCall) Context(ctx context.Context) *ChannelUpdateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ChannelUpdateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ChannelUpdateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.channelrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "apps/{appId}/channels/{label}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"appId": c.appId,
		"label": c.label,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "update.channel.update" call.
// Exactly one of *AppChannel or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *AppChannel.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ChannelUpdateCall) Do(opts ...googleapi.CallOption) (*AppChannel, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &AppChannel{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Update a channel.",
	//   "httpMethod": "PATCH",
	//   "id": "update.channel.update",
	//   "parameterOrder": [
	//     "appId",
	//     "label"
	//   ],
	//   "parameters": {
	//     "appId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "label": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "apps/{appId}/channels/{label}",
	//   "request": {
	//     "$ref": "ChannelRequest",
	//     "parameterName": "resource"
	//   },
	//   "response": {
	//     "$ref": "AppChannel"
	//   }
	// }

}

// method id "update.client.history":

type ClientHistoryCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// History: Get the update history of a single client.
func (r *ClientService) History(clientId string) *ClientHistoryCall {
	c := &ClientHistoryCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.urlParams_.Set("clientId", clientId)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ClientHistoryCall) Fields(s ...googleapi.Field) *ClientHistoryCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ClientHistoryCall) IfNoneMatch(entityTag string) *ClientHistoryCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ClientHistoryCall) Context(ctx context.Context) *ClientHistoryCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ClientHistoryCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ClientHistoryCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "client/history")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "update.client.history" call.
// Exactly one of *ClientHistoryResp or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ClientHistoryResp.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ClientHistoryCall) Do(opts ...googleapi.CallOption) (*ClientHistoryResp, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ClientHistoryResp{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Get the update history of a single client.",
	//   "httpMethod": "GET",
	//   "id": "update.client.history",
	//   "parameterOrder": [
	//     "clientId"
	//   ],
	//   "parameters": {
	//     "clientId": {
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "client/history",
	//   "response": {
	//     "$ref": "ClientHistoryResp"
	//   }
	// }

}

// method id "update.clientupdate.count":

type ClientupdateCountCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Count: Get client count for criteria.
func (r *ClientupdateService) Count() *ClientupdateCountCall {
	c := &ClientupdateCountCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// AppId sets the optional parameter "appId":
func (c *ClientupdateCountCall) AppId(appId string) *ClientupdateCountCall {
	c.urlParams_.Set("appId", appId)
	return c
}

// DateEnd sets the optional parameter "dateEnd":
func (c *ClientupdateCountCall) DateEnd(dateEnd int64) *ClientupdateCountCall {
	c.urlParams_.Set("dateEnd", fmt.Sprint(dateEnd))
	return c
}

// DateStart sets the optional parameter "dateStart":
func (c *ClientupdateCountCall) DateStart(dateStart int64) *ClientupdateCountCall {
	c.urlParams_.Set("dateStart", fmt.Sprint(dateStart))
	return c
}

// EventResult sets the optional parameter "eventResult":
func (c *ClientupdateCountCall) EventResult(eventResult string) *ClientupdateCountCall {
	c.urlParams_.Set("eventResult", eventResult)
	return c
}

// EventType sets the optional parameter "eventType":
func (c *ClientupdateCountCall) EventType(eventType string) *ClientupdateCountCall {
	c.urlParams_.Set("eventType", eventType)
	return c
}

// GroupId sets the optional parameter "groupId":
func (c *ClientupdateCountCall) GroupId(groupId string) *ClientupdateCountCall {
	c.urlParams_.Set("groupId", groupId)
	return c
}

// Oem sets the optional parameter "oem":
func (c *ClientupdateCountCall) Oem(oem string) *ClientupdateCountCall {
	c.urlParams_.Set("oem", oem)
	return c
}

// Version sets the optional parameter "version":
func (c *ClientupdateCountCall) Version(version string) *ClientupdateCountCall {
	c.urlParams_.Set("version", version)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ClientupdateCountCall) Fields(s ...googleapi.Field) *ClientupdateCountCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ClientupdateCountCall) IfNoneMatch(entityTag string) *ClientupdateCountCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ClientupdateCountCall) Context(ctx context.Context) *ClientupdateCountCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ClientupdateCountCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ClientupdateCountCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "clientupdatecount")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "update.clientupdate.count" call.
// Exactly one of *ClientCountResp or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *ClientCountResp.ServerResponse.Header or (if a response was returned
// at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ClientupdateCountCall) Do(opts ...googleapi.CallOption) (*ClientCountResp, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ClientCountResp{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Get client count for criteria.",
	//   "httpMethod": "GET",
	//   "id": "update.clientupdate.count",
	//   "parameters": {
	//     "appId": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "dateEnd": {
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "dateStart": {
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "eventResult": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "eventType": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "groupId": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "oem": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "version": {
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "clientupdatecount",
	//   "response": {
	//     "$ref": "ClientCountResp"
	//   }
	// }

}

// method id "update.clientupdate.list":

type ClientupdateListCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: List all client updates.
func (r *ClientupdateService) List() *ClientupdateListCall {
	c := &ClientupdateListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// AppId sets the optional parameter "appId":
func (c *ClientupdateListCall) AppId(appId string) *ClientupdateListCall {
	c.urlParams_.Set("appId", appId)
	return c
}

// ClientId sets the optional parameter "clientId":
func (c *ClientupdateListCall) ClientId(clientId string) *ClientupdateListCall {
	c.urlParams_.Set("clientId", clientId)
	return c
}

// DateEnd sets the optional parameter "dateEnd":
func (c *ClientupdateListCall) DateEnd(dateEnd int64) *ClientupdateListCall {
	c.urlParams_.Set("dateEnd", fmt.Sprint(dateEnd))
	return c
}

// DateStart sets the optional parameter "dateStart":
func (c *ClientupdateListCall) DateStart(dateStart int64) *ClientupdateListCall {
	c.urlParams_.Set("dateStart", fmt.Sprint(dateStart))
	return c
}

// EventResult sets the optional parameter "eventResult":
func (c *ClientupdateListCall) EventResult(eventResult string) *ClientupdateListCall {
	c.urlParams_.Set("eventResult", eventResult)
	return c
}

// EventType sets the optional parameter "eventType":
func (c *ClientupdateListCall) EventType(eventType string) *ClientupdateListCall {
	c.urlParams_.Set("eventType", eventType)
	return c
}

// GroupId sets the optional parameter "groupId":
func (c *ClientupdateListCall) GroupId(groupId string) *ClientupdateListCall {
	c.urlParams_.Set("groupId", groupId)
	return c
}

// Limit sets the optional parameter "limit":
func (c *ClientupdateListCall) Limit(limit int64) *ClientupdateListCall {
	c.urlParams_.Set("limit", fmt.Sprint(limit))
	return c
}

// Oem sets the optional parameter "oem":
func (c *ClientupdateListCall) Oem(oem string) *ClientupdateListCall {
	c.urlParams_.Set("oem", oem)
	return c
}

// Skip sets the optional parameter "skip":
func (c *ClientupdateListCall) Skip(skip int64) *ClientupdateListCall {
	c.urlParams_.Set("skip", fmt.Sprint(skip))
	return c
}

// Version sets the optional parameter "version":
func (c *ClientupdateListCall) Version(version string) *ClientupdateListCall {
	c.urlParams_.Set("version", version)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ClientupdateListCall) Fields(s ...googleapi.Field) *ClientupdateListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ClientupdateListCall) IfNoneMatch(entityTag string) *ClientupdateListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ClientupdateListCall) Context(ctx context.Context) *ClientupdateListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ClientupdateListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ClientupdateListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "clientupdates")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "update.clientupdate.list" call.
// Exactly one of *ClientUpdateList or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ClientUpdateList.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ClientupdateListCall) Do(opts ...googleapi.CallOption) (*ClientUpdateList, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ClientUpdateList{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List all client updates.",
	//   "httpMethod": "GET",
	//   "id": "update.clientupdate.list",
	//   "parameters": {
	//     "appId": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "clientId": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "dateEnd": {
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "dateStart": {
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "eventResult": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "eventType": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "groupId": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "limit": {
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "oem": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "skip": {
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "version": {
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "clientupdates",
	//   "response": {
	//     "$ref": "ClientUpdateList"
	//   }
	// }

}

// method id "update.group.delete":

type GroupDeleteCall struct {
	s          *Service
	appId      string
	id         string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Delete a group.
func (r *GroupService) Delete(appId string, id string) *GroupDeleteCall {
	c := &GroupDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.appId = appId
	c.id = id
	return c
}

// ChannelId sets the optional parameter "channelId":
func (c *GroupDeleteCall) ChannelId(channelId string) *GroupDeleteCall {
	c.urlParams_.Set("channelId", channelId)
	return c
}

// Label sets the optional parameter "label":
func (c *GroupDeleteCall) Label(label string) *GroupDeleteCall {
	c.urlParams_.Set("label", label)
	return c
}

// OemBlacklist sets the optional parameter "oemBlacklist":
func (c *GroupDeleteCall) OemBlacklist(oemBlacklist string) *GroupDeleteCall {
	c.urlParams_.Set("oemBlacklist", oemBlacklist)
	return c
}

// RolloutActive sets the optional parameter "rolloutActive":
func (c *GroupDeleteCall) RolloutActive(rolloutActive bool) *GroupDeleteCall {
	c.urlParams_.Set("rolloutActive", fmt.Sprint(rolloutActive))
	return c
}

// UpdatePercent sets the optional parameter "updatePercent":
func (c *GroupDeleteCall) UpdatePercent(updatePercent float64) *GroupDeleteCall {
	c.urlParams_.Set("updatePercent", fmt.Sprint(updatePercent))
	return c
}

// UpdatesPaused sets the optional parameter "updatesPaused":
func (c *GroupDeleteCall) UpdatesPaused(updatesPaused bool) *GroupDeleteCall {
	c.urlParams_.Set("updatesPaused", fmt.Sprint(updatesPaused))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GroupDeleteCall) Fields(s ...googleapi.Field) *GroupDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *GroupDeleteCall) Context(ctx context.Context) *GroupDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *GroupDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *GroupDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "apps/{appId}/groups/{id}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"appId": c.appId,
		"id":    c.id,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "update.group.delete" call.
// Exactly one of *Group or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Group.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *GroupDeleteCall) Do(opts ...googleapi.CallOption) (*Group, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Group{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Delete a group.",
	//   "httpMethod": "DELETE",
	//   "id": "update.group.delete",
	//   "parameterOrder": [
	//     "appId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "appId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "channelId": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "id": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "label": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "oemBlacklist": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "rolloutActive": {
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "updatePercent": {
	//       "format": "double",
	//       "location": "query",
	//       "type": "number"
	//     },
	//     "updatesPaused": {
	//       "location": "query",
	//       "type": "boolean"
	//     }
	//   },
	//   "path": "apps/{appId}/groups/{id}",
	//   "response": {
	//     "$ref": "Group"
	//   }
	// }

}

// method id "update.group.get":

type GroupGetCall struct {
	s            *Service
	appId        string
	id           string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Get a group.
func (r *GroupService) Get(appId string, id string) *GroupGetCall {
	c := &GroupGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.appId = appId
	c.id = id
	return c
}

// ChannelId sets the optional parameter "channelId":
func (c *GroupGetCall) ChannelId(channelId string) *GroupGetCall {
	c.urlParams_.Set("channelId", channelId)
	return c
}

// Label sets the optional parameter "label":
func (c *GroupGetCall) Label(label string) *GroupGetCall {
	c.urlParams_.Set("label", label)
	return c
}

// OemBlacklist sets the optional parameter "oemBlacklist":
func (c *GroupGetCall) OemBlacklist(oemBlacklist string) *GroupGetCall {
	c.urlParams_.Set("oemBlacklist", oemBlacklist)
	return c
}

// RolloutActive sets the optional parameter "rolloutActive":
func (c *GroupGetCall) RolloutActive(rolloutActive bool) *GroupGetCall {
	c.urlParams_.Set("rolloutActive", fmt.Sprint(rolloutActive))
	return c
}

// UpdatePercent sets the optional parameter "updatePercent":
func (c *GroupGetCall) UpdatePercent(updatePercent float64) *GroupGetCall {
	c.urlParams_.Set("updatePercent", fmt.Sprint(updatePercent))
	return c
}

// UpdatesPaused sets the optional parameter "updatesPaused":
func (c *GroupGetCall) UpdatesPaused(updatesPaused bool) *GroupGetCall {
	c.urlParams_.Set("updatesPaused", fmt.Sprint(updatesPaused))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GroupGetCall) Fields(s ...googleapi.Field) *GroupGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *GroupGetCall) IfNoneMatch(entityTag string) *GroupGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *GroupGetCall) Context(ctx context.Context) *GroupGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *GroupGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *GroupGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "apps/{appId}/groups/{id}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"appId": c.appId,
		"id":    c.id,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "update.group.get" call.
// Exactly one of *Group or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Group.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *GroupGetCall) Do(opts ...googleapi.CallOption) (*Group, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Group{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Get a group.",
	//   "httpMethod": "GET",
	//   "id": "update.group.get",
	//   "parameterOrder": [
	//     "appId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "appId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "channelId": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "id": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "label": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "oemBlacklist": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "rolloutActive": {
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "updatePercent": {
	//       "format": "double",
	//       "location": "query",
	//       "type": "number"
	//     },
	//     "updatesPaused": {
	//       "location": "query",
	//       "type": "boolean"
	//     }
	//   },
	//   "path": "apps/{appId}/groups/{id}",
	//   "response": {
	//     "$ref": "Group"
	//   }
	// }

}

// method id "update.group.insert":

type GroupInsertCall struct {
	s          *Service
	appId      string
	group      *Group
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Insert: Create a new group.
func (r *GroupService) Insert(appId string, group *Group) *GroupInsertCall {
	c := &GroupInsertCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.appId = appId
	c.group = group
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GroupInsertCall) Fields(s ...googleapi.Field) *GroupInsertCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *GroupInsertCall) Context(ctx context.Context) *GroupInsertCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *GroupInsertCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *GroupInsertCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.group)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "apps/{appId}/groups")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"appId": c.appId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "update.group.insert" call.
// Exactly one of *Group or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Group.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *GroupInsertCall) Do(opts ...googleapi.CallOption) (*Group, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Group{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Create a new group.",
	//   "httpMethod": "POST",
	//   "id": "update.group.insert",
	//   "parameterOrder": [
	//     "appId"
	//   ],
	//   "parameters": {
	//     "appId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "apps/{appId}/groups",
	//   "request": {
	//     "$ref": "Group",
	//     "parameterName": "resource"
	//   },
	//   "response": {
	//     "$ref": "Group"
	//   }
	// }

}

// method id "update.group.list":

type GroupListCall struct {
	s            *Service
	appId        string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: List all of the groups.
func (r *GroupService) List(appId string) *GroupListCall {
	c := &GroupListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.appId = appId
	return c
}

// Limit sets the optional parameter "limit":
func (c *GroupListCall) Limit(limit int64) *GroupListCall {
	c.urlParams_.Set("limit", fmt.Sprint(limit))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GroupListCall) Fields(s ...googleapi.Field) *GroupListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *GroupListCall) IfNoneMatch(entityTag string) *GroupListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *GroupListCall) Context(ctx context.Context) *GroupListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *GroupListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *GroupListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "apps/{appId}/groups")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"appId": c.appId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "update.group.list" call.
// Exactly one of *GroupList or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *GroupList.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *GroupListCall) Do(opts ...googleapi.CallOption) (*GroupList, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &GroupList{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List all of the groups.",
	//   "httpMethod": "GET",
	//   "id": "update.group.list",
	//   "parameterOrder": [
	//     "appId"
	//   ],
	//   "parameters": {
	//     "appId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "limit": {
	//       "default": "10",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     }
	//   },
	//   "path": "apps/{appId}/groups",
	//   "response": {
	//     "$ref": "GroupList"
	//   }
	// }

}

// method id "update.group.patch":

type GroupPatchCall struct {
	s          *Service
	appId      string
	id         string
	group      *Group
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Patch: Patch a group. This method supports patch semantics.
func (r *GroupService) Patch(appId string, id string, group *Group) *GroupPatchCall {
	c := &GroupPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.appId = appId
	c.id = id
	c.group = group
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GroupPatchCall) Fields(s ...googleapi.Field) *GroupPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *GroupPatchCall) Context(ctx context.Context) *GroupPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *GroupPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *GroupPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.group)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "apps/{appId}/groups/{id}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"appId": c.appId,
		"id":    c.id,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "update.group.patch" call.
// Exactly one of *Group or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Group.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *GroupPatchCall) Do(opts ...googleapi.CallOption) (*Group, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Group{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Patch a group. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "update.group.patch",
	//   "parameterOrder": [
	//     "appId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "appId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "id": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "apps/{appId}/groups/{id}",
	//   "request": {
	//     "$ref": "Group"
	//   },
	//   "response": {
	//     "$ref": "Group"
	//   }
	// }

}

// method id "update.group.update":

type GroupUpdateCall struct {
	s          *Service
	appId      string
	id         string
	group      *Group
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Update: Patch a group.
func (r *GroupService) Update(appId string, id string, group *Group) *GroupUpdateCall {
	c := &GroupUpdateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.appId = appId
	c.id = id
	c.group = group
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GroupUpdateCall) Fields(s ...googleapi.Field) *GroupUpdateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *GroupUpdateCall) Context(ctx context.Context) *GroupUpdateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *GroupUpdateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *GroupUpdateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.group)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "apps/{appId}/groups/{id}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"appId": c.appId,
		"id":    c.id,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "update.group.update" call.
// Exactly one of *Group or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Group.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *GroupUpdateCall) Do(opts ...googleapi.CallOption) (*Group, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Group{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Patch a group.",
	//   "httpMethod": "PATCH",
	//   "id": "update.group.update",
	//   "parameterOrder": [
	//     "appId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "appId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "id": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "apps/{appId}/groups/{id}",
	//   "request": {
	//     "$ref": "Group",
	//     "parameterName": "resource"
	//   },
	//   "response": {
	//     "$ref": "Group"
	//   }
	// }

}

// method id "update.group.percent.get":

type GroupPercentGetCall struct {
	s            *Service
	appId        string
	id           string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Get the update percentage for this group.
func (r *GroupPercentService) Get(appId string, id string) *GroupPercentGetCall {
	c := &GroupPercentGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.appId = appId
	c.id = id
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GroupPercentGetCall) Fields(s ...googleapi.Field) *GroupPercentGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *GroupPercentGetCall) IfNoneMatch(entityTag string) *GroupPercentGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *GroupPercentGetCall) Context(ctx context.Context) *GroupPercentGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *GroupPercentGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *GroupPercentGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "apps/{appId}/groups/{id}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"appId": c.appId,
		"id":    c.id,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "update.group.percent.get" call.
// Exactly one of *GroupPercent or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *GroupPercent.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *GroupPercentGetCall) Do(opts ...googleapi.CallOption) (*GroupPercent, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &GroupPercent{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Get the update percentage for this group.",
	//   "httpMethod": "GET",
	//   "id": "update.group.percent.get",
	//   "parameterOrder": [
	//     "appId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "appId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "id": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "apps/{appId}/groups/{id}",
	//   "response": {
	//     "$ref": "GroupPercent"
	//   }
	// }

}

// method id "update.group.percent.set":

type GroupPercentSetCall struct {
	s            *Service
	appId        string
	id           string
	grouppercent *GroupPercent
	urlParams_   gensupport.URLParams
	ctx_         context.Context
	header_      http.Header
}

// Set: Set the update percentage for this group.
func (r *GroupPercentService) Set(appId string, id string, grouppercent *GroupPercent) *GroupPercentSetCall {
	c := &GroupPercentSetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.appId = appId
	c.id = id
	c.grouppercent = grouppercent
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GroupPercentSetCall) Fields(s ...googleapi.Field) *GroupPercentSetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *GroupPercentSetCall) Context(ctx context.Context) *GroupPercentSetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *GroupPercentSetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *GroupPercentSetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.grouppercent)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "apps/{appId}/groups/{id}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"appId": c.appId,
		"id":    c.id,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "update.group.percent.set" call.
// Exactly one of *GroupPercent or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *GroupPercent.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *GroupPercentSetCall) Do(opts ...googleapi.CallOption) (*GroupPercent, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &GroupPercent{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Set the update percentage for this group.",
	//   "httpMethod": "POST",
	//   "id": "update.group.percent.set",
	//   "parameterOrder": [
	//     "appId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "appId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "id": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "apps/{appId}/groups/{id}",
	//   "request": {
	//     "$ref": "GroupPercent",
	//     "parameterName": "resource"
	//   },
	//   "response": {
	//     "$ref": "GroupPercent"
	//   }
	// }

}

// method id "update.group.requests.events.rollup":

type GroupRequestsEventsRollupCall struct {
	s            *Service
	appId        string
	groupId      string
	dateStart    int64
	dateEnd      int64
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Rollup: Rollup all client requests by event for this group.
func (r *GroupRequestsEventsService) Rollup(appId string, groupId string, dateStart int64, dateEnd int64) *GroupRequestsEventsRollupCall {
	c := &GroupRequestsEventsRollupCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.appId = appId
	c.groupId = groupId
	c.dateStart = dateStart
	c.dateEnd = dateEnd
	return c
}

// Resolution sets the optional parameter "resolution":
func (c *GroupRequestsEventsRollupCall) Resolution(resolution int64) *GroupRequestsEventsRollupCall {
	c.urlParams_.Set("resolution", fmt.Sprint(resolution))
	return c
}

// Versions sets the optional parameter "versions":
func (c *GroupRequestsEventsRollupCall) Versions(versions string) *GroupRequestsEventsRollupCall {
	c.urlParams_.Set("versions", versions)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GroupRequestsEventsRollupCall) Fields(s ...googleapi.Field) *GroupRequestsEventsRollupCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *GroupRequestsEventsRollupCall) IfNoneMatch(entityTag string) *GroupRequestsEventsRollupCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *GroupRequestsEventsRollupCall) Context(ctx context.Context) *GroupRequestsEventsRollupCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *GroupRequestsEventsRollupCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *GroupRequestsEventsRollupCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "apps/{appId}/groups/{groupId}/requests/events/{dateStart}/{dateEnd}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"appId":     c.appId,
		"groupId":   c.groupId,
		"dateStart": strconv.FormatInt(c.dateStart, 10),
		"dateEnd":   strconv.FormatInt(c.dateEnd, 10),
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "update.group.requests.events.rollup" call.
// Exactly one of *GroupRequestsRollup or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *GroupRequestsRollup.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *GroupRequestsEventsRollupCall) Do(opts ...googleapi.CallOption) (*GroupRequestsRollup, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &GroupRequestsRollup{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Rollup all client requests by event for this group.",
	//   "httpMethod": "GET",
	//   "id": "update.group.requests.events.rollup",
	//   "parameterOrder": [
	//     "appId",
	//     "groupId",
	//     "dateStart",
	//     "dateEnd"
	//   ],
	//   "parameters": {
	//     "appId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "dateEnd": {
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "dateStart": {
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "groupId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "resolution": {
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "versions": {
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "apps/{appId}/groups/{groupId}/requests/events/{dateStart}/{dateEnd}",
	//   "response": {
	//     "$ref": "GroupRequestsRollup"
	//   }
	// }

}

// method id "update.group.requests.versions.rollup":

type GroupRequestsVersionsRollupCall struct {
	s            *Service
	appId        string
	groupId      string
	dateStart    int64
	dateEnd      int64
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Rollup: Rollup all clients requests by version for this group.
func (r *GroupRequestsVersionsService) Rollup(appId string, groupId string, dateStart int64, dateEnd int64) *GroupRequestsVersionsRollupCall {
	c := &GroupRequestsVersionsRollupCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.appId = appId
	c.groupId = groupId
	c.dateStart = dateStart
	c.dateEnd = dateEnd
	return c
}

// Resolution sets the optional parameter "resolution":
func (c *GroupRequestsVersionsRollupCall) Resolution(resolution int64) *GroupRequestsVersionsRollupCall {
	c.urlParams_.Set("resolution", fmt.Sprint(resolution))
	return c
}

// Versions sets the optional parameter "versions":
func (c *GroupRequestsVersionsRollupCall) Versions(versions string) *GroupRequestsVersionsRollupCall {
	c.urlParams_.Set("versions", versions)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GroupRequestsVersionsRollupCall) Fields(s ...googleapi.Field) *GroupRequestsVersionsRollupCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *GroupRequestsVersionsRollupCall) IfNoneMatch(entityTag string) *GroupRequestsVersionsRollupCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *GroupRequestsVersionsRollupCall) Context(ctx context.Context) *GroupRequestsVersionsRollupCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *GroupRequestsVersionsRollupCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *GroupRequestsVersionsRollupCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "apps/{appId}/groups/{groupId}/requests/versions/{dateStart}/{dateEnd}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"appId":     c.appId,
		"groupId":   c.groupId,
		"dateStart": strconv.FormatInt(c.dateStart, 10),
		"dateEnd":   strconv.FormatInt(c.dateEnd, 10),
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "update.group.requests.versions.rollup" call.
// Exactly one of *GroupRequestsRollup or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *GroupRequestsRollup.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *GroupRequestsVersionsRollupCall) Do(opts ...googleapi.CallOption) (*GroupRequestsRollup, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &GroupRequestsRollup{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Rollup all clients requests by version for this group.",
	//   "httpMethod": "GET",
	//   "id": "update.group.requests.versions.rollup",
	//   "parameterOrder": [
	//     "appId",
	//     "groupId",
	//     "dateStart",
	//     "dateEnd"
	//   ],
	//   "parameters": {
	//     "appId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "dateEnd": {
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "dateStart": {
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "groupId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "resolution": {
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "versions": {
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "apps/{appId}/groups/{groupId}/requests/versions/{dateStart}/{dateEnd}",
	//   "response": {
	//     "$ref": "GroupRequestsRollup"
	//   }
	// }

}

// method id "update.group.rollout.get":

type GroupRolloutGetCall struct {
	s            *Service
	appId        string
	groupId      string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Get the current rollout strategy for this group.
func (r *GroupRolloutService) Get(appId string, groupId string) *GroupRolloutGetCall {
	c := &GroupRolloutGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.appId = appId
	c.groupId = groupId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GroupRolloutGetCall) Fields(s ...googleapi.Field) *GroupRolloutGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *GroupRolloutGetCall) IfNoneMatch(entityTag string) *GroupRolloutGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *GroupRolloutGetCall) Context(ctx context.Context) *GroupRolloutGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *GroupRolloutGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *GroupRolloutGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "apps/{appId}/groups/{groupId}/rollout")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"appId":   c.appId,
		"groupId": c.groupId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "update.group.rollout.get" call.
// Exactly one of *Rollout or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Rollout.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *GroupRolloutGetCall) Do(opts ...googleapi.CallOption) (*Rollout, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Rollout{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Get the current rollout strategy for this group.",
	//   "httpMethod": "GET",
	//   "id": "update.group.rollout.get",
	//   "parameterOrder": [
	//     "appId",
	//     "groupId"
	//   ],
	//   "parameters": {
	//     "appId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "groupId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "apps/{appId}/groups/{groupId}/rollout",
	//   "response": {
	//     "$ref": "Rollout"
	//   }
	// }

}

// method id "update.group.rollout.set":

type GroupRolloutSetCall struct {
	s          *Service
	appId      string
	groupId    string
	rollout    *Rollout
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Set: Set (and start) a new rollout strategy for this group.
func (r *GroupRolloutService) Set(appId string, groupId string, rollout *Rollout) *GroupRolloutSetCall {
	c := &GroupRolloutSetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.appId = appId
	c.groupId = groupId
	c.rollout = rollout
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GroupRolloutSetCall) Fields(s ...googleapi.Field) *GroupRolloutSetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *GroupRolloutSetCall) Context(ctx context.Context) *GroupRolloutSetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *GroupRolloutSetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *GroupRolloutSetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.rollout)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "apps/{appId}/groups/{groupId}/rollout")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"appId":   c.appId,
		"groupId": c.groupId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "update.group.rollout.set" call.
// Exactly one of *Rollout or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Rollout.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *GroupRolloutSetCall) Do(opts ...googleapi.CallOption) (*Rollout, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Rollout{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Set (and start) a new rollout strategy for this group.",
	//   "httpMethod": "POST",
	//   "id": "update.group.rollout.set",
	//   "parameterOrder": [
	//     "appId",
	//     "groupId"
	//   ],
	//   "parameters": {
	//     "appId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "groupId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "apps/{appId}/groups/{groupId}/rollout",
	//   "request": {
	//     "$ref": "Rollout",
	//     "parameterName": "resource"
	//   },
	//   "response": {
	//     "$ref": "Rollout"
	//   }
	// }

}

// method id "update.group.rollout.active.get":

type GroupRolloutActiveGetCall struct {
	s            *Service
	appId        string
	groupId      string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Get the active status of the rollout for this group.
func (r *GroupRolloutActiveService) Get(appId string, groupId string) *GroupRolloutActiveGetCall {
	c := &GroupRolloutActiveGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.appId = appId
	c.groupId = groupId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GroupRolloutActiveGetCall) Fields(s ...googleapi.Field) *GroupRolloutActiveGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *GroupRolloutActiveGetCall) IfNoneMatch(entityTag string) *GroupRolloutActiveGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *GroupRolloutActiveGetCall) Context(ctx context.Context) *GroupRolloutActiveGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *GroupRolloutActiveGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *GroupRolloutActiveGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "apps/{appId}/groups/{groupId}/rollout/active")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"appId":   c.appId,
		"groupId": c.groupId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "update.group.rollout.active.get" call.
// Exactly one of *RolloutActive or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *RolloutActive.ServerResponse.Header or (if a response was returned
// at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *GroupRolloutActiveGetCall) Do(opts ...googleapi.CallOption) (*RolloutActive, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &RolloutActive{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Get the active status of the rollout for this group.",
	//   "httpMethod": "GET",
	//   "id": "update.group.rollout.active.get",
	//   "parameterOrder": [
	//     "appId",
	//     "groupId"
	//   ],
	//   "parameters": {
	//     "appId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "groupId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "apps/{appId}/groups/{groupId}/rollout/active",
	//   "response": {
	//     "$ref": "RolloutActive"
	//   }
	// }

}

// method id "update.group.rollout.active.set":

type GroupRolloutActiveSetCall struct {
	s             *Service
	appId         string
	groupId       string
	rolloutactive *RolloutActive
	urlParams_    gensupport.URLParams
	ctx_          context.Context
	header_       http.Header
}

// Set: Set the active status of the rollout for this group.
func (r *GroupRolloutActiveService) Set(appId string, groupId string, rolloutactive *RolloutActive) *GroupRolloutActiveSetCall {
	c := &GroupRolloutActiveSetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.appId = appId
	c.groupId = groupId
	c.rolloutactive = rolloutactive
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GroupRolloutActiveSetCall) Fields(s ...googleapi.Field) *GroupRolloutActiveSetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *GroupRolloutActiveSetCall) Context(ctx context.Context) *GroupRolloutActiveSetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *GroupRolloutActiveSetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *GroupRolloutActiveSetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.rolloutactive)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "apps/{appId}/groups/{groupId}/rollout/active")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"appId":   c.appId,
		"groupId": c.groupId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "update.group.rollout.active.set" call.
// Exactly one of *RolloutActive or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *RolloutActive.ServerResponse.Header or (if a response was returned
// at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *GroupRolloutActiveSetCall) Do(opts ...googleapi.CallOption) (*RolloutActive, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &RolloutActive{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Set the active status of the rollout for this group.",
	//   "httpMethod": "POST",
	//   "id": "update.group.rollout.active.set",
	//   "parameterOrder": [
	//     "appId",
	//     "groupId"
	//   ],
	//   "parameters": {
	//     "appId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "groupId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "apps/{appId}/groups/{groupId}/rollout/active",
	//   "request": {
	//     "$ref": "RolloutActive",
	//     "parameterName": "resource"
	//   },
	//   "response": {
	//     "$ref": "RolloutActive"
	//   }
	// }

}

// method id "update.upstream.delete":

type UpstreamDeleteCall struct {
	s          *Service
	id         string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Delete an upstream.
func (r *UpstreamService) Delete(id string) *UpstreamDeleteCall {
	c := &UpstreamDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.id = id
	return c
}

// Label sets the optional parameter "label":
func (c *UpstreamDeleteCall) Label(label string) *UpstreamDeleteCall {
	c.urlParams_.Set("label", label)
	return c
}

// Url sets the optional parameter "url":
func (c *UpstreamDeleteCall) Url(url string) *UpstreamDeleteCall {
	c.urlParams_.Set("url", url)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UpstreamDeleteCall) Fields(s ...googleapi.Field) *UpstreamDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *UpstreamDeleteCall) Context(ctx context.Context) *UpstreamDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *UpstreamDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *UpstreamDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "upstream/{id}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "update.upstream.delete" call.
// Exactly one of *Upstream or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Upstream.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *UpstreamDeleteCall) Do(opts ...googleapi.CallOption) (*Upstream, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Upstream{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Delete an upstream.",
	//   "httpMethod": "DELETE",
	//   "id": "update.upstream.delete",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "label": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "url": {
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "upstream/{id}",
	//   "response": {
	//     "$ref": "Upstream"
	//   }
	// }

}

// method id "update.upstream.insert":

type UpstreamInsertCall struct {
	s          *Service
	upstream   *Upstream
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Insert: Insert an upstream.
func (r *UpstreamService) Insert(upstream *Upstream) *UpstreamInsertCall {
	c := &UpstreamInsertCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.upstream = upstream
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UpstreamInsertCall) Fields(s ...googleapi.Field) *UpstreamInsertCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *UpstreamInsertCall) Context(ctx context.Context) *UpstreamInsertCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *UpstreamInsertCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *UpstreamInsertCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.upstream)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "upstream")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "update.upstream.insert" call.
// Exactly one of *Upstream or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Upstream.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *UpstreamInsertCall) Do(opts ...googleapi.CallOption) (*Upstream, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Upstream{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Insert an upstream.",
	//   "httpMethod": "POST",
	//   "id": "update.upstream.insert",
	//   "path": "upstream",
	//   "request": {
	//     "$ref": "Upstream",
	//     "parameterName": "resource"
	//   },
	//   "response": {
	//     "$ref": "Upstream"
	//   }
	// }

}

// method id "update.upstream.list":

type UpstreamListCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: List all upstreams.
func (r *UpstreamService) List() *UpstreamListCall {
	c := &UpstreamListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UpstreamListCall) Fields(s ...googleapi.Field) *UpstreamListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *UpstreamListCall) IfNoneMatch(entityTag string) *UpstreamListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *UpstreamListCall) Context(ctx context.Context) *UpstreamListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *UpstreamListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *UpstreamListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "upstream")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "update.upstream.list" call.
// Exactly one of *UpstreamListResp or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *UpstreamListResp.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *UpstreamListCall) Do(opts ...googleapi.CallOption) (*UpstreamListResp, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &UpstreamListResp{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List all upstreams.",
	//   "httpMethod": "GET",
	//   "id": "update.upstream.list",
	//   "path": "upstream",
	//   "response": {
	//     "$ref": "UpstreamListResp"
	//   }
	// }

}

// method id "update.upstream.sync":

type UpstreamSyncCall struct {
	s          *Service
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Sync: Synchronize all upstreams.
func (r *UpstreamService) Sync() *UpstreamSyncCall {
	c := &UpstreamSyncCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UpstreamSyncCall) Fields(s ...googleapi.Field) *UpstreamSyncCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *UpstreamSyncCall) Context(ctx context.Context) *UpstreamSyncCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *UpstreamSyncCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *UpstreamSyncCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "upstream/sync")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "update.upstream.sync" call.
// Exactly one of *UpstreamSyncResp or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *UpstreamSyncResp.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *UpstreamSyncCall) Do(opts ...googleapi.CallOption) (*UpstreamSyncResp, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &UpstreamSyncResp{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Synchronize all upstreams.",
	//   "httpMethod": "POST",
	//   "id": "update.upstream.sync",
	//   "path": "upstream/sync",
	//   "response": {
	//     "$ref": "UpstreamSyncResp"
	//   }
	// }

}

// method id "update.upstream.update":

type UpstreamUpdateCall struct {
	s          *Service
	id         string
	upstream   *Upstream
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Update: Update an upstream.
func (r *UpstreamService) Update(id string, upstream *Upstream) *UpstreamUpdateCall {
	c := &UpstreamUpdateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.id = id
	c.upstream = upstream
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UpstreamUpdateCall) Fields(s ...googleapi.Field) *UpstreamUpdateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *UpstreamUpdateCall) Context(ctx context.Context) *UpstreamUpdateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *UpstreamUpdateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *UpstreamUpdateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.upstream)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "upstream/{id}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "update.upstream.update" call.
// Exactly one of *Upstream or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Upstream.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *UpstreamUpdateCall) Do(opts ...googleapi.CallOption) (*Upstream, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Upstream{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Update an upstream.",
	//   "httpMethod": "PUT",
	//   "id": "update.upstream.update",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "upstream/{id}",
	//   "request": {
	//     "$ref": "Upstream",
	//     "parameterName": "resource"
	//   },
	//   "response": {
	//     "$ref": "Upstream"
	//   }
	// }

}

// method id "update.util.uuid":

type UtilUuidCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Uuid: Generate a new UUID.
func (r *UtilService) Uuid() *UtilUuidCall {
	c := &UtilUuidCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UtilUuidCall) Fields(s ...googleapi.Field) *UtilUuidCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *UtilUuidCall) IfNoneMatch(entityTag string) *UtilUuidCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *UtilUuidCall) Context(ctx context.Context) *UtilUuidCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *UtilUuidCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *UtilUuidCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "util/uuid")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "update.util.uuid" call.
// Exactly one of *GenerateUuidResp or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *GenerateUuidResp.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *UtilUuidCall) Do(opts ...googleapi.CallOption) (*GenerateUuidResp, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &GenerateUuidResp{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Generate a new UUID.",
	//   "httpMethod": "GET",
	//   "id": "update.util.uuid",
	//   "path": "util/uuid",
	//   "response": {
	//     "$ref": "GenerateUuidResp"
	//   }
	// }

}

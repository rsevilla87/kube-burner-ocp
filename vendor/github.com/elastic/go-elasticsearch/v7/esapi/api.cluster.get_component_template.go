// Licensed to Elasticsearch B.V under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.
//
// Code generated from specification version 7.8.0: DO NOT EDIT

package esapi

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func newClusterGetComponentTemplateFunc(t Transport) ClusterGetComponentTemplate {
	return func(o ...func(*ClusterGetComponentTemplateRequest)) (*Response, error) {
		var r = ClusterGetComponentTemplateRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// ClusterGetComponentTemplate returns one or more component templates
//
// This API is experimental.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/indices-component-template.html.
//
type ClusterGetComponentTemplate func(o ...func(*ClusterGetComponentTemplateRequest)) (*Response, error)

// ClusterGetComponentTemplateRequest configures the Cluster Get Component Template API request.
//
type ClusterGetComponentTemplateRequest struct {
	Name []string

	Local         *bool
	MasterTimeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r ClusterGetComponentTemplateRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "GET"

	path.Grow(1 + len("_component_template") + 1 + len(strings.Join(r.Name, ",")))
	path.WriteString("/")
	path.WriteString("_component_template")
	if len(r.Name) > 0 {
		path.WriteString("/")
		path.WriteString(strings.Join(r.Name, ","))
	}

	params = make(map[string]string)

	if r.Local != nil {
		params["local"] = strconv.FormatBool(*r.Local)
	}

	if r.MasterTimeout != 0 {
		params["master_timeout"] = formatDuration(r.MasterTimeout)
	}

	if r.Pretty {
		params["pretty"] = "true"
	}

	if r.Human {
		params["human"] = "true"
	}

	if r.ErrorTrace {
		params["error_trace"] = "true"
	}

	if len(r.FilterPath) > 0 {
		params["filter_path"] = strings.Join(r.FilterPath, ",")
	}

	req, err := newRequest(method, path.String(), nil)
	if err != nil {
		return nil, err
	}

	if len(params) > 0 {
		q := req.URL.Query()
		for k, v := range params {
			q.Set(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	if len(r.Header) > 0 {
		if len(req.Header) == 0 {
			req.Header = r.Header
		} else {
			for k, vv := range r.Header {
				for _, v := range vv {
					req.Header.Add(k, v)
				}
			}
		}
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	res, err := transport.Perform(req)
	if err != nil {
		return nil, err
	}

	response := Response{
		StatusCode: res.StatusCode,
		Body:       res.Body,
		Header:     res.Header,
	}

	return &response, nil
}

// WithContext sets the request context.
//
func (f ClusterGetComponentTemplate) WithContext(v context.Context) func(*ClusterGetComponentTemplateRequest) {
	return func(r *ClusterGetComponentTemplateRequest) {
		r.ctx = v
	}
}

// WithName - the comma separated names of the component templates.
//
func (f ClusterGetComponentTemplate) WithName(v ...string) func(*ClusterGetComponentTemplateRequest) {
	return func(r *ClusterGetComponentTemplateRequest) {
		r.Name = v
	}
}

// WithLocal - return local information, do not retrieve the state from master node (default: false).
//
func (f ClusterGetComponentTemplate) WithLocal(v bool) func(*ClusterGetComponentTemplateRequest) {
	return func(r *ClusterGetComponentTemplateRequest) {
		r.Local = &v
	}
}

// WithMasterTimeout - explicit operation timeout for connection to master node.
//
func (f ClusterGetComponentTemplate) WithMasterTimeout(v time.Duration) func(*ClusterGetComponentTemplateRequest) {
	return func(r *ClusterGetComponentTemplateRequest) {
		r.MasterTimeout = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f ClusterGetComponentTemplate) WithPretty() func(*ClusterGetComponentTemplateRequest) {
	return func(r *ClusterGetComponentTemplateRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f ClusterGetComponentTemplate) WithHuman() func(*ClusterGetComponentTemplateRequest) {
	return func(r *ClusterGetComponentTemplateRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f ClusterGetComponentTemplate) WithErrorTrace() func(*ClusterGetComponentTemplateRequest) {
	return func(r *ClusterGetComponentTemplateRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f ClusterGetComponentTemplate) WithFilterPath(v ...string) func(*ClusterGetComponentTemplateRequest) {
	return func(r *ClusterGetComponentTemplateRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
//
func (f ClusterGetComponentTemplate) WithHeader(h map[string]string) func(*ClusterGetComponentTemplateRequest) {
	return func(r *ClusterGetComponentTemplateRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
//
func (f ClusterGetComponentTemplate) WithOpaqueID(s string) func(*ClusterGetComponentTemplateRequest) {
	return func(r *ClusterGetComponentTemplateRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
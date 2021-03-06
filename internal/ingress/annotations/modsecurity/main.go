/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package modsecurity

import (
	extensions "k8s.io/api/extensions/v1beta1"
	"k8s.io/ingress-nginx/internal/ingress/annotations/parser"
	"k8s.io/ingress-nginx/internal/ingress/resolver"
)

// Config contains ModSecurity Configuration items
type Config struct {
	Enable        bool   `json:"enable-modsecurity"`
	OWASPRules    bool   `json:"enable-owasp-core-rules"`
	TransactionID string `json:"modsecurity-transaction-id"`
	Snippet       string `json:"modsecurity-snippet"`
}

// Equal tests for equality between two Config types
func (modsec1 *Config) Equal(modsec2 *Config) bool {
	if modsec1 == modsec2 {
		return true
	}
	if modsec1 == nil || modsec2 == nil {
		return false
	}
	if modsec1.Enable != modsec2.Enable {
		return false
	}
	if modsec1.OWASPRules != modsec2.OWASPRules {
		return false
	}
	if modsec1.TransactionID != modsec2.TransactionID {
		return false
	}
	if modsec1.Snippet != modsec2.Snippet {
		return false
	}

	return true
}

// NewParser creates a new ModSecurity annotation parser
func NewParser(resolver resolver.Resolver) parser.IngressAnnotation {
	return modSecurity{resolver}
}

type modSecurity struct {
	r resolver.Resolver
}

// Parse parses the annotations contained in the ingress
// rule used to enable ModSecurity in a particular location
func (a modSecurity) Parse(ing *extensions.Ingress) (interface{}, error) {

	enableModSecurity, err := parser.GetBoolAnnotation("enable-modsecurity", ing)
	if err != nil {
		enableModSecurity = false
	}

	owaspRules, err := parser.GetBoolAnnotation("enable-owasp-core-rules", ing)
	if err != nil {
		owaspRules = false
	}

	transactionID, err := parser.GetStringAnnotation("modsecurity-transaction-id", ing)
	if err != nil {
		transactionID = ""
	}

	snippet, err := parser.GetStringAnnotation("modsecurity-snippet", ing)
	if err != nil {
		snippet = ""
	}

	return Config{
		Enable:        enableModSecurity,
		OWASPRules:    owaspRules,
		TransactionID: transactionID,
		Snippet:       snippet,
	}, nil
}

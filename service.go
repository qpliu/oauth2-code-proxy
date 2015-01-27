package app

import (
	"net/url"
)

type Service interface {
	ProxyPath() string
	URL() string
	BasicAuth() (username, password string, ok bool)
	AddSecret(parameters url.Values)
}

type basicAuthService struct {
	proxyPath, url, clientId, clientSecret string
}

func (s basicAuthService) ProxyPath() string {
	return s.proxyPath
}

func (s basicAuthService) URL() string {
	return s.url
}

func (s basicAuthService) BasicAuth() (string, string, bool) {
	return s.clientId, s.clientSecret, true
}

func (s basicAuthService) AddSecret(parameters url.Values) {
}

type addParameterService struct {
	proxyPath, url, parameterName, clientSecret string
}

func (s addParameterService) ProxyPath() string {
	return s.proxyPath
}

func (s addParameterService) URL() string {
	return s.url
}

func (s addParameterService) BasicAuth() (string, string, bool) {
	return "", "", false
}

func (s addParameterService) AddSecret(parameters url.Values) {
	parameters.Add(s.parameterName, s.clientSecret)
}

func BasicAuthService(proxyPath, url, clientId, clientSecret string) Service {
	return basicAuthService{proxyPath, url, clientId, clientSecret}
}

func AddParameterService(proxyPath, url, parameterName, clientSecret string) Service {
	return addParameterService{proxyPath, url, parameterName, clientSecret}
}

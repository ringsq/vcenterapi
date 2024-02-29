package vcenter

import (
	"crypto/tls"
	"fmt"
	"log/slog"

	"github.com/go-resty/resty/v2"
	"github.com/rsapc/hookcmd/models"
)

const (
	loginPath = "/rest/com/vmware/cis/session"
)

type Vcenter struct {
	baseURL string
	token   string
	client  *resty.Client
	log     models.Logger
}

func NewClient(baseURL string, username string, password string, logger models.Logger) (*Vcenter, error) {
	vc := &Vcenter{baseURL: baseURL}
	vc.log = logger
	if log, ok := logger.(*slog.Logger); ok {
		vc.log = log.With("service", "vcenter")
	}

	vc.client = resty.New()
	vc.client.SetRedirectPolicy(resty.FlexibleRedirectPolicy(5))
	vc.client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

	token, err := vc.login(username, password)
	if err != nil {
		vc.log.Error("failed to login to vcenter", "error", err)
		return vc, err
	}
	vc.token = token
	return vc, nil
}

func (vc *Vcenter) buildRequest() *resty.Request {
	return vc.client.NewRequest().SetHeader("vmware-api-session-id", vc.token)
}

func (vc *Vcenter) buildURL(path string, args ...any) string {
	urlPath := fmt.Sprintf(path, args...)
	return fmt.Sprintf("%s/api/vcenter%s", vc.baseURL, urlPath)
}

func (vc *Vcenter) get(resultObj any, path string, args ...any) error {
	url := vc.buildURL(path, args...)
	req := vc.buildRequest().SetResult(resultObj)
	resp, err := req.Get(url)
	if err != nil {
		vc.log.Error("error communicating with vcenter", "method", "GET", "url", url, "error", err)
		return err
	}
	if resp.IsError() {
		vc.log.Error("vcenter returned an error response", "method", "GET", "url", url, "status", resp.StatusCode())
		return fmt.Errorf("response error: %v", resp.Error())
	}
	return nil
}

func (vc *Vcenter) login(username string, password string) (token string, err error) {
	login := &LoginResult{}
	req := vc.client.NewRequest().SetBasicAuth(username, password).SetResult(login)
	resp, err := req.Post(fmt.Sprintf("%s%s", vc.baseURL, loginPath))
	if err != nil {
		return token, err
	}
	if resp.IsError() {
		return token, fmt.Errorf("login error: %v", resp.Error())
	}
	token = login.Value
	return token, err
}

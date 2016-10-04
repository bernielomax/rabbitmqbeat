package rabbitmq

import (
	"encoding/json"
	"fmt"
	"github.com/elastic/beats/libbeat/common"
	"io/ioutil"
	"net/http"
)

type Client struct {
	baseUrl string
	http    *http.Client
	auth    *auth
}

type auth struct {
	username string
	password string
}

type Api interface {
	Nodes() ([]common.MapStr, error)
	Overview() (common.MapStr, error)
	Queues() ([]common.MapStr, error)
}

func NewClient(url, username, password string) Api {
	c := Client{
		baseUrl: url,
		http:    http.DefaultClient,
		auth: &auth{
			username: username,
			password: password,
		},
	}
	return Api(&c)
}

func (c *Client) Nodes() ([]common.MapStr, error) {
	n, err := c.getMany("nodes")
	return n, err
}

func (c *Client) Queues() ([]common.MapStr, error) {
	q, err := c.getMany("queues")
	return q, err
}

func (c *Client) Overview() (common.MapStr, error) {
	o, err := c.getOne("overview")
	return o, err
}

func (c *Client) getOne(call string) (common.MapStr, error) {
	var q common.MapStr
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/%v", c.baseUrl, call), nil)
	if err != nil {
		return q, err
	}
	req.SetBasicAuth(c.auth.username, c.auth.password)
	resp, err := c.http.Do(req)
	if err != nil {
		return q, err
	}
	if resp.StatusCode != 200 {
		return q, fmt.Errorf(resp.Status)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return q, err
	}
	err = json.Unmarshal(body, &q)
	return q, err
}

func (c *Client) getMany(call string) ([]common.MapStr, error) {
	var q []common.MapStr
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/%v", c.baseUrl, call), nil)
	if err != nil {
		return q, err
	}
	req.SetBasicAuth(c.auth.username, c.auth.password)
	resp, err := c.http.Do(req)
	if err != nil {
		return q, err
	}
	if resp.StatusCode != 200 {
		return q, fmt.Errorf(resp.Status)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return q, err
	}
	err = json.Unmarshal(body, &q)
	return q, err
}
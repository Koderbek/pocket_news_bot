package rkn

import (
	"bytes"
	"github.com/Koderbek/pocket_news_bot/pkg/config"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

func TestRoskomsvobodaClient_List(t *testing.T) {
	testCases := []struct {
		name       string
		cfg        config.Rkn
		result     []string
		shouldFail bool
		roundTrip  func(req *http.Request) *http.Response
	}{
		{
			name:       "case-1: valid result",
			cfg:        config.Rkn{Url: "https://123test.ru", DefaultTimeout: 10},
			result:     []string{"0-00.lordfilm0.biz", "0-10.lordfilm0.biz", "0-100.lordfilm0.biz", "0-101.lordfilm0.biz"},
			shouldFail: false,
			roundTrip: func(req *http.Request) *http.Response {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewBufferString(`["0-00.lordfilm0.biz", "0-10.lordfilm0.biz", "0-100.lordfilm0.biz", "0-101.lordfilm0.biz"]`)),
					Header:     make(http.Header),
				}
			},
		},
		{
			name:       "case-2: error status code",
			cfg:        config.Rkn{Url: "https://123test.ru", DefaultTimeout: 10},
			shouldFail: true,
			roundTrip: func(req *http.Request) *http.Response {
				return &http.Response{
					StatusCode: http.StatusBadRequest,
					Header:     make(http.Header),
				}
			},
		},
		{
			name:       "case-3: unmarshal JSON",
			cfg:        config.Rkn{Url: "https://123test.ru", DefaultTimeout: 10},
			shouldFail: true,
			roundTrip: func(req *http.Request) *http.Response {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewBufferString(`, "0-10.lordfilm0.biz", "0-100.lordfilm0.biz", "0-101.lordfilm0.biz`)),
					Header:     make(http.Header),
				}
			},
		},
		{
			name:       "case-4: http error",
			cfg:        config.Rkn{Url: "~sppg:d\\|||\\d./_@ss..:gttps//:", DefaultTimeout: 10},
			shouldFail: true,
			roundTrip: func(req *http.Request) *http.Response {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewBufferString(`, "0-10.lordfilm0.biz", "0-100.lordfilm0.biz", "0-101.lordfilm0.biz`)),
					Header:     make(http.Header),
				}
			},
		},
	}

	for _, tc := range testCases {
		client := NewTestClient(tc.roundTrip)
		roskomsvobodaClient := &RoskomsvobodaClient{client, tc.cfg}
		domains, err := roskomsvobodaClient.List()
		if tc.shouldFail {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, domains, tc.result)
		}

	}
}

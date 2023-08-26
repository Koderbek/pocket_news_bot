package rkn

import netUrl "net/url"

func parseHost(srcUrl string) (string, error) {
	u, err := netUrl.Parse(srcUrl)
	if err != nil {
		return "", err
	}

	return u.Hostname(), nil
}

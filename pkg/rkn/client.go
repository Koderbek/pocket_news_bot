package rkn

type Client interface {
	IsForbidden(url string) (bool, error)
}

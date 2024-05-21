package rkn

type Client interface {
	List() ([]string, error)
}

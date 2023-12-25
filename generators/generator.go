package generators

type URLGenerator interface {
	GetRandURL() string
	GetNRandURL(n int) string
}

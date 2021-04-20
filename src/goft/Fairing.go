package goft

type Fairing interface {
	OnRequest() error
}

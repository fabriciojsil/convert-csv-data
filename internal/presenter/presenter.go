package presenter

type Presenter interface {
	Parser()
	Present(string) error
}

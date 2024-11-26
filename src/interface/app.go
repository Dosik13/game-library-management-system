package _interface

type Apper interface {
	NewApp() (Apper, error)
	Start()
}

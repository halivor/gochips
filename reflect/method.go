package main

type method struct {
	id   int32
	name string
}

func New() *method {
	return &method{}
}

func (m *method) Get() string {
	return m.name
}

func (m *method) Set(name string) string {
	m.name = name
	return m.name
}

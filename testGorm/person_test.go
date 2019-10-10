package main

import "testing"

func Test_AddPerson(T *testing.T) {
	p := Person{}
	p.Name = "1"
	p.Add()
}

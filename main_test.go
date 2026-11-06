package main
import "testing"
func TestValidation(t *testing.T){if validate([]Job{{DelaySeconds:0,Command:[]string{"/usr/bin/true"}}})!=nil{t.Fatal("valid rejected")};for _,j:=range [][]Job{nil,{{DelaySeconds:-1,Command:[]string{"x"}}},{{DelaySeconds:1,Command:nil}}}{if validate(j)==nil{t.Fatal("invalid accepted")}}}

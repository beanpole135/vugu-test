package main

import (
	"github.com/vugu/vugu"
)

type Test struct {
	A string	`json:"a"`
	B []string	`json:"b"`
	C bool		`json:"c"`
	D int		`json:"d"`
	E map[string]string	`json:"e"`
	I int		`json:"i"`
	Color string	`json:"color"`
}

func (c *Test) Toggle(event *vugu.DOMEvent) {
	c.C = !c.C
	if c.A == "" {
	  //Initialize the structure
	  c.I = 0
	  c.A = "A"
	  c.B = append([]string{}, "str1","str2","str3")
	  c.C = true
	  c.D = 34
	  c.E = make(map[string]string)
	    c.E["test1"] = "result1"
	    c.E["test2"] = "result2" 
	  c.Color = "#993322"
	}else{
	  c.D = c.D + 1
	  c.ChangeIndex()
	}
	if c.C { 
		c.Color = "#993322"
	}else{
		c.Color = "#223399"
	}

}

func (c *Test) ChangeIndex() {
  index := c.I +1
  if (index >= len(c.B) ){ index = 0 }
  c.I = index
}

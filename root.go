package main

type Root struct {
    ShowWasm bool `vugu:"data"`
    ShowGo bool   `vugu:"data"`
    ShowVugu bool `vugu:"data"`
	Show bool `vugu:"data"`
	Json Test `vugu:"data"`
}

var ROOT Root

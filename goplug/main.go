package main

// Testing Golang 1.8 features for Plugins loading

import (
	"plugin"
	"log"
)

func main() {
	// Try loading the shared library file
	// Gives back a pointer to a Plugin.
	plugin, err := plugin.Open("./true_plugin.so")
	if err != nil {
		log.Fatalf("Loading the plugin library failed: %s", err)
	}

	// Next with the name of the Symbol we want to use the plugin to look it up
	// A Symbol can be “any exported variable or function”.
	tf, err := plugin.Lookup("TrueOrNot")
	if err != nil {
		log.Fatalf("Can not find the requested symbol %s", err)
	}

	d := tf.(func(int, int) bool)(1, 2)
	log.Printf("Two different values give %t", d)

	s := tf.(func(int, int) bool)(2, 2)
	log.Printf("Two same values give %t", s)
}


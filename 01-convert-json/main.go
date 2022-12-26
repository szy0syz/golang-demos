package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Creature struct {
	Name string     `json:"Name"`
	Tags []string   `json:"Tags"`
	HP   CreatureHP `json:"HP"`
}

type CreatureHP struct {
	Value int    `json:"Value"`
	Notes string `json:"Notes"`
}

func readDirectory() {
	files, err := ioutil.ReadDir("./01-convert-json/data")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Printf("IsFile: %s\nIsFolder %t\n", file.Name(), file.IsDir())
	}
}

func readNode() {
	content, err := ioutil.ReadFile("./01-convert-json/data/1.json")
	if err != nil {
		fmt.Println("Error reading file")
	}

	var creature Creature
	// golang的"显示哲学" + 指针传递
	json.Unmarshal([]byte(content), &creature)

	fmt.Println()
	fmt.Println("Creature name: " + creature.Name)
	fmt.Printf("Creature Tags: %q\n", creature.Tags)
	fmt.Printf("Creature value: %+v Notes: %v\n", creature.HP.Value, creature.HP.Notes)
}

func main() {
	readDirectory()
	readNode()
}

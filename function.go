package main

import (
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/eiannone/keyboard"
)

func PrintAllFiles() {

	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	cmd := exec.Command("ls", "-R")
	outByte, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("string(outByte): %v\n", string(outByte))

	element := strings.SplitAfter(string(outByte), "\n\n")

	var paths []string
	var files []string
	var f []string

	for _, v := range element {

		b, a, _ := strings.Cut(v, ":")
		files = append(files, a)
		paths = append(paths, b)

	}

	for _, v := range files {

		f = append(f, strings.Split(v, "\n")...)

	}

	count := 0
	for _, v := range f {
		if v == "" {
			count++
		}
	}

	filePath := make(map[string][]string)

	///////////////// INSERT VALUE AND KEYS INTO MAP /////////////////////////////////

	keyPath := ""
	content := []string{}
	i := 0
	// stock := ""

	for i < len(paths) {

		keyPath = paths[i]

		s := strings.Split(f[i], "\n")

		content = append(content, s...)
		filePath[keyPath] = content

		content = []string{}

		i++

	}

	_, key, err := keyboard.GetKey()
	if err != nil {
		log.Fatal(err)
	}

	if key == keyboard.KeyTab {
		for _, v := range f {

			os.Stdout.Write([]byte(v))
			os.Stdout.WriteString("\t")

		}
	}

	var fichier []string
	after := ""
	before := ""

	var chemin []string

	for _, a := range element {
		before, after, _ = strings.Cut(a, ":")
		fichier = append(fichier, after)
		chemin = append(chemin, before)
	}

	carte := make(map[string][]string)

	//fmt.Printf("fichier: %v\n", fichier)
	for i, a := range chemin {
		//fmt.Printf("e: %v\n", fichier[i])
		r := strings.Split(fichier[i], "\n")
		carte[a] = r
		//	fmt.Printf("carte[a]: %v\n", carte[a])

	}

	Display(fichier, chemin)
}

package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/eiannone/keyboard"
)

func SearchFiles(selected int, fichier []string) {

	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	for {

		_, key, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}

		if key == keyboard.KeyArrowRight {
			if selected == len(fichier)-1 {
				selected = 0
				for range fichier[len(fichier)-1] {
					fmt.Print("\b \b")

				}
			} else {
				for range fichier[selected] {
					fmt.Print("\b \b")
				}
				selected++

				os.Stdout.WriteString(fichier[selected])
				_, key, err = keyboard.GetKey()
				if err != nil {
					log.Fatal(err)
				}
			}
		}

		if key == keyboard.KeyEsc {
			return
		}

		if key == keyboard.KeyArrowLeft {
			if selected == 0 {
				selected = len(fichier) - 1
				for range fichier[0] {
					fmt.Print("\b \b")
				}

			} else {
				for range fichier[selected] {

					fmt.Print("\b \b")
				}
				selected--

				os.Stdout.WriteString(fichier[selected])
				_, key, err = keyboard.GetKey()
				if err != nil {
					log.Fatal(err)
				}

			}

		}
	}
}
func Keep(s string, want string) (string, string, string) {

	f := strings.Split(s, "\n\n")
	res := ""
	before := ""
	after := ""
	element := ""

	for _, v := range f {
		if v != " " && v != "\n" {
			res += v + " "
		}
	}
	ff := strings.Fields(res)

	for i := 0; i < len(ff); i++ {
		if i > 0 && i < len(ff) && strings.Contains(ff[i], want) {
			after = ff[i+1]
			before = ff[i-1]
			element = ff[i]
		}
	}
	return before, element, after
}

func LoopSingleFile(fileArraySelected string, selected int) {
	_, key, err := keyboard.GetKey()
	if err != nil {
		log.Fatal(err)
	}

	f := strings.Split(fileArraySelected, "\n\n")
	if selected > len(f)-1 {
		selected = len(f) - 1
	}

	namef := ""

	for _, e := range f {
		if !strings.Contains(e, " ") && !strings.Contains(e, "\n") {
			namef = e
		}
	}
	name := ""
	for _, v := range namef {
		if v != rune('\n') {
			name += string(v)
		}
	}
	for {

		_, key, err = keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}

		if key == keyboard.KeyArrowLeft {
			if selected == 0 {
				erase := "->" + name
				for range erase {
					fmt.Print("\b \b")
				}

				selected = len(f) - 1

			} else {

				erase := "->" + name

				for range erase {
					fmt.Print("\b \b")
				}
				selected--
				os.Stdout.WriteString("->" + name)

			}
		}

		if key == keyboard.KeyArrowRight {

			if selected == len(f)-1 {
				erase := "->" + name
				for range erase {
					fmt.Print("\b \b")
				}

				selected = 0

			} else {
				erase := "->" + name
				for range erase {
					fmt.Print("\b \b")
				}

				selected++

				os.Stdout.WriteString("->" + name)

			}
		}

		if key == keyboard.KeyTab {
			fmt.Printf("path: %v\n", path)

			if !strings.Contains(name, ".") && !strings.Contains(name, "\n") {
				PrintAllFiles()
				return
			}
		}

		if key == keyboard.KeyEsc {
			for range f {
				fmt.Print("\b \b")
			}

			os.Exit(1)
		}
	}

}

func Copy2Clipboard(path2File string) {

	catOut, err := exec.Command("cat", path2File).Output()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	p2f := strings.NewReader(string(catOut))
	cmd := exec.Command("xsel", "-b", "-i")
	cmd.Stdin = p2f
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	os.Stdin.WriteString("File content into Clipboard...")

}

func GetFilePath() {

	wkDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("wkDir: %v\n", wkDir)

}

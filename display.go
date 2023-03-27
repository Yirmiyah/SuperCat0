package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/eiannone/keyboard"
)

var wkDir string
var path string

func Display(fichier []string, chemin []string) {

	selected := 0
	currentFolder := "Current Folder ->"
	var fSelected string
	path, _ = os.Getwd()
	/////////////////// DISPLAY LOOP ////////////////////////////
	for {

		_, key, err := keyboard.GetKey() /* Getting Keyboard Keys */
		if err != nil {
			log.Fatal(err)
		}

		if key == keyboard.KeyArrowLeft {
			if selected == 0 {
				folder, _ := getLastName2Path(chemin[selected]) /* Deleting folder's name */
				for range folder {
					fmt.Print("\b \b")
				}

				selected = len(chemin) - 1

			} else {

				folder, isWkFolder := getLastName2Path(chemin[selected]) /* Getting the folder's name */

				if isWkFolder {

					folder = currentFolder /* if we are on the current folder display current folder's name */

					for range folder {

						fmt.Print("\b \b")

					}
				} else {

					for range folder {

						fmt.Print("\b \b")

					}
				}
			}
			selected--

			folder, _ := getLastName2Path(chemin[selected])

			os.Stdout.WriteString(folder)

			if key == keyboard.KeyTab {
				folder, _ := getLastName2Path(chemin[selected])
				// path = ""

				path, _ = os.Getwd()

				path += "/" + folder
				fSelected = displayFiles(fichier, selected, folder)
				os.Stdout.WriteString(fSelected)
			}
		}

		if key == keyboard.KeyArrowRight {

			if selected == len(chemin)-1 {
				folder, _ := getLastName2Path(chemin[selected])
				for range folder {
					fmt.Print("\b \b")
				}

				selected = 0

			} else {

				folder, isWkFolder := getLastName2Path(chemin[selected])
				if isWkFolder {
					folder = currentFolder
					for range folder {
						fmt.Print("\b \b")
					}
				} else {
					for range folder {
						fmt.Print("\b \b")
					}
				}

				selected++
				folder, _ = getLastName2Path(chemin[selected])
				os.Stdout.WriteString(folder)

			}
		}

		if key == keyboard.KeyTab {
			folder, _ := getLastName2Path(chemin[selected])
			for range folder {
				fmt.Print("\b \b")
			}

			path += "/" + chemin[selected]
			fmt.Printf("chemin[selected]: %v\n", chemin[selected])

			fSelected = displayFiles(fichier, selected, folder)

			os.Stdout.WriteString(fSelected)

		}

		if key == keyboard.KeyEsc {
			return
		}
	}
}

func getLastName2Path(path string) (string, bool) {
	ans := ""
	r := ""
	wkFolderBool := false

	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == byte('/') {
			break
		} else {
			ans += string(path[i])
		}
	}

	for i := len(ans) - 1; i >= 0; i-- {
		if ans[i] == byte(':') {
			break
		} else {
			r += string(ans[i])
		}
	}

	if r == getWorkingFolder() {
		wkFolderBool = true
	}

	return "Folder-> " + r, wkFolderBool
}

func getWorkingFolder() string {
	wkDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	ans := ""
	r := ""

	for i := len(wkDir) - 1; i >= 0; i-- {
		if wkDir[i] == byte('/') {
			break
		} else {
			ans += string(wkDir[i])
		}
	}

	for i := len(ans) - 1; i >= 0; i-- {
		if ans[i] == byte(':') {
			break
		} else {
			r += string(ans[i])
		}
	}

	return ans

}

func displayFiles(fichier []string, selected int, folderName string) string {

	_, key, err := keyboard.GetKey()
	if err != nil {
		log.Fatal(err)
	}

	wkDir, _ = os.Getwd()

	f := strings.Split(fichier[selected], "\n")

	if len(f) == 4 {
		LoopSingleFile(fichier[selected], selected)
	}

	for {

		_, key, err = keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}

		if key == keyboard.KeyArrowLeft {
			if selected == 0 {

				for range f {
					fmt.Print("\b \b")
				}

				for range f {
					fmt.Print("\b \b")
				}

				selected = len(f) - 1

			} else {

				for range f {
					fmt.Print("\b \b")
				}
				selected--
				os.Stdout.WriteString("->" + f[selected])

			}
		}

		if key == keyboard.KeyArrowRight {

			if selected == len(f)-1 {

				for range f {
					fmt.Print("\b \b")
				}

				selected = 0

			} else {

				for range f {
					fmt.Print("\b \b")
				}

				selected++

				os.Stdout.WriteString("->" + f[selected])

			}
		}

		if key == keyboard.KeyTab {

			// path = ""

			// path, _ = os.Getwd()

			path += "/" + f[selected]
			if strings.Contains(f[selected], ".") && !strings.Contains(f[selected], " ") {
				fmt.Printf("path: %v\n", path)
				Copy2Clipboard(path)
				return f[selected]
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

func GettingFilesNames(filesSplit []string) {

	_, key, err := keyboard.GetKey()
	if err != nil {
		log.Fatal(err)
	}

	selected := 1

	for {

		_, key, err = keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}

		if key == keyboard.KeyArrowLeft {
			if selected == 0 {

				for range filesSplit[selected] {
					fmt.Print("\b \b")
				}

				selected = len(filesSplit) - 1

			} else {

				for range filesSplit[selected] {

					fmt.Print("\b \b")

				}
			}
			for range filesSplit[selected] {

				fmt.Print("\b \b")

			}

			selected--

			os.Stdout.WriteString(filesSplit[selected])

		}

		if key == keyboard.KeyArrowRight {

			if selected == len(filesSplit)-1 {
				for range filesSplit[selected] {
					fmt.Print("\b \b")
				}

				selected = 0

			} else {

				for range filesSplit[selected] {
					fmt.Print("\b \b")
				}
				for range filesSplit[selected] {

					fmt.Print("\b \b")

				}

				selected++

				os.Stdout.WriteString(filesSplit[selected])

			}
		}

		if key == keyboard.KeyTab {
			for range filesSplit[selected] {

				fmt.Print("\b \b")

			}
			path = ""

			path, _ = os.Getwd()

			os.Stdout.WriteString(filesSplit[selected])

		}

		if key == keyboard.KeyEsc {
			for range filesSplit[selected] {

				fmt.Print("\b \b")

			}

			os.Exit(1)
		}
	}
}

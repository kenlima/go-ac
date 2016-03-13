package main

import (
	"bufio"
	"container/list"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("/Users/jwlee/IdeaProjects/ci-autocomplete/src/main/resources/ciautocomplete/ci_autocomplete_admin.php.template")
	check(err)
	//fmt.Println(string(dat))
	_ = dat

	log.SetFlags(log.Lshortfile)
	dir := "/Users/jwlee/project/box/www/gagamel_admin"
	ignoreDirs := []string{".git"}
	allowExts := []string{".php"}

	cs := list.New()
	cstypes := list.New()

	err2 := filepath.Walk(dir, printFile(ignoreDirs, allowExts, cs))

	for e := cs.Front(); e != nil; e = e.Next() {

		file, err := os.Open(e.Value.(string))
		if err != nil {
			log.Fatal(err)
		}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			if strings.Index(line, "class ") == 0 {
				fmt.Println(e.Value)

				a := regexp.MustCompile("\\s+")
				words := a.Split(line, -1)

				cstype := cstype{Name: "", PName: ""}

				for i := 0; i < len(words); i++ {
					w := strings.Trim(words[i], " ")
					if w == "class" {
						cstype.Name = strings.Replace(strings.Trim(words[i+1], " "), "{", "", -1)
					} else if w == "extends" {
						cstype.PName = strings.Replace(strings.Trim(words[i+1], " "), "{", "", -1)
						break
					}
				}
				cstypes.PushFront(cstype)
				break

			}
		}
		file.Close()

	}

	for e := cstypes.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	if err2 != nil {
		log.Fatal(err)
	}
}

func printFile(ignoreDirs []string, allowExts []string, cs *list.List) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Print(err)
			return nil
		}

		if info.IsDir() {
			dir := filepath.Base(path)
			for _, d := range ignoreDirs {
				if d == dir {
					//	log.Println(dir)
					return filepath.SkipDir
				}
			}
		} else {
			ext := filepath.Ext(path)
			for _, e := range allowExts {
				if e == ext {
					cs.PushFront(path)
					//	fmt.Println(path)
					return nil
				}
			}
		}

		//fmt.Println(path)

		return nil
	}
}

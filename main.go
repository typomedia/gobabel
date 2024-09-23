package main

import (
	"fmt"
	"github.com/jvatic/goja-babel"
	"github.com/spf13/pflag"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	source := pflag.StringP("source", "s", "", "Source folder")
	target := pflag.StringP("target", "t", "", "Target folder")
	preset := pflag.StringP("preset", "p", "env", "Babel preset [env, react, flow]")
	pflag.Parse()

	args := pflag.Args()
	if len(args) > 0 {
		*source = args[0]
	}

	if *source == "" {
		fmt.Println("Source folder is required!")
		pflag.Usage()
		os.Exit(1)
	}

	if *target == "" {
		*target = *source
	}

	files, err := ReadDir(*source)
	if err != nil {
		log.Fatal(err)
	}

	for path, file := range files {

		if file.IsDir() {
			continue
		}

		folder := filepath.Dir(path)
		os.MkdirAll(filepath.Join(*target, folder), os.ModePerm)

		f, err := os.Open(filepath.Join(*source, path))
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		babel.Init(4) // transformer with 4 workers

		res, err := babel.Transform(f, map[string]interface{}{
			"presets": []string{
				*preset,
			},
			"plugins": []string{
				//"transform-strict-mode",
			},
		})
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(filepath.Join(*target, path))
		f, err = os.Create(filepath.Join(*target, path))
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		_, err = io.Copy(f, res)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func ReadDir(name string) (map[string]os.DirEntry, error) {
	entries := make(map[string]os.DirEntry)

	err := filepath.WalkDir(name, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if path != name { // Skip the name directory itself
			relPath, err := filepath.Rel(name, path)
			if err != nil {
				return err
			}
			entries[relPath] = d
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return entries, nil
}

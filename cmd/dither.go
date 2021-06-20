package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	dither "github.com/alkazarix/go-dither"
)

var (
	inputDir    string
	outputDir   string
	monochrome  bool
	multiplier  float64
	filterNames string
)

const helper = `
Usage:
go run <image>
Options:
  -filters string
    	name of the filters to apply, options: 'floydSteinberg', 'burkes', 'sierra', 'sierra2', 'sierra3', 'stucki', 'atkinson' (default floydSteinberg)
  -multiplier float
    	Error multiplier (default 1.18)
  -outputdir string
			Directory name, where to save the generated images (default "output")
			`

func usage() {
	fmt.Fprint(os.Stderr, helper)
	os.Exit(1)
}

func main() {

	if err := process(); err != nil {
		log.Printf("%s", err)
		usage()
	}
	log.Println("DONE")

}

func process() error {

	command := flag.NewFlagSet("command", flag.ExitOnError)
	command.StringVar(&outputDir, "outputDir", "output", "directory where to saved generated image")
	command.BoolVar(&monochrome, "monochrome", false, "generate color or greyscale dithered images")
	command.Float64Var(&multiplier, "multiplier", 1.18, "error factor")
	command.StringVar(&filterNames, "filters", "floydSteinberg", "filters to apply (comma separated)")
	command.Usage = usage

	if len(os.Args) < 2 {
		usage()
	}

	command.Parse(os.Args[2:])

	filename := os.Args[1]
	img, err := openImage(filename)
	if err != nil {
		return err
	}

	filters, err := filters(filterNames)
	if err != nil {
		return err
	}

	monochromeDir, colorDir := filepath.Join(outputDir, "mono"), filepath.Join(outputDir, "color")
	if monochrome {
		_ = os.Mkdir(monochromeDir, os.ModePerm)
	} else {
		_ = os.Mkdir(colorDir, os.ModePerm)
	}

	var wg sync.WaitGroup
	for _, filter := range filters {
		wg.Add(1)
		go func(filter *dither.Filter) {
			defer wg.Done()
			var outImg image.Image
			var directory string
			if monochrome {
				outImg = dither.Monochrome(img, *filter, float32(multiplier))
				directory = monochromeDir
			} else {
				outImg = dither.Color(img, *filter, float32(multiplier))
				directory = colorDir
			}

			writePNG(outImg, filepath.Join(directory, filter.Name+".png"))
		}(filter)
	}
	wg.Wait()

	return nil

}

func openImage(filename string) (image.Image, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil

}

func filters(filterNames string) ([]*dither.Filter, error) {
	var filters []*dither.Filter

	for _, name := range strings.Split(filterNames, ",") {
		switch strings.TrimSpace(name) {
		case "floydSteinberg":
			filters = append(filters, dither.FloydSteinberg)
		case "burkes":
			filters = append(filters, dither.Burkes)
		case "sierra":
			filters = append(filters, dither.SierraLite)
		case "sierra2":
			filters = append(filters, dither.Sierra2)
		case "sierra3":
			filters = append(filters, dither.Sierra3)
		case "stucki":
			filters = append(filters, dither.Stucki)
		case "atkinson":
			filters = append(filters, dither.Atkinson)
		}
	}

	if filters == nil {
		return nil, fmt.Errorf("invalid filter options %s", filterNames)
	}

	return filters, nil

}

func writePNG(img image.Image, fn string) {
	f, err := os.Create(fn)
	if err != nil {
		log.Fatal(err)
	} else {
		if err := png.Encode(f, img); err != nil {
			log.Fatal(err)
			usage()
		}
	}
	defer f.Close()

}

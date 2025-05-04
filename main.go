package main

import (
	"SliceIt/fileprocessing"
	"SliceIt/view"
	"flag"
	"fmt"
	"os"
)

func main() {
	const VERSION string = "0.0.1\n\n"

	mode := flag.String("mode", "", "Split or merge file")
	folder := flag.String("folder", "", "Folder for .part files")
	file := flag.String("file", "", "Name of file to split or merge")
	part_size := flag.Float64("part_size", 10.0, "Size of part (Mb)")
	use_hash := flag.Bool("hash", false, "Use sha256 checksum")
	help := flag.Bool("help", false, "View README.md")

	flag.Parse()

	if *help {
		view.ViewReadme()
	}

	if *file == "" {
		fmt.Println("Please provide a file name using -file flag")
		os.Exit(1)
	}

	fmt.Print("SliceIt ", VERSION)
	switch *mode {
	case "split":
		fileprocessing.Split_file(*file, float32(*part_size), *folder, *use_hash)
	case "merge":
		fileprocessing.Merge_file(*file, *folder, *file)
	default:
		fmt.Printf("Unknown mode: %s. Use 'split' or 'merge'.", *mode)
		os.Exit(1)
	}

}

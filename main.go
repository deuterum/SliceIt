package main

import (
	"SliceIt/fileprocessing"
	"flag"
	"fmt"
	"os"
)

func main() {
	const VERSION string = "0.0.1"
	mode := flag.String("mode", "", "Split or merge file")
	folder := flag.String("folder", "", "Folder for .part files")
	file := flag.String("file", "", "Name of file to split or merge")
	part_size := flag.Float64("part_size", 10.0, "Size of part (Mb)")

	flag.Parse()

	if *file == "" {
		fmt.Println("Please provide a file name using -file flag")
		os.Exit(1)
	}

	fmt.Print("SliceIt ", VERSION, "\n")
	switch *mode {
	case "split":
		fileprocessing.Split_file(*file, float32(*part_size), *folder)
	case "merge":
		fileprocessing.Merge_file(*file, *folder, *file)
	default:
		fmt.Printf("Unknown mode: %s. Use 'split' or 'merge'.", *mode)
		os.Exit(1)
	}

}

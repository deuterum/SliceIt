package fileprocessing

import (
	"SliceIt/view"
	"crypto/sha256"
	"fmt"
	"math"
	"os"
	"path/filepath"
)

func make_hash_file(file_name, folder string) {
	file_data, err := os.ReadFile(file_name)
	if err != nil {
		panic(err)
	}
	hash := sha256.Sum256(file_data)

	hash_file, err := os.Create(filepath.Join(folder, "file.sha256"))
	if err != nil {
		panic(err)
	}
	hash_file.Write(hash[:])

	fmt.Println("checksum file file.sha256 created in ", folder)
}

func Split_file(file_name string, chunk_size float32, folder string, use_checksum bool) {
	chunk_size_bytes_float := chunk_size * 1024 * 1024
	chunk_buffer := make([]byte, int(chunk_size_bytes_float))

	if folder == "" {
		folder = "parts"
	}

	file, err := os.Open(file_name)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file_stat, err := file.Stat()
	if err != nil {
		panic(err)
	}
	file_size := file_stat.Size()

	total_parts := int(math.Ceil(float64(file_size) / float64(chunk_size_bytes_float)))

	fmt.Println("Filename      : ", file_name)
	fmt.Println("File size     : ", file_size)
	fmt.Println("Chunk size    : ", len(chunk_buffer))
	fmt.Println("Chunks amount : ", total_parts)
	fmt.Println("Output folder : ", folder)

	err = os.MkdirAll(folder, os.ModePerm)
	if err != nil {
		panic(err)
	}

	if use_checksum {
		make_hash_file(file_name, folder)
	}

	fmt.Print("\nProgress\n")
	for part_num := 0; part_num < total_parts; part_num++ {
		n, err := file.Read(chunk_buffer)
		if err != nil {
			panic(err)
		}

		part_filename := fmt.Sprintf("%s_%d.part", filepath.Base(file_name), part_num)
		part_path := filepath.Join(folder, part_filename)

		chunk_file, err := os.Create(part_path)
		if err != nil {
			panic(err)
		}

		_, err = chunk_file.Write(chunk_buffer[:n])
		if err != nil {
			panic(err)
		}

		chunk_file.Close()
		view.Bar(part_num+1, total_parts)
	}
}

package fileprocessing

import (
	"SliceIt/view"
	"bytes"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func CheckHashFile(filename, folder, merged_file_folder string) bool {
	hashBefore, err := os.ReadFile(filepath.Join(folder, "file.sha256"))
	if err != nil {
		panic(err)
	}

	file, err := os.Open(filepath.Join(merged_file_folder, filename))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	hasher := sha256.New()
	_, err = io.Copy(hasher, file)
	if err != nil {
		panic(err)
	}
	hashAfter := hasher.Sum(nil)

	return bytes.Equal(hashBefore, hashAfter)
}

func Merge_file(output_name, folder, base_name string) {
	// Получаем путь к текущему исполняемому файлу
	execPath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	execDir := filepath.Dir(execPath)

	// Если folder — пустая строка, используем папку утилиты
	useExecDir := false
	if folder == "" {
		folder = execDir
		useExecDir = true
	}

	// Читаем файлы в папке
	files, err := os.ReadDir(folder)
	if err != nil {
		panic(err)
	}

	var partFiles []string

	for _, file := range files {
		if strings.HasPrefix(file.Name(), base_name+"_") && strings.HasSuffix(file.Name(), ".part") {
			partFiles = append(partFiles, file.Name())
		}
	}

	if len(partFiles) == 0 {
		panic("No part files found")
	}

	// Сортировка по номеру
	sort.Slice(partFiles, func(i, j int) bool {
		getPartNumber := func(name string) int {
			part := strings.TrimSuffix(strings.TrimPrefix(name, base_name+"_"), ".part")
			num, _ := strconv.Atoi(part)
			return num
		}
		return getPartNumber(partFiles[i]) < getPartNumber(partFiles[j])
	})

	// Путь к выходному файлу
	output_path := filepath.Join(execDir, output_name)
	output_file, err := os.Create(output_path)
	if err != nil {
		panic(err)
	}
	defer output_file.Close()

	fmt.Println("Filename      : ", base_name)
	fmt.Println("Chunks amount : ", len(partFiles))

	fmt.Println("\nMerging parts")
	for i, part := range partFiles {
		part_path := filepath.Join(folder, part)
		part_file, err := os.Open(part_path)
		if err != nil {
			panic(err)
		}

		_, err = io.Copy(output_file, part_file)
		if err != nil {
			panic(err)
		}
		part_file.Close()

		view.Bar(i+1, len(partFiles))
	}

	fmt.Println("\n\nFile merged successfully into : ", output_path)

	checksum_passed := CheckHashFile(output_name, folder, execDir)
	if checksum_passed {
		fmt.Println("Checksum succefully passed")
	} else {
		fmt.Println("hash check failed or hash file was not found")
	}

	if useExecDir {
		// Удаляем .part-файлы по отдельности
		for _, part := range partFiles {
			part_path := filepath.Join(folder, part)
			err := os.Remove(part_path)
			if err != nil {
				fmt.Println("Warning: failed to remove part file:", part, "-", err)
			}
		}
		os.Remove(filepath.Join(folder, "file.sha256"))
		fmt.Println("Removed individual .part files from executable folder.")
	} else {
		// Удаляем папку целиком
		err = os.RemoveAll(folder)
		if err != nil {
			fmt.Println("Warning: failed to remove folder : ", err)
		} else {
			fmt.Println("Removed chunk folder : ", folder)
		}
	}
}

# SliceIt

**SliceIt** is a simple CLI utility written in Go for splitting large files into smaller chunks and merging them back together.

## Features

- Split a file into parts of specified size (in megabytes)
- Merge previously split `.part` files into a single file
- Automatically manages output folders and cleanup

## Usage

```bash
SpliceIt -mode=<split|merge> -file=<filename> [-part_size=<MB>] [-folder=<folder_path>]
```

## Requirements

- github.com/charmbracelet/glamour

## Flags

| Flag         | Description                                                   | Required |
| ------------ | --------------------------------------------------------------| -------- |
| `-mode`      | Mode of operation: `split` or `merge`                         | Yes      |
| `-file`      | Path to the file to split or merge                            | Yes      |
| `-help`      | VIew README.md                                                | No       |
| `-part_size` | Size of each part in megabytes (default: 10)                  | No       |
| `-folder`    | Folder to save/load `.part` files (optional)                  | No       |
| `-hash`      | Use hash to verify file after splitting (default: false)      | No       |
| `-remove`    | Deletes a file after it has been split (default: false)       | No       |

## Output

- During splitting, the tool creates numbered `.part` files in the specified (or default) folder.
- During merging, the tool assembles all matching parts and optionally cleans up the temporary folder or individual part files.

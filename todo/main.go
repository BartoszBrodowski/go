package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func printWithColor(file *os.File) {
    colorRed := "\033[31m"
    colorGreen := "\033[32m"
    colorYellow := "\033[33m"
    colorBlue := "\033[34m"
    colorGray := "\033[37m"
    defer file.Close()
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        if len(line) > 0 && line[0] == '#' {
            fmt.Printf(colorGray + line)
            fmt.Println('\n')
        } else {
            if len(line) > 0 && line[0:3] == "[-]" {
            fmt.Printf(colorRed + line)
            fmt.Println('\n')
            }
            if len(line) > 0 && line[0:3] == "[+]" {
                fmt.Printf(colorYellow + line)
                fmt.Println('\n')
            }
            if len(line) > 0 && line[0:3] == "[x]" {
                fmt.Printf(colorGreen + line)
                fmt.Println('\n')
            }
            if len(line) > 0 && line[0:3] == "[ ]" {
                fmt.Printf(colorBlue + line)
                fmt.Println('\n')
            }
        }
    }
}

func printWithoutColor(file *os.File) {
    defer file.Close()
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        fmt.Println(line)
    }
}

func main() {
    colorPtr := flag.Bool("color", true, "Output colored")
    noColorPtr := flag.Bool("nocolor", false, "Output without color")
    fileNames := os.Args[2:]

    outputFile, err := os.Create("combined.txt")
    if err != nil {
        fmt.Println("Error creating output file:", err)
        return
    }
    defer outputFile.Close()

    for i, fileName := range fileNames {
        file, err := os.Open(fileName)
        if err != nil {
            fmt.Println("Error opening file:", err)
            continue
        }
        defer file.Close()

        if i > 0 {
            outputFile.WriteString("\n") // Add newline separator between files
        }

        io.Copy(outputFile, file)
        if err != nil {
            fmt.Println("Error copying file:", err)
            continue
        }
    }

    combinedFile, err := os.Open("combined.txt")
    if err != nil {
        fmt.Println("Error opening combined file:", err)
        return
    }
    defer combinedFile.Close()

    flag.Parse()

    // file, err := os.Open(fileName)
    // if err != nil {
    //     fmt.Println("Error opening file:", err)
    //     return
    // }

    // Has to be in that order to work
    if *noColorPtr {
        printWithoutColor(combinedFile)
    }
    if *colorPtr {
        printWithColor(combinedFile)
    } 
    defer combinedFile.Close()
    
}


// data := make([]byte, 1024)
//     for {
// 		n, err := file.Read(data)
// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			fmt.Println("File reading error", err)
// 			return
// 		}
// 		fmt.Println(string(data[:n]))
// 	}
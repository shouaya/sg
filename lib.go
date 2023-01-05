package main

import (
    "fmt"
    "strconv"
    "net/http"
    "os"
    "io"
    "github.com/xuri/excelize"
)

func getRows(sheet string)([][]string){
    f, err := excelize.OpenFile("steps.xlsx")
    if err != nil {
        fmt.Println(err)
        return nil
    }
    defer func() {
        // Close the spreadsheet.
        if err := f.Close(); err != nil {
            fmt.Println(err)
        }
    }()
    rows, err := f.GetRows(sheet)
    if err != nil {
        fmt.Println(err)
        return nil
    }
    return rows
}

func getCell(sheet, cellName string)(string){
    f, err := excelize.OpenFile("steps.xlsx")
    if err != nil {
        fmt.Println(err)
        return ""
    }
    defer func() {
        // Close the spreadsheet.
        if err := f.Close(); err != nil {
            fmt.Println(err)
        }
    }()
    // Get value from cell by given worksheet name and cell reference.
    cellValue, err := f.GetCellValue(sheet, cellName)
    if err != nil {
        fmt.Println(err)
        return ""
    }
    return cellValue
}

func toI(v string)(int){
    intV, err := strconv.Atoi(v)
    if err != nil {
        fmt.Println(err)
        return 0
    }
    return intV
}


// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
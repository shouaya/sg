package main

import (
    "fmt"
    "bytes"
    "mime/multipart"
    "path/filepath"
    "strconv"
    "net/http"
    "io/ioutil"
    "os"
    "io"
    "github.com/xuri/excelize"
    "encoding/json"
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

func httpGet(url string)(string){
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println(err)
        return ""
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    fmt.Println(string(body))
    fmt.Println(resp.StatusCode)
    if resp.StatusCode == 200 {
        fmt.Println("ok")
    }
    return string(body)
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

func fileUploadRequest(uri string, params map[string]string, paramName, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", uri, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, err
}

func fileUpload(url, fileName string, target interface{}) {
    path, _ := os.Getwd()
	path += "/" + fileName
	extraParams := map[string]string{}
	request, err := fileUploadRequest(url, extraParams, "file", path)
    if err != nil {
		fmt.Println(err)
	}
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	} else {
        if resp.StatusCode == http.StatusOK {
            bodyBytes, err := io.ReadAll(resp.Body)
            if err != nil {
                fmt.Println(err)
            }
            resp.Body.Close()
            bodyString := string(bodyBytes)
            fmt.Println(bodyString)
            json.Unmarshal(bodyBytes, target)
        }        
	}
}
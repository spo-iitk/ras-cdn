# API Documentation

```sh
    [POST]      /cdn/upload             - Uploads a file to the server
    [DELETE]    /cdn/delete             - Deletes a file from the server
    [GET]       /cdn/view/<filename>    - Views a file from the server
    [GET]       /cdn/view/all           - List of all files on the server
    [POST]      /cdn/zip                - Zips a list of files, returns the filename
    [GET]       /cdn/zip/<filename>     - Zips a list of files and downloads it
    [DELETE]    /cdn/zip                - Deletes a zip from the server
    [DELETE]    /cdn/zip/all            - Deletes all zip files available on the server
```

---

### Body

```go

type Delete*Request struct {
	Filename string `json:"filename"`
}

type ZipRequest struct {
	Files   []string `json:"files"`
	OutFile string   `json:"outfile"`
}

```


# API Documentation

```sh
    [POST]      /api/upload             - Uploads a file to the server
    [DELETE]    /api/delete             - Deletes a file from the server
    [GET]       /api/view/<filename>    - Views a file from the server
    [GET]       /api/view/all           - List of all files on the server
    [POST]      /api/zip                - Zips a list of files, returns the filename
    [GET]       /api/zip/<filename>     - Zips a list of files and downloads it
    [DELETE]    /api/zip                - Deletes a zip from the server
    [DELETE]    /api/zip/all            - Deletes all zip files available on the server
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


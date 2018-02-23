# html2pdf

html2pdf use wkhtmltopdf to generate PDF. It wrapper of the wkhtmltopdf using fasthttp template for html render.

This repo still under development. We accept any pull request. ^\_^

## Installation

```bash
  // dependency
  $ go get github.com/magicwebes/html2pdf
  $ go get github.com/valyala/fasttemplate

  // Downloading wkhtmltopdf from this website and install to the computer or server
  https://wkhtmltopdf.org/downloads.html
```

## Quick Start

### Convert html file to pdf and download

```go
  filepath := "public/views/index.html"
  data := map[string]interface{}{
    "message": "hello",
  }
  outputPath := "public/views/index.pdf"

  if err := html2pdf.ParseHTML(filepath, data).ToFile(outputPath); err != nil {
      log.Println("error", err)
  }
```

## Advance Usage

###

```go
  filepath := "public/views/data.html"
  data := map[string]interface{}{
    "message": "hello",
  }
  outputPath := "public/views/data.pdf"

  if err := html2pdf.ParseHTML(filepath, data).
    SetConfig(&html2pdf.Config{
        Orientation: html2pdf.OrientationLandscape,
        PageSize: html2pdf.PageSizeA4,
    }).
    ToFile(outputPath); err != nil {
      log.Println("error", err)
  }
```

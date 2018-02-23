package html2pdf

import (
	"errors"
	"io"
	"io/ioutil"
	"strings"

	"github.com/magicwebes/html2pdf/wkhtmltopdf"
	"github.com/valyala/fasttemplate"
)

var (
	errFileNotFound          = errors.New("html2pdf: The file not found")
	errWhileExecuteDate      = errors.New("html2pdf: Error while execute data to template file")
	errWhileGeneratingNewPDF = errors.New("html2pdf: Error while generating new pdf")
)

// https://gist.github.com/paulsturgess/cfe1a59c7c03f1504c879d45787699f5

// PDF :
type PDF struct {
	template *wkhtmltopdf.PDFGenerator
	data     []byte
	err      error
}

// ParseHTML :
func ParseHTML(path string, data map[string]interface{}) *PDF {
	p := new(PDF)

	html, err := ioutil.ReadFile(path)
	if err != nil {
		p.err = errFileNotFound
		return p
	}

	t := fasttemplate.New(string(html), "{{", "}}")
	s := t.ExecuteFuncString(func(w io.Writer, tag string) (int, error) {
		trimTag := strings.TrimSpace(tag)
		tagData, err := getBytes(data[trimTag])
		if err != nil {
			return 0, errWhileExecuteDate
		}
		return w.Write(tagData)
	})
	pdf, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		p.err = errWhileGeneratingNewPDF
		return p
	}
	pdf.PageSize.Set(wkhtmltopdf.PageSizeA4)
	pdf.Dpi.Set(300)
	pdf.NoCollate.Set(false)

	p.data = []byte(s)
	p.template = pdf

	return p
}

// Config :
type Config struct {
	Title       string
	Orientation string
	PageSize    string
	Grayscale   bool
}

// SetConfig :
func (p *PDF) SetConfig(c *Config) *PDF {
	if c != nil {
		if len(c.Title) > 0 {
			p.SetTitle(c.Title)
		}
		if len(c.Orientation) > 0 {
			p.SetOrientation(c.Orientation)
		}
		if len(c.PageSize) > 0 {
			p.SetPageSize(c.PageSize)
		}
		if c.Grayscale == true {
			p.Grayscale()
		}
	}

	return p
}

// SetPageSize :
func (p *PDF) SetPageSize(size string) *PDF {
	p.template.PageSize.Set(size)
	return p
}

// SetOrientation :
func (p *PDF) SetOrientation(orientation string) *PDF {
	p.template.Orientation.Set(orientation)
	return p
}

// SetTitle :
func (p *PDF) SetTitle(title string) *PDF {
	p.template.Title.Set(title)
	return p
}

// SetMargin :
func (p *PDF) SetMargin(top, right, bottom, left uint) *PDF {
	p.template.MarginTop.Set(top)
	p.template.MarginRight.Set(right)
	p.template.MarginBottom.Set(bottom)
	p.template.MarginLeft.Set(left)
	return p
}

// Grayscale :
func (p *PDF) Grayscale() *PDF {
	p.template.Grayscale.Set(true)
	return p
}

// ToFile :
func (p *PDF) ToFile(outputPath string) error {
	if p.err != nil {
		return p.err
	}
	p.template.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(string(p.data))))

	if err := p.template.Create(); err != nil {
		return err
	}

	return p.template.WriteFile(outputPath)
}

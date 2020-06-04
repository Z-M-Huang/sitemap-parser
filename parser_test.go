package sitemap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetIndex(t *testing.T) {
	index, err := GetIndex("https://raw.githubusercontent.com/Z-M-Huang/sitemap-parser/master/test/sample_index.xml")
	assert.Empty(t, err)
	assert.NotEmpty(t, index)
}

func TestIndexBadURL(t *testing.T) {
	index, err := GetIndex("abc123")
	assert.NotEmpty(t, err)
	assert.Empty(t, index)
}

func TestIndexInvalidXML(t *testing.T) {
	index, err := GetIndex("https://raw.githubusercontent.com/Z-M-Huang/sitemap-parser/master/test/invalid.xml")
	assert.NotEmpty(t, err)
	assert.Empty(t, index)
}

func TestGetSitemap(t *testing.T) {
	index, err := GetSitemap("https://raw.githubusercontent.com/Z-M-Huang/sitemap-parser/master/test/sample_sitemap.xml")
	assert.Empty(t, err)
	assert.NotEmpty(t, index)
}

func TestSitemapBadURL(t *testing.T) {
	index, err := GetSitemap("abc123")
	assert.NotEmpty(t, err)
	assert.Empty(t, index)
}

func TestSitemapInvalidXML(t *testing.T) {
	index, err := GetSitemap("https://raw.githubusercontent.com/Z-M-Huang/sitemap-parser/master/test/invalid.xml")
	assert.NotEmpty(t, err)
	assert.Empty(t, index)
}

func TestGetSitemapGZ(t *testing.T) {
	index, err := GetSitemapGZ("https://github.com/Z-M-Huang/sitemap-parser/raw/master/test/sample_sitemap.xml.gz")
	assert.Empty(t, err)
	assert.NotEmpty(t, index)
}

func TestSitemapGZBadURL(t *testing.T) {
	index, err := GetSitemapGZ("abc123")
	assert.NotEmpty(t, err)
	assert.Empty(t, index)
}

func TestSitemapGZInvalidXML(t *testing.T) {
	index, err := GetSitemapGZ("https://raw.githubusercontent.com/Z-M-Huang/sitemap-parser/master/test/invalid.xml")
	assert.NotEmpty(t, err)
	assert.Empty(t, index)
}

func TestGetSitemaps(t *testing.T) {
	index, _ := GetIndex("https://raw.githubusercontent.com/Z-M-Huang/sitemap-parser/master/test/sample_index.xml")
	s, err := index.GetSitemaps()
	assert.Empty(t, err)
	assert.NotEmpty(t, s)
}

func TestGetSitemapsInvalidSitemap(t *testing.T) {
	var elements []Element
	elements = append(elements, Element{
		Loc: "https://raw.githubusercontent.com/Z-M-Huang/sitemap-parser/master/test/invalid.xml",
	})
	index := &Index{
		Elements: elements,
	}
	_, err := index.GetSitemaps()
	assert.NotEmpty(t, err)
}

func TestGetSitemapsInvalidSitemapGZ(t *testing.T) {
	var elements []Element
	elements = append(elements, Element{
		Loc: "https://raw.githubusercontent.com/Z-M-Huang/sitemap-parser/master/test/invalid.xml.gz",
	})
	index := &Index{
		Elements: elements,
	}
	_, err := index.GetSitemaps()
	assert.NotEmpty(t, err)
}

func TestGetSitemapsInvalidExtension(t *testing.T) {
	var elements []Element
	elements = append(elements, Element{
		Loc: "https://raw.githubusercontent.com/Z-M-Huang/sitemap-parser/master/test/invalid.xml.invalid",
	})
	index := &Index{
		Elements: elements,
	}
	_, err := index.GetSitemaps()
	assert.NotEmpty(t, err)
}

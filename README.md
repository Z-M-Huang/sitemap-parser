# sitemap-parser
Golang sitemap parser for .xml and .gz

[![Build Status](https://travis-ci.com/Z-M-Huang/sitemap-parser.svg?branch=master)](https://travis-ci.com/Z-M-Huang/sitemap-parser)[![codecov](https://codecov.io/gh/Z-M-Huang/sitemap-parser/branch/master/graph/badge.svg)](https://codecov.io/gh/Z-M-Huang/sitemap-parser)

# Install
`go get github.com/Z-M-Huang/sitemap-parser`

# Example


### Get Index
```go
	index, err := sitemap.GetIndex("https://raw.githubusercontent.com/Z-M-Huang/sitemap-parser/master/test/sample_index.xml")
```

### Get Sitemap
```go
	index, err := sitemap.GetSitemap("https://raw.githubusercontent.com/Z-M-Huang/sitemap-parser/master/test/sample_index.xml")
```

### Get Sitemap in GZ
```go
	index, err := sitemap.GetSitemapGZ("https://github.com/Z-M-Huang/sitemap-parser/raw/master/test/sample_sitemap.xml.gz"")
```

### Get All Sitemaps in Index
```go
	index, _ := sitemap.GetIndex("https://raw.githubusercontent.com/Z-M-Huang/sitemap-parser/master/test/sample_index.xml")
	sitemaps, err := index.GetSitemaps()
```


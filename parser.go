package sitemap

import (
	"compress/gzip"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

//Index sitemap index
type Index struct {
	XMLName  xml.Name  `xml:"sitemapindex"`
	Elements []Element `xml:"sitemap"`
}

//Sitemap sitemap data
type Sitemap struct {
	XMLName  xml.Name  `xml:"urlset"`
	Elements []Element `xml:"url"`
}

//Element single sitemap element
type Element struct {
	Loc        string    `xml:"loc"`
	LastMod    time.Time `xml:"lastmod"`
	ChangeFreq string    `xml:"changefreq"`
	Priority   float32   `xml:"priority"`
}

//GetIndex get sitemap index from URL
func GetIndex(url string) (*Index, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	ret := &Index{}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = xml.Unmarshal(respBytes, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func getSitemapFromBytes(body []byte) (*Sitemap, error) {
	ret := &Sitemap{}
	err := xml.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

//GetSitemap get sitemap from URL
func GetSitemap(url string) (*Sitemap, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return getSitemapFromBytes(respBytes)
}

//GetSitemapGZ get sitemaps from .gz URL
func GetSitemapGZ(url string) (*Sitemap, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	reader, err := gzip.NewReader(resp.Body)
	if err != nil {
		return nil, err
	}
	bodyBytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return getSitemapFromBytes(bodyBytes)
}

//GetSitemaps loads all sitemaps from index
func (i *Index) GetSitemaps() ([]*Sitemap, error) {
	var sitemaps []*Sitemap
	sitemapChan := make(chan *Sitemap, 1)
	errChan := make(chan error, 1)

	for _, e := range i.Elements {
		if strings.HasSuffix(e.Loc, ".xml") {
			//get .xml
			go func() {
				s, err := GetSitemap(e.Loc)
				if err != nil {
					errChan <- err
				} else {
					sitemapChan <- s
				}
				close(sitemapChan)
				close(errChan)
			}()
		} else if strings.HasSuffix(e.Loc, ".gz") {
			//get .gz
		} else {
			return sitemaps, fmt.Errorf("Invalid sitemap loc: %s", e.Loc)
		}
	}

	for j := 0; j < len(i.Elements); j++ {
		if err, open := <-errChan; open {
			return sitemaps, err
		}
		sitemaps = append(sitemaps, <-sitemapChan)
	}
	return sitemaps, nil
}

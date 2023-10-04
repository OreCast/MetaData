package main

import (
	"log"

	oreConfig "github.com/OreCast/common/config"
)

// MetaData represents meta-data object
type MetaData struct {
	ID          string   `json:"id"`
	Site        string   `json:"site" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Bucket      string   `json:"bucket" binding:"required"`
	Tags        []string `json:"tags"`
}

// global list of existing meta-data records
// should be replaced with permistent MongoDB storage
var _metaData []MetaData

// helper function to return existing meta-data
func metadata(site string) []MetaData {
	// so far we will return our global _metaData list
	if oreConfig.Config.MetaData.WebServer.Verbose > 0 {
		log.Println("metadata for site=", site)
	}
	if site == "" {
		return _metaData
	}
	var out []MetaData
	for _, r := range _metaData {
		if oreConfig.Config.MetaData.WebServer.Verbose > 0 {
			log.Printf("MetaData record %+v matching site %s", r, site)
		}
		if r.Site == site {
			out = append(out, r)
		}
	}
	return out
}

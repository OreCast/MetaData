package main

import (
	"log"

	oreConfig "github.com/OreCast/common/config"
	oreMongo "github.com/OreCast/common/mongo"
	"gopkg.in/mgo.v2/bson"
)

// MetaData represents meta-data object
type MetaData struct {
	ID          string   `json:"id"`
	Site        string   `json:"site" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Bucket      string   `json:"bucket" binding:"required"`
	Tags        []string `json:"tags"`
}

// Record converts MetaData to MongoDB record
func (m *MetaData) Record() oreMongo.Record {
	rec := make(oreMongo.Record)
	rec["id"] = m.ID
	rec["site"] = m.Site
	rec["description"] = m.Description
	rec["bucket"] = m.Bucket
	rec["tags"] = m.Tags
	return rec
}

// insert MetaData record to MongoDB
func (m *MetaData) mongoInsert() {
	var records []oreMongo.Record
	records = append(records, m.Record())
	oreMongo.Insert(
		oreConfig.Config.MetaData.MongoDB.DBName,
		oreConfig.Config.MetaData.MongoDB.DBColl,
		records)
}

// remove MetaData record from MongoDB
func (m *MetaData) mongoRemove() {
	spec := bson.M{"id": m.ID}
	oreMongo.Remove(
		oreConfig.Config.MetaData.MongoDB.DBName,
		oreConfig.Config.MetaData.MongoDB.DBColl,
		spec)
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

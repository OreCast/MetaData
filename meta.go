package main

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
	// example of adding meta-data
	/*
		var out []MetaData
		data := MetaData{
			Site:        "SiteA",
			Description: "this site provides access to mineral waste",
			Tags:        []string{"waste", "minerals"},
		}
		out = append(out, data)
		data = MetaData{
			Site:        "SiteB",
			Description: "this site provides access to metal waste",
			Tags:        []string{"waste", "metal"},
		}
		out = append(out, data)
		data = MetaData{
			Site:        "SiteC",
			Description: "this site provides access to glass waste",
			Tags:        []string{"waste", "glass"},
		}
		out = append(out, data)
		return out
	*/
	// so far we will return our global _metaData list
	if site == "" {
		return _metaData
	}
	var out []MetaData
	for _, r := range _metaData {
		if r.Site == site {
			out = append(out, r)
		}
	}
	return out
}

package main

// MetaData represents meta-data object
type MetaData struct {
	Site        string   `json:"site"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}

// helper function to return existing meta-data
func metadata() []MetaData {
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
}

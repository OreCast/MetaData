# MetaData Service
Meta Data service for OreCast. It provides the following functionality:
- RESTful APIs with JSON data format
- it keeps information about participated sites like, site description, etc.

### OreCast APIs
	r.GET("/meta", MetaHandler)
	r.GET("/meta/:site", MetaSiteHandler)

	// all POST methods ahould be authorized
	authorized := r.Group("/")
	authorized.Use(authz.TokenMiddleware(oreConfig.Config.Authz.ClientId, oreConfig.Config.MetaData.Verbose))
	{
		authorized.POST("/meta", MetaPostHandler)
		authorized.DELETE("/meta/:mid", MetaDeleteHandler)

#### public APIs
- `/meta` get all meta data records
- `/meta/:site` get meta data record for a given site

#### Example
```
# get all sites records
curl http://localhost:8300/meta
```

#### protected APIs
- `/meta` post new meta data record
- `/meta/:mid` delete meta data record for a given meta-data ID

#### Example
```
# record.json
{
    "site":"Cornell", 
    "description": "waste minerals", 
    "tags": ["waste", "minerals"],
    "bucket": "waste"
}

# inject new record
curl -v -X POST -H "Content-type: application/json" \
    -H "Authorization: Bearer $token" \
    -d@./record.json \
    http://localhost:8300/meta
```

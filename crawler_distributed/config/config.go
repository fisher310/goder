package config

const (
	ItemSaverPort   = 1234
	ElasticIndex    = "dating_profile"
	ElasticAddress  = "http://10.252.19.55:9200"
	ItemSaverRPC    = "ItemSaverService.Save"
	CrawlServiceRpc = "CrawlService.Process"

	ParseCity     = "CityParser"
	ParseCityList = "CityListParser"
	ParseProfile  = "ProfileParser"
	NilParser     = "NilParser"

	WorkerPort0 = 9000
)

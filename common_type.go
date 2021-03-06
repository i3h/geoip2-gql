package main

import "github.com/graphql-go/graphql"

type SimpleData struct {
	Continent string  `json:"Continent"`
	Country   string  `json:"Country"`
	City      string  `json:"City"`
	Latitude  float64 `json:"Latitude"`
	Longitude float64 `json:"Longitude"`
	TimeZone  string  `json:"TimeZone"`
	IsEU      bool    `json:"IsEU"`
	ASN       uint    `json:"ASN"`
	ORG       string  `json:"ORG"`
}

var namesType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Names",
		Fields: graphql.Fields{
			"de": &graphql.Field{
				Type: graphql.String,
			},
			"en": &graphql.Field{
				Type: graphql.String,
			},
			"es": &graphql.Field{
				Type: graphql.String,
			},
			"fr": &graphql.Field{
				Type: graphql.String,
			},
			"ja": &graphql.Field{
				Type: graphql.String,
			},
			"pt": &graphql.Field{
				Type: graphql.String,
			},
			"ru": &graphql.Field{
				Type: graphql.String,
			},
			"zh": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var cityType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "City",
		Fields: graphql.Fields{
			"GeoNameID": &graphql.Field{
				Type: graphql.String,
			},
			"Names": &graphql.Field{
				Type: namesType,
			},
		},
	},
)

var continentType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Continent",
		Fields: graphql.Fields{
			"Code": &graphql.Field{
				Type: graphql.String,
			},
			"GeoNameID": &graphql.Field{
				Type: graphql.String,
			},
			"Names": &graphql.Field{
				Type: namesType,
			},
		},
	},
)

var countryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Country",
		Fields: graphql.Fields{
			"GeoNameID": &graphql.Field{
				Type: graphql.String,
			},
			"IsInEuropeanUnion": &graphql.Field{
				Type: graphql.Boolean,
			},
			"IsoCode": &graphql.Field{
				Type: graphql.String,
			},
			"Names": &graphql.Field{
				Type: namesType,
			},
		},
	},
)

var locationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Location",
		Fields: graphql.Fields{
			"AccuracyRadius": &graphql.Field{
				Type: graphql.String,
			},
			"Latitude": &graphql.Field{
				Type: graphql.String,
			},
			"Longitude": &graphql.Field{
				Type: graphql.String,
			},
			"MetroCode": &graphql.Field{
				Type: graphql.String,
			},
			"TimeZone": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var postalType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Postal",
		Fields: graphql.Fields{
			"Code": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var subdivisionsType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Subdivisions",
		Fields: graphql.Fields{
			"GeoNameID": &graphql.Field{
				Type: graphql.String,
			},
			"IsoCode": &graphql.Field{
				Type: graphql.String,
			},
			"Names": &graphql.Field{
				Type: namesType,
			},
		},
	},
)

var traitsType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Traits",
		Fields: graphql.Fields{
			"IsAnonymousProxy": &graphql.Field{
				Type: graphql.Boolean,
			},
			"IsSatelliteProvider": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

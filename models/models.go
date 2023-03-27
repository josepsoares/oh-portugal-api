package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Country struct {
	gorm.Model

	NamePt          string         `json:"name_pt"`
	DemonymsPt      pq.StringArray `json:"demonyms_pt"`
	DemonymsEn      pq.StringArray `json:"demonyms:en"`
	Languages       pq.StringArray `json:"languages"`
	Population      int            `json:"population"`
	Area            int            `json:"area"`
	Continents      pq.StringArray `json:"continents"`
	Capital         string         `json:"capital"`
	Region          string         `json:"region"`
	Subregion       string         `json:"subregion"`
	Borders         pq.StringArray `json:"borders"`
	CurrencyName    string         `json:"currency_name"`
	CurrencySymbol  string         `json:"currency_symbol"`
	Gini            string         `json:"gini"`
	UnMember        bool           `json:"un_member"`
	Independent     bool           `json:"independent"`
	PoliticalSystem string         `json:"political_system"`
	Timezones       pq.StringArray `json:"timezone"`
	StartOfWeek     string         `json:"start_of_week"`
	DrivingSide     string         `json:"driving_side"`
	Latitude        string         `json:"latitude"`
	Longitude       string         `json:"longitude"`
	GoogleMaps      string         `json:"google_maps"`
	OpenStreetMaps  string         `json:"open_street_maps"`
	Idd             string         `json:"idd"`
	Tld             string         `json:"tld"`
	Cca2            string         `json:"cca2"`
	Ccn3            string         `json:"ccn3"`
	Cca3            string         `json:"cca3"`
	Cioc            string         `json:"cioc"`
	PostalCode      string         `json:"postal_code"`
	MapImage        string         `json:"map_image"`
	GlobeImage      string         `json:"globe_image"`
	FlagPng         string         `json:"flag_png"`
	FlagSvg         string         `json:"flag_svg"`
	FlagEmoji       string         `json:"flag_emoji"`
	CoatOfArmsPng   string         `json:"coat_of_arms_png"`
	CoatOfArmsSvg   string         `json:"coat_of_arms_svg"`

	Regions []Region
}

type Region struct {
	gorm.Model

	NamePt            string         `json:"name_pt"`
	NameEn            string         `json:"name_en"`
	Description       string         `json:"description"`
	Population        int            `json:"population"`
	PopulationDensity int            `json:"population_density"`
	Area              int            `json:"area"`
	Autonomous        bool           `json:"autonomous"`
	SubRegions        pq.StringArray `json:"sub_regions"`
	Districts         pq.StringArray `json:"districts"`
	Municipalities    pq.StringArray `json:"municipalities"`
	Freguesias        pq.StringArray `json:"freguesias"`
	Image             string         `json:"image"`

	Islands                  []Island
	Rivers                   []*River `gorm:"many2many:region_rivers;"`
	Lagoons                  []Lagoon
	UnescoWorldHeritageSites []UnescoWorldHeritageSite

	CountryID uint
}

type Island struct {
	gorm.Model

	NamePt      string         `json:"name_pt"`
	NameEn      string         `json:"name_en"`
	Description string         `json:"description"`
	Population  int            `json:"population"`
	Images      pq.StringArray `gorm:"type:text[]" json:"images"`

	RegionID uint
}

type River struct {
	gorm.Model

	NamePt         string         `json:"name_pt"`
	NameEn         string         `json:"name_en"`
	Description    string         `json:"description"`
	Length         int            `json:"length"`
	Source         string         `json:"source"`
	SourceCountry  string         `json:"source_country"`
	SourceAltitude int            `json:"source_altitude"`
	Estuary        string         `json:"estuary"`
	AverageFlow    int            `json:"average_flow"`
	WatershedArea  int            `json:"watershed_area"`
	Images         pq.StringArray `gorm:"type:text[]" json:"images"`

	Regions []*Region `gorm:"many2many:region_rivers;"`
}

type Lagoon struct {
	gorm.Model

	NamePt    string         `json:"name_pt"`
	NameEn    string         `json:"name_eng"`
	Area      string         `json:"area"`
	Depth     string         `json:"depth"`
	Latitude  string         `json:"latitude"`
	Longitude string         `json:"longitude"`
	Images    pq.StringArray `gorm:"type:text[]" json:"images"`

	RegionID uint
}

type Mountain struct {
	gorm.Model

	NamePt      string         `json:"name_pt"`
	NameEn      string         `json:"name_en"`
	Description string         `json:"description"`
	Altitude    string         `json:"altitude"`
	Images      pq.StringArray `gorm:"type:text[]" json:"images"`

	Region string
}

type UnescoWorldHeritageSite struct {
	gorm.Model

	NamePt            string         `json:"name_pt"`
	NameEn            string         `json:"name_en"`
	Description       string         `json:"description"`
	Integrity         string         `json:"integrity"`
	Authenticity      string         `json:"authenticity"`
	DateOfInscription string         `json:"date_of_inscription"`
	Latitude          string         `json:"latitude"`
	Longitude         string         `json:"longitude"`
	Images            pq.StringArray `gorm:"type:text[]" json:"images"`

	RegionID uint
}

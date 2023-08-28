package models

import (
	"github.com/lib/pq"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type OrgMember struct {
	Member bool   `json:"member"`
	Date   string `json:"date"`
}

type CountryDemonyms struct {
	Pt pq.StringArray `gorm:"type:text[]" json:"pt"`
	En pq.StringArray `gorm:"type:text[]" json:"en"`
}

type CountryLanguages struct {
	Official   string         `json:"official"`
	Recognized pq.StringArray `gorm:"type:text[]" json:"recognized"`
}

type CountryCapital struct {
	Name      string  `json:"name"`
	NameEn    string  `json:"name_en"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type CountryCurrency struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

type CountryGini struct {
	Value int    `json:"value"`
	Date  string `json:"date"`
}

type CountryMaps struct {
	GoogleMaps string `json:"google_maps"`
	OpenStreet string `json:"open_street"`
}

type CountryPostalCode struct {
	Format string `json:"format"`
	Regex  string `json:"regex"`
}

type CountryFlag struct {
	Png   string `json:"png"`
	Svg   string `json:"svg"`
	Emoji string `json:"emoji"`
}

type CountryCoatOfArms struct {
	Png string `json:"png"`
	Svg string `json:"svg"`
}

type CountryNationalAnthem struct {
	Name  string `json:"name"`
	Track string `json:"track"`
}

type Country struct {
	gorm.Model

	Name              string                                    `json:"name"`
	NameEn            *string                                   `json:"name_en"`
	NameOfficial      string                                    `json:"name_official"`
	NameOfficialEn    *string                                   `json:"name_official_en"`
	Demonyms          datatypes.JSONType[CountryDemonyms]       `json:"demonyms"`
	Languages         datatypes.JSONType[CountryLanguages]      `json:"languages"`
	FoundationDate    bool                                      `json:"foundation_date"`
	Independent       bool                                      `json:"independent"`
	PoliticalSystem   *pq.StringArray                           `json:"political_system"`
	Population        int                                       `json:"population"`
	PopulationDensity int                                       `json:"population_density"`
	Area              int                                       `json:"area"`
	Continents        pq.StringArray                            `gorm:"type:text[]" json:"continents"`
	Capital           datatypes.JSONType[CountryCapital]        `json:"capital"`
	Region            string                                    `json:"region"`
	Subregion         *string                                   `json:"subregion"`
	Borders           *pq.StringArray                           `gorm:"type:text[]" json:"borders"`
	Currency          datatypes.JSONType[CountryCurrency]       `json:"currency"`
	Gini              datatypes.JSONType[CountryGini]           `json:"gini"`
	Nato              datatypes.JSONType[OrgMember]             `json:"nato"`
	UnitedNations     datatypes.JSONType[OrgMember]             `json:"united_nations"`
	EuropeanUnion     datatypes.JSONType[OrgMember]             `json:"european_union"`
	G7                datatypes.JSONType[OrgMember]             `json:"g7"`
	Timezones         pq.StringArray                            `gorm:"type:text[]" json:"timezones"`
	StartOfWeek       string                                    `json:"start_of_week"`
	DrivingSide       string                                    `json:"driving_side"`
	Maps              datatypes.JSONType[CountryMaps]           `json:"maps"`
	Idd               string                                    `json:"idd"`
	Tld               string                                    `json:"tld"`
	Cca2              string                                    `json:"cca2"`
	Ccn3              int                                       `json:"ccn3"`
	Cca3              string                                    `json:"cca3"`
	Cioc              string                                    `json:"cioc"`
	PostalCode        datatypes.JSONType[CountryPostalCode]     `json:"postal_code"`
	Flag              datatypes.JSONType[CountryFlag]           `json:"flag"`
	CoatOfArms        datatypes.JSONType[CountryCoatOfArms]     `json:"coat_of_arms"`
	OfficialWebsites  *pq.StringArray                           `json:"official_website"`
	NationalAnthem    datatypes.JSONType[CountryNationalAnthem] `json:"national_anthem"`

	Regions []Region
}

// TODO => Add hero image prop to most of the structs

type Region struct {
	gorm.Model

	Name              string         `json:"name"`
	NameEn            *string        `json:"name_en"`
	Description       string         `json:"description"`
	Population        int            `json:"population"`
	PopulationDensity string         `json:"population_density"`
	Area              string         `json:"area"`
	Autonomous        bool           `json:"autonomous"`
	SubRegions        pq.StringArray `gorm:"type:text[]" json:"sub_regions"`
	Districts         pq.StringArray `gorm:"type:text[]" json:"districts"`
	Municipalities    pq.StringArray `gorm:"type:text[]" json:"municipalities"`
	Freguesias        pq.StringArray `gorm:"type:text[]" json:"freguesias"`
	Images            pq.StringArray `gorm:"type:text[]" json:"images"`

	Islands                  []Island
	Rivers                   []*River `gorm:"many2many:region_rivers;"`
	Lagoons                  []Lagoon
	UnescoWorldHeritageSites []UnescoWorldHeritageSite

	CountryId uint
}

type Island struct {
	gorm.Model

	Name              string         `json:"name"`
	NameEn            *string        `json:"name_en"`
	Description       *string        `json:"description"`
	Population        int            `json:"population"`
	PopulationDensity string         `json:"population_density"`
	Area              string         `json:"area"`
	Latitude          float64        `json:"latitude"`
	Longitude         float64        `json:"longitude"`
	Images            pq.StringArray `gorm:"type:text[]" json:"images"`

	RegionId uint
}

type Mountain struct {
	gorm.Model

	Name        string         `json:"name"`
	NameEn      *string        `json:"name_en"`
	Description *string        `json:"description"`
	Altitude    string         `json:"altitude"`
	Latitude    float64        `json:"latitude"`
	Longitude   float64        `json:"longitude"`
	Images      pq.StringArray `gorm:"type:text[]" json:"images"`

	RegionId uint
}

// TODO => Revise properties
type NationalPark struct {
	gorm.Model

	Name        string         `json:"name"`
	NameEn      *string        `json:"name_en"`
	Description string         `json:"description"`
	Latitude    *float64       `json:"latitude"`
	Longitude   *float64       `json:"longitude"`
	Images      pq.StringArray `gorm:"type:text[]" json:"images"`

	RegionId uint
}

type River struct {
	gorm.Model

	Name           string         `json:"name"`
	NameEn         *string        `json:"name_en"`
	Description    string         `json:"description"`
	Length         *string        `json:"length"`
	National       bool           `json:"national"`
	Source         string         `json:"source"`
	SourceAltitude string         `json:"source_altitude"`
	Estuary        *string        `json:"estuary"`
	AverageFlow    *string        `json:"average_flow"`
	Images         pq.StringArray `gorm:"type:text[]" json:"images"`

	Regions []*Region `gorm:"many2many:region_rivers;"`
}

type Lagoon struct {
	gorm.Model

	Name      string         `json:"name"`
	NameEn    *string        `json:"name_en"`
	Area      *string        `json:"area"`
	Depth     *string        `json:"depth"`
	Latitude  *float64       `json:"latitude"`
	Longitude *float64       `json:"longitude"`
	Images    pq.StringArray `gorm:"type:text[]" json:"images"`

	RegionId uint
}

// TODO => Revise properties
type Beach struct {
	gorm.Model

	Name        string         `json:"name"`
	NameEn      *string        `json:"name_en"`
	Description string         `json:"description"`
	Images      pq.StringArray `gorm:"type:text[]" json:"images"`

	RegionId uint
}

type UnescoWorldHeritageSite struct {
	gorm.Model

	Name            string         `json:"name"`
	NameEn          *string        `json:"name_en"`
	Description     string         `json:"description"`
	Integrity       string         `json:"integrity"`
	Authenticity    string         `json:"authenticity"`
	InscriptionDate string         `json:"inscription_date"`
	InscriptionYear int            `json:"inscription_year"`
	ApprovedDate    string         `json:"approved_date"`
	ApprovedYear    int            `json:"approved_year"`
	Latitude        *float64       `json:"latitude"`
	Longitude       *float64       `json:"longitude"`
	Images          pq.StringArray `gorm:"type:text[]" json:"images"`

	RegionId uint
}

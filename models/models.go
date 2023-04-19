package models

import (
	"github.com/lib/pq"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type NamePtEn struct {
	Pt string  `json:"pt"`
	En *string `json:"en"`
}

type OrgMember struct {
	Member bool   `json:"member"`
	Date   string `json:"date"`
}

type CountryName struct {
	Official datatypes.JSONType[NamePtEn] `json:"official"`
	Common   datatypes.JSONType[NamePtEn] `json:"common"`
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
	Name      datatypes.JSONType[NamePtEn] `json:"name"`
	Latitude  string                       `json:"latitude"`
	Longitude string                       `json:"longitude"`
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
	Demonyms          datatypes.JSONType[CountryDemonyms]       `json:"demonyms"`
	Languages         datatypes.JSONType[CountryLanguages]      `json:"languages"`
	FoundationDate    bool                                      `json:"foundation_date"`
	Independent       bool                                      `json:"independent"`
	PoliticalSystem   *string                                   `json:"political_system"`
	Population        int                                       `json:"population"`
	PopulationDensity int                                       `json:"population_density"`
	Area              int                                       `json:"area"`
	Continents        pq.StringArray                            `gorm:"type:text[]" json:"continents"`
	Capital           datatypes.JSONType[CountryCapital]        `json:"capital"`
	Region            string                                    `json:"region"`
	Subregion         *string                                   `json:"subregion"`
	Borders           pq.StringArray                            `gorm:"type:text[]" json:"borders"`
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
	OfficialWebsite   *string                                   `json:"official_website"`
	NationalAnthem    datatypes.JSONType[CountryNationalAnthem] `json:"national_anthem"`

	Regions []Region
}

type Region struct {
	gorm.Model

	Name              string         `json:"name"`
	NameEn            *string        `json:"name_en"`
	Description       string         `json:"description"`
	Population        int            `json:"population"`
	PopulationDensity int            `json:"population_density"`
	Area              int            `json:"area"`
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

	CountryID uint
}

type Island struct {
	gorm.Model

	Name              string         `json:"name"`
	NameEn            *string        `json:"name_en"`
	Description       *string        `json:"description"`
	Population        int            `json:"population"`
	PopulationDensity int            `json:"population_density"`
	Area              int            `json:"area"`
	Latitude          string         `json:"latitude"`
	Longitude         string         `json:"longitude"`
	Images            pq.StringArray `gorm:"type:text[]" json:"images"`

	RegionID uint
}

type River struct {
	gorm.Model

	Name           string         `json:"name"`
	NameEn         *string        `json:"name_en"`
	Description    string         `json:"description"`
	Length         int            `json:"length"`
	National       bool           `json:"national"`
	Source         string         `json:"source"`
	SourceAltitude int            `json:"source_altitude"`
	Estuary        string         `json:"estuary"`
	AverageFlow    int            `json:"average_flow"`
	WatershedArea  int            `json:"watershed_area"`
	Images         pq.StringArray `gorm:"type:text[]" json:"images"`

	Regions []*Region `gorm:"many2many:region_rivers;"`
}

type Lagoon struct {
	gorm.Model

	Name      string         `json:"name"`
	NameEn    *string        `json:"name_en"`
	Area      string         `json:"area"`
	Depth     string         `json:"depth"`
	Latitude  string         `json:"latitude"`
	Longitude string         `json:"longitude"`
	Images    pq.StringArray `gorm:"type:text[]" json:"images"`

	RegionID uint
}

type Mountain struct {
	gorm.Model

	Name        string         `json:"name"`
	NameEn      *string        `json:"name_en"`
	Description string         `json:"description"`
	Altitude    string         `json:"altitude"`
	Latitude    string         `json:"latitude"`
	Longitude   string         `json:"longitude"`
	Images      pq.StringArray `gorm:"type:text[]" json:"images"`

	RegionID uint
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
	ApprovedYear    string         `json:"approved_year"`
	Latitude        string         `json:"latitude"`
	Longitude       string         `json:"longitude"`
	Images          pq.StringArray `gorm:"type:text[]" json:"images"`

	RegionID uint
}

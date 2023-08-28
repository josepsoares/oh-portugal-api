package models

type CountryDemonymsJSON struct {
	Pt []string `json:"pt"`
	En []string `json:"en"`
}

type CountryLanguagesJSON struct {
	Official   string   `json:"official"`
	Recognized []string `json:"recognized"`
}

type CountriesJSON struct {
	Name              string                `json:"name"`
	NameEn            *string               `json:"name_en"`
	NameOfficial      string                `json:"name_official"`
	NameOfficialEn    *string               `json:"name_official_en"`
	Demonyms          CountryDemonymsJSON   `json:"demonyms"`
	Languages         CountryLanguagesJSON  `json:"languages"`
	FoundationDate    bool                  `json:"foundation_date"`
	Independent       bool                  `json:"independent"`
	PoliticalSystem   *[]string             `json:"political_system"`
	Population        int                   `json:"population"`
	PopulationDensity int                   `json:"population_density"`
	Area              int                   `json:"area"`
	Continents        []string              `gorm:"type:text[]" json:"continents"`
	Capital           CountryCapital        `json:"capital"`
	Region            string                `json:"region"`
	Subregion         *string               `json:"subregion"`
	Borders           []string              `gorm:"type:text[]" json:"borders"`
	Currency          CountryCurrency       `json:"currency"`
	Gini              CountryGini           `json:"gini"`
	Nato              OrgMember             `json:"nato"`
	UnitedNations     OrgMember             `json:"united_nations"`
	EuropeanUnion     OrgMember             `json:"european_union"`
	G7                OrgMember             `json:"g7"`
	Timezones         []string              `gorm:"type:text[]" json:"timezones"`
	StartOfWeek       string                `json:"start_of_week"`
	DrivingSide       string                `json:"driving_side"`
	Maps              CountryMaps           `json:"maps"`
	Idd               string                `json:"idd"`
	Tld               string                `json:"tld"`
	Cca2              string                `json:"cca2"`
	Ccn3              int                   `json:"ccn3"`
	Cca3              string                `json:"cca3"`
	Cioc              string                `json:"cioc"`
	PostalCode        CountryPostalCode     `json:"postal_code"`
	Flag              CountryFlag           `json:"flag"`
	CoatOfArms        CountryCoatOfArms     `json:"coat_of_arms"`
	OfficialWebsites  []string              `json:"official_website"`
	NationalAnthem    CountryNationalAnthem `json:"national_anthem"`
}

type IslandsJSON []struct {
	Name              string   `json:"name"`
	NameEn            *string  `json:"name_en"`
	Description       *string  `json:"description"`
	Population        int      `json:"population"`
	PopulationDensity string   `json:"population_density"`
	Area              string   `json:"area"`
	Latitude          float64  `json:"latitude"`
	Longitude         float64  `json:"longitude"`
	Images            []string `json:"images"`
	RegionId          uint     `json:"region_id"`
}

type MountainsJSON []struct {
	Name        string   `json:"name"`
	NameEn      *string  `json:"name_en"`
	Description *string  `json:"description"`
	Altitude    string   `json:"altitude"`
	Latitude    float64  `json:"latitude"`
	Longitude   float64  `json:"longitude"`
	Images      []string `json:"images"`
	RegionId    uint     `json:"region_id"`
}

type RiversJSON []struct {
	Name           string   `json:"name"`
	NameEn         *string  `json:"name_en"`
	Description    string   `json:"description"`
	Length         string   `json:"length"`
	National       bool     `json:"national"`
	Source         string   `json:"source"`
	SourceAltitude string   `json:"source_altitude"`
	Estuary        string   `json:"estuary"`
	AverageFlow    *string  `json:"average_flow"`
	Images         []string `json:"images"`
	RegionsIds     []int    `json:"regions_ids"`
}

type LagoonsJSON []struct {
	Name      string   `json:"name"`
	NameEn    *string  `json:"name_en"`
	Area      *string  `json:"area"`
	Depth     *string  `json:"depth"`
	Latitude  *float64 `json:"latitude"`
	Longitude *float64 `json:"longitude"`
	Images    []string `json:"images"`
	RegionId  uint     `json:"region_id"`
}

type UnescoWorldHeritageSitesJSON []struct {
	Name            string   `json:"name"`
	NameEn          *string  `json:"name_en"`
	Description     string   `json:"description"`
	Integrity       string   `json:"integrity"`
	Authenticity    string   `json:"authenticity"`
	InscriptionDate string   `json:"inscription_date"`
	InscriptionYear int      `json:"inscription_year"`
	ApprovedDate    string   `json:"approved_date"`
	ApprovedYear    int      `json:"approved_year"`
	Latitude        *float64 `json:"latitude"`
	Longitude       *float64 `json:"longitude"`
	Images          []string `json:"images"`
	RegionId        uint     `json:"region_id"`
}

type RegionsJSON []struct {
	Name              string   `json:"name"`
	NameEn            *string  `json:"name_en"`
	Description       *string  `json:"description"`
	Population        int      `json:"population"`
	PopulationDensity string   `json:"population_density"`
	Area              string   `json:"area"`
	Autonomous        bool     `json:"autonomous"`
	SubRegions        []string `json:"sub_regions"`
	Districts         []string `json:"districts"`
	Municipalities    []string `json:"municipalities"`
	Freguesias        []string `json:"freguesias"`
	Images            []string `json:"images"`
	CountryId         uint     `json:"country_id"`
}

CREATE TABLE IF NOT EXISTS countries (
  country_id INT primary key GENERATED ALWAYS AS IDENTITY,
  name VARCHAR(255),
  name_en VARCHAR(255),
  name_official VARCHAR(255),
  name_official_en VARCHAR(255),
  demonyms JSON,
  languages TEXT[],
  foundation_date BOOLEAN,
  independent BOOLEAN,
  political_system TEXT[],
  population INTEGER,
  population_density INTEGER,
  area INTEGER,
  continents TEXT[],
  capital JSON,
  region VARCHAR(255),
  subregion VARCHAR(255),
  borders TEXT[],
  currency JSON,
  gini JSON,
  nato JSON,
  united_nations JSON,
  european_union JSON,
  g7 JSON,
  timezones TEXT[],
  start_of_week VARCHAR(255),
  driving_side VARCHAR(255),
  maps JSON,
  idd VARCHAR(255),
  tld VARCHAR(255),
  cca2 VARCHAR(255),
  ccn3 INTEGER,
  cca3 VARCHAR(255),
  cioc VARCHAR(255),
  postal_code JSON,
  flag JSON,
  coat_of_arms JSON,
  official_websites TEXT[],
  national_anthem JSON
)

CREATE TABLE IF NOT EXISTS regions (
  region_id INT primary key GENERATED ALWAYS AS IDENTITY,
  name VARCHAR(255) NOT NULL,
  name_en VARCHAR(255),
  description VARCHAR(255),
  population INTEGER,
  population_density VARCHAR(255),
  area VARCHAR(255),
  autonomous BOOLEAN,
  sub_regions TEXT[],
  districts TEXT[],
  municipalities TEXT[],
  freguesias TEXT[],
  hero_image VARCHAR(255),
  images TEXT[],
  country_id INTEGER REFERENCES countries(id)
);

CREATE TABLE IF NOT EXISTS islands (
  island_id INT primary key GENERATED ALWAYS AS IDENTITY,
  name VARCHAR(250) NOT NULL,
  name VARCHAR(255)
  name_en VARCHAR(255)
  description VARCHAR(255)
  population INTEGER
  population_density VARCHAR(255)
  area VARCHAR(255)
  latitude DOUBLE PRECISION,
  longitude DOUBLE PRECISION,
  hero_image varchar(250) NOT NULL,
  images TEXT[],
  region_id INTEGER REFERENCES regions(id)
);

CREATE TABLE IF NOT EXISTS mountains (
  mountain_id INT primary key GENERATED ALWAYS AS IDENTITY,
  name VARCHAR(255) NOT NULL,
  name_en VARCHAR(255),
  description VARCHAR(255) NOT NULL,
  altitude VARCHAR(255) NOT NULL,
  latitude DOUBLE PRECISION NOT NULL,
  longitude DOUBLE PRECISION NOT NULL,
  images TEXT[],
  region_id INTEGER REFERENCES regions(id)
);

CREATE TABLE IF NOT EXISTS rivers (
  river_id INT primary key GENERATED ALWAYS AS IDENTITY,
  name VARCHAR(255),
  name_en VARCHAR(255),
  description VARCHAR(255),
  length VARCHAR(255),
  national BOOLEAN,
  source VARCHAR(255),
  source_altitude VARCHAR(255),
  estuary VARCHAR(255),
  average_flow VARCHAR(255),
  images TEXT[],
);

CREATE TABLE lagoons (
  lagoon_id INT primary key GENERATED ALWAYS AS IDENTITY,
  name VARCHAR(255),
  name_en VARCHAR(255),
  area VARCHAR(255),
  depth VARCHAR(255),
  latitude DOUBLE PRECISION,
  longitude DOUBLE PRECISION,
  images TEXT[],
  region_id INTEGER REFERENCES regions(id)
);

CREATE TABLE IF NOT EXISTS unesco_world_heritage_sites (
  id INT primary key GENERATED ALWAYS AS IDENTITY,
  name VARCHAR(255) NOT NULL,
  name_en VARCHAR(255),
  description VARCHAR(255) NOT NULL,
  integrity VARCHAR(255) NOT NULL,
  authenticity VARCHAR(255) NOT NULL,
  inscription_date DATE,
  approved_date DATE,
  latitude DOUBLE PRECISION,
  longitude DOUBLE PRECISION,
  images TEXT[],
  region_id INTEGER REFERENCES regions(id)
);

CREATE TABLE river_countries (
  id INT primary key GENERATED ALWAYS AS IDENTITY,
  river_id INTEGER REFERENCES rivers(id),
  country_id INTEGER REFERENCES countries(id)
);

CREATE INDEX idx_river_countries_river_id ON river_countries(river_id);
CREATE INDEX idx_river_countries_country_id ON river_countries(country_id);

CREATE TABLE river_regions (
  id INT primary key GENERATED ALWAYS AS IDENTITY,
  region_id INTEGER REFERENCES regions(id),
  country_id INTEGER REFERENCES countries(id)
);

CREATE INDEX idx_river_regions_river_id ON river_countries(river_id);
CREATE INDEX idx_river_regions_region_id ON river_countries(country_id);
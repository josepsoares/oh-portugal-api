# iberiapi

![postgresql](https://img.shields.io/badge/PostgreSQL-4169E1.svg?style=for-the-badge&logo=PostgreSQL&logoColor=white) ![golang](https://img.shields.io/badge/Go-00ADD8.svg?style=for-the-badge&logo=Go&logoColor=white) ![swagger](https://img.shields.io/badge/Swagger-85EA2D.svg?style=for-the-badge&logo=Swagger&logoColor=black)

## About

A simple golang REST API developed with the [fiber framework](https://gofiber.io/) and documented with [swagger](https://swagger.io/) that returns general info, mostly about geography, about countries from the iberian peninsual (e.g regions, mountains, rivers etc.). The data returned from the API is stored in a postgresql database.

The following recipes from the fiber framework repo were used to streamline the development of this project.

- [gorm-postgres recipe](https://github.com/gofiber/recipes/tree/master/gorm-postgres)
- [swagger recipe](https://github.com/gofiber/recipes/tree/master/swagger)

## Requirements

### with docker

- [docker](https://www.docker.com/)
- [docker-compose](https://docs.docker.com/compose/)

### without docker

- [postgresql](https://www.postgresql.org/)
- [golang](https://go.dev/)

## Author

[josepsoares](https://josepsoares.vercel.app)

# TVDB V4 API for Go

Documentation of TheTVDB API V4 [can be found here](https://thetvdb.github.io/v4-api/).

## Installation

    $ go get github.com/us3r64/tvdb

Running unit tests requires environment keys: TVDB_APIKEY, TVDB_PIN(for a user-supported key)

## Implemented endpoints

There are just a few endpoints implemented so far. Pull requests are welcome at https://github.com/us3r64/tvdb.

---

* `POST /login`
* `GET  /search`
* `GET  /series/{id}`
* `GET  /updates`

## License

The package is available as open source under the terms of the [MIT License](http://opensource.org/licenses/MIT).
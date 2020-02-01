[![Build Status](https://travis-ci.org/as-ideas/happy-stars-go.svg?branch=master)](https://travis-ci.org/as-ideas/happy-stars-go)

Happy Stars API
===============

sample application to play around with CRUD methods. Please us the following endpoints
to modify the galaxy of happy stars

universes endpoint
-----------------
- `GET /api/universes` returns a list of all universes
- `POST /api/universes` adds a new universe with json payload: 
```
{
    id: "id",           /* must be unique */ 
    name: "name",       /* human friendly name*/
    max_size: 123       /* star capacity [0 - 1000000]*/
}
```
- `GET /api/universes/<id>` returns the universe with id 'id'
- `DELETE /api/universes/<id>` removes the universe with id 'id'


info endpoint 
-------------

- `GET /api/colors/values` returns the list of available colors

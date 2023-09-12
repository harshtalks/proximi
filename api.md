# Proximi

proximi is an api service that delivers nearby services to the users. Using Proximi, one can find out nearby businesses with specified range.

## Tech Stack

- Go
- REST APIs
- Postgres
- ORM
- GeoHashing
- JWT based custom Auth

## Features

- Rate Limiting
- Authentication
- Pagination
- Geocoding
- Distance from the business.

## Procedure

- First of all, signin/signup from /auth endpoints to get the auth token (Authenticate Yourself)
- once received the token, copy the token and click on Authorize button to login urself.
- the format is `Bearer <your-token>`
- once upon verification of the token, you will be able to access protected routes such as /api endpoints

## Important

we have applied a rate limiter to make our service always available, and keep our downtime as zero.
check header to see the limit.

### Checkout our better API DOC

[https://bump.sh/harshtalks/doc/proximi/](https://bump.sh/harshtalks/doc/proximi/)

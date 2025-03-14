# CMPR
Simple script for comparing if an actor appears in both movies.

Currently only supports movie querying.

## Usage

- Supply your own API token in the `.env` file through TMDB

```bash
cmpr 'Movie1' 'Movie2'
```
If your movie has multiple words in its title, single quotation marks ('') are **mandatory**.

This will display all actors who appear in both movies and their respective roles.

<br>

`cmpr Irishman Goodfellas`
will result in a list of common actors found, such as:

```bash
Robert De Niro
as Frank Sheeran in The Irishman
as James Conway in Goodfellas
-------------------------------------------------------------
Joe Pesci
as Russell Bufalino in The Irishman
as Tommy DeVito in Goodfellas
-------------------------------------------------------------
Welker White
as Josephine "Jo" Hoffa in The Irishman
as Lois Byrd in Goodfellas
-------------------------------------------------------------
```

## TODO
- tidy
- add tv series
- add init for token checking
- add proper error display
Error fetching cast for `bleack movie`: error fetching movie id: no movies found for query: bleack+movieexit status 1
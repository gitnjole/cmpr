# CMPR
Simple script for comparing if an actor appears in both movies.

Currently only supports movie querying.

## Usage

Supply your own API token in the env file

```bash
cmpr 'Movie1' 'Movie2'`
```

This will display all actors who appeared in both movies and their respective roles.

```bash
cmpr Irishman Goodfellas
```

Will result in a list of common actors found, such as:

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

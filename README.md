# go-graphql-api
GraphQL API for MapDB

This API handles requests for registering, signing in, creating a place, listing places, and listing users.

## What is MapDB?
MapDB allows users to save places on a map for viewing later, when a user launches the app they will be presented with a map with all their locations on it. Users will also be given the option to add new places to the map.

## What does this do?
This is a Go Application that handles API requests, instead of having multiple endpoints `/api/v1/login` and `/api/v1/signup` for example, we have one endpoint `/query` and the data sent to it will change the data sent back. See [GraphQL](https://graphql.org/) for more info.


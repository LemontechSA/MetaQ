package store

import "lemontech.com/metaq/domain"

// ENV stores the current environment
var ENV domain.ENV

// DBs contains the current DBs
var DBs = domain.Databases{}

// Data contains the last results obtained
var Data = domain.Datasets{}

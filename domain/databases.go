package domain

// Database is a struct to represent a single database
type Database struct {
	Name     string
	Selected bool
}

// Databases is a struct to represent available databases
type Databases []Database

package queryhelper

import (
	"github.com/dminGod/D30-HectorDA/endpoint/cassandra_helper"
	"github.com/dminGod/D30-HectorDA/endpoint/presto_helper"
)

// PrepareInsertQuery is used to parse Application metadata
// and return the corresponding INSERT query
func PrepareInsertQuery(metaInput map[string]interface{}) []string {

	// get the endpoint
	databaseType := metaInput["databaseType"].(string)

	var query []string

	if databaseType == "cassandra" {
		query = cassandra_helper.InsertQueryBuild(metaInput)
	}

	return query
}

// PrepareSelectQuery is used to parse Application metadata
// and return the corresponding SELECT query
func PrepareSelectQuery(metaInput map[string]interface{}) []string {
	// get the endpoint
	databaseType := metaInput["databaseType"].(string)

	var query []string

	if databaseType == "cassandra" {

		query = []string{cassandra_helper.SelectQueryBuild(metaInput)}
	} else if databaseType == "presto" {

		query = []string{presto_helper.FindIDQueryBuild(metaInput)}
	} else if databaseType == "cassandra_stratio" {
		query = []string{cassandra_helper.StratioSelectQueryBuild(metaInput)}
	}

	return query
}

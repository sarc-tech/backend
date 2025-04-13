package utils


var POSTGRES_HOST = getEnv("POSTGRES_HOST", "localhost")
var POSTGRES_PORT = getEnv("POSTGRES_PORT", "5432")
var POSTGRES_DB = getEnv("POSTGRES_DB", "sarc")
var POSTGRES_USER = getEnv("POSTGRES_USER", "sarc")
var POSTGRES_PASSWORD = getEnv("POSTGRES_PASSWORD", "sarc")

var PORT = getEnv("PORT", "8082")
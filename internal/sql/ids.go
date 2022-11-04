package sql

type DatabaseId int

type GenericDatabasePrincipalId int

type UserId GenericDatabasePrincipalId

type DatabaseRoleId GenericDatabasePrincipalId

const EmptyDatabasePrincipalId GenericDatabasePrincipalId = -1

type LoginId string

type AADObjectId string

type SchemaId int

type DatabasePrincipalId interface {
	UserId | DatabaseRoleId | GenericDatabasePrincipalId
}

type GenericServerPrincipalId int

type ServerRoleId GenericServerPrincipalId

const EmptyServerPrincipalId GenericServerPrincipalId = -1

type NumericObjectId interface {
	DatabaseId | DatabasePrincipalId | SchemaId
}

type StringObjectId interface {
	LoginId
}

type ObjectId interface {
	NumericObjectId | StringObjectId
}

type DatabasePrincipalType int

const (
	UNONOWN DatabasePrincipalType = iota
	SQL_USER
	DATABASE_ROLE
	AZUREAD_USER
)

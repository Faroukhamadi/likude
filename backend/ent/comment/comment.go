// Code generated by ent, DO NOT EDIT.

package comment

import (
	"time"
)

const (
	// Label holds the string label denoting the comment type in the database.
	Label = "comment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldPoints holds the string denoting the points field in the database.
	FieldPoints = "points"
	// EdgePost holds the string denoting the post edge name in mutations.
	EdgePost = "post"
	// EdgeReplies holds the string denoting the replies edge name in mutations.
	EdgeReplies = "replies"
	// Table holds the table name of the comment in the database.
	Table = "comments"
	// PostTable is the table that holds the post relation/edge.
	PostTable = "comments"
	// PostInverseTable is the table name for the Post entity.
	// It exists in this package in order to avoid circular dependency with the "post" package.
	PostInverseTable = "posts"
	// PostColumn is the table column denoting the post relation/edge.
	PostColumn = "post_comments"
	// RepliesTable is the table that holds the replies relation/edge. The primary key declared below.
	RepliesTable = "comment_replies"
	// RepliesInverseTable is the table name for the Reply entity.
	// It exists in this package in order to avoid circular dependency with the "reply" package.
	RepliesInverseTable = "replies"
)

// Columns holds all SQL columns for comment fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldContent,
	FieldPoints,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "comments"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"post_comments",
}

var (
	// RepliesPrimaryKey and RepliesColumn2 are the table columns denoting the
	// primary key for the replies relation (M2M).
	RepliesPrimaryKey = []string{"comment_id", "reply_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

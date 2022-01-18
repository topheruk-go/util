package sqli

import "fmt"

type Keyword int

const (
	Abort Keyword = iota
	Action
	Add
	After
	All
	Alter
	Always
	Analyze
	And
	As
	Asc
	Attach
	AuotIncrement
	Before
	Begin
	Between
	By
	Cascade
	Case
	Cast
	Check
	Collate
	Column
	Commit
	Conflict
	Constraint
	Create
	Cross
	Current
	CurrentDate
	CurrentTime
	CurrentTimestamp
	Database
	Default
	Deferrable
	Deferred
	Delete
	Desc
	Detach
	Distinct
	Do
	Drop
	Each
	Else
	End
	Escape
	Except
	Exclude
	Exclusive
	Exists
	Explain
	Fail
	Filter
	First
	Following
	For
	Foreign
	From
	Full
	Generated
	Glob
	Group
	Groups
	Having
	If
	Ignore
	Immediate
	In
	Index
	Indexed
	Initially
	Inner
	Insert
	Instead
	Intersect
	Into
	Is
	IsNull
	Join
	Key
	Last
	Left
	Like
	Limit
	Match
	Materialized
	Natural
	No
	Not
	Nothing
	NotNull
	Null
	Nulls
	Of
	Offset
	On
	Or
	Order
	Others
	Outer
	Over
	Partition
	Plan
	Pragma
	Preceding
	Primary
	Query
	Raise
	Range
	Recursive
	References
	RegExp
	ReIndex
	Release
	Rename
	Replace
	Restrict
	Returning
	Right
	Rollback
	Row
	Rows
	SavePoint
	Select
	Set
	Table
	Temp
	Temporary
	Then
	Ties
	To
	Transaction
	Trigger
	Unbounded
	Union
	Unique
	Update
	Using
	Vaccum
	Values
	View
	Virtual
	When
	Where
	Window
	With
	Without
)

func (k Keyword) String() string {
	switch k {
	case Abort:
		return "ABORT"
	case Action:
		return "ACTION"
	case Add:
		return "ADD"
	case After:
		return "AFTER"
	case All:
		return "ALL"
	case Alter:
		return "ALTER"
	case Always:
		return "ALWAYS"
	case Analyze:
		return "ANALYZE"
	case And:
		return "AND"
	case As:
		return "AS"
	case Asc:
		return "ASC"
	case Attach:
		return "ATTACH"
	case AuotIncrement:
		return "AUTOINCREMENT"
	case Before:
		return "BEFORE"
	case Begin:
		return "BEGIN"
	case Between:
		return "BETWEEN"
	case By:
		return "BY"
	case Cascade:
		return "CASCADE"
	case Case:
		return "CASE"
	case Cast:
		return "CAST"
	case Check:
		return "CHECK"
	case Collate:
		return "COLLATE"
	case Column:
		return "COLUMN"
	case Commit:
		return "COMMIT"
	case Conflict:
		return "CONFLICT"
	case Constraint:
		return "CONSTRAINT"
	case Create:
		return "CREATE"
	case Cross:
		return "CROSS"
	case Current:
		return "CURRENT"
	case CurrentDate:
		return "CURRENT_DATE"
	case CurrentTime:
		return "CURRENT_TIME"
	case CurrentTimestamp:
		return "CURRENT_TIMESTAMP"
	case Database:
		return "DATABASE"
	case Default:
		return "DEFAULT"
	case Deferrable:
		return "DEFERRABLE"
	case Deferred:
		return "DEFERRED"
	case Delete:
		return "DELETE"
	case Desc:
		return "DESC"
	case Detach:
		return "DETACH"
	case Distinct:
		return "DISTINCT"
	case Do:
		return "DO"
	case Drop:
		return "DROP"
	case Each:
		return "EACH"
	case Else:
		return "ELSE"
	case End:
		return "END"
	case Escape:
		return "ESCAPE"
	case Except:
		return "EXCEPT"
	case Exclude:
		return "EXCLUDE"
	case Exclusive:
		return "EXCLUSIVE"
	case Exists:
		return "EXISTS"
	case Explain:
		return "EXPLAIN"
	case Fail:
		return "FAIL"
	case Filter:
		return "FILTER"
	case First:
		return "FIRST"
	case Following:
		return "FOLLOWING"
	case For:
		return "FOR"
	case Foreign:
		return "FOREIGN"
	case From:
		return "FROM"
	case Full:
		return "FULL"
	case Generated:
		return "GENERATED"
	case Glob:
		return "GLOB"
	case Group:
		return "GROUP"
	case Groups:
		return "GROUPS"
	case Having:
		return "HAVING"
	case If:
		return "IF"
	case Ignore:
		return "IGNORE"
	case Immediate:
		return "IMMEDIATE"
	case In:
		return "IN"
	case Index:
		return "INDEX"
	case Indexed:
		return "INDEXED"
	case Initially:
		return "INITIALLY"
	case Inner:
		return "INNER"
	case Insert:
		return "INSERT"
	case Instead:
		return "INSTEAD"
	case Intersect:
		return "INTERSECT"
	case Into:
		return "INTO"
	case Is:
		return "IS"
	case IsNull:
		return "ISNULL"
	case Join:
		return "JOIN"
	case Key:
		return "KEY"
	case Last:
		return "LAST"
	case Left:
		return "LEFT"
	case Like:
		return "LIKE"
	case Limit:
		return "LIMIT"
	case Match:
		return "MATCH"
	case Materialized:
		return "MATERIALIZED"
	case Natural:
		return "NATURAL"
	case No:
		return "NO"
	case Not:
		return "NOT"
	case Nothing:
		return "NOTHING"
	case NotNull:
		return "NOTNULL"
	case Null:
		return "NULL"
	case Nulls:
		return "NULLS"
	case Of:
		return "OF"
	case Offset:
		return "OFFSET"
	case On:
		return "ON"
	case Or:
		return "OR"
	case Order:
		return "ORDER"
	case Others:
		return "OTHERS"
	case Outer:
		return "OUTER"
	case Over:
		return "OVER"
	case Partition:
		return "PARTITION"
	case Plan:
		return "PLAN"
	case Pragma:
		return "PRAGMA"
	case Preceding:
		return "PRECEDING"
	case Primary:
		return "PRIMARY"
	case Query:
		return "QUERY"
	case Raise:
		return "RAISE"
	case Range:
		return "RANGE"
	case Recursive:
		return "RECURSIVE"
	case References:
		return "REFERENCES"
	case RegExp:
		return "REGEXP"
	case ReIndex:
		return "REINDEX"
	case Release:
		return "RELEASE"
	case Rename:
		return "RENAME"
	case Replace:
		return "REPLACE"
	case Restrict:
		return "RESTRICT"
	case Returning:
		return "RETURNING"
	case Right:
		return "RIGHT"
	case Rollback:
		return "ROLLBACK"
	case Row:
		return "ROW"
	case Rows:
		return "ROWS"
	case SavePoint:
		return "SAVEPOINT"
	case Select:
		return "SELECT"
	case Set:
		return "SET"
	case Table:
		return "TABLE"
	case Temp:
		return "TEMP"
	case Temporary:
		return "TEMPORARY"
	case Then:
		return "THEN"
	case Ties:
		return "TIES"
	case To:
		return "TO"
	case Transaction:
		return "TRANSACTION"
	case Trigger:
		return "TRIGGER"
	case Unbounded:
		return "UNBOUNDED"
	case Union:
		return "UNION"
	case Unique:
		return "UNIQUE"
	case Update:
		return "UPDATE"
	case Using:
		return "USING"
	case Vaccum:
		return "VACUUM"
	case Values:
		return "VALUES"
	case View:
		return "VIEW"
	case Virtual:
		return "VIRTUAL"
	case When:
		return "WHEN"
	case Where:
		return "WHERE"
	case Window:
		return "WINDOW"
	case With:
		return "WITH"
	case Without:
		return "WITHOUT"
	default:
		return fmt.Sprintf("%d", int(k))
	}
}

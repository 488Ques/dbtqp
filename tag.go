package dbtqp

type TagEntity interface {
	IsTag() bool // Return true if it's a single tag
	SetNext(te TagEntity)
	SetRelationship(rel string)
}

// Data structure of a single tag
type Tag struct {
	Name         string
	Next         TagEntity
	Relationship string
	Negate       bool
}

func (t *Tag) IsTag() bool {
	return true
}

func (t *Tag) SetNext(te TagEntity) {
	t.Next = te
}

func (t *Tag) SetRelationship(rel string) {
	t.Relationship = rel
}

func NewTag(name string, negate bool) *Tag {
	return &Tag{
		Name:         name,
		Relationship: "AND", // Default relationship
		Negate:       negate,
	}
}

// Data structure of a group of tags
type TagGroup struct {
	Tags         []*Tag
	Next         TagEntity
	Relationship string
	Negate       bool
}

func (tg *TagGroup) IsTag() bool {
	return false
}

func (tg *TagGroup) SetNext(te TagEntity) {
	tg.Next = te
}

func (tg *TagGroup) SetRelationship(rel string) {
	tg.Relationship = rel
}

func NewTagGroup(tags []*Tag, negate bool) *TagGroup {
	return &TagGroup{
		Tags:         tags,
		Relationship: "AND", // Default relationship
		Negate:       negate,
	}
}

// Data structure to hold search query information
type TagQuery []TagEntity

package intercom

import "fmt"

// TagService handles interactions with the API through a TagRepository.
type TagService struct {
	Repository TagRepository
}

// Tag represents an Tag in Intercom.
type Tag struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// TagList, an object holding a list of Tags
type TagList struct {
	Tags []Tag `json:"tags,omitempty"`
}

// List all Tags for the App
func (t *TagService) List() (TagList, error) {
	return t.Repository.list()
}

// Save a new Tag for the App.
func (t *TagService) Save(tag *Tag) (Tag, error) {
	return t.Repository.save(tag)
}

// Delete a Tag
func (t *TagService) Delete(id string) error {
	return t.Repository.delete(id)
}

// Tag Users or Companies using a TaggingList.
func (t *TagService) Tag(taggingList *TaggingList) (Tag, error) {
	return t.Repository.tag(taggingList)
}

func (l *TagList) AddressableList(id string) *AddressableList {
	if l == nil {
		return nil
	}

	data := make([]Addressable, len(l.Tags))
	for i, t := range l.Tags {
		data[i] = t.Addressable()
	}

	return &AddressableList{
		Type:       "list",
		Data:       data,
		Url:        fmt.Sprintf("/contacts/%s/tags", id),
		TotalCount: int64(len(data)),
		HasMore:    false,
	}
}

func (t *Tag) Addressable() Addressable {
	return Addressable{
		Type: "tag",
		ID:   t.ID,
		Url:  fmt.Sprintf("tags/%s", t.ID),
	}
}

func (t Tag) String() string {
	return fmt.Sprintf("[intercom] tag { id: %s name: %s }", t.ID, t.Name)
}

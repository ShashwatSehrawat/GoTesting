package main

type Dictionary map[string]string
type DictionaryErr string

const (
	ErrNotFound         = DictionaryErr("word is not in dictionary")
	ErrWordExist        = DictionaryErr("word already exist")
	ErrWordDoesNotExist = DictionaryErr("word is not in dictionary")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}

func (d Dictionary) Update(word, newDefinition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = newDefinition
		return nil
	default:
		return err
	}

}
func (d Dictionary) Search(text string) (string, error) {
	value, ok := d[text]
	if !ok {
		return "", ErrNotFound
	} else {
		return value, nil
	}
}
func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)
	switch err {
	case ErrNotFound:
		d[key] = value
		return nil
	case nil:
		return ErrWordExist
	default:
		return err
	}
}

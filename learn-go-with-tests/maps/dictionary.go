package maps

func Search(dictionary map[string]string, word string) string {
	return dictionary[word]
}

type Dictionary map[string]string

//func (d Dictionary) Search(word string) (string, error) {
//	return d[word], nil
//}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

//func (d Dictionary) Add(word, definiton string) {
//
//	d[word] = definiton
//}

//func (d Dictionary) Add(word, definiton string) error {
//
//	d[word] = definiton
//	return nil
//}

//var (
//	ErrNotFound   = errors.New("could not find the word you were looking for")
//	ErrWordExists = errors.New("cannot add word because it already exists")
//)

func (d Dictionary) Add(word, definition string) error {

	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err

	}

	return nil
}

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Update(word, definition string) error {

	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}

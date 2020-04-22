package main

// map is a reference type

type Dictionary map[string]string

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("can not add word because it alreay exists")
	ErrWordDoesNotExist = DictionaryErr("can not update word because it doesn't exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (definition string, err error) {
	definition, ok := d[word]
	if !ok {
		err = ErrNotFound
	}
	return
}

func (d Dictionary) Add(word, definition string) (err error) {
	_, e := d.Search(word)
	err = e
	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		err = ErrWordExists
	}
	return
}

func (d Dictionary) Update(word, definition string) (err error) {
	_, e := d.Search(word)
	err = e
	switch err {
	case nil:
		d[word] = definition
	case ErrNotFound:
		err = ErrWordDoesNotExist
	}
	return
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}

package maps

const (
	ErrUnknownKey = DictError("not found")
	ErrKeyExists = DictError("key exists")
)

type DictError string

func (e DictError) Error() string {
	return string(e)
}

type Dictionary map[string]string 

func (d Dictionary) Search(key string) (string, error) {

	value, exists := d[key]

	if !exists {
		return "", ErrUnknownKey
	}

	return value, ErrKeyExists
}

func (d Dictionary) Add(key string, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrUnknownKey:
		d[key] = value
	case ErrKeyExists:
		return ErrKeyExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(key string, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrUnknownKey:
		return ErrUnknownKey
	case ErrKeyExists:
		d[key] = value
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(key string) error {
	_, err := d.Search(key)

	switch err {
	case ErrUnknownKey:
		return ErrUnknownKey
	case ErrKeyExists:
		delete(d, key)
	default:
		return err
	}

	return nil
}

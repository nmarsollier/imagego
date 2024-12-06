package db

/**
 * @title ImageDao
 * @desc ImageDao is the interface that defines the methods to interact with the image storage
 */
type ImageDao interface {
	Get(key string) (string, error)
	Set(key string, value string) (string, error)
}

func Get(deps ...interface{}) ImageDao {
	for _, o := range deps {
		if client, ok := o.(ImageDao); ok {
			return client
		}
	}

	return getDynamoDb()
}

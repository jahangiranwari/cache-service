package cache

func Save(key string, data string) {
	client.Set(key, data, Expiry)
}

func Query(key string) string {
	if value, err := client.Get(key).Result(); err == nil {
		return value
	}
	return ""
}

func FlushDB() {
	client.FlushAll()
}

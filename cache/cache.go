package cache

type Cache interface{

    Get(key string) 
    Set(key string, value string)
    Exists(key string)
    Del(key string)
    Keys()
    Flush()
}



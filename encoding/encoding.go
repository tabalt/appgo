package encoding

type Encode interface{

    Read(fileName string) 
    Write(fileName string, content string)
}
package field

//结构体成员访问的方法列表
//方法																说明
//Field(i int) StructField											根据索引返回索引对应的结构体字段的信息，当值不是结构体或索引超界时发生宕机
//NumField() int													返回结构体成员字段数量，当类型不是结构体或索引超界时发生宕机
//FieldByName(name string) (StructField, bool)						根据给定字符串返回字符串对应的结构体字段的信息，没有找到时 bool 返回 false，当类型不是结构体或索引超界时发生宕机
//FieldByIndex(index []int) StructField								多层成员访问时，根据 []int, 提供的每个结构体的字段索引，返回字段的信息，没有找到时返回零值。当类型不是结构体或索引超界时发生宕机
//FieldByNameFunc(match func(string) bool) (StructField,bool)		根据匹配函数匹配需要的字段，当值不是结构体或索引超界时发生宕机

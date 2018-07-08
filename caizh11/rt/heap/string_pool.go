package heap

import "unicode/utf16"

var internedStrings = map[string]*Object{}

/*
JString（）函数根据Go字符串返回相应的Java字符串实例。如果
Java字符串已经在池中，直接返回即可，否则先把Go字符串（UTF8
格式）转换成Java字符数组（UTF16格式），然后创建一个Java字符串
实例，把它的value变量设置成刚刚转换而来的字符数组，最后把
Java字符串放入池中。注意，这里其实是跳过了String的构造函数，
直接用hack的方式创建实例。
 */
func JString(loader *ClassLoader, goStr string) *Object {
	if internedStr, ok := internedStrings[goStr]; ok {
		return internedStr
	}
	chars := stringToUtf16(goStr)
	jChars := &Object{loader.LoadClass("[C"), chars,nil}
	jStr := loader.LoadClass("java/lang/String").NewObject()
	jStr.SetRefVar("value", "[C", jChars)
	internedStrings[goStr] = jStr
	return jStr
}

func stringToUtf16(s string) []uint16 {
	runes := []rune(s) // utf32
	return utf16.Encode(runes)
}



func GoString(jStr *Object) string {
	charArr := jStr.GetRefVar("value", "[C")
	return utf16ToString(charArr.Chars())
}

func utf16ToString(s []uint16) string {
	runes := utf16.Decode(s) // utf8
	return string(runes)
}
func InternString(jStr *Object) *Object {
	goStr := GoString(jStr)
	if internedStr, ok := internedStrings[goStr]; ok {
		return internedStr
	}
	internedStrings[goStr] = jStr
	return jStr
}

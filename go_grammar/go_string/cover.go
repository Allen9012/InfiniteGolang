package main

import "unicode/utf8"

func main() {
	//url := "https://i0.hdslb.com/bfs/tvcover/extEp_SHVRS8974932_snm_1744852145_8009b9e12c06c8455a82596091eecfda.jpg"
	//print(len(url))

	title := "一二三四五六七八九十一二三四五六七八九十一二三四五六七八九十一二三四五六七八九十一二三四五六七八九十一二三四五六七八九十一二三四五六七八九十"
	print(utf8.RuneCountInString(title))
}

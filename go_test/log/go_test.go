package main

import (
	"github.com/c2fo/testify/assert"
	"regexp"
	"strings"
	"testing"
)

var (
	EmptyUnicodeReg = []*regexp.Regexp{
		regexp.MustCompile(`[\x{2060}]+`),  // word joiner
		regexp.MustCompile(`[\x{202e}]+`),  // right-to-left override
		regexp.MustCompile(`[\x{202d}]+`),  // left-to-right override
		regexp.MustCompile(`[\x{200b}]+`),  // zeroWithChar
		regexp.MustCompile(`[\x{1f6ab}]+`), // no_entry_sign
	}
	// 标题里不能出现换行和回车
	NewlineUnicodeRegForTitle = []*regexp.Regexp{
		regexp.MustCompile(`[\n]+`), // newline
		regexp.MustCompile(`[\r]+`), // newline
	}
	NocharReg = []*regexp.Regexp{
		regexp.MustCompile(`[\p{Hangul}]+`),  // kr
		regexp.MustCompile(`[\p{Tibetan}]+`), // tibe
		regexp.MustCompile(`[\p{Arabic}]+`),  // arabic
	}
)

func checkTitleReg(title string) (ct string, ok bool) {
	ct = strings.TrimSpace(title)
	for _, reg := range NocharReg {
		if reg.MatchString(ct) {
			return
		}
	}
	for _, reg := range EmptyUnicodeReg {
		ct = reg.ReplaceAllString(ct, "")
	}
	for _, reg := range NewlineUnicodeRegForTitle {
		ct = reg.ReplaceAllString(ct, "")
	}
	ok = true
	return
}

func TestService_checkTitleReg(t *testing.T) {
	//Convey("checkTitleReg", t, func() {
	// 正常情况
	title := "正常标题"
	ct, ok := checkTitleReg(title)
	t.Log(title, ct, ok)

	// 包含非法字符
	title = "  ( ง `ω´ ).道自己很帅”的臭屁小狗的感觉！"
	ct, ok = checkTitleReg(title)
	t.Log(title, ct, ok)
	assert.Equal(t, title, ct)
	// 包含空字符
	title = "   "
	ct, ok = checkTitleReg(title)
	t.Log(title, ct, ok)

	// 包含换行符
	title = "包含\n换行符"
	ct, ok = checkTitleReg(title)
	t.Log(title, ct, ok)
}

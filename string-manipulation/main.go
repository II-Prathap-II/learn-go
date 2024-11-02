package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("basename looping", basename("a/b.c.go"))
	fmt.Println("basename lastIndexing", basenameWithLib("a/b.c.go"))

	fmt.Println("inserting commas without buffer", insertCommas("12345"))
	fmt.Println("inserting commas with buffer", insertCommasUsingBuffer("12345678"))
	fmt.Println("inserting commas with buffer", insertCommasUsingBufferWriteString("12345678"))

	fmt.Println("convert integer to string using bytes.buffer", intToString(123))
}

func basename(s string) string {
	for i := len(s) - 1; i > 0; i-- {

		// note that single quotes is used here against using "/" in the argument for strings.LastIndex
		if s[i] != '/' {
			continue
		}

		s = s[i+1:]
		break
	}

	for i := len(s) - 1; i > 0; i-- {

		// note that single quotes is used here against using "." in the argument for strings.LastIndex
		if s[i] != '.' {
			continue
		}

		s = s[0:i]
		break
	}
	return s
}

func basenameWithLib(s string) string {

	// note that double quotes is used here against using '/' in the if condition in loop block
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]

	// note that double quotes is used here against using '.' in the if condition in loop block
	dot := strings.LastIndex(s, ".")
	s = s[0:dot]
	return s
}

func insertCommas(num string) string {
	if len(num) <= 3 {
		return num
	}

	return insertCommas(num[0:len(num)-3]) + "," + num[len(num)-3:]
}

func intToString(num int) string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "%d", num) // https://pkg.go.dev/fmt#Fprintf
	return buf.String()
}

func insertCommasUsingBuffer(num string) string {
	// This code works, but instead of writing byte by byte,
	// buf.WriteString can be used to write in groups of 3 to the buffer
	// implement this in insertCommasUsingBufferWriteString

	var buf bytes.Buffer

	if len(num) <= 3 {
		return num
	}

	offset := len(num) % 3
	for i := 0; i < offset; i++ {
		buf.WriteByte(num[i])
	}

	num = num[offset:]
	count := 0
	for i := 0; i < len(num); i++ {
		if count%3 == 0 {
			if offset == 0 {
				offset++
				count++
				buf.WriteByte(num[i])
				continue
			}

			buf.WriteByte(',')
		}

		count++
		buf.WriteByte(num[i])
	}

	return buf.String()
}

func insertCommasUsingBufferWriteString(num string) string {
	var buf bytes.Buffer

	offset := len(num) % 3
	if offset > 0 {
		buf.WriteString(num[:offset])
		if len(num) > offset {
			buf.WriteByte(',')
		}
	}

	// Write the rest of the numbers in groups of three
	for i := offset; i < len(num); i += 3 {
		buf.WriteString(num[i : i+3])
		if i+3 < len(num) {
			buf.WriteByte(',')
		}
	}

	return buf.String()
}

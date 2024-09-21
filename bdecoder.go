package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func unmarshal(data *bufio.Reader) (interface{}, error) {
	ch, err := data.ReadByte()
	if err != nil {
		return nil, err
	}
	if ch == 'i' {
		intBytes, err := data.ReadSlice('e')
		if err != nil {
			return nil, err
		}
		integer, err := strconv.ParseInt(string(intBytes)[:len(intBytes)-1], 10, 64)
		if err != nil {
			return nil, err
		}
		return integer, nil

	} else if ch == 'd' {
		fmt.Println(string(ch))
	} else if ch == 'l' {
		var list []interface{}
		for {
			ch, err := data.ReadByte()
			if err != nil {
				return nil, err
			}
			if ch == 'e' {
				break
			}
			data.UnreadByte()
			item, err := unmarshal(data)
			if err != nil {
				return nil, err
			}
			list = append(list, item)
		}
		return list, nil

	} else {
		data.UnreadByte()
		intBytes, err := data.ReadSlice(':')
		if err != nil {
			return nil, err
		}
		integer, err := strconv.ParseInt(string(intBytes)[:len(intBytes)-1], 10, 64)
		if err != nil {
			return nil, err
		}
		buf := make([]byte, integer)
		_, err = io.ReadFull(data, buf)
		if err != nil {
			return nil, err
		}
		return string(buf), nil
	}
	return nil, nil

}
func main() {
	file, err := os.Open("test2.torrent")
	if err != nil {
		fmt.Println(err)
		return
	}
	reader := bufio.NewReader(file)
	//fmt.Println(reader)
	first, err := unmarshal(reader)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(first)
	second, err := unmarshal(reader)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(second)

}

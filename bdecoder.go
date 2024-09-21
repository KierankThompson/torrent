package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func unmarshal(data *bufio.Reader) (interface{}, error) {
	ch, err := data.ReadByte()
	if err != nil {
		return nil, err
	}
	if ch == 'i' {
		intBytes, err2 := data.ReadSlice('e')
		if err2 != nil {
			return nil, err
		}
		integer, err2 := strconv.ParseInt(string(intBytes)[:len(intBytes)-1], 10, 64)
		if err2 != nil {
			return nil, err2
		}
		return integer, nil

	} else if ch == 'd' {
		fmt.Println(string(ch))
	} else if ch == 'l' {
		fmt.Println(string(ch))
	} else {
		data.UnreadByte()
		intBytes, err2 := data.ReadSlice(':')
		if err2 != nil {
			return nil, err
		}
		integer, err2 := strconv.ParseInt(string(intBytes)[:len(intBytes)-1], 10, 64)
		if err2 != nil {
			return nil, err2
		}

		fmt.Println("none!")
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
	integer, err := unmarshal(reader)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(integer)
	_, err = unmarshal(reader)
	if err != nil {
		fmt.Println(err)
		return
	}

}

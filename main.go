package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"os"
)

func main() {
	var file = flag.String("f", "", "tablespace_file")
	flag.Parse()

	if *file == "" {
		fmt.Printf("Please input tablespace file")
		return
	}

	f, err := os.Open(*file)
	if err != nil {
		fmt.Errorf("open file %s error :%s", *file, err.Error())
		return
	}

	info, err := f.Stat()
	if err != nil {
		fmt.Errorf("get file %s info error :%s", *file, err.Error())
		return
	}
	var res = make(map[uint16]int)
	var i int64 = 0
	b := make([]byte, INNODB_PAGE_SIZE)
	for ; i < info.Size()/INNODB_PAGE_SIZE; i++ {
		if _, err = f.Read(b); err != nil {
			fmt.Errorf("read file %s error :%s", *file, err.Error())
			return
		}
		page_no := binary.BigEndian.Uint32(b[FIL_PAGE_OFFSET:])
		page_type := binary.BigEndian.Uint16(b[FIL_PAGE_TYPE:])
		if page_type == 17855 {
			level := binary.BigEndian.Uint16(b[FIL_PAGE_DATA+PAGE_LEVEL:])
			fmt.Printf("page offset %d, page type <%s>, page level <%d>\n", page_no, innodb_page_type[page_type], level)
		} else {
			fmt.Printf("page offset %d, page type <%s>\n", page_no, innodb_page_type[page_type])
		}

		if _, ok := res[page_type]; ok {
			res[page_type]++
		} else {
			res[page_type] = 1
		}
	}
	fmt.Printf("\n =================================== \n")
	fmt.Printf("Total number of pages: %d\n", info.Size()/INNODB_PAGE_SIZE)
	for k, v := range res {
		fmt.Printf("%s : %d\n", innodb_page_type[k], v)
	}
	return
}

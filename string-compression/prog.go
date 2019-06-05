package main

import (
	"fmt"
	"bytes"
	"strconv"
)

func os(cc rune, count int, os * bytes.Buffer) {
	fmt.Println("dbg:os",cc,count)
	count1,err1 := os.WriteString(string(cc))
	if err1 != nil {
		fmt.Println("dbg:os:err write 1",err1)
	}else {
		fmt.Println("dbg:os:wrote",count1)
	}
	val := strconv.Itoa(count)
	count2,err2 := os.WriteString(val)
	if err2 != nil {
		fmt.Println("dbg:os:err write 2",err2)
	}else {
		fmt.Println("dbg:os:wrote",count2)
	}

}

func compress(s string) string {
	fmt.Println("compress orig:",s,len(s))
	var outputBuffer bytes.Buffer
	var cc rune = -1
	xcount := 0
	for _,v := range s {
		fmt.Println("dbg",v)
		if cc == -1 {
			// this is the start so nothing in front
			cc = v
			xcount = 1
			continue
		} else if cc==v {
			// in this case this is another character
			xcount++
		} else {
			// in this case the cc did not match
			// output the character and the xcount
			os(cc,xcount,&outputBuffer)
			cc=v
			xcount = 1
		}
	}
	// check if we needed to close something out
	if xcount>0 {
		os(cc,xcount,&outputBuffer)
	}
	if len(outputBuffer.String())>len(s) {
		return s
	} 
	return outputBuffer.String()
}

func testBuffer() {
	var buffer bytes.Buffer
	for i := 0;i<3;i++ {
		// does not convert to ascii
		// rep like you would think :()
		buffer.WriteString(string(i))
	}
	fmt.Println(buffer.String())
}

func main() {
	//testBuffer()

	//res:= compress("aaa")
	//fmt.Println("res from compress",res)

	//res := compress("aab")
	//fmt.Println("res from compress",res)

	//res := compress("aabcccc")
	//fmt.Println("res from compress",res,len(res))
	res := compress("aabccccddddd")
	fmt.Println("res from compress",res,len(res))

}
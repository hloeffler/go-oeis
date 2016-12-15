package main

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"strings"
)

type Data struct {
	all []Sequence
}

type Sequence struct {
	ANumber string
	Name    string
	Seq     []*big.Int
}

func main() {
	fmt.Println("hello")
	data, err := getDataFromFiles()
	if err != nil {
		panic(err)
	}
	fmt.Print(len(data.all))
	fmt.Printf("%#v\n", data.all[23234])

	fmt.Println()

	fmt.Println("search:")
	find(data, "0 1 4 9 16 25")

}
func find(data Data, search string) {

	a := strings.Split(search, " ")
	for _, s := range data.all {

		//fmt.Println("search in", s.ANumber)
		skip := false
		match := false

		for i, c := range a {
			if skip == true {
				break
			}
			bi := new(big.Int)
			x, _ := bi.SetString(c, 10)

			//fmt.Printf("comparing %v with %v\n", s.Seq[i], x)

			if s.Seq[i].Cmp(x) != 0 {
				match = false
				skip = true
				break
			} else {

				match = true
			}

		}
		if match == true {
			fmt.Printf("%#v", s)
			fmt.Println("match")
			os.Exit(0)
		}
	}
}

func getDataFromFiles() (Data, error) {
	returnData := Data{}
	returnData.all = []Sequence{}
	names := make(map[string]string)
	_ = names

	dat, err := ioutil.ReadFile("names")
	if err != nil {
		return returnData, err
	}

	lines := strings.Split(string(dat), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		if line[0] == '#' {
			//fmt.Printf("skip: %s\n", line)
			continue
		}
		a := strings.SplitN(string(line), " ", 2)

		if len(a) < 2 {
			fmt.Print("WARN: too short: ", line)
			panic(line)
		}

		names[a[0]] = a[1]
	}

	{
		dat, err := ioutil.ReadFile("stripped")
		if err != nil {
			return returnData, err
		}

		lines := strings.Split(string(dat), "\n")
		for _, line := range lines {
			if line == "" {
				continue
			}
			if line[0] == '#' {
				//fmt.Printf("skip: %s\n", line)
				continue
			}
			a := strings.SplitN(string(line), " ", 2)

			if len(a) < 2 {
				fmt.Print("WARN: too short: ", line)
				panic(line)
			}

			name := a[0]
			seq := a[1]
			seq = seq[1:]

			s := strings.Split(seq, ",")
			s = s[:len(s)-1]

			intSeq := make([]*big.Int, len(s))
			if len(s) > 1 {

				for i := range s {
					bi := new(big.Int)
					x, b := bi.SetString(s[i], 10)
					if b == false && x == nil {
						fmt.Println(name)
						fmt.Printf("%#v", intSeq)
						break
					}
					intSeq[i] = x
				}
			} else {
				intSeq = nil
			}

			S := Sequence{ANumber: name, Name: names[name], Seq: intSeq}
			returnData.all = append(returnData.all, (S))
		}
	}
	return returnData, nil
}

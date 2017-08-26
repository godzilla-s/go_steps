package main

type Tata struct {
	A 		int
	B  		string
	C 		bool
}

func  gobEncode() {
	a := Tata {
		A: 10,
		B: "ABCDEFG",
		C: true,
	}

	fgob, _ := os.OpenFile("test.gob", os.O_RDWR|os.O_CREATE, 0755)
	defer fgob.Close()

	en := gob.NewEncoder(fgob)
	if err := en.Encode(a); err != nil {
		fmt.Println(err)
	}
}

func gobDecode() {
	var d  Tata
	fg, _ := os.Open("test.gob")
	D := gob.NewDecoder(fg)
	D.Decode(&d)
	fmt.Println(d)
}

func main() {
  gobEncode()
  gobDecode()
}

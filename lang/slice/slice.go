package main

import "fmt"

func main() {
	const (
		LevelDebug = iota
		LevelInfo
		LevelWarning
		LevelError
		LevelCritical
		LevelFatal
	)
	var levelPrefix = [...]string{
		LevelDebug:    "DEBUG",
		LevelInfo:     "INFO",
		LevelWarning:  "WARNING",
		LevelError:    "ERROR",
		LevelCritical: "CRITICAL",
		LevelFatal:    "FATAL",
	}

	for index, value := range levelPrefix {
		fmt.Printf("levelPreix[%d]=%s\n", index, value)
	}

	//Go's arrays are values. An array variable denotes the entire array;
	//it is not a pointer to the first array element (as would be the case in C).
	//This means that when you assign or pass around an array value you will
	//make a copy of its contents. (To avoid the copy you could pass a pointer to the array,
	//but then that's a pointer to an array, not an array.) One way to think about arrays
	//is as a sort of struct but with indexed rather than named fields: a fixed-size composite value.
	a := [...]string{"a", "b", "c"}
	var b [3]string = a
	fmt.Printf("Array 'a' element[0] address %p, value is %s\n", &(a[0]), a[0])
	fmt.Printf("Array 'b' element[0] address %p, value is %s\n", &(b[0]), b[0])

	sliceA := make([]string, 3, 50)
	sliceA[0] = "a"
	var sliceB []string = sliceA
	fmt.Printf("sliceA element[0] address %p, value is %s \n", &sliceA[0], sliceA[0])
	fmt.Printf("sliceB element[0] address %p, value is %s \n", &sliceB[0], sliceB[0])

	sliceC := sliceA[:10]
	fmt.Printf("sliceC element[0] address %p, value is %s \n", &sliceC[0], sliceC[0])
	sliceC[9] = "9"
	fmt.Printf("sliceC %v\n", sliceC)
	//growing slice with copy
	sliceD := make([]string, 30, 60)
	copy(sliceD, sliceC)
	sliceD[29] = "29"
	fmt.Printf("sliceD %v\n", sliceD)
	sliceE := sliceD[30:60]
	sliceE[0] = "30"
	fmt.Printf("sliceE %v\n", sliceE)
	//
	sliceA[1] = "b"
	sliceA[2] = "c"
	sliceE[29] = "59"
	fmt.Printf("sliceA %v\n", sliceA)
	fmt.Printf("sliceA[0] address %p\n", &sliceA[0])
	fmt.Printf("sliceE %v\n", sliceE)
	fmt.Printf("sliceE[0] address %p\n", &sliceE[0])
	sliceF := append(sliceA, sliceE...)
	fmt.Printf("sliceA append sliceE %v\n", sliceF)
	fmt.Printf("sliceF[0] address %p\n", &sliceF[0])
	fmt.Printf("sliceF[3] address %p\n", &sliceF[3])

}

package main

type Euro struct{}

type Doller struct{}

func convertEmptySlice(euros []Euro) []Doller {
	//creates the output slice
	dollers := make([]Doller, 0)

	for _, euro := range euros {
		//convert euro to doller and append to the output slice
		dollers = append(dollers, euroToDoller(euro))
	}
	return dollers
}

func euroToDoller(euro Euro) Doller {
	return Doller{}
}

func convertGivenCapacity(euros []Euro) []Doller {
	n := len(euros)
	//intialize with 0 length and given
	dollers := make([]Doller, 0, n)

	for _, euro := range euros {
		//convert euro to doller and append to the output slice
		dollers = append(dollers, euroToDoller(euro))
	}
	return dollers
}

func convertGivenLength(euros []Euro) []Doller {
	n := len(euros)
	//intialize with given length
	// if capacity is not passed then it is same as its length
	dollers := make([]Doller, n)
	for i, euro := range euros {
		//set element i of the slice
		dollers[i] = euroToDoller(euro)
	}
	return dollers
}

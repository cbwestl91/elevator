import "encoding/json"

func StringToByte(input string)(output []byte){
	output, _ := json.Marshall(input)
	
	return output
}

func ByteToString(input []byte)(output string){
	json.Unmarshall(input, &output)
	
	return output
}

func IntToByte(input int)(output []byte){
	output, _ := json.Marshall(input)
	
	return output
}

func ByteToInt(input []byte)(output int){
	json.Unmarshall(input, &output)
	
	return output
}

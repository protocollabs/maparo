package shared

import "encoding/json"
import "fmt"

// channel result

type ChResult struct {
	Json []byte
	ConnObj WriterAnswer
}

type DataObj struct {
	// type is currently 1 byte in type field: 0x02 etc.
	Type uint64
	Id string
	MsmtId uint64
	Seq uint64
	/* not mandatory atm:
	Timestamp time.Date
	Secret string
	SeqRp uint64
	Modules []string => could be also [][]string
	Os string
	Arch string
	Measurement [][]string
	ControlProtocol [][]string
	*/
}

func NewDataObj() *DataObj {
	dataObj := new(DataObj)
	return dataObj
}

func (data *DataObj) TransformJson(jsonData []byte) {
	err := json.Unmarshal(jsonData, data)
	if err != nil {
		fmt.Printf("Cannot parse JSON %s\n", err)
	}
}
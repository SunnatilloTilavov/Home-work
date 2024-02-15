package main

import (
	"fmt"
)

type Aeraport struct {
	PlaneType        string
	FlightDate int
	FromCity     string
	ToCity  string
	FlightTime int
}

func (r*Aeraport)Flight () string {
	if r.FlightDate==8{
		return r.PlaneType
	}
	return ""

}
func main() {
	airplane1:=Aeraport{
	PlaneType  : "A733",
	FlightDate :8,
	FromCity   :  "Tashkent",
	ToCity : "Samarkand",
	FlightTime :90,
	}
	airplane2:=Aeraport{
		PlaneType  : "A743",
		FlightDate :10,
		FromCity   :  "Tashkent1",
		ToCity : "Samarkand1",
		FlightTime :100,
	}
	airplane3:=Aeraport{
		PlaneType  : "A7323",
		FlightDate :9,
		FromCity   :  "Tashkent12",
		ToCity : "Samarkand3",
		FlightTime :930,
	}
	airplane4:=Aeraport{
		PlaneType  : "A7373",
		FlightDate :8,
		FromCity   :  "Tashkent3",
		ToCity : "Samarkand1",
		FlightTime :900,
	}
	airplane5:=Aeraport{
		PlaneType  : "A773",
		FlightDate :8,
		FromCity   :  "Tashkent2",
		ToCity : "Samarkand1",
		FlightTime :900,
	}
	airplane6:=Aeraport{
		PlaneType  : "A373",
		FlightDate :9,
		FromCity   :  "Tashkent4",
		ToCity : "Samarkand12",
		FlightTime :900,
	}
	airplane7:=Aeraport{
		PlaneType  : "A737",
		FlightDate :1,
		FromCity   :  "Tashkentq",
		ToCity : "Samarkandw",
		FlightTime :900,
	}
	airplane8:=Aeraport{
		PlaneType  : "A75703",
		FlightDate :8,
		FromCity   :  "Tashkent",
		ToCity : "Samarkand",
		FlightTime :900,
	}
	airplane9:=Aeraport{
		PlaneType  : "A766373",
		FlightDate :8,
		FromCity   :  "Tashkent",
		ToCity : "Samarkand",
		FlightTime :100,
	}
	airplane10:=Aeraport{
		PlaneType  : "A71233",
		FlightDate :8,
		FromCity   :  "Tashkent",
		ToCity : "Samarkand",
		FlightTime :900,
	}
	a:=airplane1.Flight ()+" "+airplane2.Flight ()+" "+airplane3.Flight ()+" "+airplane4.Flight ()+" "+airplane5.Flight ()+" "+airplane6.Flight ()+" "+airplane7.Flight ()+" "+airplane8.Flight ()+" "+airplane9.Flight ()+" "+airplane10.Flight ()
fmt.Print(a)
}

	
	

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

func (r*Aeraport)Flighttime () string {
	if r.FlightTime>=2&&r.FlightTime<=3&&r.ToCity=="Tashkent"{
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
	FlightTime :2,
	}
	airplane2:=Aeraport{
		PlaneType  : "A743",
		FlightDate :10,
		FromCity   :  "Samarqand",
		ToCity : "Tashkent",
		FlightTime :3,
	}
	airplane3:=Aeraport{
		PlaneType  : "A7323",
		FlightDate :9,
		FromCity   :  "Tashkent",
		ToCity : "Samarkand",
		FlightTime :1,
	}
	airplane4:=Aeraport{
		PlaneType  : "A7373",
		FlightDate :8,
		FromCity   :  "Tashkent",
		ToCity : "Samarkand",
		FlightTime :5,
	}
	airplane5:=Aeraport{
		PlaneType  : "A773",
		FlightDate :8,
		FromCity   :  "Tashkent",
		ToCity : "Samarkand",
		FlightTime :6,
	}
	airplane6:=Aeraport{
		PlaneType  : "A373",
		FlightDate :2,
		FromCity   :  "Tashkent",
		ToCity : "Samarkand",
		FlightTime :4,
	}
	airplane7:=Aeraport{
		PlaneType  : "A737",
		FlightDate :1,
		FromCity   :  "Samarqand",
		ToCity : "Tashkent",
		FlightTime :3,
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
		FromCity   :  "Samarqand",
		ToCity : "Tashkent",
		FlightTime :12,
	}
	airplane10:=Aeraport{
		PlaneType  : "A71233",
		FlightDate :8,
		FromCity   :  "Tashkent",
		ToCity : "Samarkand",
		FlightTime :6,
	}
	a:=airplane1.Flighttime ()+" "+airplane2.Flighttime ()+" "+airplane3.Flighttime ()+" "+airplane4.Flighttime ()+" "+airplane5.Flighttime ()+" "+airplane6.Flighttime ()+" "+airplane7.Flighttime ()+" "+airplane8.Flighttime ()+" "+airplane9.Flighttime ()+" "+airplane10.Flighttime ()
fmt.Print(a)
}

	
	

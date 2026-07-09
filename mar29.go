package main

import "fmt"

func chem() {
	elements := make(map[string]string)

	elements["H"] = "hydrogen"
	elements["He"] = "helium"
	elements["Li"] = "lithium"
	elements["B"] = "beryllium"
	elements["Bo"] = "boron"
	elements["C"] = "carbon"
	elements["Ni"] = "nitrogen"
	elements["O"] = "oxygen"
	elements["F"] = "Flourine"
	elements["Ne"] = "neon"
	fmt.Println(elements["Li"])

	fellows := make(map[string]string)

	fellows["A"] = "abraham"
	fellows["B"] = "benjamin"
	fellows["C"] = "caro"
	fellows["D"] = "daniel"
	fellows["E"] = "ella"

	fmt.Println(fellows["D"])

	guy, ok := fellows["A"]
	fmt.Println(guy, ok)

	// name, ok := elements["Un"]
	// fmt.Println(name, ok)

	elements1 := map[string]map[string]string{
		"H": {
			"name":  "hydrogen",
			"state": "gas",
		},
		"He": {
			"name":  "helium",
			"state": "gas",
		},
	}
	if el, okk := elements1["H"]; okk {
		fmt.Println(el["name"], el["state"])
	}
}

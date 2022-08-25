package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

/*
This package describes working with writers and readers. Also it describes ho to read and write JSON Data.
*/

func main() {
	reader := strings.NewReader(`{"Name":"Kayak","Category":"Watersports","Price":279, "Offer": "10"}`)
	decoder := json.NewDecoder(reader)

	for {
		val := DiscountedProduct{}
		err := decoder.Decode(&val)
		if err != nil {
			if err != io.EOF {
				Printfln("Error: %v", err.Error())
			}
			break
		} else {
			Printfln("Name: %v, Category: %v, Price: %v, Discount: %v", val.Name, val.Category, val.Price, val.Discount)
		}
	}
}

func MakeMap[K comparable, V any](K, V) map[K]V {
	return make(map[K]V)
}

func decodeMaps() {
	reader := strings.NewReader(`{"Kayak" : 279, "Lifejacket" : 49.95}`)
	//m := MakeMap("", 0.0)
	//m := map[string]interface{}{}
	m := map[string]float64{}

	decoder := json.NewDecoder(reader)

	err := decoder.Decode(&m)
	if err != nil {
		Printfln("Error: %v", err.Error())
	} else {
		Printfln("Map: %T, %v", m, m)
		for k, v := range m {
			Printfln("Key: %v, Value: %v", k, v)
		}
	}
}

func decodeArrayJSON() {
	reader := strings.NewReader(`[10,20,30] ["Kayak","Lifejacket",279]`)
	vals := []interface{}{}
	decoder := json.NewDecoder(reader)
	for {
		var decodedVal interface{}
		err := decoder.Decode(&decodedVal)
		if err != nil {
			if err != io.EOF {
				Printfln("Error: %v", err.Error())
			}
			break
		}
		vals = append(vals, decodedVal)
	}
	for _, val := range vals {
		Printfln("Decoded (%T): %v", val, val)
	}

	//reader := strings.NewReader(`[10,20,30]["Kayak","Lifejacket",279]`)
	//ints := []int{}
	//mixed := []interface{}{}
	//vals := []interface{}{&ints, &mixed}
	//decoder := json.NewDecoder(reader)
	//for i := 0; i < len(vals); i++ {
	//	err := decoder.Decode(vals[i])
	//	if err != nil {
	//		Printfln("Error: %v", err.Error())
	//		break
	//	}
	//}
	//Printfln("Decoded (%T): %v", ints, ints)
	//Printfln("Decoded (%T): %v", mixed, mixed)
}

func decodeJSONSpecificValues() {
	reader := strings.NewReader(`true "Hello" 99.99 200`)
	var bval bool
	var sval string
	var fpval float64
	var ival int
	vals := []interface{}{&bval, &sval, &fpval, &ival}
	decoder := json.NewDecoder(reader)
	for i := 0; i < len(vals); i++ {
		err := decoder.Decode(vals[i])
		if err != nil {
			Printfln("Error: %v", err.Error())
			break
		}
	}
	Printfln("Decoded (%T): %v", bval, bval)
	Printfln("Decoded (%T): %v", sval, sval)
	Printfln("Decoded (%T): %v", fpval, fpval)
	Printfln("Decoded (%T): %v", ival, ival)
}

func decodeJSONNumbers() {
	reader := strings.NewReader(`true "Hello" 99.99 200`)
	vals := []interface{}{}
	decoder := json.NewDecoder(reader)
	decoder.UseNumber()
	for {
		var decodedVal interface{}
		err := decoder.Decode(&decodedVal)
		if err != nil {
			if err != io.EOF {
				Printfln("Error: %v", err.Error())
			}
			break
		}
		vals = append(vals, decodedVal)
	}
	for _, val := range vals {
		if num, ok := val.(json.Number); ok {
			if ival, err := num.Int64(); err == nil {
				Printfln("Decoded Integer: %v", ival)
			} else if fpval, err := num.Float64(); err == nil {
				Printfln("Decoded Floating Point: %v", fpval)
			} else {
				Printfln("Decoded String: %v", num.String())
			}
		} else {
			Printfln("Decoded (%T): %v", val, val)
		}
	}
}

func decodeJSONData() {
	reader := strings.NewReader(`true "Hello" 99.99 200`)
	vals := []interface{}{}
	decoder := json.NewDecoder(reader)
	for {
		var decodedVal interface{}
		err := decoder.Decode(&decodedVal)
		if err != nil {
			if err != io.EOF {
				Printfln("Error: %v", err.Error())
			}
			break
		}
		vals = append(vals, decodedVal)
	}
	for _, val := range vals {
		Printfln("Decoded (%T): %v", val, val)
	}
}

func funcName() {
	var writer strings.Builder
	encoder := json.NewEncoder(&writer)
	dp := DiscountedProduct{
		Product:  &Kayak,
		Discount: 10.50,
	}

	namedItems := []Named{&dp, &Person{PersonName: "Alice"}}
	encoder.Encode(namedItems)

	fmt.Print(writer.String())
}

func encodeJSON() {
	var b = true
	var str = "Hello"
	var fval = 99.99
	var ival = 200
	var pointer = &ival

	var writer strings.Builder
	encoder := json.NewEncoder(&writer)

	for _, val := range []interface{}{b, str, fval, ival, pointer} {
		encoder.Encode(val)
	}
	fmt.Print(writer.String())
}

func writeReplaced() {
	text := "It was a boat. A small boat."
	subs := []string{"boat", "duck", "small", "huge"}
	var writer strings.Builder
	replacer := strings.NewReplacer(subs...)
	replacer.WriteString(&writer, text)
	fmt.Println(writer.String())
}

func writeFormatted() {
	var writer strings.Builder
	template := "Name: %s, Category: %s, Price: $%.2f"
	fmt.Fprintf(&writer, template, "Kayak", "Watersports", float64(279))
	fmt.Println(writer.String())
}

func scanFromReader() {
	reader := strings.NewReader("Kayak Watersports $279.00")
	var name, category string
	var price float64
	scanTemplate := "%s %s $%f"
	_, err := fmt.Fscanf(reader, scanTemplate, &name, &category, &price)
	if err != nil {
		Printfln("Error: %v", err.Error())
	} else {
		Printfln("Name: %v", name)
		Printfln("Category: %v", category)
		Printfln("Price: %.2f", price)
	}
}

func scanSingle() {
	reader := strings.NewReader("Kayak Watersports $279.00")

	for {
		var str string
		_, err := fmt.Fscan(reader, &str)
		if err != nil {
			if err != io.EOF {
				Printfln("Error: %v", err.Error())
			}
			break
		}
		Printfln("Value: %v", str)
	}
}

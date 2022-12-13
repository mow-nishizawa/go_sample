package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Fruits struct {
	Name           string
	Money          int
	Quantity       int
	QuantityPerYen int
}

const (
	APPLE  = 100
	ORANGE = 50
	GRAPE  = 200
	PEACH  = 150
)

func RequestParseHandler(w http.ResponseWriter, r *http.Request) {
	apple_quantity := formValueInt(r, "apple_num", 0)
	orange_quantity := formValueInt(r, "orange_num", 0)
	grape_quantity := formValueInt(r, "grape_num", 0)
	peach_quantity := formValueInt(r, "peach_num", 0)

	t, err := template.ParseFiles("templates/request_parse.html")
	if err != nil {
		log.Fatal(err)
	}

	if apple_quantity == 0 && orange_quantity == 0 && grape_quantity == 0 && peach_quantity == 0 {
		if err := t.Execute(w, nil); err != nil {
			log.Fatal(err)
		}
		return
	}

	am := APPLE * apple_quantity
	om := ORANGE * orange_quantity
	gm := GRAPE * grape_quantity
	pm := PEACH * peach_quantity

	var fruits []Fruits
	if apple_quantity != 0 {
		fruits = append(fruits, Fruits{"りんご", am, apple_quantity, APPLE})
	}
	if orange_quantity != 0 {
		fruits = append(fruits, Fruits{"みかん", om, orange_quantity, ORANGE})
	}
	if grape_quantity != 0 {
		fruits = append(fruits, Fruits{"ぶどう", gm, grape_quantity, GRAPE})
	}
	if peach_quantity != 0 {
		fruits = append(fruits, Fruits{"もも", pm, peach_quantity, PEACH})
	}

	d := struct {
		Fruits []Fruits
		Total  int
	}{
		fruits,
		am + om + gm + pm,
	}

	if err := t.Execute(w, d); err != nil {
		log.Fatal(err)
	}
}

func formValueInt(r *http.Request, k string, defaultValue int) int {
	if x := strings.TrimSpace(r.FormValue(k)); len(x) > 0 {
		if v, err := strconv.Atoi(x); err == nil {
			return v
		}
	}
	return defaultValue
}

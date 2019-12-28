package burgundy

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strconv"
)

func processHeaders(item Subject, order ColOrder) Headers {
	headers := Headers{}
	v := reflect.ValueOf(item)
	if v.Type().Kind() != reflect.Ptr {
		v = reflect.New(reflect.TypeOf(item)) // create new pointer
	}
	e := v.Elem()
	for i := 0; i < len(order); i++ {
		headers = append(headers, e.Type().Field(order[i]).Name)
	}
	return headers
}

func processRow(item Subject, order ColOrder) Row {
	row := Row{}
	v := reflect.ValueOf(item)
	if v.Type().Kind() == reflect.Ptr {
		v = v.Elem() // create new pointer
	}
	// e := reflect.Indirect(v).Elem()
	for i := 0; i < len(order); i++ {
		fmt.Printf("appending %v for %+v\n", v.Field(order[i]).Interface(), v.Field(order[i]))
		row = append(row, v.Field(order[i]).Interface())
	}
	fmt.Printf("row:%v\n", row)
	return row
}

//TODO: will want to make it so that this can handle tags with values higher than len of fields
func processOrder(item Subject) ColOrder {
	order := ColOrder{}

	v := reflect.TypeOf(item)
	// if v.Kind() != reflect.Ptr {
	// 	v = reflect.New(reflect.TypeOf(item)) // create new pointer
	// }
	// e := v.Elem()

	fields := make(FieldIndexes, 0)

	for i := 0; i < v.NumField(); i++ {
		tagIndex := -1
		tag := v.Field(i).Tag.Get("col")
		if tag == "-" {
			continue
		} else if i, err := strconv.Atoi(tag); err == nil {
			tagIndex = i
		}
		fields = append(fields, FieldIndexing{
			Index:    i,
			TagIndex: tagIndex,
		})
	}

	sort.Sort(fields)

	for _, fieldIdx := range fields {
		order = append(order, fieldIdx.Index)
	}

	return order
}

func Report(items interface{}, e Reporter) ([]byte, error) {
	return Process(items, e)
}

/// Process takes the given items and reflects the structure of them to find the headers and
/// proper order of the columns. It then grabs all the data creating a row per item and passes
/// this all on to the star of the show the Anchor... I mean Reporter.
func Process(items interface{}, e Reporter) ([]byte, error) {
	t := reflect.ValueOf(items)
	if t.Kind() != reflect.Slice {
		return []byte{}, errors.New("Must pass a slice of objects.")
	}

	input := make([]Row, t.Len())
	headers := make(Headers, 0)

	if t.Len() <= 0 {
		return []byte{}, nil
	}

	order := processOrder(t.Index(0).Interface())
	fmt.Printf("order:%v\n", order)

	for i := 0; i < t.Len(); i++ {
		item := t.Index(i).Interface()
		if i == 0 {
			headers = processHeaders(item, order)
		}
		input[i] = processRow(item, order)
	}

	return e.Process(headers, input)
}

/// DataBlockForProcessing takes the [headers] and [rows] and creates a 2D array of strings
/// if [dontAddHeaders] isn't true the first row will be the headers, and the rest of the
/// row data will be converted to strings and put in place. If the type of the data in a given
/// field/cell adheres to the [fmt.Stringer] interface that will be used in the conversion.
func DataBlockForProcessing(headers Headers, rows []Row, dontAddHeaders bool) [][]string {
	output_data := [][]string{}
	if !dontAddHeaders {
		output_data = append(output_data, headers)
	}
	for _, row := range rows {
		new_row := make([]string, len(row))
		for i, field := range row {
			if str, ok := field.(fmt.Stringer); ok {
				new_row[i] = str.String()
			} else {
				new_row[i] = fmt.Sprintf("%v", field)
			}
		}
		output_data = append(output_data, new_row)
	}
	return output_data
}

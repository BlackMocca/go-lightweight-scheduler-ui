package constants

import "reflect"

var (
	CHUNK = func(dataSlice reflect.Value, chunkSize int) []reflect.Value {
		batchs := make([]reflect.Value, 0)
		for i := 0; i < dataSlice.Len(); i += chunkSize {
			end := i + chunkSize
			if end > dataSlice.Len() {
				end = dataSlice.Len()
			}

			batch := dataSlice.Slice(i, end)
			batchs = append(batchs, batch)
		}
		return batchs
	}
)

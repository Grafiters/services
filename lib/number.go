package lib

import (
    "reflect"
    "math"
    "strconv"
    "strings"
)

func IsNumber(value interface{}) bool {
    kind := reflect.TypeOf(value).Kind()
    switch kind {
    case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
        return true
    case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
        return true
    case reflect.Float32, reflect.Float64:
        return true
    default:
        return false
    }
}

func ToFixed(num float64, precision int) string {
    output := strconv.FormatFloat(math.Round(num*math.Pow(10, float64(precision)))/math.Pow(10, float64(precision)), 'f', precision, 64)
    if precision > 0 {
        output = strings.TrimSuffix(output, "0")
        output = strings.TrimSuffix(output, ".")
    }
    return output
}

func ToFixedWithPercent(num float64, precision int) string {
    output := strconv.FormatFloat(math.Round(num*math.Pow(10, float64(precision)))/math.Pow(10, float64(precision)), 'f', precision, 64)
    if precision > 0 {
        output = strings.TrimSuffix(output, "0")
        output = strings.TrimSuffix(output, ".")
    }
    strings.TrimSuffix(output, "%") 

    return output
}
package utils

type any interface{}

func Map[T any, U any](input []T, operator func(T, int) U) []U {
    output := make([]U, len(input))
    for i, element := range input {
        output[i] = operator(element, i)
    }
    return output
}

func Filter[T any](input []T, predicate func(T) bool) []T {
    output := []T{}
    for _, element := range input {
        if predicate(element) {
            output = append(output, element)
        }
    }
    return output
}

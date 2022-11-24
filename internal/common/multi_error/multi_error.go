package multi_error

import multierr "github.com/hashicorp/go-multierror"

// Append добавляет ошибки errors в агреггирующую ошибку aggregateError.
func Append(aggregateError error, errors ...error) error {
	for _, err := range errors {
		if err != nil {
			if aggregateError == nil {
				aggregateError = err
			} else {
				aggregateError = multierr.Append(aggregateError, err)
			}
		}
	}
	return aggregateError
}

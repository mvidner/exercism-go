package erratum

// Call *o* to get a resource and use it to Frob *input*.
// Close before returning.
// Retry TransientErrors.
// If Frob panics with FrobError, Defrob the resource with a defrobTag
// In any case, recover from any Frob panics and return them.
func Use(o ResourceOpener, input string) (err error) {
	var rez Resource
	for {
		rez, err = o()
		if err == nil {
			break // no error
		}
		switch err.(type) {
		case TransientError:
			continue // keep retrying ResourceOpener
		default:
			return // other error
		}
	}
	defer rez.Close()

	err = safeFrob(rez, input)
	return
}

func safeFrob(rez Resource, input string) (err error) {
	defer func() {
		r := recover()
		frobErr, ok := r.(FrobError)
		if ok {
			rez.Defrob(frobErr.defrobTag)
		}
		err, _ = r.(error) // maybe nil which is good
	}()
	rez.Frob(input)
	return
}

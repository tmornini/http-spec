package main

func getRequestFromFile(s *state) (*request, error) {
	message, err := getMessageFromFile(s)

	if err != nil {
		return nil, err
	}

	return &request{message}, nil
}

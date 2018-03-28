package main

func getResponseFromFile(state *state) (*response, error) {
	message, err := getMessageFromFile(state)

	if err != nil {
		return nil, err
	}

	return &response{message}, nil
}

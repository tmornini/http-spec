package main

import (
	"errors"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

func loadMatchers(context *context) error {
	// load default matchers
	context.Matchers[":date"] = "(Mon|Tue|Wed|Thu|Fri|Sat|Sun), (0\\d|1\\d|2\\d|3[01]) (Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec) 2\\d{3} (0\\d|1\\d|2[0-3]):(0\\d|1\\d|2\\d|3\\d|4\\d|5\\d):(0\\d|1\\d|2\\d|3\\d|4\\d|5\\d) (A|M|N|Y|Z|UT|GMT|[A-Z]{3}|[+-](0\\d|1[012]))"
	context.Matchers[":b62:22"] = "[0-9A-Za-z]{22}"
	context.Matchers[":iso8601:µs:z"] = "\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}[.]\\d{6}Z"
	context.Matchers[":uuid"] = "[[:xdigit:]]{8}-[[:xdigit:]]{4}-[[:xdigit:]]{4}-[[:xdigit:]]{4}-[[:xdigit:]]{12}"

	// load custom matchers if provided
	var ymlData map[string]string

	ymlFile, err := ioutil.ReadFile("/matchers.yml")
	// ignore error (perhaps they didn't provide a custom matcher file)
	if err == nil {
		err = yaml.Unmarshal(ymlFile, &ymlData)

		if err != nil {
			return err
		}

		for key, value := range ymlData {
			trueKey := ":" + key

			_, ok := context.Matchers[trueKey]

			if ok {
				return errors.New(key + " is a default matcher and cannot be used")
			}

			context.Matchers[trueKey] = value
		}
	}

	return nil
}

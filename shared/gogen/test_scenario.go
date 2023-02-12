package gogen

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestScenario[REQUEST, RESPONSE, OUTPORT any] struct {
	Name           string
	InportRequest  REQUEST
	InportResponse *RESPONSE
	Outport        OUTPORT
	ExpectedError  error
}

func RunTestcaseScenarios[REQUEST, RESPONSE, OUTPORT any](t *testing.T, f func(o OUTPORT) Inport[REQUEST, RESPONSE], scenarioList ...TestScenario[REQUEST, RESPONSE, OUTPORT]) {

	t.Parallel()

	for _, tt := range scenarioList {

		t.Run(tt.Name, func(t *testing.T) {

			res, err := f(tt.Outport).Execute(context.Background(), tt.InportRequest)

			if err != nil {
				assert.Equal(t, tt.ExpectedError, err, "Testcase name %s", tt.Name)
				return
			}

			assert.Equal(t, tt.InportResponse, res, "Testcase name %s", tt.Name)

		})

	}

}

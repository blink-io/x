package failsafe

import "github.com/failsafe-go/failsafe-go"

var (
	ErrExecutionCanceled = failsafe.ErrExecutionCanceled

	Run              = failsafe.Run
	RunWithExecution = failsafe.RunWithExecution
)

package incorrect_replay

import (
	"context"
	"time"

	"go.temporal.io/sdk/workflow"

	// TODO(cretz): Remove when tagged
	_ "go.temporal.io/sdk/contrib/tools/workflowcheck/determinism"
)

func TestWorkflow(ctx workflow.Context) (string, error) {

	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute * 5,
	})

	var res1 string
	if err := workflow.ExecuteActivity(ctx, ActivityC, "qwe").Get(ctx, &res1); err != nil {
		return "", err
	}

	var res2 string
	if err := workflow.ExecuteActivity(ctx, ActivityB, &res1).Get(ctx, &res2); err != nil {
		return "", err
	}

	var res3 string
	if err := workflow.ExecuteActivity(ctx, ActivityD, res2).Get(ctx, &res3); err != nil {
		return "", err
	}

	return "Completed Successfully", nil
}

func ActivityA(ctx context.Context, name string) (string, error) {
	return "Hello A " + name + "!", nil
}

func ActivityB(ctx context.Context, name string) (string, error) {
	return "Hello B " + name + "!", nil
}

func ActivityC(ctx context.Context, name string) (string, error) {
	return "Hello B " + name + "!", nil
}

func ActivityD(ctx context.Context, name string) (string, error) {
	return "Hello D " + name + "!", nil
}

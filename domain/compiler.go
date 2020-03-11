package domain

import (
	"github.com/visket-lang/playground/infra"
	"github.com/visket-lang/playground/model"
)

func CompileProgram(job model.Job) model.Result {
	filename, err := infra.SaveProgram(job.Program)
	if err != nil {
		return model.Result{
			Error: err.Error(),
		}
	}

	binname, err := infra.CompileCode(filename)
	if err != nil {
		return model.Result{
			CompilerError: err.Error(),
		}
	}

	output, err := infra.RunCode(binname, job.Stdin)

	if err != nil {
		// TODO fix
		if err.Error() == "signal: killed" {
			return model.Result{
				CompilerError: "Timeout",
			}
		}

		return model.Result{
			Error: err.Error(),
		}
	}

	return model.Result{
		Stdout: output,
	}
}

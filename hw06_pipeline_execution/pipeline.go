package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	out := in

	for _, stage := range stages {
		if stage != nil {
			out = execute(stage, done, out)
		}
	}

	return out
}

func execute(stage Stage, done In, out In) Out {
	result := make(Bi)

	go func() {
		defer close(result)
		for {
			select {
			case <-done:
				return
			case v, ok := <-out:
				if !ok {
					return
				}
				result <- v
			}
		}
	}()

	return stage(result)
}

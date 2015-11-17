package worker

import (
	"github.com/qor/qor/admin"
	"github.com/qor/qor/roles"
)

type Job struct {
	Name       string
	Handler    func(record interface{}) error
	Resource   admin.Resource
	Permission roles.Permission
	Queue      Queue
	Worker     *Worker
}

func (job *Job) Run(argument interface{}) error {
	return job.Handler(argument)
}

func (job *Job) GetQueue() Queue {
	if job.Queue != nil {
		return job.Queue
	}
	return job.Worker.Queue
}

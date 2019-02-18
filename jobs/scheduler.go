package jobs

import (
	"time"

	"github.com/wayn3h0/gop/jobs/expression"
)

// Scheduler represents a job scheduler.
type Scheduler struct {
	jobs    jobs
	running bool
	add     chan *job
	stop    chan bool
}

func (s *Scheduler) run() {
	go func() {
		now := time.Now()
		for _, job := range s.jobs {
			job.Next = job.Expression.Next(now)
		}

		for {
			s.jobs.Sort()

			var effective time.Time
			if len(s.jobs) == 0 || s.jobs[0].Next.IsZero() {
				effective = now.AddDate(10, 0, 0)
			} else {
				effective = s.jobs[0].Next
			}

			select {
			case now = <-time.After(effective.Sub(now)):
				for _, job := range s.jobs {
					if job.Next != effective {
						break
					}

					go job.Function()

					job.Previous = job.Next
					job.Next = job.Expression.Next(effective)
				}
				continue

			case job := <-s.add:
				s.jobs = append(s.jobs, job)
				job.Next = job.Expression.Next(now)

			case <-s.stop:
				return
			}

			now = time.Now()
		}
	}()
}

// Schedule adds a job to the scheduler.
func (s *Scheduler) Schedule(fn func(), expr expression.Expression) {
	if fn != nil {
		job := &job{
			Function:   fn,
			Expression: expr,
		}

		if !s.running {
			s.jobs = append(s.jobs, job)
		} else {
			s.add <- job
		}
	}
}

// Start starts the scheduler for scheduling tasks.
func (s *Scheduler) Start() {
	s.running = true
	s.run()
}

// Stop stops scheduler.
func (s *Scheduler) Stop() {
	s.stop <- true
	s.running = false
}

// NewScheduler returns a new scheduler.
func NewScheduler() *Scheduler {
	return &Scheduler{
		jobs:    nil,
		running: false,
		add:     make(chan *job),
		stop:    make(chan bool),
	}
}

var (
	// DefaultScheduler represents the default jobs scheduler.
	DefaultScheduler = NewScheduler()
)

// Schedule adds a job to default scheduler.
// This is short for DefaultScheduler.Schedule.
func Schedule(fn func(), expr expression.Expression) {
	DefaultScheduler.Schedule(fn, expr)
}

// Start starts the default scheduler.
// This is short for DefaultScheduler.Start.
func Start() {
	DefaultScheduler.Start()
}

// Stop stops the default scheduler.
// This is short for DefaultScheduler.Stop.
func Stop() {
	DefaultScheduler.Stop()
}

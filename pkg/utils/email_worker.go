package utils

import (
	"fmt"
	"sync"
)

type EmailJob struct {
	Email string
	Otp   string
}

// Worker pool
func StartEmailWorkers(workerCount int, jobs <-chan EmailJob, stop <-chan struct{}, metrics *Metrics, wg *sync.WaitGroup, config SMTPConfig) {
	wg.Add(workerCount)

	for i := 1; i <= workerCount; i++ {
		go func(id int) {
			defer wg.Done()

			for {
				select {
				case job, ok := <-jobs:
					if !ok {
						fmt.Println("worker", id, "jobs channel closed")
						return
					}

					// send email
					if err := SendOTPEmail(job.Email, job.Otp, config); err != nil {
						fmt.Println("email failed:", err)
						metrics.Failed()
						continue
					}

					metrics.Sent()

				case <-stop:
					fmt.Println("worker", id, "received stop signal")
					return
				}
			}
		}(i)
	}
}

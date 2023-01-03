package notifier

type Notifier interface {
	Send(isWorker bool) error
	Get() string
}

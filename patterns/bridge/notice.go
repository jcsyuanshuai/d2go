package bridge

type MsgSender interface {
	Send(msg string) error
}

type EmailSender struct {
	emails []string
}

func NewEmailSender(emails []string) *EmailSender {
	return &EmailSender{emails: emails}
}

func (s *EmailSender) Send(msg string) error {
	return nil
}

type WechatSender struct {
	names []string
}

func NewWechatSender(names []string) *WechatSender {
	return &WechatSender{names: names}
}

func (s *WechatSender) Send(msg string) error {
	return nil
}

type Notification interface {
	Notify(msg string) error
}

type WarnNotification struct {
	sender MsgSender
}

func NewWarnNotification(sender MsgSender) *WarnNotification {
	return &WarnNotification{sender: sender}
}

func (n *WarnNotification) Notify(msg string) error {
	return n.sender.Send(msg)
}

type FailNotification struct {
	sender MsgSender
}

func NewFailNotification(sender MsgSender) *FailNotification {
	return &FailNotification{sender: sender}
}

func (n *FailNotification) Notify(msg string) error {
	return n.sender.Send(msg)
}

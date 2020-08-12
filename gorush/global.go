package gorush

import (
	"github.com/appleboy/gorush/config"
	"github.com/appleboy/gorush/storage"

	"github.com/egorovandreyrm/go-fcm"
	"github.com/sideshow/apns2"
	"github.com/sirupsen/logrus"
)

var (
	// PushConf is gorush config
	PushConf config.ConfYaml
	// QueueNotification is chan type
	QueueNotification chan PushNotification
	// ApnsClient is apns client
	ApnsClient *apns2.Client
	// FCMClient is apns client
	FCMClient *fcm.Client
	// LogAccess is log server request log
	LogAccess *logrus.Logger
	// LogError is log server error log
	LogError *logrus.Logger
	// StatStorage implements the storage interface
	StatStorage storage.Storage
	// MaxConcurrentIOSPushes pool to limit the number of concurrent iOS pushes
	MaxConcurrentIOSPushes chan struct{}
)

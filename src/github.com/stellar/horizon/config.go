package horizon

import (
	throttled "gopkg.in/throttled/throttled.v2"
	"github.com/Sirupsen/logrus"
)

// Config is the configuration for horizon.  It get's populated by the
// app's main function and is provided to NewApp.
type Config struct {
	DatabaseUrl            string
	StellarCoreDatabaseUrl string
	StellarCoreUrl         string
	RubyHorizonUrl         string
	Port                   int
	Autopump               bool
	RateLimit              throttled.RateQuota
	RedisUrl               string
	LogLevel               logrus.Level
	SentryDSN              string
	LogglyHost             string
	LogglyToken            string
	FriendbotSecret        string
}

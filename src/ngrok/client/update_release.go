// +build release autoupdate

package client

import (
	"ngrok/client/mvc"
	"ngrok/log"
	"ngrok/version"
	"time"

	"gopkg.in/inconshreveable/go-update.v0"
	"gopkg.in/inconshreveable/go-update.v0/check"
)

const (
	appId          = "ap_pJSFC5wQYkAyI0FIVwKYs9h1hW"
	updateEndpoint = "https://api.equinox.io/1/Updates"
)

const publicKey = `-----BEGIN CERTIFICATE-----
MIIFJzCCBA+gAwIBAgISA7ERPGbA+0PkmGrXOK+h6SVJMA0GCSqGSIb3DQEBCwUA
MDIxCzAJBgNVBAYTAlVTMRYwFAYDVQQKEw1MZXQncyBFbmNyeXB0MQswCQYDVQQD
EwJSMzAeFw0yMTA0MDMxMTI1NDBaFw0yMTA3MDIxMTI1NDBaMBwxGjAYBgNVBAMT
EXd3dy5zcHJvZ3JvdXAuY29tMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKC
AQEAmfdAvh3m66og2SkvkHVZvMQL26xtJlBAduBpLm0TKE/AnFvGLRZtaJQcUjeF
0ODRgNNjft9zvlUoyXDjK5sNcUhKswO0Fc+a4vE2xx5+AcbUV6qg4Fh76uUjjJVs
zXNnp+KAVXDeXnpt8heAaCP0jxLvj053kMM/r2Nce7x46XSuyw88p5Y+Zk85uPtC
WfqAN/kbtzHi0XJhe7Tu1bMSeNFK9ZwteRm3g+mDdB+zOUPyKq9fAkFkFwHGS++C
SQjB9x9KYeC/hNo1frlrMhp3f5gYUm69g3FnQvFOi96jGDLbXVXmN27wBgVkEWG0
tfmuEU2mcnNqLCC4yBsYXxCjpQIDAQABo4ICSzCCAkcwDgYDVR0PAQH/BAQDAgWg
MB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAMBgNVHRMBAf8EAjAAMB0G
A1UdDgQWBBQFR55RxBYk9cqwSUVEIuQQ4EK9sDAfBgNVHSMEGDAWgBQULrMXt1hW
y65QCUDmH6+dixTCxjBVBggrBgEFBQcBAQRJMEcwIQYIKwYBBQUHMAGGFWh0dHA6
Ly9yMy5vLmxlbmNyLm9yZzAiBggrBgEFBQcwAoYWaHR0cDovL3IzLmkubGVuY3Iu
b3JnLzAcBgNVHREEFTATghF3d3cuc3Byb2dyb3VwLmNvbTBMBgNVHSAERTBDMAgG
BmeBDAECATA3BgsrBgEEAYLfEwEBATAoMCYGCCsGAQUFBwIBFhpodHRwOi8vY3Bz
LmxldHNlbmNyeXB0Lm9yZzCCAQMGCisGAQQB1nkCBAIEgfQEgfEA7wB2AFzcQ5L+
5qtFRLFemtRW5hA3+9X6R9yhc5SyXub2xw7KAAABeJeyLi0AAAQDAEcwRQIgWcz7
CYmVfGWaNcGRgXmR9z6Lf2QN4W84OyduipbuJ+UCIQCYE68wpu0z0ZfmRg3IrUUR
MiogUs4tAVvbNmzsB5YZVQB1AH0+8viP/4hVaCTCwMqeUol5K8UOeAl/LmqXaJl+
IvDXAAABeJeyLngAAAQDAEYwRAIgSBFC+gkDAX3Ym2zzcSWv38loJCnRCCTpUFZX
1rCnDQECIBYNlpv4D1CkUrgGJyMdeUF2bfEL69N2w0fup71NESiWMA0GCSqGSIb3
DQEBCwUAA4IBAQCpc4Ppqe7ZTaPAXDX0ZyJLoIhKeJZxz4ca3x+fZXK2cT78OFO5
r3mz4n+2/xpsgArGyyUfEFYhMH0kDOpSGQsWl0LAqP1aJT3hllBaSKrneQGyI44+
kYgCOEjDry+P0XLfhOz2aU0KZyPi1VjCpFiz4hTm7OE/qELnOzDT59CLptg2cAfs
tv29ASf2pBsYVBNCPySwt5uIdN3RIP1tvlK+difg4e8U5hlbkyOallk2OC+cvwUn
ftwErwXobUnOTwWkU4Txdxg9z1guOJ/+kn7mTUXmtPVXEJ7cBmQqpGTTpqbnSFoE
qCivjI3KPu52vwJGNNNbLZJs/6tCmKMp2S2M
-----END CERTIFICATE-----`

func autoUpdate(s mvc.State, token string) {
	up, err := update.New().VerifySignatureWithPEM([]byte(publicKey))
	if err != nil {
		log.Error("Failed to create update with signature: %v", err)
		return
	}

	update := func() (tryAgain bool) {
		log.Info("Checking for update")
		params := check.Params{
			AppId:      appId,
			AppVersion: version.MajorMinor(),
			UserId:     token,
		}

		result, err := params.CheckForUpdate(updateEndpoint, up)
		if err == check.NoUpdateAvailable {
			log.Info("No update available")
			return true
		} else if err != nil {
			log.Error("Error while checking for update: %v", err)
			return true
		}

		if result.Initiative == check.INITIATIVE_AUTO {
			if err := up.CanUpdate(); err != nil {
				log.Error("Can't update: insufficient permissions: %v", err)
				// tell the user to update manually
				s.SetUpdateStatus(mvc.UpdateAvailable)
			} else {
				applyUpdate(s, result)
			}
		} else if result.Initiative == check.INITIATIVE_MANUAL {
			// this is the way the server tells us to update manually
			log.Info("Server wants us to update manually")
			s.SetUpdateStatus(mvc.UpdateAvailable)
		} else {
			log.Info("Update available, but ignoring")
		}

		// stop trying after a single download attempt
		// XXX: improve this so the we can:
		// 1. safely update multiple times
		// 2. only retry after temporary errors
		return false
	}

	// try to update immediately and then at a set interval
	for {
		if tryAgain := update(); !tryAgain {
			break
		}

		time.Sleep(updateCheckInterval)
	}
}

func applyUpdate(s mvc.State, result *check.Result) {
	err, errRecover := result.Update()
	if err == nil {
		log.Info("Update ready!")
		s.SetUpdateStatus(mvc.UpdateReady)
		return
	}

	log.Error("Error while updating ngrok: %v", err)
	if errRecover != nil {
		log.Error("Error while recovering from failed ngrok update, your binary may be missing: %v", errRecover.Error())
	}

	// tell the user to update manually
	s.SetUpdateStatus(mvc.UpdateAvailable)
}

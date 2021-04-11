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

const publicKey = `-----BEGIN PUBLIC KEY-----
MIIFFDCCAvwCCQCkbN0RG/o15DANBgkqhkiG9w0BAQUFADBMMQswCQYDVQQGEwJV
UzETMBEGA1UECBMKQ2FsaWZvcm5pYTESMBAGA1UEChMJbmdyb2suY29tMRQwEgYD
VQQDFAsqLm5ncm9rLmNvbTAeFw0xMzA2MDMwMzUxNTZaFw0yMzA2MDEwMzUxNTZa
MEwxCzAJBgNVBAYTAlVTMRMwEQYDVQQIEwpDYWxpZm9ybmlhMRIwEAYDVQQKEwlu
Z3Jvay5jb20xFDASBgNVBAMUCyoubmdyb2suY29tMIICIjANBgkqhkiG9w0BAQEF
AAOCAg8AMIICCgKCAgEA6QryXeKl8AWWa9uG2UbSOpooH74zLkXs3FZfk9gKvqki
zXXQCRmtU6Dyn0+OuS3sE/rRmZAsSjkQG/YDtdE/SgL4dV6S62qiQngPokjR0USh
PC4Hwb8TjM9W5Cd+owVzMQ0vl0AYhQk8Yc/0vX+zDOwmRWGjNKPq422usF9CJFc/
8QY+ODJDHun8VVAkq3XfcPXgytHIqxvSJnYgDouFCA+GTsKp/65S5cigSlIrQZbH
775cTWhCjvYnq6gzyrk3RiGdb1IGuIJftMJxuJyJVbfTFtqgMGTmjHZxiLvM7dz7
j/bmrz4PvnhbQSZZLhsvP1o8mxnoNMpo/To5tHp/Ts6b5FQNL7FHpmOVLAoQ3FdX
VryTjoSjiE3JLDGZINQ9MFEPgPzR8mrzqFo/6e7uB4AYlKoQo01Kzx+YmVwRKEtr
VCTRZRcl66+gMkcX0ryoVnggjIWWu4d8uAh3jf++Kd0Djb/l7DCPpEgSJwZYTjCL
Z6SxiBwQ4o9dEQadMGgk3tlDFCBsrHoq7NyzvXHP0BF2HKb8KREBEKCIuQj9RChW
g07zmOpjngWs0CXaYly+TDP+5DZCMGD9kmXkQY9q/zqqvMt+T+/TBK9lwUsoi2Uc
v93wS+TNu06aRopqPo9YZr38ka3xKPiO964pk2BoFN57g767G8k9TbhkBxitvFUC
AwEAATANBgkqhkiG9w0BAQUFAAOCAgEARf/hVYntzUwFUgQrWD0l/UaCBgrlvxVC
yUa8Isj3vezAhFSyZntEL+ELFv8vvQbtBGHH/tCn76WuqjwOjVL23yxkaJsrnR9+
TRNFnVeB8157+IF6HKzLCL/HIAiQ0kw/2OSLD0lZnAkg24A0/9SHcpI5GA0WlThE
4GqgcUiN9m+mL8jWG3gj+SXC7IcVS3vAvS1J7Kz0NzTh1dYkQNWJlauO2Qn95T99
plkPPh87yZO9a9bxpX9PUJkTJzOwUkZISRZEbTA0CfspUpq/phzuTViy7o2fr+To
xVVa2aKT912vlQadWw5oqEK7xyxTqsYwV7CUljtCnhpS7wOZhwzI3qUk9HKH0Rt9
/HQsANuSikZosNvdM3/hv3c5DRUOwKiKdbgZCyqQf8XSeJRM1iQvcHo8U8kEJEEt
dmftn+0gH3RPsV028+7XD6wraqfc7dzNzLnq2rDLSAfG3T7pp0n19JyUieWsTR45
pGcwNpXDk+weqQKknwbNoha3o51Xq/1nhRtqrUPIEgOnBWBCV8vkhXMh9Vbe+oAm
LmweN1Mr6MjrHWddVn+JGcB5p+AMWr8zE+bhpPCAnupFR98z3fbOleMCWq6Q+hMP
MUThqJROHamHRIJ3Iz8whIrj7EKkDBKLEfE1ExS3B9VQ+YdHy6sjexdmCIQgElq4
4SMuY/JkZlo=
-----END PUBLIC KEY-----`

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

package crypter

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

var testsign string = "ZFKrbOKa5V3adcSsb4VIjcE0knRPipNzXReUEc/i2rJZ4erUMQ6wqBLxV+LsMO+fX7jOpn/3qbSJdBbdEWRmquvQzWmFiVYRhsQecKwM7qu7dB8pSmmryD1BefcHGk8+w7l692LZHSyu/fuAvN9T4weohg+/e+Pcop3VXXaUQRPcYQjZpoo6ZMUD89y4ziPJa2iihBeC83Nr8tJl/D7iNPr5P9xxzvrVJZFBzwK6iRL9v2aA4buPdOAfLfEMZIhEWb3/dmd1CoXZfDJvtrtzYKwece8s0wAZU+ZmWhbr6l9kdqfPPDpD8QLsAKsxFRbKjVmAAR6Am4fKYc9/Nkc3mm8Yli1c6DFGlpZ7J2GtKi3/bhw19PUErRDlScg6tMwQXakafh14uRSgb+TkAtZQNQJScQLu3CgpPm5QhDI2KeY4ElvO+d8RjunGv0yv6mc8zwFEWNsql80aYplaQ9sog3FcFAWJ6s1CfBlyt1cXWm8vLVw84lU8L17BZHtux8HJBo48Uy3JkaszuvrpvSHQwQvJr/27T01zocE9L0AsjYSFFMCDuWMVQeSPvk9vLbn2C4SIRR52wPQLC3geV6RAFmAEnmonaXmP86ic8DlpQUg9wKxqyIGVSnn/NzyFBHHBWQ+LyKoD2I0ecg9EbmLR19NPV2fVCoat2kNoUmJBB7MmQfraUANGII5YzB1gfuEC2uiQbajAmOpv0bfTtqPaOqgvg/dRr7IhMwubfkdQDEh+Z2MvM/5n3+S97Ev0gFyJGZBGiPN0V237xbQuQVnfMI+VgLnbXDvLJHWt8Tut1Fw+PWJ+7ks2ZW4aswaK+CEV48DPty0RJ9qC9o4V3vrvZZKtWM2eQV4ONREbR1SeloNifzrJYbcNEsBVKRn5sbNVEeLrNzKoIde6IxLrmDR0F5QbLhe2gCKANB8reN+XAhZmq4Hev3+C0Ss6He4DGak5W/bQsfquJlh6OP9kzVAGhG8pBNg0OwpWv0BcQ70L0iW1+zeR/kPhE3REFjjAYHK2TuTOIpq6yZEG9pyiSd8ssgrx8lpxJjsB2bDmvQuGPNjGjNwWF89axnh7X4o+JvlDUAoem8V9j9H2+g1TAOO6uSf4UNa2g/IbkV6EkHc9o/UN/KA6U2rHJ4O7HgY+ZccxQkNbmoiks17bA7myY42LFUEog4VzyJT1cqtEGlrWXdjVAO/sUZSrteaHGbm0tOhNypZVh/cI3jn9YvKdMpsSRWQrgiycZRAuXJEJ65zXp0HJ+hNw6KXcXKnHlncLyWbljHZWhn6rodLSrdjfAuo/bwWuMi+N7dhKxeEcC39Bv3ZgjwBaY7ubHudBArOj12CCMzCGA76C0kGp1Itjk3AG9A"
var testmes string = "123"

func TestSign(t *testing.T) {
	godotenv.Load("../.env")
	private, _ := os.LookupEnv("MARKET_PRIVATEKEY")
	crypter, _ := Get(private)
	sign, err := crypter.Sign(testmes)
	if err != nil {
		t.Error(err)
	}
	if sign != testsign {
		t.Error("incorrect signature: ", sign)
	}
}

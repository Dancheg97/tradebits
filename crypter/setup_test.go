package crypter

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

var testpub string = "MIIECgKCBAEApAqr6+MbKqDAAiU+8XTu55krjllICL1tPS8TheV8xLiWetxCE2t+mYx9qyl7a2C9Cy5y4Eu/t/v9WHQ2p0T2U6Jt+wbgJQtsOoXuBXU8LKmOwMVxNJLgqCBiyn+YPWPDqEENSO8E+wawolKfvg+jfryvcTWztmMw609VCrghhpVz0ezPlmhS5mBK/AGSOMV/v4Zph/F30+TLLFa0LjhSZHH5WDmaDXybZlDr+lGOfjg3agG1fbOIAdkcuHIfI/24vF4jovgOnYLZGbOO0792gBVHqEps/b7gH5PMoZi7AJCFW66Y0NaXgM+oF1ifp6CNQZfscypSfVJMGYf66ZB4h6VFTQlXQGUJaHzCorMa9HKBszuMm6jNKqK62PinzSMVtZDvRB9/Fd6lxgBzdSO3entnYMshL3ez5zsQAesUkPgrbDyz9lxYt/5rOMEkgG2f/9HSVOsHV5scrOT08HXdpa9HfqeJWc5iK/Kny9+0cazRyXNM2JB83XbvrG4xPsvO/00mtQEnPCAMRkAIf93Fexd1V5ZH0hpXMCY2PoN6B0UkKerRX4cp3ldeWaJyBdfJHXi1uTLqFs35pnOiSfyDFrCeKlmRpCCHIicl/qxatdAsnbMl73p7bVyY7cQ6clCS4SfEsICNae2R/QtN8tp4gHO5ZaZmOSb+4qFA+IzCS5XSoDTgNqbPG/a8OO626bwkwqEMQtEC5mm+ckDtaqiD9iGXLFzTM4dEW0j8ba8MCn/FGR9i5Nfaf56kFDnxW26w+fVQti+zzEIowmMZq42Ifm6dNASXo2d1cLN/QiXLqZMEQ1xL6ATzHkXx4LjrRcCSJm8hKYipQ6ag/4XQEE2nCvnrzqzd//G69bR/MUzu7N7WCaNOmuTor7ve3YdaphCiQ2dZGq1lIEFMzSvtPfbxnjBVOGY/TEMh8t01W/CYLqbv8/OrXSC4v5ILyoXlR1IWuXGEdiH5Ihd0rxr8bONO1JMGCuAcxkTbZs9ONbnPzsJkdRCaUcrK4SYCP3wzvlKraS8yIbB5HEUKvF3XIbQ0y9qQa4qtBiVplieoLXRtBkDmk9Gqi/W5PwJwirBQWsra6bYog/EaG2uuy98z3myGLxn6rbkr1hCB3sCloFYKsXLWC+ezh7wjDHkn+obIOPRrseI9OSAxFKdYimyisAe68Tueg2UJLRArsT2/V6VcSoeeN198BQiBCB3dFKcp0BMdqi0GYoeR10g8hYvPd1qJ8gAq0VXC1O8PblVwSwNWGNTzvbw299anKHxbLURSPtViu0u0YHunwJZbs0UWE/OaZhp67Nhw/VN/FBp0w7xoK7RHW8qwC7hb62CZpwd8541agOf3TJVgKLFP8oXJ6ESPDwIDAQAB"

func TestSetup(t *testing.T) {
	godotenv.Load("../.env")
	private, _ := os.LookupEnv("MARKET_PRIVATEKEY")
	crypter, err := Get(private)
	if err != nil {
		t.Error(err)
	}
	if crypter.priv == nil {
		t.Error("private key should not be nil")
	}
	if crypter.pub != testpub {
		t.Error("Public key is not matching with actual")
	}
}

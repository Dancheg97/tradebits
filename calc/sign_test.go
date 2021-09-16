package calc

import (
	"encoding/base64"
	"testing"
)

func TestSign(t *testing.T) {
	keyBase64 := `MIIEpAIBAAKCAQEAix1rNd69NMH2URBBgFpu02VQ3pBQ8NQkztfIwNE5fk1Lhu5kl+Vz0egN9Bn+AsGcDGs4Q8A5HHejvik1AC1Xo3ZdpzrYYGW2bxzCy61cd9ZYALfRkmjfmUbbEpqd/APJq94Xvl785figUMQqTqZCakNXVRabC1sN8vkvPVF4wwk/wTj73+eDA//L4t/5Vo9tQscKSAtc+2uu8Rt9jTcMUQQCF3ek2O0hTbWQJkGgR8rGFqzWGl1depZGK7rpUtw0cnAT2aY9WK3Ohn0nRGoNaeF7Imil9oHeyWmkRYMoOANheW9XDnT3G2Ni51lX2evAYjOntZBBQsVD2EQE/xWmTQIDAQABAoIBAQCB/u6lO0wBFdMPsyoSP+9aGU31fFQ2h3cR1nChRaH9FUujffeAPYj6OGd5mCRV/QfOTw8XuHNl16KqWyVM0hIYDM69K27wvXs5311kKM7EqQY5PYxXzSpMt/qVKujAi+WTdgkn+ulGGQvuDTEfp5teYykjuOkPNnoVd6ceHJjFXufzeZuEnqTbYUcFjNcv0pw1SbNVVHvTvFGqYFT5Vk0QNOJCSBtBcvrZJTiVPsxX82aXIyNkgQSgbSSEFIL0okqA2d45dpvmCJnafx3sq7ew8JtVZxJDIhcxOXT+Y7epR7HdmRdAIBtgVsjFbg8UCEKjHH/AQ7WraKTbByr7HyVBAoGBAM9Xl8JnRaPW69xTRQUD3vd56Q+SiTOI9DuqqrylV5wj+HyUiidL66p9LFpskUGv8yXrB52Fl1ZpyXaRv9BlPtsLlTf6Wf+JlhCyUdxGRe0zgSSDUJQi62IpuJHpZ9yxBePTdsd1de4Ux8Zzq+IQc8JI3OLFi29X3Evu/KMcVBJdAoGBAKvC94EEyAoMVXEud1tanJQzQxiOOUAGBNcEIk99YpUFuNbbiLQTOQjXWXO/0din/MCCr9DM+HPv7pqdz08/1YbE2olkJ2D9tmXEpev05nIEK7mjUYpeyYeqpTHhEmMFI5N40hqJ1ln0+IyYb1G4RAiz3jiYtHbThT14o49foYSxAoGASYQQPhCp44a/yDYKIdYWh37FQQ0+8nPbzRAdS5gAcU9buqhcN8S2gZOyVzlpY3YEa2xjHdoRJ8WIdi8FaNmNKNN5jAiLQozBFZh0Z/rOrmT1pEGfH3cmKH8Mf0fX2Ks3oGkJE321W74oe1fmGNimgruc/hb28nqflJjfIOJZeM0CgYEAlQ3/zYYBtr6XUz4YtD2BDhciTef2hLn0b0eWItdTurEZg9NF4cegwixn40v+Zn75l3KR8qg9oBcCQElcw5BI08h3Qa2ZgtEJ/WzW9Fo1tZIrA9FYDLVPPJ14+D0ywJ2QDwVVh1RAoRY2r4QFbqxy1f6JhSH30fqjcQIFJoG94DECgYADNNax1vfsCiyot3lIHsj3fDZxwDJdoNY9wzWrERAdKJG1MbA594AF/VVeRwOailh+3XLADmwHBYcZIW2IHVgnzVzTTV8gcX5mKPbtv+761QrckOzy2orrkrhnlrLvbhMtwMtekGu95g+jkiFrMZJ/0q0EA0eCGaLywcreuWYxQQ`
	keyBytes, _ := base64.RawStdEncoding.DecodeString(keyBase64)
	mes := []byte{1, 2, 3}
	sign, _ := Sign([][]byte{mes, mes}, keyBytes)
	if len(sign) == 256 {
		return
	}
	t.Error("signature error")
}

func TestBadKeySign(t *testing.T) {
	badKey := []byte{1, 2, 3, 4}
	mes := [][]byte{badKey, badKey}
	_, err := Sign(mes, badKey)
	if err == nil {
		t.Error("should get an error cuz key is invalid")
	}
}

package calc

import (
	"crypto/x509"
	"encoding/base64"
	"reflect"
	"testing"
)

func TestHash(t *testing.T) {
	hash := Hash([]byte{0})
	expected := []byte{184, 36, 77, 2, 137, 129, 214, 147, 175, 123, 69, 106, 248, 239, 164, 202, 214, 61, 40, 46, 25, 255, 20, 148, 44, 36, 110, 80, 217, 53, 29, 34, 112, 74, 128, 42, 113, 195, 88, 11, 99, 112, 222, 76, 235, 41, 60, 50, 74, 132, 35, 52, 37, 87, 212, 229, 195, 132, 56, 240, 227, 105, 16, 238}
	if reflect.DeepEqual(hash, expected) {
		return
	}
	t.Error("taking blake2b hash")
}

func TestGetKeys(t *testing.T) {
	keys := Gen()
	lenSum := len(keys.PersPriv) + len(keys.PersPub) + len(keys.MesPriv) + len(keys.MesPub)
	if lenSum < 4330 || lenSum > 4340 {
		t.Error("failed to generate correct keys")
	}
}

func TestNumToBytes(t *testing.T) {
	number := uint64(1823879123)
	bytes := NumberToBytes(number)
	if len(bytes) != 8 {
		t.Error("byte length of the number should be 8")
	}
}

func TestGenerateRandomBytes(t *testing.T) {
	bytes := Rand()
	if bytes != nil {
		return
	}
	t.Error("failed to generate random bytes")
}

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

func TestVerify(t *testing.T) {
	keyBase64 := `MIIEpAIBAAKCAQEAix1rNd69NMH2URBBgFpu02VQ3pBQ8NQkztfIwNE5fk1Lhu5kl+Vz0egN9Bn+AsGcDGs4Q8A5HHejvik1AC1Xo3ZdpzrYYGW2bxzCy61cd9ZYALfRkmjfmUbbEpqd/APJq94Xvl785figUMQqTqZCakNXVRabC1sN8vkvPVF4wwk/wTj73+eDA//L4t/5Vo9tQscKSAtc+2uu8Rt9jTcMUQQCF3ek2O0hTbWQJkGgR8rGFqzWGl1depZGK7rpUtw0cnAT2aY9WK3Ohn0nRGoNaeF7Imil9oHeyWmkRYMoOANheW9XDnT3G2Ni51lX2evAYjOntZBBQsVD2EQE/xWmTQIDAQABAoIBAQCB/u6lO0wBFdMPsyoSP+9aGU31fFQ2h3cR1nChRaH9FUujffeAPYj6OGd5mCRV/QfOTw8XuHNl16KqWyVM0hIYDM69K27wvXs5311kKM7EqQY5PYxXzSpMt/qVKujAi+WTdgkn+ulGGQvuDTEfp5teYykjuOkPNnoVd6ceHJjFXufzeZuEnqTbYUcFjNcv0pw1SbNVVHvTvFGqYFT5Vk0QNOJCSBtBcvrZJTiVPsxX82aXIyNkgQSgbSSEFIL0okqA2d45dpvmCJnafx3sq7ew8JtVZxJDIhcxOXT+Y7epR7HdmRdAIBtgVsjFbg8UCEKjHH/AQ7WraKTbByr7HyVBAoGBAM9Xl8JnRaPW69xTRQUD3vd56Q+SiTOI9DuqqrylV5wj+HyUiidL66p9LFpskUGv8yXrB52Fl1ZpyXaRv9BlPtsLlTf6Wf+JlhCyUdxGRe0zgSSDUJQi62IpuJHpZ9yxBePTdsd1de4Ux8Zzq+IQc8JI3OLFi29X3Evu/KMcVBJdAoGBAKvC94EEyAoMVXEud1tanJQzQxiOOUAGBNcEIk99YpUFuNbbiLQTOQjXWXO/0din/MCCr9DM+HPv7pqdz08/1YbE2olkJ2D9tmXEpev05nIEK7mjUYpeyYeqpTHhEmMFI5N40hqJ1ln0+IyYb1G4RAiz3jiYtHbThT14o49foYSxAoGASYQQPhCp44a/yDYKIdYWh37FQQ0+8nPbzRAdS5gAcU9buqhcN8S2gZOyVzlpY3YEa2xjHdoRJ8WIdi8FaNmNKNN5jAiLQozBFZh0Z/rOrmT1pEGfH3cmKH8Mf0fX2Ks3oGkJE321W74oe1fmGNimgruc/hb28nqflJjfIOJZeM0CgYEAlQ3/zYYBtr6XUz4YtD2BDhciTef2hLn0b0eWItdTurEZg9NF4cegwixn40v+Zn75l3KR8qg9oBcCQElcw5BI08h3Qa2ZgtEJ/WzW9Fo1tZIrA9FYDLVPPJ14+D0ywJ2QDwVVh1RAoRY2r4QFbqxy1f6JhSH30fqjcQIFJoG94DECgYADNNax1vfsCiyot3lIHsj3fDZxwDJdoNY9wzWrERAdKJG1MbA594AF/VVeRwOailh+3XLADmwHBYcZIW2IHVgnzVzTTV8gcX5mKPbtv+761QrckOzy2orrkrhnlrLvbhMtwMtekGu95g+jkiFrMZJ/0q0EA0eCGaLywcreuWYxQQ`
	keyBytes, _ := base64.RawStdEncoding.DecodeString(keyBase64)
	mes := []byte{1, 2, 3}
	sign, _ := Sign([][]byte{mes, mes}, keyBytes)
	priv, _ := x509.ParsePKCS1PrivateKey(keyBytes)
	pubBytes := x509.MarshalPKCS1PublicKey(&priv.PublicKey)
	verified := Verify(
		[][]byte{mes, mes},
		pubBytes,
		sign,
	)
	if verified == nil {
		return
	}
	t.Error("failed to verify sign")
}

func TestBadKeyVerification(t *testing.T) {
	badKey := []byte{1, 2, 3, 4}
	sign := []byte{1, 2, 3, 4}
	mes := [][]byte{badKey, badKey}
	err := Verify(mes, badKey, sign)
	if err == nil {
		t.Error("should be an error here, cuz key bytes are invalid")
	}
}

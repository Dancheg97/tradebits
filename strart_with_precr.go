package main

import (
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"strings"
	"sync_tree/calc"
	"sync_tree/market"
	"sync_tree/user"
)

var nicotinAdr = `lproB8nRtPn4kjmRQtdpongRukRgFjmNe2PXOD7mZw9JxO/by1fdGDbV1vYUX/Hzce8oF9z7rTPYa1No5ZkGwA==`

var alcoholAdr = `EMQqIoItMrkILx5YBJqvHRr3loDQtBIy3AURtlKE9+OLmWGvITE9wFEoMZAH/+3QyMj1iZ8qWJChuomFTSw/wg==`

var alcoKeyString = `-----BEGIN RSA PRIVATE KEY-----
MIIJKgIBAAKCAgEAjdgFBLDvmGQvl3o87UZAxUCuGwKqym92Iv4FGujlcr95oPxk
IoXtLYnZUNXBAPdDRqeag8dhTbftCuywd99xJybV4uoqlOSghd2lu5nrtixrcIF/
x3KH+5brTzLjGaTV5WsK9hLXgnEBdPDJ6FadCo1Se8/6D5oAvKQYoqNNgZqiJBYs
f9rO0bPQqwqSaghWcfGVO/tbJY0OUfhOlhqokeGsvFa3c9g5Ca0kHvpY8u/yF33x
iCl7yadza68yxHu5X5/ciqEOldzsyK5SauyVANGHeIio5X9yVHKiEUeNLkRh3Zd1
QLZvLT6iHyBbhFzN+aI7nWUZx8mqTQPcOoJdUgXpwzsbq2ODgqZf6vLG+vonoW8A
+IrtpyP3t/SdgzwtTt9j8ssK08ma4K+mC5vHdyNt2zd0bT05ZUl2X0hGtcGgkusB
mfXOgV0CD6KynI7E25kNSFm5puCxPiZ5tisyGdrfgBM5lK8tWl5AZPBbPmNNnriU
KOIrC72EsXnxV5F3ZX626t1YUKMpfmkr9H4yicSjHYtZDVInF1VeNOW3eMUUgiw0
kyEZZZLQigUd1IuLkKEvoW9R1gtXKXy9e3xv8v0uhvdLL7UCEA8bDCAmzJ8fOTQL
IfxeAd3bM7JcmnzdglKgGTQU16tIzhizQHOKGM3gNf2W1f3TJyGU8lseA50CAwEA
AQKCAgB6hUqoZ7E9v+RAAs1X4uQVvi3llOMz4x13UlXUjbGL8Yq9HOXd0DMzr6uw
IEMLc4mAIUvdYj5RZHrzgclYcgAKYohD7iMD44wotXDv4B9zh2DymFloTSzai90L
1xfFUdWUymZ4cYQH31bzzP2KnC4gLDxASEBmc48JYvEIouhS99/wLdsYQUd+acTn
yUT5I+SPna5bVl2/zNy+vy4aiEJtaDLNqFXkUbPLAOx/eZ2Vkt9ZnWchRzvuzp5r
OLUwdPKFf89/S2kO1/sLyU/Vni+Cuno7CkhhnDmEalRFermfal/iiW41jJ/xw/VD
z65dYdBlPdYiuW6QXUimJ5x2SZ8S6ZBElZvgdgKxRTp2haYXvcl54Jzs1oRM5d59
lltHI50318qcHeTBTIJ5t6UMuqaVdVSOjSi98pwFmNTSd6JyLeZO2PmbuLXFGRLD
ZSbyXE2b/hSe/LDup/EsnKfKcbxHWEnoADTmb0JPyNL6EyNxaVT6oVrV8i5OMzjj
x4mkfsCkDBZ5qq3+5Q2kJI2Djf9fpIqQDhxGcrai52fTIa1PQD40Y5IO9OGYjyyM
Ib3vWYCI+R8j3afiBEvvFk0wlQJpIPqeQETFIPgr0MOsmod1ATRtzzPlYcfC93mT
ExzgZzerDZMnyo/mf9Zu878CJ0/xbd7GYc4/UR9T9PiC9eOdwQKCAQEA4ieFAqUk
y8jsoOgZ1hjQFGhrDiPCtIHSjdhuGDvGRHhbLyGRtvWComcQyF6OU09GSbsyWsYn
AL+KEgUP95m8iOpeBnjgp4OHFOUiSrBJXBoJPAY6qCZ8hX7Mpe7X6e0+WDbLOP+y
h+UFZ8ryDzLy+2XdhkFArv+hKorZCOAbNc3mYQiUiCQekfYBDDjSn7t2gyxmtIuz
lAuFX04zzB43/81J2i0NpA08asiTHbtmAX7iXaLEMNo7N539pGlAbFwKT7B/AABt
EtbXMJmsL8izv29fG6rzmK0bnUYbnvyWzMJJlecTT5xl7lseYsPct1usU70mr7mB
11SLsp/KTp6wDQKCAQEAoJAfTpjpoWWP257PDcyi//hwP8s1r7SjAht0w55aJd0j
mgjwIgxuDuMkqdT9JqoNek0Mu5Ag6Zsd1HinETvChfswrTvkZ+w6dlAbmi1VcukV
LS8UYXIMJe9MjTUbmM2JnUoWDHNeCiyJttKhUWbcpeGF733wfAyM3AOvlfLp+E2a
+D5TpoLMeFno61yHAJTQrfcSL3o5fW9D4T0hH9G7UQlHyZvz7TSxEtgJSj2VhzJO
5nWQz1RdSG8Kc6F/lmlFLt2vR9ntqfIclAkIEhpv+CQrK/jRav+R0iDn629NQozN
HHCSBlD/Fiu7ikJo1IiA6yNTUO/VTLM7coTg25Yt0QKCAQEAl6fAvmFYJ3/5jn2W
tce14bmT8MtySfG07IHycjXDYy8ClryAJpzEmXtb+cHO13cVv2PVBKJUjV/kiujy
KCvkDjx8CQMOQ8cj0pNFR6inySd8gc/7jb8hSbQaMDvWdnwtjqJP7U0qv46I7W1e
nWqhHozU3K7tXWz2+rxZhmYiWRZWI0PJtvHhviel1DhWpd7kPlqMFQ6JwgFn0t1y
RrMNAnD70V7dAuINXd5bKVj9sosP6ZK41bnCdC1yZcpe9UtTbtUAlVOz0SpY67Jh
kTmZsApTUQ52hC+xkOnuo9LmnTC8f9WRUZWjlQ9tdPW7EQpcqVrykmxn8ao0c/hL
hiVMEQKCAQEAhzVWZ8V02CjCX9r32VL94DfBrrbc2r1ul5OvXBu5JjF2jXwBW2rE
M10rvqhRFYIWH2tK0ZepX470K7v+VBDS2iiQ8RCCYtEW2hLTa53JWx7Ualw++WAk
wr0k5bRXLPHIdN+yeYGdK0weHAp7NiU8oaOsOeT/4jZfV9bywX7xb9tKKMmzf4Cd
1qs/7RLZzK+Emzp34es3Bl3v93iWsEjZSIBcJ3ZgjtvUUV4DVLVPea3vqzpNW4pb
k2eOkZjp7Ctxyul7dTIbZQUw86g3V6dTqs4qXLkHAOd+UEjNhAESOnx/LPfOv+rt
JCpnnyIVTQQ2K8xvmLSeFhzIbsNe5wssQQKCAQEArRP4J+xH9LovUQKf5HBlKCtS
arUXLXiSjm65DQEalrHBZTzKOr/wDy55gWmNlwqMMU5sFQHvfxAOA/7yQqbUrM/S
CFyzGm/6Lz7pFe8XDLuxtaxLscC00HmeBM8QrnLZjfVQnZEf6TS4Zv6IqQTf1NRt
91JW6aVDAF3aHep2afbbXlROFthVDd/PUcFo8fwSP5oibVf6ELfyzuyI4Y0cvM1p
HTU1kxGzAd1bTJNr0V1xkqcMc2EGjIoDOCgXxfLJf1tAYt/b0cooqFOc4XVq1pYq
NXENNyn3UtWXR+4ni6jmpud2Xn2xeJKEiI1MtMz7DfOa/yolB8AomFa6Zhvc/A==
-----END RSA PRIVATE KEY-----|-----BEGIN RSA PUBLIC KEY-----
MIICCgKCAgEAjdgFBLDvmGQvl3o87UZAxUCuGwKqym92Iv4FGujlcr95oPxkIoXt
LYnZUNXBAPdDRqeag8dhTbftCuywd99xJybV4uoqlOSghd2lu5nrtixrcIF/x3KH
+5brTzLjGaTV5WsK9hLXgnEBdPDJ6FadCo1Se8/6D5oAvKQYoqNNgZqiJBYsf9rO
0bPQqwqSaghWcfGVO/tbJY0OUfhOlhqokeGsvFa3c9g5Ca0kHvpY8u/yF33xiCl7
yadza68yxHu5X5/ciqEOldzsyK5SauyVANGHeIio5X9yVHKiEUeNLkRh3Zd1QLZv
LT6iHyBbhFzN+aI7nWUZx8mqTQPcOoJdUgXpwzsbq2ODgqZf6vLG+vonoW8A+Irt
pyP3t/SdgzwtTt9j8ssK08ma4K+mC5vHdyNt2zd0bT05ZUl2X0hGtcGgkusBmfXO
gV0CD6KynI7E25kNSFm5puCxPiZ5tisyGdrfgBM5lK8tWl5AZPBbPmNNnriUKOIr
C72EsXnxV5F3ZX626t1YUKMpfmkr9H4yicSjHYtZDVInF1VeNOW3eMUUgiw0kyEZ
ZZLQigUd1IuLkKEvoW9R1gtXKXy9e3xv8v0uhvdLL7UCEA8bDCAmzJ8fOTQLIfxe
Ad3bM7JcmnzdglKgGTQU16tIzhizQHOKGM3gNf2W1f3TJyGU8lseA50CAwEAAQ==
-----END RSA PUBLIC KEY-----|-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAsCNPGn5mU4Yd5XrFPTQfi/mWvMDH7xv7d7YANik3SjHUOG89
zW975zFtwq8TbXsmqvvwATIJ84waYKOPm6fMHtZXKxsLGZWkMPXIUFjyubnfHwOC
n/YHi2OOwBv/8VEoT7sek0zP/RgBhu1HG8NUT/mrSJOvwmMu/QYuwDYArx4GYmoz
5gUw0/3FkcV6PaknbHrAR4tTXrrM2DU31DGR/MVIShsVaDs9l1V1qzkTQX/u5zTM
1W/2YcMyC//5G2UjA3qG40ex4OzvoZVJbi/RMQWTOBI9Cb7FRxLyn5jtFDX7hiGN
EyglphweOgVb+zHAzEtemcZ6wzy50C5mgqZqdwIDAQABAoIBAQCt9gGPgh1gCssG
NY4OV6d8oMJeo+wTXUiswAVPeL+RuMsH/d5FQwFVonzsVevTKllTcEZnd9taPhpj
LceNDVnNbxZijXAWEbQ7YtQ3ftIj2s7iG55sjBqBcY+wbzcGSIiOzplxE+Er8AS6
FSkw42cnIcgogyyIyVrELdIIpo1gZZwesYWiEBCf8vkXUSrnqJezbG1ZeEW+UVwx
oLeLlrhzqtY1aNTpprdrKvPqXlelDS3Fp4M1HeHLjVM5iTbJQBIJUJz8jr5CP9dr
U+xPsXJEPWVoKGRctFSFMFzF22HW7nQ8JA+Ap76JcpZmSob7+hcgHTA5r5nWdvk1
2zSjGMJ5AoGBAN9RtKg9FoqNs/EY2I+y7VTCKTDYHomXJoJprrRxk7+TNPCyc+El
JAtTpUtGg342cFlzXe1LoTuwpz0dzQ94XJru//w1xNKePu9pG7I91v+pmZ3qscQQ
3Hfsbii96DxvV9E2UngYM/84EdmxivwgfMB0889exPbLp6x2Q4kjZkd1AoGBAMnq
CK8QOmHqP15fbBAfq/5IIyoLHSr0EsuyB+4w1Tt6Mp6JmsSj3WODcEFX6eNTDEAY
QzD6iuZt5Gumgu4u78qcVzzPQ88nyu771Z1cO7+0aHWZQnmn8tdlW8QoYe8Q0QGP
btk5Zj3IpnxcBPOAy8tEam+K/nbM6GIDoskY4Fi7AoGAEZYUnP8ajFS2a/Eh9McJ
a5hSuBb86KbvkGuUJ3+ePALys1XKTFupZ/7xUGlAJ6tP06uDJRHEo7e4o9Z59JLs
8AKMRCdT1aZCnIXxeWrnSIDpqQb3ctJFGmiTEEBMTp4SE1aszepSBfI7baJOqKZC
TMGpBZaX+jt1e95gWSdhgZUCgYBVZ/n+WsXnLtdZCtHBzDPLHiZJnWPcDFt677Yt
CVjYjKqmppvhe1Kkyi8MZKxY7ILtBZY+PGvOZkNNSaRuLUpekFzSKvWKxVEwXzOP
GTlosF3TCF226dRtScFUKqSqI5nUqCACICZgEeRiC+ZIPX12YF0JC2jxRWEyPEl5
+FzFCwKBgApQJN4uYc40vq3qiXgjUrvn5ir2Bmw9Uq139As/jBgRq2pF4mw1PsKN
fJZ1ePsMzQvSvN4o7NJyGkdH3okkkfXQyVzAJ5jEU18WsDUOZaHfrBvaDsTbm6eI
ectYiZxpZPQoSmaoltD+FcZltXU2hS1NiXe8dkg2NMzgzdD2LJvn
-----END RSA PRIVATE KEY-----|-----BEGIN RSA PUBLIC KEY-----
MIIBCgKCAQEAsCNPGn5mU4Yd5XrFPTQfi/mWvMDH7xv7d7YANik3SjHUOG89zW97
5zFtwq8TbXsmqvvwATIJ84waYKOPm6fMHtZXKxsLGZWkMPXIUFjyubnfHwOCn/YH
i2OOwBv/8VEoT7sek0zP/RgBhu1HG8NUT/mrSJOvwmMu/QYuwDYArx4GYmoz5gUw
0/3FkcV6PaknbHrAR4tTXrrM2DU31DGR/MVIShsVaDs9l1V1qzkTQX/u5zTM1W/2
YcMyC//5G2UjA3qG40ex4OzvoZVJbi/RMQWTOBI9Cb7FRxLyn5jtFDX7hiGNEygl
phweOgVb+zHAzEtemcZ6wzy50C5mgqZqdwIDAQAB
-----END RSA PUBLIC KEY-----`

var nicoKeyString = `-----BEGIN RSA PRIVATE KEY-----
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
-----END RSA PRIVATE KEY-----|-----BEGIN RSA PUBLIC KEY-----
MIICCgKCAgEAhunQ5HoEjjrDKK3kWbnHEc7QDa2FN4o6Z3yPULQxbZ8ZMjy1aipW
LghpoGDlLOXG/T7zhPBe9Vsz1g1I3wR+U6Jd3C4M7B3kJPzOWXUaXuIYGtPVttE4
MBNDdNUxv2P5F3ZJiSBhNAbGIDgxEFvizGJcJnGYNUNnAu+B7VYdNoUS+mAGKmXm
cZHOOgkbZrs3pRFmNcnLiVf1vx1BFML9U8rMajDw2Ikgn9UBNmYnQmWaycf1fxF4
kRyuyV9z1O3O9Ez2yzJWw3FbzLpv3wDimmuDElfe+2DMZuc8JkovKqnPadDlXsO7
GF5qaA8tAJM64FSJz3uEaPzJKeJtEc6VvpIY8PGsuF0nYHQLp2PVsy2iTDBhADGx
J6fW2k2tuWQpyB4B9M9PujAZ4vnGrYoHrfEmfJunJvPdvIKaCXQm1R0a0DbIfXAR
/Avp6VkTUB5oSLm8TVOqzhp741I5iXFGhJFuYvbyW5TxNZoDKmwMzhYQtZEK1Jxx
ItfvuoJnNpmLqbNc2QmxAVtkCM2cgtvTxuKMwh1tGgWfqrS4b26K5Da2NTsnAF+u
++14V98sJfh8IRhmRYmLx/Sin2vCp+xqJA05AAQpBFXID/3KUa6nZ2EI+44mlJ2Y
Z3Tj+Qkp9XZZ62D/lfPCGIp/5DNANr0jCBZc57hBJHuvdHma0Y6rErMCAwEAAQ==
-----END RSA PUBLIC KEY-----|-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAneEpLS5VrIjJ9cwMZ+/BZdVkA59wO3C4R8hwP9iwmCnP8UiS
dVVMnI8WSWIBi9DNDfNZmOOcXVAdNOKW/XBMfapdRtNSdjZnoGhzdv24oIt4ks6P
br5/XhgLwm6MS5YcUshXFiY8Wp5D212ZJxdxasx9pNUSXrzPBi6hGlcxt/tQ1D6I
tOR3qGVJd6FgCXVXTRMaWU1xieQuVz+SycS6rA6ab0a/ZzoOoeo9Hj24K3wVyY/E
yTIAVeCqd1mH2HObdGrPa+oFrhaaiuBpTJujOMkfwOk8YD9yHbUnPizhp0Xzl65F
lnU6C5oNpcyDkyZXB9Ga2TT0gbRGODRKWr5XBQIDAQABAoIBAA8Mkmc7HnGGzHGA
lMIniXDSiWkK/qC72uUExyZe2l52xdqDB2rlyZq2zPILxIf2TEWDJ8w6DgKbC12q
J6ha1v+3iJKivB6GixQwQKlcAAhVoxWjWOq+f9eRRTWwYfEmiuadMCftVx1wLpmv
b1TZmDcM7QEUSl77cmRtd2/SVevqg1zeAry4zb0kZqsRV/sAm8yTAbA45LlMqfRu
XjXxi19niEmhjZe2pYklbvobLS4Snu8+YhdEatfHYQOjDXNHDN6QwCZLNLBEuquu
+1dLiqQpRLBxKa33whhcdpgYsXe6vWr8p01pCdtSNhpdCMf1bQ/ybV3CE5XfkfKE
Xk22w2ECgYEA1xIFZ5Lx5GHfbI5TbVo86u3uXtE99euJz+c+AJENH1bmYlAC/Iu8
7sL7URP1RQfg+4bqnVKTqBeD7nM8nEYnnUlD+SSyuqHwKxIe9d6zB+tMDfDSKKGa
09WTRQVB1mOfUI4zAoO+oXoQqBNf/8p+3WR9BSxcMqMnzaDe1FNrzC0CgYEAu+ze
ITm8Da8Y59/rbk5SALXnyPakCVGAD01Iew4mPSnYccYZdgJzXHKAmYamZ4gYn32k
kRUQ61pCo4zw/IYF/lklfxJWcDzRHkYo9uv3hp7BRBDYFh29KZuVxqk7TNB9UsTg
67Oq3WZM5LVVO3722hbABUFWwXYqyGp9N/tlBTkCgYEAjPXoyOvM0w067OGT2JZO
WwQ7ObYW0Dnih3tebeykFyhTs5gRfQjxeeQd6BQWq3nwgar51scqs/9Fn2G0FhD9
Id/FR2RAKIa+7NvhovClkskrfe+bLrpMyg6viYtajXOaFag90qYaqadhh2e7geFs
qRqVvnBGtYGuMt3/blaGzrkCgYBpoEVqH21Xs6coDW/i3BqzU4soJZkls61q7GFk
6RsNHcwOhu4gqlf1ClbF9Czr3E8atiAde926q0zxaHK7PQl5YNn2hnQXOdfAUGcO
Aqtp+yld/Km4JJcCH/UbaFtwXW9Tal2RnbbYInvtwQ5bFo3hHxBm48bDjouVgmVH
ZgHt4QKBgFvxCQzhKcSc7vwhNa+/B19XK6r8AS81jeBR4spuiDlviilTMHCUYX4n
ie1DR+PEZGsWv+Q0gzQTsUGZAPWBrktWsY9cst0R6Kto5zYUa22NjUjHGZirQy2G
FvrvXDHJZB/8AfEl5yKWW5QT/R/+MjRMX1kLI0n2KEB6N0U3DHV0
-----END RSA PRIVATE KEY-----|-----BEGIN RSA PUBLIC KEY-----
MIIBCgKCAQEAneEpLS5VrIjJ9cwMZ+/BZdVkA59wO3C4R8hwP9iwmCnP8UiSdVVM
nI8WSWIBi9DNDfNZmOOcXVAdNOKW/XBMfapdRtNSdjZnoGhzdv24oIt4ks6Pbr5/
XhgLwm6MS5YcUshXFiY8Wp5D212ZJxdxasx9pNUSXrzPBi6hGlcxt/tQ1D6ItOR3
qGVJd6FgCXVXTRMaWU1xieQuVz+SycS6rA6ab0a/ZzoOoeo9Hj24K3wVyY/EyTIA
VeCqd1mH2HObdGrPa+oFrhaaiuBpTJujOMkfwOk8YD9yHbUnPizhp0Xzl65FlnU6
C5oNpcyDkyZXB9Ga2TT0gbRGODRKWr5XBQIDAQAB
-----END RSA PUBLIC KEY-----`

func stringToKeyBytes(key string) []byte {
	r := strings.NewReader(key)
	pemBytes, _ := ioutil.ReadAll(r)
	block, _ := pem.Decode(pemBytes)
	return block.Bytes
}

func createNewUsers() {
	alcoContent, _ := ioutil.ReadFile("_Alcohol.pem")
	alcoText := string(alcoContent)
	alcoSplitted := strings.Split(alcoText, "|")
	alcoAdress := calc.Hash(stringToKeyBytes(alcoSplitted[1]))
	alcoMesKey := stringToKeyBytes(alcoSplitted[3])
	user.Create(alcoAdress, alcoMesKey, "Alcohol")
	alco := user.Get(alcoAdress)
	defer alco.Save()
	if alco.Balance == 0 {
		alco.Balance = 50000
	}
	marketAdress := string([]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 63})
	if alco.Balances[marketAdress] == 0 {
		alco.Balances[marketAdress] = 10000
	}
	fmt.Println("alco wallet created with", alco.Balance, "balance")
	nicoContent, _ := ioutil.ReadFile("_Nicotin.pem")
	nicoText := string(nicoContent)
	nicoSplitted := strings.Split(nicoText, "|")
	nicoAdress := calc.Hash(stringToKeyBytes(nicoSplitted[1]))
	nicoMesKey := stringToKeyBytes(nicoSplitted[3])
	user.Create(nicoAdress, nicoMesKey, "Nicotin")
	nico := user.Get(nicoAdress)
	defer nico.Save()
	if nico.Balance == 0 {
		nico.Balance = 50000
	}
	fmt.Println("nico wallet created with", nico.Balance, "balance")
}

func createStartMarket() {
	market.Create(
		[]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 63},
		"Bitcoin Ftem",
		[]byte{0, 1, 2, 3, 4, 2, 8},
		"There are many variations of passages of Lorem Ipsum available, but the majority have suffered alteration in some form, by injected humour, or randomised words which don't look even slightly believable. If you are going to use a passage of Lorem Ipsum, you need to be sure there isn't anything embarrassing hidden in the middle of text. All the Lorem Ipsum generators on the Internet tend to repeat predefined chunks as necessary, making this the first true generator on the Internet. It uses a dictionary of over 200 Latin words, combined with a handful of model sentence structures, to generate Lorem Ipsum which looks reasonable. The generated Lorem Ipsum is therefore always free from repetition, injected humour, or non-characteristic words etc.",
		"https://image.flaticon.com/icons/png/512/1490/1490849.png",
	)
	market.Create(
		[]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64},
		"Sber ruble Ftem",
		[]byte{0, 1, 2, 3, 4, 2, 8},
		"There are many variations of passages of Lorem Ipsum available, but the majority have suffered alteration in some form, by injected humour, or randomised words which don't look even slightly believable. If you are going to use a passage of Lorem Ipsum, you need to be sure there isn't anything embarrassing hidden in the middle of text. All the Lorem Ipsum generators on the Internet tend to repeat predefined chunks as necessary, making this the first true generator on the Internet. It uses a dictionary of over 200 Latin words, combined with a handful of model sentence structures, to generate Lorem Ipsum which looks reasonable. The generated Lorem Ipsum is therefore always free from repetition, injected humour, or non-characteristic words etc.",
		"https://image.flaticon.com/icons/png/512/1490/1490839.png",
	)
}

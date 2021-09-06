package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	encodeString := "C/nnnbP828eITQZEgaqrMEI/og2cqShHTLdQCBkvC0hG6rJ308QbtE/uw1uLxXIAhX7rIJTaVRyPqs3R12lF67mOS8QIFpNlSasUGO+62NU40NOSaqLjsHfbg3OaZQSjF2m/N9mnrCzJzOwf7OwQ4g9Ms6HrNDQIFG3MGABhrAyVaXTvIihCinM+5AI/cyliPR92D0eG/qAur8bsyywxmAnFGOvVUxx8Q5w2slWnwJEjAUAiKyaHOSTks02p4DWrMRon+yWkg6/vDWnO5vHPDs5AoTGiqCai2FrnZnZNKs3+bCbWCOqTaVlOoimiE+k6+Baihv3xf1caLa1Q5Y5Hx5jAxEolMG6UPbqCwwTfrMTCo9voECUUCWf+BkECLb8urAR5zxlZVOG9lEmP4hJ/9myOmDvcLZD5gMA0skzPFEiKBqABPamr7hEPb8KddBl7znlMgil2APXvIV9SpDHC+cg5M6yMS7HeS0PmolvLHxpDA7KNjvWCTbsKPFCpT7yD13FYbzQA8Pp03Hk6nsts19/QUetA4SZhIwFAIismhzn6uwfSkbnhASmMMv5UT1c7f9GHr3elEBl2aHo6Wp8Fddjlzr510yiMKU7GSgggmp4aO1IQaqtMTLcZpLDTS9DZ7Qj7OFWmYmAV7E8hIg4Ol+sM9NaTznhZj9ve8br1mRK4uWYIWHkkzlqp80m16y91W72pTzmwjqazrc4dvbyWUw=="
	// 演示base64编码
	decodeBytes, err := base64.StdEncoding.DecodeString(encodeString)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(decodeBytes))

}

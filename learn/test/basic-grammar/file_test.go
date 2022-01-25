package golearning_test

import (
	"fmt"
	"net/url"
	"path/filepath"
	"strings"
	"testing"
)

func TestFileName(t *testing.T) {
	file := "/Users/wangjian/Downloads/debug/data server zip/rest(tr3-start0)/subject1_pre_45frames_withrecon.zip"
	dir := filepath.Base(file)
	t.Logf("1: %s \n", dir)
}

func TestURlParse(t *testing.T) {
	signedUrl := "https://test2-ng-temp.neuralgalaxy.cn/drnobNFL2/s/1268/f/subject1_pre_45frames_withrecon.zip?OSSAccessKeyId=STS.NV5xtNirrpJm9LwUARHzBKNjf&Expires=1642676602&Signature=hcmscZE%2FMw8AVkjJxBQjlP%2BAMhk%3D&security-token=CAIShwJ1q6Ft5B2yfSjIr5WAM876hK1Tx4iGO2rGsUEHRPVupIvBhDz2IH5FenlsAO0YvvU2n29Q6%2FwalrNPRppdQlbNd9pM6Zda9wmvOdOb65Dp4LZc2MUk38NMhFmpsvXJasDVEfn%2FGJ70GX2m%2BwZ3xbzlD0bAO3WuLZyOj7N%2Bc90TRXPWRDFaBdBQVGAAwY1gQhm3D%2Fu2NQPwiWf9FVdhvhEG6Vly8qOi2MaRmHG85R%2FYsrZF99qge8D%2FMpcxbM0iD4%2BPsbYoJvab4kl58ANX8ap6tqtA9Arcs8uVa1sruEvXbbqIq4AxfF4gPfdiRvQctp7gmOZk4KnYkIj60RFJMPGx2tdbhGxlqxqAAS8nlMmjwlPuIDWm71iakBD6bQ6EgleC%2FZCuutFMGEWbFpk5zfbYvZCKdqFudv6AwDwn8%2FmY8g0ZSh7add3P9TTnvxJHTop1BInPjnoGjIzMBYYgHuRztJ5%2BusDylEVXyl%2FiDNZiX2JfFtSLZdVOuvDC0GRt%2Fos%2FRrGzcyc3Xpbf"
	urlValue, _ := url.ParseQuery(strings.Split(signedUrl, "?")[1])
	fmt.Printf("%#v\n", urlValue)
}

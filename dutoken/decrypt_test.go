/*
 * The MIT License (MIT)
 *
 * Copyright (c) 2014 Manuel Martínez-Almeida
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE.
 */

package dutoken

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestAndDecrypt(t *testing.T) {
	fmt.Println(AndDecrypt("AY0lDqGtikyfa35CB3SsNDNnyD//jjOtcGd5ZJBPO7Iw45gtCrkdUrnOWw5EBIMzYCwqwzYu1h8ZK9DdcJwLckgN5q19LoFBV3mmUZSjFqt4M/Xe52Y81Q=="))
	toString := hex.EncodeToString([]byte{49, 141, 213, 157, 240, 58, 201, 62})
	fmt.Println(toString)
	fmt.Println(TokenEncrypt("I4V2K3IPBPTRMQRLJ"))
	// sZK8+er5CS+A70OgaHjcNxlDRmDukA3LyrUEEnYdnL9jxU7VJUMt5Wv47iebe1RueR5fuEs+77Mi7hsM4f/YiTGN1Z3wOsk+
	// 04CVPFSuQgrioGLcKoolWH0LZv4qD76wHDFOGgwFgrPbeCoAyEegCSUTGgfcm12OW8+ycSQf0SyHa4WD0TTrlgAAAAAAAAAAVoR6/peZc5s=

}

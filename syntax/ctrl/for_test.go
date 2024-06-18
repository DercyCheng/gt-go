package main

import (
	"fmt"
	"testing"
)

func TestGoFor(t *testing.T) {
	// 在 Go 1.21 和 Go 1.22 里面运行
	for i := 0; i < 10; i++ {
		defer func() {
			fmt.Printf("%d, 地址 %p \n", i, &i)
		}()
	}
	// 1.22 结果
	//9, 地址 0x1400012a0f8
	//8, 地址 0x1400012a0f0
	//7, 地址 0x1400012a0e8
	//6, 地址 0x1400012a0e0
	//5, 地址 0x1400012a0d8
	//4, 地址 0x1400012a0d0
	//3, 地址 0x1400012a0c8
	//2, 地址 0x1400012a0c0
	//1, 地址 0x1400012a0b8
	//0, 地址 0x1400012a0b0

	//1.21 运行的结果
	//=== RUN   TestGoFor
	//10, 地址 0x140000a60a8
	//10, 地址 0x140000a60a8
	//10, 地址 0x140000a60a8
	//10, 地址 0x140000a60a8
	//10, 地址 0x140000a60a8
	//10, 地址 0x140000a60a8
	//10, 地址 0x140000a60a8
	//10, 地址 0x140000a60a8
	//10, 地址 0x140000a60a8
	//10, 地址 0x140000a60a8
}

func BenchmarkFor(b *testing.B) {
	src := make([]byte, 1024*1024)
	for i := 0; i < len(src); i++ {
		src[i] = 'a'
	}
	b.Run("for_loop", func(b *testing.B) {
		dst := make([]byte, 1024*1024)
		for i := 0; i < len(src); i++ {
			dst[i] = src[i]
			// 1. 先找到 src[i] 的起始地址
			// 2. dst[i] 的起始地址
			// 3. 将 src[i] 当前地址里面的数据，复制过去 dst_ptr[i] 上（根据你的类型信息，就能知道，一个元素多大）
			// 4. i ++
			// 5. 下一个循环
		}
	})
	b.Run("copy", func(b *testing.B) {
		dst := make([]byte, 1024*1024)
		copy(dst, src)
		// 这个实现
		// 1. copy 先找到 src 的起始地址，记为 src_ptr
		// 2. dst 的起始地址，记为 dst_ptr
		// 3. 将 src 当前地址里面的数据，复制过去 dst_ptr 上（根据你的类型信息，就能知道，一个元素多大）
		// 在这个例子里面，就是将 (src_ptr, src_ptr + 1) 的数据复制到 (dst_ptr, dst+ 1) 上
		// 4. 将 src_ptr = src_ptr + len(type), 这里 len(byte) = 1, dst_ptr = dst_ptr + 1
	})
}

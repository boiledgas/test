package bench

import "testing"

var deep byte = 5

func BenchmarkFnSliceCpy_100(b *testing.B) {
	benchmark(100, deep, b)
}

func BenchmarkFnSliceCpy_200(b *testing.B) {
	benchmark(200, deep, b)
}

func BenchmarkFnSliceCpy_300(b *testing.B) {
	benchmark(300, deep, b)
}

func BenchmarkFnSliceCpy_400(b *testing.B) {
	benchmark(400, deep, b)
}

func BenchmarkFnSliceCpy_500(b *testing.B) {
	benchmark(500, deep, b)
}

func BenchmarkFnSliceCpy_600(b *testing.B) {
	benchmark(600, deep, b)
}

func BenchmarkFnSliceCpy_700(b *testing.B) {
	benchmark(700, deep, b)
}

func BenchmarkFnSliceCpy_800(b *testing.B) {
	benchmark(800, deep, b)
}

func BenchmarkFnSliceCpy_900(b *testing.B) {
	benchmark(900, deep, b)
}

func BenchmarkFnSliceCpy_1000(b *testing.B) {
	benchmark(1000, deep, b)
}

func BenchmarkFnArrayCpy_100(b *testing.B) {
	var buf [100]byte
	for i := 0; i < len(buf); i++ {
		buf[i] = byte(i)
	}

	for i := 0; i < b.N; i++ {
		benchmark_recursive_100(0, deep, buf)
	}
}

func BenchmarkFnArrayCpy_200(b *testing.B) {
	var buf [200]byte
	for i := 0; i < len(buf); i++ {
		buf[i] = byte(i)
	}

	for i := 0; i < b.N; i++ {
		benchmark_recursive_200(0, deep, buf)
	}
}

func BenchmarkFnArrayCpy_300(b *testing.B) {
	var buf [300]byte
	for i := 0; i < len(buf); i++ {
		buf[i] = byte(i)
	}

	for i := 0; i < b.N; i++ {
		benchmark_recursive_300(0, deep, buf)
	}
}

func BenchmarkFnArrayCpy_400(b *testing.B) {
	var buf [400]byte
	for i := 0; i < len(buf); i++ {
		buf[i] = byte(i)
	}

	for i := 0; i < b.N; i++ {
		benchmark_recursive_400(0, deep, buf)
	}
}

func BenchmarkFnArrayCpy_500(b *testing.B) {
	var buf [500]byte
	for i := 0; i < len(buf); i++ {
		buf[i] = byte(i)
	}

	for i := 0; i < b.N; i++ {
		benchmark_recursive_500(0, deep, buf)
	}
}

func BenchmarkFnArrayCpy_600(b *testing.B) {
	var buf [600]byte
	for i := 0; i < len(buf); i++ {
		buf[i] = byte(i)
	}

	for i := 0; i < b.N; i++ {
		benchmark_recursive_600(0, deep, buf)
	}
}

func BenchmarkFnArrayCpy_700(b *testing.B) {
	var buf [700]byte
	for i := 0; i < len(buf); i++ {
		buf[i] = byte(i)
	}

	for i := 0; i < b.N; i++ {
		benchmark_recursive_700(0, deep, buf)
	}
}

func BenchmarkFnArrayCpy_800(b *testing.B) {
	var buf [800]byte
	for i := 0; i < len(buf); i++ {
		buf[i] = byte(i)
	}

	for i := 0; i < b.N; i++ {
		benchmark_recursive_800(0, deep, buf)
	}
}

func BenchmarkFnArrayCpy_900(b *testing.B) {
	var buf [900]byte
	for i := 0; i < len(buf); i++ {
		buf[i] = byte(i)
	}

	for i := 0; i < b.N; i++ {
		benchmark_recursive_900(0, deep, buf)
	}
}

func BenchmarkFnArrayCpy_1000(b *testing.B) {
	var buf [1000]byte
	for i := 0; i < len(buf); i++ {
		buf[i] = byte(i)
	}

	for i := 0; i < b.N; i++ {
		benchmark_recursive_1000(0, deep, buf)
	}
}

func benchmark(l uint16, deep byte, b *testing.B) {
	buf := make([]byte, l)
	for i := 0; i < len(buf); i++ {
		buf[i] = byte(i)
	}

	for i := 0; i < b.N; i++ {
		benchmark_recursive(0, deep, buf)
	}
}

func benchmark_recursive(i byte, deep byte, b []byte) {
	if i < deep {
		b_copy := make([]byte, len(b))
		copied := copy(b_copy, b)
		if copied < len(b) {
			panic("not copied")
		}

		benchmark_recursive(i+1, deep, b_copy)
	}
}

func benchmark_recursive_100(i byte, deep byte, b [100]byte) {
	if i < deep {
		benchmark_recursive_100(i+1, deep, b)
	}
}

func benchmark_recursive_200(i byte, deep byte, b [200]byte) {
	if i < deep {
		benchmark_recursive_200(i+1, deep, b)
	}
}

func benchmark_recursive_300(i byte, deep byte, b [300]byte) {
	if i < deep {
		benchmark_recursive_300(i+1, deep, b)
	}
}

func benchmark_recursive_400(i byte, deep byte, b [400]byte) {
	if i < deep {
		benchmark_recursive_400(i+1, deep, b)
	}
}

func benchmark_recursive_500(i byte, deep byte, b [500]byte) {
	if i < deep {
		benchmark_recursive_500(i+1, deep, b)
	}
}

func benchmark_recursive_600(i byte, deep byte, b [600]byte) {
	if i < deep {
		benchmark_recursive_600(i+1, deep, b)
	}
}

func benchmark_recursive_700(i byte, deep byte, b [700]byte) {
	if i < deep {
		benchmark_recursive_700(i+1, deep, b)
	}
}

func benchmark_recursive_800(i byte, deep byte, b [800]byte) {
	if i < deep {
		benchmark_recursive_800(i+1, deep, b)
	}
}

func benchmark_recursive_900(i byte, deep byte, b [900]byte) {
	if i < deep {
		benchmark_recursive_900(i+1, deep, b)
	}
}

func benchmark_recursive_1000(i byte, deep byte, b [1000]byte) {
	if i < deep {
		benchmark_recursive_1000(i+1, deep, b)
	}
}

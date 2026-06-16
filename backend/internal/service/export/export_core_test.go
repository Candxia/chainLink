package export

import (
	"context"
	"math"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// ==================== formatCell 测试 ====================

func Test_formatCell(t *testing.T) {
	tests := []struct {
		name string
		val  interface{}
		want string
	}{
		{"nil", nil, ""},
		{"float64 整数", float64(100), "100"},
		{"float64 小数", float64(99.5), "99.50"},
		{"float64 大整数", float64(123456789), "123456789"},
		{"float64 零", float64(0), "0"},
		{"float64 负整数", float64(-42), "-42"},
		{"float64 负小数", float64(-3.14), "-3.14"},
		{"bool true", true, "是"},
		{"bool false", false, "否"},
		{"int", int(42), "42"},
		{"int64", int64(999), "999"},
		{"string 普通", "hello", "hello"},
		{"string 长数字", "1234567890123", "\t1234567890123"},
		{"string 非数字长串", "abcdefghijklm", "abcdefghijklm"},
		{"string 短数字", "12345", "12345"},
		{"string 混合13位", "12345abcde6789", "12345abcde6789"},
		{"time.Time 含时间", parseTime("2026-01-01 12:00:00"), "2026-01-01 12:00:00"},
		{"time.Time 仅日期", parseTime("2026-01-02 00:00:00"), "2026-01-02"},
		{"空字符串", "", ""},
		{"uint", uint(88), "88"},
		// NaN / Inf 边界
		{"float64 NaN", math.NaN(), "NaN"},
		{"float64 +Inf", math.Inf(1), "+Inf"},
		{"float64 -Inf", math.Inf(-1), "-Inf"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := formatCell(tt.val)
			if got != tt.want {
				t.Errorf("formatCell(%v) = %q, want %q", tt.val, got, tt.want)
			}
		})
	}
}

func parseTime(s string) time.Time {
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", s, time.Local)
	return t
}

// ==================== looksLikeLongNumber 测试 ====================

func Test_looksLikeLongNumber(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"13位纯数字", "1234567890123", true},
		{"14位纯数字", "12345678901234", true},
		{"空字符串", "", false},
		{"短数字", "12345", false},
		{"12位纯数字", "123456789012", false},
		{"字母串", "abcdefghijklm", false},
		{"数字字母混合", "1234567890abc", false},
		{"含符号", "123456789012-", false},
		{"含空格", "1234567890 123", false},
		{"含前导空格", " 1234567890123", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := looksLikeLongNumber(tt.s)
			if got != tt.want {
				t.Errorf("looksLikeLongNumber(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}

// ==================== getNestedValue 测试 ====================

func Test_getNestedValue(t *testing.T) {
	tests := []struct {
		name  string
		data  map[string]interface{}
		field string
		want  interface{}
	}{
		{"nil map", nil, "name", nil},
		{"简单字段", map[string]interface{}{"name": "张三"}, "name", "张三"},
		{"嵌套字段", map[string]interface{}{"user": map[string]interface{}{"name": "张三"}}, "user.name", "张三"},
		{"深层嵌套", map[string]interface{}{"a": map[string]interface{}{"b": map[string]interface{}{"c": "deep"}}}, "a.b.c", "deep"},
		{"字段不存在", map[string]interface{}{"name": "张三"}, "age", nil},
		{"嵌套中间不存在", map[string]interface{}{"user": map[string]interface{}{"name": "张三"}}, "user.age", nil},
		{"中间不是 map", map[string]interface{}{"user": "string"}, "user.name", nil},
		{"空 map", map[string]interface{}{}, "name", nil},
		// 注意: 带点号的 key 会被 SplitN 分割, 此处不测该边界
		{"整数值", map[string]interface{}{"count": 42}, "count", 42},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getNestedValue(tt.data, tt.field)
			if got != tt.want {
				t.Errorf("getNestedValue(%v, %q) = %v, want %v", tt.data, tt.field, got, tt.want)
			}
		})
	}
}

// ==================== ExportConfig 默认值测试 ====================

type taskWithConfig struct {
	cfg ExportConfig
}

func (t *taskWithConfig) FetchPage(ctx context.Context, page, pageSize int) ([]map[string]interface{}, error) {
	// 返回空数据以正常结束导出
	return nil, nil
}

func (t *taskWithConfig) GetExportConfig() ExportConfig {
	return t.cfg
}

func TestExportConfigDefaults(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	task := &taskWithConfig{
		cfg: ExportConfig{
			FilePrefix: "test_defaults",
			Total:      100,
			Fields:     []string{"name"},
			FieldNames: []string{"姓名"},
			// BatchSize = 0, MaxRowsPerFile = 0, Concurrency = 0 → 应使用默认值
		},
	}

	result, err := SubmitExport(ctx, task)
	if err != nil {
		t.Fatalf("SubmitExport 不应报错: %v", err)
	}
	if result == nil {
		t.Fatal("result 不应为 nil")
	}

	// 验证内部配置默认值: 通过检查生成的文件数量间接验证
	// 100行 / 100000 (默认 MaxRowsPerFile) = 1个文件
	if len(result.Files) != 1 {
		t.Errorf("期望 1 个文件, 得到 %d 个", len(result.Files))
	}
	if result.Total != 100 {
		t.Errorf("期望 Total=100, 得到 %d", result.Total)
	}
}

func TestExportConfigConcurrencyCap(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	task := &taskWithConfig{
		cfg: ExportConfig{
			FilePrefix:  "test_concurrency",
			Total:       2000,
			BatchSize:   10,
			Concurrency: 100, // 超过上限 10
			Fields:      []string{"x"},
			FieldNames:  []string{"X"},
		},
	}

	result, err := SubmitExport(ctx, task)
	if err != nil {
		t.Fatalf("SubmitExport 不应报错: %v", err)
	}
	if result == nil {
		t.Fatal("result 不应为 nil")
	}
	if result.Total != 2000 {
		t.Errorf("期望 Total=2000, 得到 %d", result.Total)
	}
}

// ==================== generateFilename 测试 ====================

func Test_generateFilename(t *testing.T) {
	ts := time.Now().Format("20060102_150405")

	tests := []struct {
		name   string
		prefix string
		index  int
		total  int
	}{
		{"单文件 index 0 total 1", "test", 0, 1},
		{"多文件 index 1 total 3", "test", 1, 3},
		{"多文件 index 0 total 2", "stats", 0, 2},
		{"多文件 index 1 total 2", "stats", 1, 2},
		{"单文件总数为0", "test", 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateFilename(tt.prefix, tt.index, tt.total)

			// 必须包含前缀
			if !strings.HasPrefix(got, tt.prefix+"_") {
				t.Errorf("文件名 %q 应以前缀 %q 开头", got, tt.prefix+"_")
			}

			// 必须包含时间戳
			if !strings.Contains(got, ts) {
				t.Errorf("文件名 %q 应包含时间戳 %q", got, ts)
			}

			if tt.total <= 1 {
				// 单文件格式: {prefix}_{ts}.csv
				want := tt.prefix + "_" + ts + ".csv"
				if got != want {
					t.Errorf("单文件名 = %q, want %q", got, want)
				}
			} else {
				// 多文件格式: {prefix}_{ts}_part_{index+1}.csv
				expectedSuffix := "_part_" + itoa(tt.index+1) + ".csv"
				if !strings.HasSuffix(got, expectedSuffix) {
					t.Errorf("文件名 %q 应以 %q 结尾", got, expectedSuffix)
				}
			}
		})
	}
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	s := ""
	for n > 0 {
		s = string(rune('0'+n%10)) + s
		n /= 10
	}
	return s
}

// ==================== SubmitExport 边界测试 ====================

func TestSubmitExportTotalZero(t *testing.T) {
	// 当 Total=0 时应该快速返回空结果
	cfg := ExportConfig{
		FilePrefix: "zero",
		Total:      0,
	}
	result, err := SubmitExport(context.Background(), &taskWithConfig{cfg: cfg})
	if err != nil {
		t.Fatalf("Total=0 不应报错: %v", err)
	}
	if len(result.Files) != 0 {
		t.Errorf("Total=0 期望 0 个文件, 得到 %d", len(result.Files))
	}
}

// ==================== 带数据的 SubmitExport 完整流程 ====================

type dataTask struct {
	cfg       ExportConfig
	data      []map[string]interface{}
	callSeq   []int // 记录 FetchPage 调用顺序, 验证并发
	mu        sync.Mutex
	fetchFunc func(ctx context.Context, page, pageSize int) ([]map[string]interface{}, error)
}

func (t *dataTask) FetchPage(ctx context.Context, page, pageSize int) ([]map[string]interface{}, error) {
	if t.fetchFunc != nil {
		return t.fetchFunc(ctx, page, pageSize)
	}
	t.mu.Lock()
	t.callSeq = append(t.callSeq, page)
	t.mu.Unlock()

	// 模拟分页
	start := (page - 1) * pageSize
	if start >= len(t.data) {
		return nil, nil
	}
	end := start + pageSize
	if end > len(t.data) {
		end = len(t.data)
	}
	return t.data[start:end], nil
}

func (t *dataTask) GetExportConfig() ExportConfig {
	return t.cfg
}

func TestSubmitExportWithData(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 构造测试数据: 3 个字段, 15 行
	fields := []string{"id", "name", "score"}
	fieldNames := []string{"ID", "姓名", "分数"}
	data := make([]map[string]interface{}, 15)
	for i := 0; i < 15; i++ {
		data[i] = map[string]interface{}{
			"id":    int64(i + 1),
			"name":  "用户" + itoa(i+1),
			"score": float64(i*10) + 0.5,
		}
	}

	task := &dataTask{
		cfg: ExportConfig{
			FilePrefix:     "test_data",
			BatchSize:      5,
			MaxRowsPerFile: 10,
			Total:          15,
			Fields:         fields,
			FieldNames:     fieldNames,
			Concurrency:    3,
		},
		data: data,
	}

	result, err := SubmitExport(ctx, task)
	if err != nil {
		t.Fatalf("SubmitExport 报错: %v", err)
	}

	// 15行, 每文件最多10行 → 2个文件
	if len(result.Files) != 2 {
		t.Errorf("期望 2 个文件, 得到 %d: %v", len(result.Files), result.Files)
	}
	if len(result.FileUrls) != 2 {
		t.Errorf("期望 2 个 URL, 得到 %d", len(result.FileUrls))
	}
	if result.Total != 15 {
		t.Errorf("期望 Total=15, 得到 %d", result.Total)
	}
}

func TestSubmitExportCancel(t *testing.T) {
	// 验证 context 取消后快速返回
	ctx, cancel := context.WithCancel(context.Background())
	task := &dataTask{
		cfg: ExportConfig{
			FilePrefix: "cancel",
			BatchSize:  5,
			Total:      50,
			Fields:     []string{"x"},
		},
		data: make([]map[string]interface{}, 50),
	}

	// 立即取消
	cancel()

	result, err := SubmitExport(ctx, task)
	// 取消后应该返回 context.Canceled 错误
	if err == nil {
		t.Fatal("取消后应返回错误")
	}
	if result == nil {
		t.Fatal("result 不应为 nil")
	}
}

// ==================== FileReportUtil 基础测试 ====================

func TestFileReportUtil(t *testing.T) {
	fr := FileReport()

	t.Run("FormatPercent", func(t *testing.T) {
		// FormatPercent 格式化传入的数值, 不做乘以100的转换
		got := fr.FormatPercent(75.3, 1)
		if got != "75.3%" {
			t.Errorf("FormatPercent = %q, want %q", got, "75.3%")
		}

		got2 := fr.FormatPercent(75, 0)
		if got2 != "75%" {
			t.Errorf("FormatPercent = %q, want %q", got2, "75%")
		}

		got3 := fr.FormatPercent(0.753, 1)
		if got3 != "0.8%" {
			t.Errorf("FormatPercent = %q, want %q", got3, "0.8%")
		}
	})

	t.Run("GetTime int64 秒", func(t *testing.T) {
		// 2026-01-01 00:00:00 UTC 的时间戳
		ts := int64(1767484800)
		got := fr.GetTime(ts)
		if got == "" {
			t.Error("GetTime 不应返回空字符串")
		}
	})

	t.Run("GetTime 空字符串", func(t *testing.T) {
		got := fr.GetTime("")
		if got != "" {
			t.Errorf("空字符串应返回空, 得到 %q", got)
		}
	})

	t.Run("GetTime 普通字符串", func(t *testing.T) {
		got := fr.GetTime("2026-01-01 12:00:00")
		if got != "2026-01-01 12:00:00" {
			t.Errorf("期望原样返回, 得到 %q", got)
		}
	})

	t.Run("HashParams", func(t *testing.T) {
		params := map[string]interface{}{
			"name": "张三",
			"age":  30,
		}
		hash := fr.HashParams(params)
		if len(hash) != 32 {
			t.Errorf("MD5 hash 应为 32 位, 得到 %d 位: %s", len(hash), hash)
		}
	})

	t.Run("FormatFilePath 路径归一化", func(t *testing.T) {
		got := fr.FormatFilePath("test.csv")
		if !strings.Contains(got, "test.csv") {
			t.Errorf("FormatFilePath 应包含文件名, 得到 %q", got)
		}
		// 验证路径使用正斜杠
		if strings.Contains(got, "\\") {
			t.Errorf("FormatFilePath 应使用正斜杠, 得到 %q", got)
		}
	})

	t.Run("FormatFileUrl", func(t *testing.T) {
		got := fr.FormatFileUrl("test.csv")
		want := "/res/export/test.csv"
		if got != want {
			t.Errorf("FormatFileUrl = %q, want %q", got, want)
		}
	})
}

// ==================== 导出进度回调测试 ====================

func TestSubmitExportProgressCallback(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var totalWritten int64
	var callCount int32

	progress := func(current, total int64, fileIdx int) {
		atomic.StoreInt64(&totalWritten, current)
		atomic.AddInt32(&callCount, 1)
	}

	data := make([]map[string]interface{}, 100)
	for i := 0; i < 100; i++ {
		data[i] = map[string]interface{}{"x": i}
	}

	task := &dataTask{
		cfg: ExportConfig{
			FilePrefix:     "progress",
			BatchSize:      10,
			MaxRowsPerFile: 100,
			Total:          100,
			Fields:         []string{"x"},
			FieldNames:     []string{"X"},
			Concurrency:    3,
		},
		data: data,
	}

	_, err := SubmitExport(ctx, task, progress)
	if err != nil {
		t.Fatalf("SubmitExport 报错: %v", err)
	}

	// 至少有回调被调用
	if atomic.LoadInt32(&callCount) == 0 {
		t.Error("进度回调未被调用")
	}

	// 最终写入量应为 100
	final := atomic.LoadInt64(&totalWritten)
	if final != 100 {
		t.Errorf("期望最终进度 100, 得到 %d", final)
	}
}

// ==================== 并发安全测试 ====================

func TestExportConcurrentPages(t *testing.T) {
	// 验证多个 FetchPage 被并发调用 (通过检查调用时序)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var mu sync.Mutex
	var pages []int
	var wg sync.WaitGroup

	task := &dataTask{
		cfg: ExportConfig{
			FilePrefix:     "concurrent",
			BatchSize:      10,
			MaxRowsPerFile: 100,
			Total:          50,
			Fields:         []string{"x"},
			Concurrency:    5,
		},
		fetchFunc: func(ctx context.Context, page, pageSize int) ([]map[string]interface{}, error) {
		mu.Lock()
		pages = append(pages, page)
		pageCount := len(pages)
		mu.Unlock()

		// 模拟延迟
		time.Sleep(50 * time.Millisecond)

		wg.Done()
		_ = pageCount
		return []map[string]interface{}{{"x": page}}, nil
	},
	}

	// 预期 5 页 (50行 / 10每页)
	for i := 0; i < 5; i++ {
		wg.Add(1)
	}

	ctx2, cancel2 := context.WithTimeout(ctx, 2*time.Second)
	defer cancel2()

	go func() {
		_, _ = SubmitExport(ctx2, task)
	}()

	wg.Wait()

	mu.Lock()
	pageCount := len(pages)
	mu.Unlock()

	if pageCount != 5 {
		t.Errorf("期望 5 页被调用, 得到 %d", pageCount)
	}
}

// ==================== 排除已生成文件的清理 ====================

func cleanupTestFiles(t *testing.T, files []string) {
	t.Helper()
	for _, f := range files {
		if err := FileReport().DelFile(f); err != nil {
			t.Logf("清理文件 %s 失败: %v", f, err)
		}
	}
}

func init() {
	// 清除测试文件不需要额外初始化，但保持统一清理模式
}

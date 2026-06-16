package export

import (
	"context"
	"crypto/md5"
	"encoding/csv"
	"encoding/hex"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

// ExportConfig 导出配置
type ExportConfig struct {
	FilePrefix      string   // 导出文件前缀
	Limit           int      // 每次查询行数
	MaxRowsPerFile  int      // 每个文件最大行数
	Total           int64    // 总行数
	Fields          []string // 字段键列表
	FieldNames      []string // 字段显示名称
	FileExt         string   // 文件扩展名（默认 .csv）
}

// ExportTask 导出任务接口
type ExportTask interface {
	// FetchPage 分页获取导出数据
	FetchPage(ctx context.Context, page, pageSize int) ([]map[string]interface{}, error)
	// GetExportConfig 获取导出配置
	GetExportConfig() ExportConfig
}

// ExportResult 导出结果
type ExportResult struct {
	Files    []string // 生成的文件路径列表
	FileUrls []string // 可访问的文件 URL 列表
	Total    int64    // 总记录数
}

// SubmitExport 提交导出任务
func SubmitExport(ctx context.Context, task ExportTask) (*ExportResult, error) {
	cfg := task.GetExportConfig()
	result := &ExportResult{
		Total: cfg.Total,
	}

	if cfg.Total == 0 {
		return result, nil
	}

	if cfg.FileExt == "" {
		cfg.FileExt = ".csv"
	}

	totalRows := int(cfg.Total)
	rowsPerFile := cfg.MaxRowsPerFile
	if rowsPerFile <= 0 {
		rowsPerFile = 50000
	}
	fileCount := int(math.Ceil(float64(totalRows) / float64(rowsPerFile)))

	for fileIdx := 0; fileIdx < fileCount; fileIdx++ {
		startRow := fileIdx * rowsPerFile
		endRow := (fileIdx + 1) * rowsPerFile
		if endRow > totalRows {
			endRow = totalRows
		}

		filename := generateFilename(cfg.FilePrefix, fileIdx, fileCount, cfg.FileExt)
		filepath := GetExportFilePath(filename)

		if err := writeCSV(ctx, task, cfg, startRow, endRow, filepath); err != nil {
			return nil, fmt.Errorf("写入导出文件失败: %w", err)
		}

		result.Files = append(result.Files, filepath)
		result.FileUrls = append(result.FileUrls, GetExportFileUrl(filename))
	}

	return result, nil
}

// writeCSV 写入 CSV 文件
func writeCSV(ctx context.Context, task ExportTask, cfg ExportConfig, startRow, endRow int, outPath string) error {
	dir := filepath.Dir(outPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	f, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer f.Close()

	// 写入 UTF-8 BOM 以便 Excel 正确识别中文
	f.WriteString("\xEF\xBB\xBF")

	writer := csv.NewWriter(f)
	defer writer.Flush()

	// 写入表头
	if len(cfg.FieldNames) > 0 {
		if err := writer.Write(cfg.FieldNames); err != nil {
			return err
		}
	}

	// 计算分页参数
	totalPages := int(math.Ceil(float64(endRow-startRow) / float64(cfg.Limit)))
	for page := 1; page <= totalPages; page++ {
		data, err := task.FetchPage(ctx, page, cfg.Limit)
		if err != nil {
			g.Log().Warningf(ctx, "导出分页查询失败 page=%d: %v", page, err)
			break
		}
		if len(data) == 0 {
			break
		}

		for _, row := range data {
			record := make([]string, len(cfg.Fields))
			for i, field := range cfg.Fields {
				val := getNestedValue(row, field)
				record[i] = fmt.Sprintf("%v", val)
			}
			if err := writer.Write(record); err != nil {
				return err
			}
		}
	}

	return nil
}

// generateFilename 生成导出文件名
func generateFilename(prefix string, index, total int, ext string) string {
	ts := time.Now().Format("20060102_150405")
	if total <= 1 {
		return fmt.Sprintf("%s_%s%s", prefix, ts, ext)
	}
	return fmt.Sprintf("%s_%s_part_%d%s", prefix, ts, index+1, ext)
}

// getNestedValue 从 map 中获取字段值，支持点号分隔的嵌套字段
func getNestedValue(data map[string]interface{}, field string) interface{} {
	parts := strings.SplitN(field, ".", 2)
	if len(parts) == 1 {
		return data[field]
	}
	if sub, ok := data[parts[0]].(map[string]interface{}); ok {
		return getNestedValue(sub, parts[1])
	}
	return nil
}

// ==================== 文件路径工具 ====================

// GetExportDir 获取导出文件存放目录
func GetExportDir() string {
	dir := g.Cfg().MustGet(context.Background(), "server.exportDir", "resource/export").String()
	return dir
}

// GetExportFilePath 获取导出文件的完整路径
func GetExportFilePath(filename string) string {
	dir := GetExportDir()
	return filepath.Join(dir, filename)
}

// GetExportFileUrl 获取导出文件的访问 URL
func GetExportFileUrl(filename string) string {
	return "/res/export/" + filename
}

// ==================== 文件报告工具 ====================

type FileReportUtil struct{}

func FileReport() *FileReportUtil {
	return &FileReportUtil{}
}

// FormatFilePath 格式化文件路径
func (f *FileReportUtil) FormatFilePath(filename string) string {
	dir := GetExportDir()
	re := regexp.MustCompile(`[/\\]+`)
	return re.ReplaceAllString(filepath.Join(dir, filename), "/")
}

// FormatFileUrl 格式化文件 URL
func (f *FileReportUtil) FormatFileUrl(filename string) string {
	return GetExportFileUrl(filename)
}

// FormatPercent 格式化百分比值
func (f *FileReportUtil) FormatPercent(val interface{}, digits int) string {
	v := gconv.Float64(val)
	format := "%." + fmt.Sprintf("%d", digits) + "f%%"
	return fmt.Sprintf(format, v)
}

// GetTime 时间戳转日期字符串
func (f *FileReportUtil) GetTime(ts int64) string {
	if ts == 0 {
		return ""
	}
	return gtime.NewFromTimeStamp(ts).Format("Y-m-d H:i:s")
}

// HashParams 将参数排序后计算 MD5 哈希（用于缓存键）
func (f *FileReportUtil) HashParams(params interface{}) string {
	m := gconv.Map(params)
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var buf strings.Builder
	for _, k := range keys {
		buf.WriteString(fmt.Sprintf("%s:%v,", k, m[k]))
	}
	hash := md5.Sum([]byte(buf.String()))
	return hex.EncodeToString(hash[:])
}

// FileSize 获取文件大小的人类可读格式
func (f *FileReportUtil) FileSize(filename string) string {
	info, err := os.Stat(filename)
	if err != nil {
		return "0 B"
	}
	return humanize.Bytes(uint64(info.Size()))
}

// DelFile 删除文件
func (f *FileReportUtil) DelFile(filepath string) error {
	return os.Remove(filepath)
}

package main

import (
	"bytes"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/Andrew-M-C/go.util/log"
	"github.com/Andrew-M-C/go.util/unsafe"
)

func main() {
	handleNameAbbreviation()
}

func init() {
	log.SetLevel(log.NoLog, log.DebugLevel)
}

// 处理无歧义缩写
func handleNameAbbreviation() {
	const fileSrc = "../../常用无歧义缩写.md"
	lines, err := readTextFile(fileSrc)
	if err != nil {
		log.Fatalf("读取文件 '%s' 错误: %v", fileSrc, err)
	}

	// 读取数据
	names := make([]naming, 0, len(lines))
	textStarts := false
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if !strings.HasPrefix(line, "|") {
			continue
		}
		if !textStarts {
			if strings.HasPrefix(line, "|:") || strings.HasPrefix(line, "|-") {
				textStarts = true
			}
			continue
		}
		cols := spiltAndTrimSpace(line, "|")
		if len(cols) < 4 {
			log.Errorf("不合法的行: '%s'", line)
			continue
		}
		names = append(names, naming{
			Short:   cols[1],
			Full:    cols[2],
			Comment: cols[3],
		})
	}

	// 按照字母序重新排序
	sort.Slice(names, func(i, j int) bool {
		if names[i].Short == "" {
			return false
		}
		if names[j].Short == "" {
			return true
		}
		return strings.ToLower(names[i].Short) <= strings.ToLower(names[j].Short)
	})

	// 输出到文件
	buff := bytes.Buffer{}
	buff.WriteString("# 常用无歧义缩写\n\n")
	buff.WriteString("| 缩写 | 完整词汇 | 备注 |\n")
	buff.WriteString("|:---:|:---:|:---|\n")

	for _, name := range names {
		line := fmt.Sprintf("| %s | %s | %s |", name.Short, name.Full, name.Comment)
		buff.WriteString(line)
		buff.WriteRune('\n')
	}

	err = os.WriteFile(fileSrc, buff.Bytes(), 0644)
	if err != nil {
		log.Fatalf("写入文件 '%s' 错误: %v", err)
		return
	}

	log.Infof("已更新文件 '%s', 共写入 %d 行简写词", fileSrc, len(names))
}

type naming struct {
	Short   string
	Full    string
	Comment string
}

func spiltAndTrimSpace(s string, sep string) []string {
	res := strings.Split(s, sep)
	for i, part := range res {
		res[i] = strings.TrimSpace(part)
	}
	return res
}

func readTextFile(path string) ([]string, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(unsafe.BtoS(b), "\n")
	for i, s := range lines {
		lines[i] = strings.TrimSuffix(s, "\r")
	}
	return lines, nil
}

package logger

import (
	"fmt"
	"time"
)

// Timer 性能监控工具
type Timer struct {
	startTime time.Time
	name      string
}

// NewTimer 创建新的计时器
func NewTimer(name string) *Timer {
	return &Timer{
		startTime: time.Now(),
		name:      name,
	}
}

// Elapsed 获取经过的时间
func (t *Timer) Elapsed() time.Duration {
	return time.Since(t.startTime)
}

// ElapsedMs 获取经过的时间（毫秒）
func (t *Timer) ElapsedMs() int64 {
	return t.Elapsed().Milliseconds()
}

// ElapsedString 获取格式化的耗时字符串
func (t *Timer) ElapsedString() string {
	elapsed := t.Elapsed()

	if elapsed < time.Millisecond {
		return fmt.Sprintf("%.2fμs", float64(elapsed.Microseconds()))
	} else if elapsed < time.Second {
		return fmt.Sprintf("%.2fms", float64(elapsed.Milliseconds()))
	} else {
		return fmt.Sprintf("%.2fs", elapsed.Seconds())
	}
}

// Log 记录执行时间（Info级别）
func (t *Timer) Log() {
	Info(fmt.Sprintf("[性能][%s] 耗时: %s", t.name, t.ElapsedString()))
}

// LogSlow 记录执行时间，如果超过阈值则警告
func (t *Timer) LogSlow(threshold time.Duration) {
	elapsed := t.Elapsed()
	if elapsed > threshold {
		Warn(fmt.Sprintf("[性能][%s] 耗时过长: %s (阈值: %s)", t.name, t.ElapsedString(), threshold.String()))
	} else {
		t.Log()
	}
}

// LogError 记录执行时间（Error级别）
func (t *Timer) LogError(message string) {
	Error(fmt.Sprintf("[性能][%s] %s | 耗时: %s", t.name, message, t.ElapsedString()))
}

// LogWithMsg 记录执行时间和自定义消息
func (t *Timer) LogWithMsg(level string, message string) {
	fullMessage := fmt.Sprintf("[性能][%s] %s | 耗时: %s", t.name, message, t.ElapsedString())

	switch level {
	case "debug":
		Debug(fullMessage)
	case "info":
		Info(fullMessage)
	case "warn":
		Warn(fullMessage)
	case "error":
		Error(fullMessage)
	default:
		Info(fullMessage)
	}
}

// LogSlowWithThreshold 记录执行时间，如果超过阈值则记录为警告
func (t *Timer) LogSlowWithThreshold(message string, thresholdMs int64) {
	elapsed := t.ElapsedMs()
	if elapsed > thresholdMs {
		Warn(fmt.Sprintf("[性能][%s] %s | 耗时: %dms (警告阈值: %dms)", t.name, message, elapsed, thresholdMs))
	} else {
		Info(fmt.Sprintf("[性能][%s] %s | 耗时: %dms", t.name, message, elapsed))
	}
}

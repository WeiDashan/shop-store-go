package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path/filepath"
	"time"
)

func InitLogger() *zap.SugaredLogger {
	logMode := zapcore.DebugLevel
	if !viper.GetBool("mode.develop") {
		logMode = zapcore.InfoLevel
	}
	core := zapcore.NewCore(getEncoder(), zapcore.NewMultiWriteSyncer(getWriteSyncer(), zapcore.AddSync(os.Stdout)), logMode)
	return zap.New(core).Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Local().Format(time.DateTime))
	}
	return zapcore.NewJSONEncoder(encoderConfig)
}
func getWriteSyncer() zapcore.WriteSyncer {
	separator := string(filepath.Separator)
	rootDir, _ := os.Getwd()
	logFilePath := rootDir + separator + "log" + separator + time.Now().Format(time.DateOnly) + ".txt"
	fmt.Println(logFilePath)

	// lumberjack
	lumberjackSyncer := &lumberjack.Logger{
		Filename:   logFilePath,
		MaxSize:    viper.GetInt("log.MaxSize"),    // 日志文件最大的尺寸(M)，超限后开始自动分割
		MaxBackups: viper.GetInt("log.MaxBackups"), // 保留旧文件的最大个数
		MaxAge:     viper.GetInt("log.MaxAge"),     // 保留旧文件的最大天数
		Compress:   false,
	}

	return zapcore.AddSync(lumberjackSyncer)

}

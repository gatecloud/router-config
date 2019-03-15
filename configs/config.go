package configs

import (
	"errors"
	"os"
	"reflect"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type LocalConfig struct {
	Port             string `json:"port"`
	DbEngine         string `json:"db_engine"`
	DbConn           string `json:"db_conn"`
	Production       bool   `json:"production"`
	AWSS3Key         string `json:"aws_s3_key"`
	AWSS3Secret      string `json:"aws_s3_secret"`
	AWSS3Region      string `json:"aws_s3_region"`
	AWSS3ImageDomain string `json:"aws_s3_image_domain"`
	LegalFileExt     string `json:"legal_file_ext"`
}

// Configuration is proxy's global configuration
var Configuration LocalConfig

func InitConfig() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	rfConfig := reflect.ValueOf(&Configuration).Elem()
	if !rfConfig.IsValid() {
		return errors.New("reflect configuration error")
	}

	for i := 0; i < rfConfig.NumField(); i++ {
		field := rfConfig.Field(i)
		key := rfConfig.Type().Field(i).Tag.Get("json")
		value := os.Getenv(key)
		switch field.Kind() {
		case reflect.Int:
			v, err := strconv.ParseInt(value, 0, 64)
			if err != nil {
				return err
			}
			field.SetInt(int64(v))
		case reflect.Float64:
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return err
			}
			field.SetFloat(v)
		case reflect.String:
			field.SetString(value)
		case reflect.Bool:
			v, err := strconv.ParseBool(value)
			if err != nil {
				return err
			}
			field.SetBool(v)
		case reflect.Ptr:
			v, err := time.LoadLocation(value)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(v))
		}
	}
	return nil
}

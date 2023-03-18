package app

import (
	"github.com/core-go/log"
	mid "github.com/core-go/log/middleware"
	"github.com/core-go/sql"
)

type Root struct {
	Server     ServerConf    `mapstructure:"server"`
	Sql        sql.Config    `mapstructure:"sql"`
	Log        log.Config    `mapstructure:"log"`
	MiddleWare mid.LogConfig `mapstructure:"middleware"`
}
type ServerConf struct {
	Name              string         `yaml:"name" mapstructure:"name" json:"name,omitempty" gorm:"column:name" bson:"name,omitempty" dynamodbav:"name,omitempty" firestore:"name,omitempty"`
	Port              *int64         `yaml:"port" mapstructure:"port" json:"port,omitempty" gorm:"column:port" bson:"port,omitempty" dynamodbav:"port,omitempty" firestore:"port,omitempty"`
}

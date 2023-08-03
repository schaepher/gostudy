package gostudy

import (
	"fmt"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDB struct {
}

func (m *MongoDB) Aggregate(p mongo.Pipeline, c string) (rs []map[string]interface{}, err error) {
	return
}

func TestMongo(t *testing.T) {
	GetMongoDataFiled(time.Now(), time.Now(), &MongoDB{})
}

func GetMongoDataFiled(start, end time.Time, mgo *MongoDB) int64 {
	pipeMatch := bson.D{
		{
			"$match",
			bson.D{
				{"date", time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, time.Local).Unix()},
				{"provider", 8},
			},
		},
	}

	// 2. 对象转数组，便于后续将数据部分独立出来
	pipeProjectRootToArray := bson.D{{
		"$project", bson.D{
			{"date", 1},
			{"data", bson.D{{"$objectToArray", "$$ROOT"}}},
		}}}

	// 3. 把数组里不是数据的字段过滤掉
	pipeProjectFilterData := bson.D{{
		"$project", bson.D{
			{"date", 1},
			{"data", bson.D{
				{"$filter", bson.D{
					{"input", "$data"},
					{"as", "item"},
					{"cond", bson.D{{
						"$and", bson.A{
							bson.D{{"$gte", bson.A{"$$item.k", "0000"}}},
							bson.D{{"$lte", bson.A{"$$item.k", "2359"}}},
						},
					}}},
				},
				}},
			}},
	}}

	// 4. 每一分钟一条数据
	pipeUnwind := bson.D{{"$unwind", "$data"}}

	// 5. 小时分钟的格式转换为时间戳
	pipeDateToTimestamp := bson.D{
		{"$project", bson.D{
			{"timestamp", bson.D{
				{"$toInt", bson.D{
					{"$add", bson.A{
						"$date",
						bson.D{
							{"$multiply", bson.A{
								bson.D{
									{"$floor", bson.D{
										{"$divide", bson.A{
											bson.D{{"$toInt", "$data.k"}},
											100},
										}},
									},
								},
								3600},
							}},
						bson.D{
							{"$multiply", bson.A{
								bson.D{
									{"$mod", bson.A{
										bson.D{{"$toInt", "$data.k"}},
										100},
									}},
								60},
							}}},
					}},
				}},
			},
			{"date", 1},
			{"data", "$data.v"}},
		},
	}

	// 6. 使用起止时间缩小数据集
	pipeMatchTimeRange := bson.D{
		{"$match", bson.D{
			{"timestamp", bson.D{
				{"$gte", start.Unix()},
				{"$lt", end.Unix()}},
			}},
		},
	}

	pipeSumDataFiled := bson.D{{
		"$group", bson.D{
			{"_id", "$date"},
			{"DataFiled", bson.D{{"$sum", "$data.DataFiled"}}},
		},
	}}

	pipeline := mongo.Pipeline{
		pipeMatch,
		pipeProjectRootToArray,
		pipeProjectFilterData,
		pipeUnwind,
		pipeDateToTimestamp,
		pipeMatchTimeRange,
		pipeSumDataFiled,
	}

	rs, err := mgo.Aggregate(pipeline, "collection")
	if err != nil {
		fmt.Println(err)
		return 0
	}

	all := int64(0)
	for _, element := range rs {
		all += element["DataFiled"].(int64)
	}

	return all
}

package helpers

import (
	"fmt"
	moptions "go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
	"net/url"
	"strconv"
	"strings"
)

type EnumOperator string

const (
	Equal              EnumOperator = "="
	GreaterThan        EnumOperator = ">"
	LessThan           EnumOperator = "<"
	GreaterThanOrEqual EnumOperator = ">="
	LessThanOrEqual    EnumOperator = "<="
	In                 EnumOperator = "in"
	Contains           EnumOperator = "contains"
)

const (
	defaultLimit   = 10
	defaultPage    = 1
	defaultSort    = "asc"
	defaultSortDir = "created_at"
)

type Filter struct {
	Field    string
	Operator EnumOperator
	Value    any
}

type QueryResult struct {
	Query       bson.M
	FindOptions *moptions.FindOptions
}

func GenerateQuery(filters []Filter, options bson.M) QueryResult {
	query := bson.M{}
	mongoOptions := moptions.Find()

	for _, filter := range filters {
		switch filter.Operator {
		case "=":
			query[filter.Field] = filter.Value
		case ">":
			query[filter.Field] = bson.M{"$gt": filter.Value}
		case "<":
			query[filter.Field] = bson.M{"$lt": filter.Value}
		case ">=":
			query[filter.Field] = bson.M{"$gte": filter.Value}
		case "<=":
			query[filter.Field] = bson.M{"$lte": filter.Value}
		case "in":
			query[filter.Field] = bson.M{"$in": filter.Value}
		case "contains":
			query[filter.Field] = bson.M{"$regex": fmt.Sprintf(".*%s.*", filter.Value)}
		default:
		}
	}
	page := GetOptions[int64](options, "page", defaultPage)
	limit := GetOptions[int64](options, "limit", defaultLimit)
	sort := GetOptions[string](options, "sort", defaultSort)
	sortBy := GetOptions[string](options, "dir", defaultSortDir)

	mongoOptions.SetSkip(page * limit)
	mongoOptions.SetLimit(limit)

	sortQ := bson.M{}
	if strings.ToLower(sort) == "desc" {
		sortQ[sortBy] = -1
	} else {
		sortQ[sortBy] = 1
	}
	mongoOptions.SetSort(sortQ)
	return QueryResult{Query: query, FindOptions: mongoOptions}
}

func SetPaginationQuery(query url.Values, optionsRepo map[string]interface{}) {
	//paginate
	pageQuery := query.Get("page")
	pageInt, err := strconv.Atoi(pageQuery)
	if err != nil || pageInt <= 0 {
		pageInt = defaultPage
	}

	limitQuery := query.Get("limit")
	limitInt, err := strconv.Atoi(limitQuery)
	if err != nil || limitInt <= 0 {
		limitInt = defaultLimit
	}
	optionsRepo["limit"] = int64(limitInt)
	optionsRepo["page"] = (int64(pageInt) - 1) * int64(limitInt)

	// Validate allowed_sort
	if sort := query.Get("sort"); sort != "" {
		optionsRepo["sort"] = sort
		if dir := query.Get("dir"); dir != "" {
			optionsRepo["dir"] = dir
		}
	}
}

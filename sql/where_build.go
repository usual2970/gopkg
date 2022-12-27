package sql

import (
	"fmt"
	"strings"
)

type NullType byte

const (
	_ NullType = iota
	// IsNull the same as `is null`
	IsNull
	// IsNotNull the same as `is not null`
	IsNotNull
)

// WhereBuild build where
func WhereBuild(where map[string]interface{}) (string, []interface{}, error) {
	whereSQL := ""
	vals := []interface{}{}
	var err error
	for k, v := range where {
		ks := strings.Split(k, " ")
		if len(ks) > 2 {
			return "", nil, fmt.Errorf("Error in query condition: %s. ", k)
		}

		if whereSQL != "" {
			whereSQL += " AND "
		}

		switch len(ks) {
		case 1:
			switch v.(type) {
			case NullType:
				if v == IsNotNull {
					whereSQL += fmt.Sprint(k, " IS NOT NULL")
				} else {
					whereSQL += fmt.Sprint(k, " IS NULL")
				}
			case map[string]interface{}:
				if k == "raw" {
					whereSQL += v.(map[string]interface{})["sql"].(string)
					vals = append(vals, v.(map[string]interface{})["bind"].([]interface{})...)
				}
			default:
				whereSQL += fmt.Sprint(k, "=?")
				vals = append(vals, v)
			}
		case 2:
			k = ks[0]
			switch ks[1] {
			case "=":
				whereSQL += fmt.Sprint(k, "=?")
				vals = append(vals, v)
			case ">":
				whereSQL += fmt.Sprint(k, ">?")
				vals = append(vals, v)
			case ">=":
				whereSQL += fmt.Sprint(k, ">=?")
				vals = append(vals, v)
			case "<":
				whereSQL += fmt.Sprint(k, "<?")
				vals = append(vals, v)
			case "<=":
				whereSQL += fmt.Sprint(k, "<=?")
				vals = append(vals, v)
			case "!=":
				whereSQL += fmt.Sprint(k, "!=?")
				vals = append(vals, v)
			case "<>":
				whereSQL += fmt.Sprint(k, "!=?")
				vals = append(vals, v)
			case "in":
				whereSQL += fmt.Sprint(k, " in (?)")
				vals = append(vals, v)
			case "like":
				whereSQL += fmt.Sprint(k, " like ?")
				vals = append(vals, v)
			}
		}
	}
	return whereSQL, vals, err
}

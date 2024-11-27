package golwf_sql

import (
	"fmt"
	"strings"
)

type WhereCondition struct {
	Field        string
	Operator     string
	Value        interface{}
	Values       []interface{}
	Logic        string
	SubCondition []WhereCondition
}

type QueryParams struct {
	Where   []WhereCondition
	Having  []WhereCondition
	OrderBy []string
	GroupBy []string
	Limit   *int
	Offset  *int
}

func BuildWhereCondition(params QueryParams) (string, interface{}) {
	var query strings.Builder
	var values []interface{}

	if len(params.Where) > 0 {
		query.WriteString(" WHERE ")
		buildConditions(&query, params.Where, &values, "")
	}

	if len(params.GroupBy) > 0 {
		query.WriteString(" GROUP BY ")
		query.WriteString(strings.Join(params.GroupBy, ","))
	}

	if len(params.Having) > 0 {
		query.WriteString(" HAVING ")
		buildConditions(&query, params.Having, &values, "")
	}

	if len(params.OrderBy) > 0 {
		query.WriteString(" ORDER BY ")
		query.WriteString(strings.Join(params.OrderBy, ","))
	}

	if params.Limit != nil {
		query.WriteString(fmt.Sprintf(" LIMIT %d", *params.Limit))
	}

	if params.Offset != nil {
		query.WriteString(fmt.Sprintf(" OFFSET %d", *params.Offset))
	}
	
	return query.String(), values
}

func buildConditions(query *strings.Builder, conditions []WhereCondition, values *[]interface{}, indent string) {
	for i, cond := range conditions {
		if i > 0 {
			query.WriteString(" " + cond.Logic + " ")
		}

		if len(cond.SubCondition) > 0 {
			query.WriteString("(")
			buildConditions(query, cond.SubCondition, values, indent+" ")
			query.WriteString(")")
		}

		if cond.Operator == "IN" || cond.Operator == "NOT IN" {
			query.WriteString(cond.Field + " " + cond.Operator + "(?")
			*values = append(*values, cond.Values[0])
			for _, v := range cond.Values[1:] {
				query.WriteString(",?")
				*values = append(*values, v)
			}
			query.WriteString(")")
		} else if cond.Operator == "BETWEEN" {
			if len(cond.Values) == 2 {
				query.WriteString(cond.Field + " BETWEEN ? AND ?")
				vals := cond.Values
				*values = append(*values, vals[0], vals[1])
			} else {
				panic("For the \"BETWEEN\" operator two condition values are required")
			}
		} else {
			query.WriteString(cond.Field + " " + cond.Operator + "?")
			*values = append(*values, cond.Value)
		}
	}
}

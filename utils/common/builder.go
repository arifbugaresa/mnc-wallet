package common

import "github.com/doug-martin/goqu/v9"

func BuildDatasetGetListWithParams(dataset *goqu.SelectDataset, param DefaultListRequest) *goqu.SelectDataset {
	if param.Search.Field != "" && param.Search.Value != "" {
		dataset = dataset.Where(
			goqu.I(param.Search.Field).ILike("%" + param.Search.Value + "%"),
		)
	}

	if param.Sort.Field != "" && param.Sort.Order != "" {
		if param.Sort.Order == "asc" {
			dataset = dataset.Order(goqu.I(param.Sort.Field).Asc())
		} else {
			dataset = dataset.Order(goqu.I(param.Sort.Field).Desc())
		}
	}

	if param.Page != 0 && param.Limit != 0 {
		offset := (param.Page - 1) * param.Limit
		dataset = dataset.Limit(uint(param.Limit)).Offset(uint(offset))
	}

	return dataset
}

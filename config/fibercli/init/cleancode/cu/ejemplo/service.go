package ejemplo

import (
	"%FIBERCLI%/infra"
)

type Service struct {
	r *infra.RepositorySql
}

func (t Service) NewService(r *infra.RepositorySql) error {
	t.r = r
	return nil
}
func (t Service) CreateFilter(awsaccount_id uint, region string) Dto {
	return Dto{
		Awsaccount_id: awsaccount_id,
		Region:        region,
	}
}
func (t Service) CreateWithField(filter interface{}, v interface{}) error {
	return t.r.CreateField(filter, v)
}
func (t Service) ListAllFor1(awsid uint, region string, v interface{}) error {
	filters := t.CreateFilter(awsid, region)
	return t.r.ListAllFor(filters, v)
}

package zube

import (
	"reflect"
	"testing"

	"github.com/platogo/zube-cli/zube/models"
)

func TestGetMembersByNames(t *testing.T) {
	type args struct {
		names   []string
		members []models.Member
	}

	members := []models.Member{
		{Person: models.Person{Name: "dan"}},
		{Person: models.Person{Name: "pete"}},
		{Person: models.Person{Name: "john"}},
	}

	tests := []struct {
		name string
		args args
		want []models.Member
	}{
		{"check returns members for given names", args{[]string{"dan", "pete"}, members}, []models.Member{
			{Person: models.Person{Name: "dan"}},
			{Person: models.Person{Name: "pete"}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMembersByNames(tt.args.names, tt.args.members); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMembersByNames() = %v, want %v", got, tt.want)
			}
		})
	}
}

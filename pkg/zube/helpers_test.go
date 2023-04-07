package zube

import (
	"reflect"
	"testing"

	"github.com/platogo/zube/models"
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

func TestCardUrl(t *testing.T) {
	type args struct {
		account *models.Account
		project *models.Project
		card    *models.Card
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"check that returns correct url",
			args{&models.Account{Slug: "platogo"}, &models.Project{Slug: "server"}, &models.Card{Number: 1234}},
			"https://zube.io/platogo/server/c/1234"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CardUrl(tt.args.account, tt.args.project, tt.args.card); got != tt.want {
				t.Errorf("CardUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

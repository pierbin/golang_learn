package mockActivities_test

import (
	"learnGo/customizePackages/mockActivities/infra"

	"learnGo/customizePackages/mockActivities"

	"reflect"
	"testing"
)

func TestApp_Apply(t *testing.T) {
	type deps struct {
		applicationRepo mockActivities.ApplicationRepo
		emailSender     mockActivities.EmailSender
	}

	// mock applicationRepo interface
	applicationRepoEmpty := infra.ApplicationRepoMock{
		FindAllFunc: func(tutor string) ([]mockActivities.Application, error) { return nil, nil },
		SaveFunc:    func(application mockActivities.Application) error { return nil },
	}

	// mock emailSender interface
	emailSenderSuccess := infra.EmailSenderMock{
		SendFunc: func(application mockActivities.Application) error { return nil },
	}

	tests := []struct {
		name    string
		deps    deps
		tutor   string
		student string
		what    string
		want    mockActivities.Application
		wantErr bool
	}{
		{
			name:    "first student",
			deps:    deps{applicationRepo: applicationRepoEmpty, emailSender: emailSenderSuccess},
			tutor:   "some@tutor.com",
			student: "Student Full name",
			what:    "soccer",
			want:    mockActivities.Application{TutorEmail: "some@tutor.com", Cost: 10, ActivityName: "soccer", Student: "Student Full name"},
			wantErr: false,
		},
		{
			name:    "wrong activity",
			deps:    deps{applicationRepo: applicationRepoEmpty, emailSender: emailSenderSuccess},
			tutor:   "some@tutor.com",
			student: "Student Full name",
			what:    "undefined",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &mockActivities.App{ApplicationRepo: tt.deps.applicationRepo, EmailSender: tt.deps.emailSender}
			got, err := a.Apply(tt.tutor, tt.student, tt.what)
			if (err != nil) != tt.wantErr {
				t.Errorf("Apply() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Apply() got = %v, want %v", got, tt.want)
			}
		})
	}
}

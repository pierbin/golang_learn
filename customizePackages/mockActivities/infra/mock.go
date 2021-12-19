package infra

import "learnGo/customizePackages/mockActivities"

type EmailSenderMock struct {
	SendFunc func(application mockActivities.Application) error
}

func (e EmailSenderMock) Send(application mockActivities.Application) error {
	return e.SendFunc(application)
}

type ApplicationRepoMock struct {
	FindAllFunc func(tutor string) ([]mockActivities.Application, error)
	SaveFunc    func(application mockActivities.Application) error
}

func (a ApplicationRepoMock) FindAll(tutor string) ([]mockActivities.Application, error) {
	return a.FindAllFunc(tutor)
}

func (a ApplicationRepoMock) Save(application mockActivities.Application) error {
	return a.SaveFunc(application)
}

//emailSenderSuccess := infra.EmailSenderMock{
//		SendFunc: func (application mockActivities.Application) error {
//		return nil
//	},
//}

//emailSenderFail := infra.EmailSenderMock{
//		SendFunc: func (application mockActivities.Application) error {
//		return mockActivities.ErrInternal
//	},
//}

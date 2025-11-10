/*
Package models provides data types for representing CVs in memory.
*/
package models

import "fmt"

type Month int

const (
	January Month = 1
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

func (m Month) String() string {
	switch m {
	case 1:
		return "January"
	case 2:
		return "February"
	case 3:
		return "March"
	case 4:
		return "April"
	case 5:
		return "May"
	case 6:
		return "June"
	case 7:
		return "July"
	case 8:
		return "August"
	case 9:
		return "September"
	case 10:
		return "October"
	case 11:
		return "November"
	case 12:
		return "December"
	}
	return ""
}

type Date struct {
	Year  int   `xml:"year,attr"`
	Month Month `xml:"month,attr"`
}

type ExperienceData struct {
	Business  string   `xml:"business,attr"`
	Title     string   `xml:"title,attr"`
	StartDate Date     `xml:"start-date"`
	EndDate   Date     `xml:"end-date"`
	Details   []string `xml:"detail"`
}

type EducationData struct {
	School    string   `xml:"school,attr"`
	Degree    string   `xml:"degree,attr"`
	Details   []string `xml:"detail"`
	StartDate Date     `xml:"start-date"`
	EndDate   Date     `xml:"end-date"`
}

type CertificationData struct {
	Name string `xml:"name,attr"`
}

type HeaderData struct {
	Name    string `xml:"name,attr"`
	Email   string `xml:"email,attr"`
	Github  string `xml:"github,attr"`
	Site    string `xml:"site,attr"`
	Phone 	string `xml:"phone,attr"`
	Summary string `xml:"summary"`
}

type SkillsData struct {
	Group  string   `xml:"group,attr"`
	Skills []string `xml:"skill"`
}

type ResumeContentData struct {
	Header         HeaderData          `xml:"header"`
	Experience     []ExperienceData    `xml:"experience"`
	Education      []EducationData     `xml:"education"`
	Certifications []CertificationData `xml:"certification"`
	Skills         []SkillsData        `xml:"skills"`
}

type HeaderView struct {
	Name    string
	Email   string
	Github  string
	Site    string
	Phone		string
	Summary string
}

func MakeHeaderView(data HeaderData) HeaderView {
	return HeaderView{
		Name:    data.Name,
		Email:   data.Email,
		Github:  data.Github,
		Site:    data.Site,
		Phone:	 data.Phone,
		Summary: data.Summary,
	}
}

type ExperienceView struct {
	Business  string
	Title     string
	StartDate string
	EndDate   string
	Details   []string
}

func (d Date) String() string {
	if d == (Date{}) {
		return "Present"
	}

	return fmt.Sprint(d.Month.String(), " ", d.Year)
}

func MakeExperienceView(data ExperienceData) ExperienceView {
	return ExperienceView{
		Business:  data.Business,
		Title:     data.Title,
		StartDate: data.StartDate.String(),
		EndDate:   data.EndDate.String(),
		Details:   data.Details,
	}
}

type EducationView struct {
	School    string
	Degree    string
	Details   []string
	StartDate string
	EndDate   string
}

func MakeEducationView(data EducationData) EducationView {
	return EducationView{
		School:    data.School,
		Degree:    data.Degree,
		Details:   data.Details,
		StartDate: data.StartDate.String(),
		EndDate:   data.EndDate.String(),
	}
}

type CertificationView struct {
	Name string
}

func MakeCertificationView(data CertificationData) CertificationView {
	return CertificationView{
		Name: data.Name,
	}
}

type SkillsView struct {
	Group  string
	Skills []string
}

func MakeSkillsView(data SkillsData) SkillsView {
	return SkillsView{
		Group:  data.Group,
		Skills: data.Skills,
	}
}

type ResumeContentView struct {
	Header         HeaderView
	Experience     []ExperienceView
	Education      []EducationView
	Certifications []CertificationView
	Skills         []SkillsView
}

func MakeResumeContentView(data ResumeContentData) ResumeContentView {
	experienceView := []ExperienceView{}
	for _, e := range data.Experience {
		eView := MakeExperienceView(e)
		experienceView = append(experienceView, eView)
	}

	educationView := []EducationView{}
	for _, e := range data.Education {
		eView := MakeEducationView(e)
		educationView = append(educationView, eView)
	}

	certificationsView := []CertificationView{}
	for _, c := range data.Certifications {
		cView := MakeCertificationView(c)
		certificationsView = append(certificationsView, cView)
	}

	skillsView := []SkillsView{}
	for _, s := range data.Skills {
		sView := MakeSkillsView(s)
		skillsView = append(skillsView, sView)
	}

	return ResumeContentView{
		Header:         MakeHeaderView(data.Header),
		Experience:     experienceView,
		Education:      educationView,
		Certifications: certificationsView,
		Skills:         skillsView,
	}
}
